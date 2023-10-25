package ibcratelimit

import (
	"encoding/json"
	"strings"

	errorsmod "cosmossdk.io/errors"

	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	transfertypes "github.com/cosmos/ibc-go/v6/modules/apps/transfer/types"
	clienttypes "github.com/cosmos/ibc-go/v6/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/v6/modules/core/04-channel/types"
	porttypes "github.com/cosmos/ibc-go/v6/modules/core/05-port/types"
	"github.com/cosmos/ibc-go/v6/modules/core/exported"

	"github.com/provenance-io/provenance/x/ibcratelimit/keeper"
	"github.com/provenance-io/provenance/x/ibcratelimit/osmosis/osmoutils"
	"github.com/provenance-io/provenance/x/ibcratelimit/types"
)

var (
	_ porttypes.Middleware = &IBCMiddleware{}
)

type IBCMiddleware struct {
	app            porttypes.IBCModule
	keeper         *keeper.Keeper
	channel        porttypes.ICS4Wrapper
	accountKeeper  *authkeeper.AccountKeeper
	bankKeeper     *bankkeeper.BaseKeeper
	ContractKeeper *wasmkeeper.PermissionedKeeper
}

func NewIBCMiddleware(app porttypes.IBCModule,
	channel porttypes.ICS4Wrapper,
	keeper *keeper.Keeper,
	accountKeeper *authkeeper.AccountKeeper,
	contractKeeper *wasmkeeper.PermissionedKeeper,
	bankKeeper *bankkeeper.BaseKeeper) IBCMiddleware {
	return IBCMiddleware{
		app:            app,
		keeper:         keeper,
		channel:        channel,
		accountKeeper:  accountKeeper,
		bankKeeper:     bankKeeper,
		ContractKeeper: contractKeeper,
	}
}

func (im *IBCMiddleware) WithIBCModule(app porttypes.IBCModule) *IBCMiddleware {
	im.app = app
	return im
}

// OnChanOpenInit implements the IBCModule interface
func (im IBCMiddleware) OnChanOpenInit(ctx sdk.Context,
	order channeltypes.Order,
	connectionHops []string,
	portID string,
	channelID string,
	channelCap *capabilitytypes.Capability,
	counterparty channeltypes.Counterparty,
	version string,
) (string, error) {
	return im.app.OnChanOpenInit(
		ctx,
		order,
		connectionHops,
		portID,
		channelID,
		channelCap,
		counterparty,
		version,
	)
}

// OnChanOpenTry implements the IBCModule interface
func (im *IBCMiddleware) OnChanOpenTry(
	ctx sdk.Context,
	order channeltypes.Order,
	connectionHops []string,
	portID,
	channelID string,
	channelCap *capabilitytypes.Capability,
	counterparty channeltypes.Counterparty,
	counterpartyVersion string,
) (string, error) {
	return im.app.OnChanOpenTry(ctx, order, connectionHops, portID, channelID, channelCap, counterparty, counterpartyVersion)
}

// OnChanOpenAck implements the IBCModule interface
func (im *IBCMiddleware) OnChanOpenAck(
	ctx sdk.Context,
	portID,
	channelID string,
	counterpartyChannelID string,
	counterpartyVersion string,
) error {
	// Here we can add initial limits when a new channel is open. For now, they can be added manually on the contract
	return im.app.OnChanOpenAck(ctx, portID, channelID, counterpartyChannelID, counterpartyVersion)
}

// OnChanOpenConfirm implements the IBCModule interface
func (im *IBCMiddleware) OnChanOpenConfirm(
	ctx sdk.Context,
	portID,
	channelID string,
) error {
	// Here we can add initial limits when a new channel is open. For now, they can be added manually on the contract
	return im.app.OnChanOpenConfirm(ctx, portID, channelID)
}

// OnChanCloseInit implements the IBCModule interface
func (im *IBCMiddleware) OnChanCloseInit(
	ctx sdk.Context,
	portID,
	channelID string,
) error {
	// Here we can remove the limits when a new channel is closed. For now, they can remove them  manually on the contract
	return im.app.OnChanCloseInit(ctx, portID, channelID)
}

// OnChanCloseConfirm implements the IBCModule interface
func (im *IBCMiddleware) OnChanCloseConfirm(
	ctx sdk.Context,
	portID,
	channelID string,
) error {
	// Here we can remove the limits when a new channel is closed. For now, they can remove them  manually on the contract
	return im.app.OnChanCloseConfirm(ctx, portID, channelID)
}

func ValidateReceiverAddress(packet exported.PacketI) error {
	var packetData transfertypes.FungibleTokenPacketData
	if err := json.Unmarshal(packet.GetData(), &packetData); err != nil {
		return err
	}
	if len(packetData.Receiver) >= 4096 {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "IBC Receiver address too long. Max supported length is %d", 4096)
	}
	return nil
}

// OnRecvPacket implements the IBCModule interface
func (im *IBCMiddleware) OnRecvPacket(
	ctx sdk.Context,
	packet channeltypes.Packet,
	relayer sdk.AccAddress,
) exported.Acknowledgement {
	if err := ValidateReceiverAddress(packet); err != nil {
		return osmoutils.NewEmitErrorAcknowledgement(ctx, types.ErrBadMessage, err.Error())
	}

	contract := im.keeper.GetContractAddress(ctx)
	if contract == "" {
		// The contract has not been configured. Continue as usual
		return im.app.OnRecvPacket(ctx, packet, relayer)
	}

	err := CheckAndUpdateRateLimits(ctx, im.ContractKeeper, "recv_packet", contract, packet)
	if err != nil {
		if strings.Contains(err.Error(), "rate limit exceeded") {
			return osmoutils.NewEmitErrorAcknowledgement(ctx, types.ErrRateLimitExceeded)
		}
		fullError := errorsmod.Wrap(types.ErrContractError, err.Error())
		return osmoutils.NewEmitErrorAcknowledgement(ctx, fullError)
	}

	// if this returns an Acknowledgement that isn't successful, all state changes are discarded
	return im.app.OnRecvPacket(ctx, packet, relayer)
}

// OnAcknowledgementPacket implements the IBCModule interface
func (im *IBCMiddleware) OnAcknowledgementPacket(
	ctx sdk.Context,
	packet channeltypes.Packet,
	acknowledgement []byte,
	relayer sdk.AccAddress,
) error {
	var ack channeltypes.Acknowledgement
	if err := json.Unmarshal(acknowledgement, &ack); err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrUnknownRequest, "cannot unmarshal ICS-20 transfer packet acknowledgement: %v", err)
	}

	if osmoutils.IsAckError(acknowledgement) {
		err := im.RevertSentPacket(ctx, packet) // If there is an error here we should still handle the ack
		if err != nil {
			ctx.EventManager().EmitEvent(
				sdk.NewEvent(
					types.EventBadRevert,
					sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
					sdk.NewAttribute(types.AttributeKeyFailureType, "acknowledgment"),
					sdk.NewAttribute(types.AttributeKeyPacket, string(packet.GetData())),
					sdk.NewAttribute(types.AttributeKeyAck, string(acknowledgement)),
				),
			)
		}
	}

	return im.app.OnAcknowledgementPacket(ctx, packet, acknowledgement, relayer)
}

// OnTimeoutPacket implements the IBCModule interface
func (im *IBCMiddleware) OnTimeoutPacket(
	ctx sdk.Context,
	packet channeltypes.Packet,
	relayer sdk.AccAddress,
) error {
	err := im.RevertSentPacket(ctx, packet) // If there is an error here we should still handle the timeout
	if err != nil {
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				types.EventBadRevert,
				sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
				sdk.NewAttribute(types.AttributeKeyFailureType, "timeout"),
				sdk.NewAttribute(types.AttributeKeyPacket, string(packet.GetData())),
			),
		)
	}
	return im.app.OnTimeoutPacket(ctx, packet, relayer)
}

// RevertSentPacket Notifies the contract that a sent packet wasn't properly received
func (im *IBCMiddleware) RevertSentPacket(
	ctx sdk.Context,
	packet exported.PacketI,
) error {
	contract := im.keeper.GetContractAddress(ctx)
	if contract == "" {
		// The contract has not been configured. Continue as usual
		return nil
	}

	return UndoSendRateLimit(
		ctx,
		im.ContractKeeper,
		contract,
		packet,
	)
}

// SendPacket implements the ICS4 interface and is called when sending packets.
// This method retrieves the contract from the middleware's parameters and checks if the limits have been exceeded for
// the current transfer, in which case it returns an error preventing the IBC send from taking place.
// If the contract param is not configured, or the contract doesn't have a configuration for the (channel+denom) being
// used, transfers are not prevented and handled by the wrapped IBC app
func (im *IBCMiddleware) SendPacket(
	ctx sdk.Context,
	chanCap *capabilitytypes.Capability,
	sourcePort string,
	sourceChannel string,
	timeoutHeight clienttypes.Height,
	timeoutTimestamp uint64,
	data []byte,
) (sequence uint64, err error) {
	contract := im.keeper.GetContractAddress(ctx)
	if contract == "" {
		// The contract has not been configured. Continue as usual
		return im.channel.SendPacket(ctx, chanCap, sourcePort, sourceChannel, timeoutHeight, timeoutTimestamp, data)
	}

	// We need the full packet so the contract can process it. If it can't be cast to a channeltypes.Packet, this
	// should fail. The only reason that would happen is if another middleware is modifying the packet, though. In
	// that case we can modify the middleware order or change this cast so we have all the data we need.
	// TODO Verify we don't need destination port or channel
	packet := channeltypes.NewPacket(
		data,
		sequence,
		sourcePort,
		sourceChannel,
		"",
		"",
		timeoutHeight,
		timeoutTimestamp,
	)

	err = CheckAndUpdateRateLimits(ctx, im.ContractKeeper, "send_packet", contract, packet)
	if err != nil {
		return 0, errorsmod.Wrap(err, "rate limit SendPacket failed to authorize transfer")
	}

	return im.channel.SendPacket(ctx, chanCap, sourcePort, sourceChannel, timeoutHeight, timeoutTimestamp, data)
}

// WriteAcknowledgement implements the ICS4 Wrapper interface
func (im *IBCMiddleware) WriteAcknowledgement(
	ctx sdk.Context,
	chanCap *capabilitytypes.Capability,
	packet exported.PacketI,
	ack exported.Acknowledgement,
) error {
	return im.channel.WriteAcknowledgement(ctx, chanCap, packet, ack)
}

func (im *IBCMiddleware) GetAppVersion(ctx sdk.Context, portID, channelID string) (string, bool) {
	return im.channel.GetAppVersion(ctx, portID, channelID)
}
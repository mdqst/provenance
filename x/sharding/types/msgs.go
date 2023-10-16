package types

import (
	fmt "fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	_ sdk.Msg = &MsgReadRequest{}
	_ sdk.Msg = &MsgWriteRequest{}
	_ sdk.Msg = &MsgUpdateRequest{}
)

// NewMsgRead creates a new NewMsgRead
func NewMsgRead(authority string, owner, name, color, spots, full, group, sharded bool, iterations uint64) *MsgReadRequest {
	return &MsgReadRequest{
		Authority:   authority,
		OwnerRead:   owner,
		NameRead:    name,
		ColorRead:   color,
		SpotsRead:   spots,
		FullRead:    full,
		GroupRead:   group,
		ShardedRead: sharded,
		Iterations:  iterations,
	}
}

// GetSigners indicates that the message must have been signed by the parent.
func (msg MsgReadRequest) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(msg.Authority)}
}

// ValidateBasic runs stateless validation checks on the message.
func (msg MsgReadRequest) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Authority); err != nil {
		return fmt.Errorf("invalid authority address: %w", err)
	}
	return nil
}

// NewMsgWrite creates a new MsgWriteRequest
func NewMsgWrite(authority string, owner, name, color, spots, full, group, sharded bool, iterations uint64) *MsgWriteRequest {
	return &MsgWriteRequest{
		Authority:    authority,
		OwnerWrite:   owner,
		NameWrite:    name,
		ColorWrite:   color,
		SpotsWrite:   spots,
		FullWrite:    full,
		GroupWrite:   group,
		ShardedWrite: sharded,
		Iterations:   iterations,
	}
}

// GetSigners indicates that the message must have been signed by the parent.
func (msg MsgWriteRequest) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(msg.Authority)}
}

// ValidateBasic runs stateless validation checks on the message.
func (msg MsgWriteRequest) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Authority); err != nil {
		return fmt.Errorf("invalid authority address: %w", err)
	}
	return nil
}

// NewMsgUpdate creates a new MsgUpdateRequest
func NewMsgUpdate(authority string) *MsgUpdateRequest {
	return &MsgUpdateRequest{
		Authority: authority,
	}
}

// GetSigners indicates that the message must have been signed by the parent.
func (msg MsgUpdateRequest) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(msg.Authority)}
}

// ValidateBasic runs stateless validation checks on the message.
func (msg MsgUpdateRequest) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Authority); err != nil {
		return fmt.Errorf("invalid authority address: %w", err)
	}
	return nil
}

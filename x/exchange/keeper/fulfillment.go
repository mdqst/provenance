package keeper

import (
	"errors"
	"fmt"

	"github.com/gogo/protobuf/proto"

	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	"github.com/provenance-io/provenance/x/exchange"
)

// sumAssetsAndPrice gets the sum of assets, and the sum of prices of the provided orders.
func sumAssetsAndPrice(orders []*exchange.Order) (sdk.Coins, sdk.Coins) {
	var totalAssets, totalPrice sdk.Coins
	for _, order := range orders {
		totalAssets = totalAssets.Add(order.GetAssets())
		totalPrice = totalPrice.Add(order.GetPrice())
	}
	return totalAssets, totalPrice
}

// validateAcceptingOrdersAndCanUserSettle returns an error if the market isn't active or doesn't allow user settlement.
func validateAcceptingOrdersAndCanUserSettle(store sdk.KVStore, marketID uint32) error {
	if err := validateMarketIsAcceptingOrders(store, marketID); err != nil {
		return err
	}
	if !isUserSettlementAllowed(store, marketID) {
		return fmt.Errorf("market %d does not allow user settlement", marketID)
	}
	return nil
}

// FillBids settles one or more bid orders for a seller.
func (k Keeper) FillBids(ctx sdk.Context, msg *exchange.MsgFillBidsRequest) error {
	if err := msg.ValidateBasic(); err != nil {
		return err
	}

	marketID := msg.MarketId
	store := k.getStore(ctx)

	if err := validateAcceptingOrdersAndCanUserSettle(store, marketID); err != nil {
		return err
	}
	seller := sdk.MustAccAddressFromBech32(msg.Seller)
	if err := k.validateUserCanCreateAsk(ctx, marketID, seller); err != nil {
		return err
	}
	if err := validateCreateAskFees(store, marketID, msg.AskOrderCreationFee, msg.SellerSettlementFlatFee); err != nil {
		return err
	}

	orders, oerrs := k.getBidOrders(store, marketID, msg.BidOrderIds, msg.Seller)
	if oerrs != nil {
		return oerrs
	}

	totalAssets, totalPrice := sumAssetsAndPrice(orders)
	if !exchange.CoinsEquals(totalAssets, msg.TotalAssets) {
		return fmt.Errorf("total assets %q does not equal sum of bid order assets %q", msg.TotalAssets, totalAssets)
	}

	var totalSellerFee sdk.Coins
	if msg.SellerSettlementFlatFee != nil {
		totalSellerFee = totalSellerFee.Add(*msg.SellerSettlementFlatFee)
	}

	var errs []error
	feeAddrIdx := exchange.NewIndexedAddrAmts()
	assetsAddrIdx := exchange.NewIndexedAddrAmts()
	priceAddrIdx := exchange.NewIndexedAddrAmts()
	settlement := &exchange.Settlement{FullyFilledOrders: make([]*exchange.FilledOrder, 0, len(msg.BidOrderIds))}
	for _, order := range orders {
		bidOrder := order.GetBidOrder()
		buyer := bidOrder.Buyer
		assets := bidOrder.Assets
		price := bidOrder.Price
		buyerFees := bidOrder.BuyerSettlementFees

		assetsAddrIdx.Add(buyer, assets)
		priceAddrIdx.Add(buyer, price)
		feeAddrIdx.Add(buyer, buyerFees...)
		settlement.FullyFilledOrders = append(settlement.FullyFilledOrders, exchange.NewFilledOrder(order, price, buyerFees))
	}

	for _, price := range totalPrice {
		sellerRatioFee, rerr := calculateSellerSettlementRatioFee(store, marketID, price)
		if rerr != nil {
			errs = append(errs, fmt.Errorf("error calculating seller settlement ratio fee: %w", rerr))
		}
		if sellerRatioFee != nil {
			totalSellerFee = totalSellerFee.Add(*sellerRatioFee)
		}
	}

	if len(errs) > 0 {
		return errors.Join(errs...)
	}

	feeAddrIdx.Add(msg.Seller, totalSellerFee...)

	settlement.Transfers = []*exchange.Transfer{
		{
			Inputs:  []banktypes.Input{{Address: msg.Seller, Coins: msg.TotalAssets}},
			Outputs: assetsAddrIdx.GetAsOutputs(),
		},
		{
			Inputs:  priceAddrIdx.GetAsInputs(),
			Outputs: []banktypes.Output{{Address: msg.Seller, Coins: totalPrice}},
		},
	}
	settlement.FeeInputs = feeAddrIdx.GetAsInputs()

	if err := k.closeSettlement(ctx, store, marketID, settlement); err != nil {
		return err
	}

	// Collected last so that it's easier for a seller to fill bids without needing those funds first.
	// Collected separately so it's not combined with the seller settlement fees in the events.
	if msg.AskOrderCreationFee != nil {
		if err := k.CollectFee(ctx, marketID, seller, sdk.Coins{*msg.AskOrderCreationFee}); err != nil {
			return fmt.Errorf("error collecting create-ask fee %q: %w", msg.AskOrderCreationFee, err)
		}
	}

	return nil
}

// FillAsks settles one or more ask orders for a buyer.
func (k Keeper) FillAsks(ctx sdk.Context, msg *exchange.MsgFillAsksRequest) error {
	if err := msg.ValidateBasic(); err != nil {
		return err
	}

	marketID := msg.MarketId
	store := k.getStore(ctx)

	if err := validateAcceptingOrdersAndCanUserSettle(store, marketID); err != nil {
		return err
	}
	buyer := sdk.MustAccAddressFromBech32(msg.Buyer)
	if err := k.validateUserCanCreateBid(ctx, marketID, buyer); err != nil {
		return err
	}
	if err := validateCreateBidFees(store, marketID, msg.BidOrderCreationFee, msg.TotalPrice, msg.BuyerSettlementFees); err != nil {
		return err
	}

	orders, oerrs := k.getAskOrders(store, marketID, msg.AskOrderIds, msg.Buyer)
	if oerrs != nil {
		return oerrs
	}

	totalAssets, totalPrice := sumAssetsAndPrice(orders)
	if !exchange.CoinsEquals(totalPrice, sdk.Coins{msg.TotalPrice}) {
		return fmt.Errorf("total price %q does not equal sum of ask order prices %q", msg.TotalPrice, totalPrice)
	}

	var errs []error
	assetsAddrIdx := exchange.NewIndexedAddrAmts()
	priceAddrIdx := exchange.NewIndexedAddrAmts()
	feeAddrIdx := exchange.NewIndexedAddrAmts()
	settlement := &exchange.Settlement{FullyFilledOrders: make([]*exchange.FilledOrder, 0, len(msg.AskOrderIds))}
	for _, order := range orders {
		askOrder := order.GetAskOrder()
		seller := askOrder.Seller
		assets := askOrder.Assets
		price := askOrder.Price
		sellerFees := askOrder.GetSettlementFees()

		sellerRatioFee, rerr := calculateSellerSettlementRatioFee(store, marketID, price)
		if rerr != nil {
			errs = append(errs, fmt.Errorf("error calculating seller settlement ratio fee for order %d: %w",
				order.OrderId, rerr))
		}
		if sellerRatioFee != nil {
			sellerFees = sellerFees.Add(*sellerRatioFee)
		}

		assetsAddrIdx.Add(seller, assets)
		priceAddrIdx.Add(seller, price)
		feeAddrIdx.Add(seller, sellerFees...)
		settlement.FullyFilledOrders = append(settlement.FullyFilledOrders, exchange.NewFilledOrder(order, price, sellerFees))
	}

	if len(errs) > 0 {
		return errors.Join(errs...)
	}

	// Done after the loop so that it's always last like it has to be in FillBids.
	feeAddrIdx.Add(msg.Buyer, msg.BuyerSettlementFees...)

	settlement.Transfers = []*exchange.Transfer{
		{
			Inputs:  assetsAddrIdx.GetAsInputs(),
			Outputs: []banktypes.Output{{Address: msg.Buyer, Coins: totalAssets}},
		},
		{
			Inputs:  []banktypes.Input{{Address: msg.Buyer, Coins: sdk.Coins{msg.TotalPrice}}},
			Outputs: priceAddrIdx.GetAsOutputs(),
		},
	}
	settlement.FeeInputs = feeAddrIdx.GetAsInputs()

	if err := k.closeSettlement(ctx, store, marketID, settlement); err != nil {
		return err
	}

	// Collected last so that it's easier for a seller to fill asks without needing those funds first.
	// Collected separately so it's not combined with the buyer settlement fees in the events.
	if msg.BidOrderCreationFee != nil {
		if err := k.CollectFee(ctx, marketID, buyer, sdk.Coins{*msg.BidOrderCreationFee}); err != nil {
			return fmt.Errorf("error collecting create-ask fee %q: %w", msg.BidOrderCreationFee, err)
		}
	}

	return nil
}

// SettleOrders attempts to settle all the provided orders.
func (k Keeper) SettleOrders(ctx sdk.Context, marketID uint32, askOrderIDs, bidOrderIds []uint64, expectPartial bool) error {
	store := k.getStore(ctx)
	if err := validateMarketExists(store, marketID); err != nil {
		return err
	}

	askOrders, aoerr := k.getAskOrders(store, marketID, askOrderIDs, "")
	bidOrders, boerr := k.getBidOrders(store, marketID, bidOrderIds, "")
	if aoerr != nil || boerr != nil {
		return errors.Join(aoerr, boerr)
	}

	ratioGetter := func(denom string) (*exchange.FeeRatio, error) {
		return getSellerSettlementRatio(store, marketID, denom)
	}

	settlement, err := exchange.BuildSettlement(askOrders, bidOrders, ratioGetter)
	if err != nil {
		return err
	}

	if !expectPartial && settlement.PartialOrderFilled != nil {
		return fmt.Errorf("settlement resulted in unexpected partial order %d", settlement.PartialOrderFilled.GetOrderID())
	}
	if expectPartial && settlement.PartialOrderFilled == nil {
		return errors.New("settlement unexpectedly resulted in all orders fully filled")
	}

	return k.closeSettlement(ctx, store, marketID, settlement)
}

// closeSettlement does all the processing needed to complete a settlement.
// It releases all the holds, does all the transfers, collects the fees, deletes/updates the orders, and emits events.
func (k Keeper) closeSettlement(ctx sdk.Context, store sdk.KVStore, marketID uint32, settlement *exchange.Settlement) error {
	// Release the holds!!!!
	var errs []error
	for _, order := range settlement.FullyFilledOrders {
		if err := k.releaseHoldOnOrder(ctx, order); err != nil {
			errs = append(errs, err)
		}
	}
	if settlement.PartialOrderFilled != nil {
		if err := k.releaseHoldOnOrder(ctx, settlement.PartialOrderFilled); err != nil {
			errs = append(errs, err)
		}
	}
	if len(errs) > 0 {
		return errors.Join(errs...)
	}

	// Transfer all the things!!!!
	for _, transfer := range settlement.Transfers {
		if err := k.DoTransfer(ctx, transfer.Inputs, transfer.Outputs); err != nil {
			errs = append(errs, err)
		}
	}

	// Collect all the fees (not as exciting).
	if err := k.CollectFees(ctx, marketID, settlement.FeeInputs); err != nil {
		errs = append(errs, err)
	}

	if len(errs) > 0 {
		return errors.Join(errs...)
	}

	// Update the partial order if there was one.
	if settlement.PartialOrderLeft != nil {
		if err := k.setOrderInStore(store, *settlement.PartialOrderLeft); err != nil {
			return fmt.Errorf("could not update partial %s order %d: %w",
				settlement.PartialOrderLeft.GetOrderType(), settlement.PartialOrderLeft.OrderId, err)
		}
	}
	// Delete all the fully filled orders.
	for _, order := range settlement.FullyFilledOrders {
		deleteAndDeIndexOrder(store, *order.GetOriginalOrder())
	}

	// And emit all the needed events.
	events := make([]proto.Message, 0, len(settlement.FullyFilledOrders)+1)
	for _, order := range settlement.FullyFilledOrders {
		events = append(events, exchange.NewEventOrderFilled(order))
	}
	if settlement.PartialOrderFilled != nil {
		events = append(events, exchange.NewEventOrderPartiallyFilled(settlement.PartialOrderFilled))
	}
	k.emitEvents(ctx, events)

	return nil
}

package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/libs/log"

	epochkeeper "github.com/provenance-io/provenance/x/epoch/keeper"
	"github.com/provenance-io/provenance/x/reward/types"
)

const StoreKey = types.ModuleName

type Keeper struct {
	storeKey    sdk.StoreKey
	cdc         codec.BinaryCodec
	epochKeeper epochkeeper.Keeper
}

func NewKeeper(
	cdc codec.BinaryCodec,
	key sdk.StoreKey,
	epochKeeper epochkeeper.Keeper,
) Keeper {
	return Keeper{
		storeKey:    key,
		cdc:         cdc,
		epochKeeper: epochKeeper,
	}
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", "x/"+types.ModuleName)
}

// SetRewardProgram sets the reward program in the keeper
func (k Keeper) SetRewardProgram(ctx sdk.Context, rewardProgram types.RewardProgram) error {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&rewardProgram)
	store.Set(types.GetRewardProgramKey(int64(rewardProgram.Id)), bz)
	return nil
}

// GetRewardProgram returns a RewardProgram by id if it exists nil if it does not
func (k Keeper) GetRewardProgram(ctx sdk.Context, id int64) (*types.RewardProgram, error) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetRewardProgramKey(id)
	bz := store.Get(key)
	if len(bz) == 0 {
		return nil, nil
	}

	var rewardProgram types.RewardProgram
	if err := k.cdc.Unmarshal(bz, &rewardProgram); err != nil {
		return nil, err
	}

	return &rewardProgram, nil
}

// IterateRewardPrograms iterates all reward programs with the given handler function.
func (k Keeper) IterateRewardPrograms(ctx sdk.Context, handle func(rewardProgram types.RewardProgram) (stop bool)) error {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.RewardProgramKeyPrefix)

	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		record := types.RewardProgram{}
		if err := k.cdc.Unmarshal(iterator.Value(), &record); err != nil {
			return err
		}
		if handle(record) {
			break
		}
	}
	return nil
}

// SetRewardClaim sets the reward program in the keeper
func (k Keeper) SetRewardClaim(ctx sdk.Context, rewardProgram types.RewardClaim) error {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&rewardProgram)
	store.Set(types.AddrRewardClaimsKey([]byte(rewardProgram.Address), rewardProgram.SharesPerEpochPerReward), bz)
	return nil
}

// GetRewardClaim returns a RewardClaim by id if it exists nil if it does not
func (k Keeper) GetRewardClaim(ctx sdk.Context, addr []byte, epochId int64, rewardsId int64) (*types.RewardClaim, error) {
	store := ctx.KVStore(k.storeKey)
	key := types.AddrRewardClaimsKey(addr, epochId, rewardsId)
	bz := store.Get(key)
	if len(bz) == 0 {
		return nil, nil
	}

	var rewardClaim types.RewardClaim
	if err := k.cdc.Unmarshal(bz, &rewardClaim); err != nil {
		return nil, err
	}

	return &rewardClaim, nil
}

// IterateRewardClaims  iterates all reward claims with the given handler function.
func (k Keeper) IterateRewardClaims(ctx sdk.Context, handle func(rewardClaim types.RewardClaim) (stop bool)) error {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.RewardClaimKeyPrefix)

	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		record := types.RewardClaim{}
		if err := k.cdc.Unmarshal(iterator.Value(), &record); err != nil {
			return err
		}
		if handle(record) {
			break
		}
	}
	return nil
}

// SetEpochRewardDistribution sets the EpochRewardDistribution in the keeper
func (k Keeper) SetEpochRewardDistribution(ctx sdk.Context, epochRewardDistribution types.EpochRewardDistribution) error {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&epochRewardDistribution)
	store.Set(types.GetEpochRewardDistributionKey(epochRewardDistribution.Id), bz)
	return nil
}

// GetEpochRewardDistribution returns a EpochRewardDistribution by id if it exists nil if it does not
func (k Keeper) GetEpochRewardDistribution(ctx sdk.Context, id string) (*types.EpochRewardDistribution, error) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetEpochRewardDistributionKey(id)
	bz := store.Get(key)
	if len(bz) == 0 {
		return nil, nil
	}

	var epochRewardDistribution types.EpochRewardDistribution
	if err := k.cdc.Unmarshal(bz, &epochRewardDistribution); err != nil {
		return nil, err
	}

	return &epochRewardDistribution, nil
}

// IterateEpochRewardDistributions  iterates all epoch reward distributions with the given handler function.
func (k Keeper) IterateEpochRewardDistributions(ctx sdk.Context, handle func(epochRewardDistribution types.EpochRewardDistribution) (stop bool)) error {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.EpochRewardDistributionKeyPrefix)

	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		record := types.EpochRewardDistribution{}
		if err := k.cdc.Unmarshal(iterator.Value(), &record); err != nil {
			return err
		}
		if handle(record) {
			break
		}
	}
	return nil
}

// SetEligibilityCriteria sets the EpochRewardDistribution in the keeper
func (k Keeper) SetEligibilityCriteria(ctx sdk.Context, eligibilityCriteria types.EligibilityCriteria) error {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&eligibilityCriteria)
	store.Set(types.GetEligibilityCriteriaKey(eligibilityCriteria.Name), bz)
	return nil
}

// GetEligibilityCriteria returns a EligibilityCriteria by name if it exists nil if it does not
func (k Keeper) GetEligibilityCriteria(ctx sdk.Context, name string) (*types.EligibilityCriteria, error) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetEligibilityCriteriaKey(name)
	bz := store.Get(key)
	if len(bz) == 0 {
		return nil, nil
	}

	var eligibilityCriteria types.EligibilityCriteria
	if err := k.cdc.Unmarshal(bz, &eligibilityCriteria); err != nil {
		return nil, err
	}

	return &eligibilityCriteria, nil
}

// IterateEligibilityCriterias  iterates all reward eligibility criterions with the given handler function.
func (k Keeper) IterateEligibilityCriterias(ctx sdk.Context, handle func(eligibilityCriteria types.EligibilityCriteria) (stop bool)) error {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.EligibilityCriteriaKeyPrefix)

	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		record := types.EligibilityCriteria{}
		if err := k.cdc.Unmarshal(iterator.Value(), &record); err != nil {
			return err
		}
		if handle(record) {
			break
		}
	}
	return nil
}

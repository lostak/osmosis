package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/osmosis-labs/osmosis/v7/x/gamm/whitelist-models/basic"
	"github.com/osmosis-labs/osmosis/v7/x/gamm/types"
)

func (k Keeper) MarshalWhitelist(whitelist types.WhitelistI) ([]byte, error) {
	return k.cdc.MarshalInterface(whitelist)
}

func (k Keeper) UnmarshalWhitelist(bz []byte) (types.WhitelistI, error) {
	var acc types.PoolI
	return acc, k.cdc.UnmarshalInterface(bz, &acc)
}

func (k Keeper) GetWhitelist(ctx sdk.Context, poolId uint64) (types.WhitelistI, error) {
	store := ctx.KVStore(k.storeKey)
	whitelistKey := types.GetKeyPrefixWhitelist(poolId)
	if !store.Has(whitelist) {
		return nil, fmt.Errorf("whitelist with ID %d does not exist", poolId)
	}

	bz := store.Get(whitelistKey)

	whitelist, err := k.UnmarshalWhitelist(bz)
	if err != nil {
		return nil, err
	}

	return whitelist, nil
}

func (k Keeper) iterator(ctx sdk.Context, prefix []byte) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, prefix)
}

func (k Keeper) GetWhitelists(ctx sdk.Context) (res []types.PoolI, err error) {
	iter := k.iterator(ctx, types.KeyPrefixWhitelist)
	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		bz := iter.Value()

		whitelist, err := k.UnmarshalWhitelist(bz)
		if err != nil {
			return nil, err
		}

		res = append(res, whitelist)
	}

	return res, nil
}

func (k Keeper) SetWhitelist(ctx sdk.Context, whitelist types.WhitelistI) error {
	bz, err := k.MarshalWhitelist(whitelist)
	if err != nil {
		return err
	}

	store := ctx.KVStore(k.storeKey)
	whitelistKey := types.GetKeyPrefixWhitelist(whitelist.GetId())
	store.Set(whitelistKey, bz)

	return nil
}

func (k Keeper) DeleteWhitelist(ctx sdk.Context, poolId uint64) error {
	store := ctx.KVStore(k.storeKey)
	whitelistKey := types.GetKeyPrefixWhitelist(poolId)
	if !store.Has(poolKey) {
		return fmt.Errorf("whitelist with pool ID %d does not exist", poolId)
	}

	store.Delete(whitelistKey)
	return nil
}


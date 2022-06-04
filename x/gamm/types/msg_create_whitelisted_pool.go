package types

import sdk "github.com/cosmos/cosmos-sdk/types"

// CreatePoolMsg defines an interface that every CreatePool transaction should implement.
// The gamm logic will use this to create a pool and associated governor.
type CreateWhitelistedPoolMsg interface {
	CreatePoolMsg
	
	// CreateWhitelist creates a new whitelist for the pool id with the sender as the governor
	CreateWhitelist(ctx sdk.Context, poolID uint64) (WhitelistI, error)

	// CreateWhitelistedPool creates a new pool and a new associated whitelist where 
	// the sender is the governor 
	CreateWhitelistedPool(ctx sdk.Context, poolID uint64) (PoolI, WhitelistI, error)
}

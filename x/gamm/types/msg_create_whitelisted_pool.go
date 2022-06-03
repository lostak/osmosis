package types

import sdk "github.com/cosmos/cosmos-sdk/types"

// CreatePoolMsg defines an interface that every CreatePool transaction should implement.
// The gamm logic will use this to create a pool and associated governor.
type CreatePoolWithGovMsg interface {
	CreatePoolMsg
	
	// CreateGovernor makes the msg signer the governor for the pool 
	// with an empty whitelist and exclusivity initially true. 
	CreateGovernor(ctx sdk.Context, poolID uint64, sender string) (governor, error)
}

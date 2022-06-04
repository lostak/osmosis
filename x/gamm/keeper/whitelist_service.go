package keeper 

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/osmosis-labs/osmosis/v7/x/gamm/types"
)

// TODO 
func (k Keeper) CreateWhitelist(ctx sdk.Context, poolID uint64) error {

}

// TODO
func (k Keeper) MsgCreateWhitelistedBalancerPool(ctx sdk.Context, msg types.MsgCreateWhitelistedBalancerPool) error {
	
}
package balancer

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/osmosis-labs/osmosis/v7/x/gamm/types"
)

const (
	TypeMsgCreateBalancerPool = "create_balancer_pool"
	TypeMsgCreateWhitelistedBalancerPool = "create_whitelisted_balancer_pool"
)

var (
	// TODO figure out what this actually means
	// Is this the glue for PoolI?
	_ sdk.Msg             = &MsgCreateBalancerPool{}
	_ types.CreatePoolMsg = &MsgCreateBalancerPool{}
	
	_ sdk.Msg 			  			 = &TypeMsgCreateWhitelistedBalancerPool{}
	_ types.CreateWhitelistedPoolMsg = &TypeMsgCreateWhitelistedBalancerPool{}
)

func NewMsgCreateWhitelistedBalancerPool(
	sender sdk.AccAddress,
	poolParams PoolParams,
	poolAssets []PoolAsset,
	futurePoolGovernor string,
	exclusive bool,
) MsgCreateWhitelistedBalancerPool {
	return MsgCreateWhitelistedBalancerPool{
		Sender:				sender.String(),
		PoolParams:			&poolParams,
		PoolAssets:			poolAssets,
		FuturePoolGovernor:	futurePoolGovernor,
	}
}

func (msg MsgCreateWhitelistedBalancerPool) Route() string { return types.RouterKey }
func (msg MsgCreateWhitelistedBalancerPool) Type() string { return TypeMsgCreateWhitelistedBalancerPool }
func (msg MsgCreateWhitelistedBalancerPool) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid sender address (%s)", err)
	}

	err = ValidateUserSpecifiedPoolAssets(msg.PoolAssets)
	if err != nil {
		return err
	}

	err = msg.PoolParams.Validate(msg.PoolAssets)
	if err != nil {
		return err
	}

	// validation for future owner
	if err = types.ValidateFutureGovernor(msg.FuturePoolGovernor); err != nil {
		return err
	}
}

func (msg MsgCreateWhitelistedBalancerPool) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgCreateWhitelistedBalancerPool) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

func NewMsgCreateBalancerPool(
	sender sdk.AccAddress,
	poolParams PoolParams,
	poolAssets []PoolAsset,
	futurePoolGovernor string,
) MsgCreateBalancerPool {
	return MsgCreateBalancerPool{
		Sender:             sender.String(),
		PoolParams:         &poolParams,
		PoolAssets:         poolAssets,
		FuturePoolGovernor: futurePoolGovernor,
	}
}

func (msg MsgCreateBalancerPool) Route() string { return types.RouterKey }
func (msg MsgCreateBalancerPool) Type() string  { return TypeMsgCreateBalancerPool }
func (msg MsgCreateBalancerPool) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid sender address (%s)", err)
	}

	err = ValidateUserSpecifiedPoolAssets(msg.PoolAssets)
	if err != nil {
		return err
	}

	err = msg.PoolParams.Validate(msg.PoolAssets)
	if err != nil {
		return err
	}

	// validation for future owner
	if err = types.ValidateFutureGovernor(msg.FuturePoolGovernor); err != nil {
		return err
	}

	return nil
}

func (msg MsgCreateBalancerPool) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgCreateBalancerPool) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

/// Implement the CreatePoolMsg interface

func (msg MsgCreateBalancerPool) PoolCreator() sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return sender
}

func (msg MsgCreateBalancerPool) Validate(ctx sdk.Context) error {
	return msg.ValidateBasic()
}

func (msg MsgCreateBalancerPool) InitialLiquidity() sdk.Coins {
	var coins sdk.Coins
	for _, asset := range msg.PoolAssets {
		coins = append(coins, asset.Token)
	}
	if coins == nil {
		panic("Shouldn't happen")
	}
	coins = coins.Sort()
	return coins
}

func (msg MsgCreateBalancerPool) CreatePool(ctx sdk.Context, poolID uint64) (types.PoolI, error) {
	poolI, err := NewBalancerPool(poolID, *msg.PoolParams, msg.PoolAssets, msg.FuturePoolGovernor, ctx.BlockTime())
	return &poolI, err
}

/// Implement the CreateWhitelistedPoolMsg interface

func (msg MsgCreateWhitelistedBalancerPool) PoolCreator() sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return sender
}

func (msg MsgCreateWhitelistedBalancerPool) Validate(ctx sdk.Context) error {
	return msg.ValidateBasic()
}

func (msg MsgCreateWhitelistedBalancerPool) InitialLiquidity() sdk.Coins {
	var coins sdk.Coins
	for _, asset := range msg.PoolAssets {
		coins = append(coins, asset.Token)
	}
	if coins == nil {
		panic("Shouldn't happen")
	}
	coins = coins.Sort()
	return coins
}

func (msg MsgCreateWhitelistedBalancerPool) CreatePool(ctx sdk.Context, poolID uint64) (types.PoolI, error) {
	poolI, err := NewBalancerPool(poolID, *msg.PoolParams, msg.PoolAssets, msg.FuturePoolGovernor, ctx.BlockTime())
	return &poolI, err
}

func (msg MsgCreateWhitelistedBalancerPool) CreateWhitelist(ctx sdk.Context, poolID uint64, exclusive bool) (types.WhitelistI, error) {
	// create and return basic whitelist
	whitelistI, err := NewBasicWhitelist(poolID, msg.Sender, exclusive) 
	return &whitelistI, err 
}

// create and return a balancer pool and a basic whitelist
fucn (msg MsgCreateWhitelistedBalancerPool) CreateWhitelistedPool(ctx sdk.Context, poolID uint64, exclusive bool) (pool types.PoolI, whitelist types.WhitelistI, err error) {
	// create pool
	pool, err = msg.CreatePool(ctx, poolID)
	if err := nil {
		return &pool, &whitelist, err
	}
	// create whitelist
	whitelist, err = msg.CreateWhitelist(ctx sdk.Context, poolID uint64, exclusive)
	if err != nil {
		return &pool, &whitelist, err
	}
	return &pool, &whitelist, nil
}

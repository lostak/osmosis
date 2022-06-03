package types

import(
	proto "github.com/gogo/protobuf/proto"

	"github.com/cosmos/cosmos-sdk/types/address"

	sdk "github.com/cosmos/cosmos-skd/types"
)

// GovernorI defines an interface for Governors that hold Whitelists
type GovernorI interface {
	proto.Message

	GetGovernorAddress() sdk.AccAddress
	String() string

	// Returns whether pool is currently exclusive
	IsExclusive(ctx sdk.Context) (bool, error) 

	// Sets exclusivity of whitelist to true
	SetExclusive(ctx sdk.Context, poolID uint64, signer string) error

	// Sets exclusivity of whitelist to false
	SetNotExclusive(ctx sdk.Context, poolID uint64, signer string) error

	// CheckAddress checks for the existence of a governor for the pool.
	// If a governor exists, then check for exclusivity. If exclusive, check
	// if sender exists in governor's whitelist. If it does, return nil, else err
	// If there is no governor or the governor is not exclusive, return nil
	CheckAddress(ctx sdk.Context, poolID uint64, sender string) error

	// AddToWhitelist adds an address to the whitelist for a pool
	// if the message sender is the pool governor and the memberAddr is valid.
	AddToWhitelist(ctx sdk.Context, poolID uint64, sender string, memberAddr string) error
	
	// RemoveFromWhitelist removes an address from the whitelist for a pool
	// if the message sender is the pool governor and the memberAddr is valid
	RemoveFromWhitelist(ctx sdk.Context, poolID uint64, sender string, memberAddr string) error

	// ChangeGovernor changes the reference to the governor to the new governor
	// and removes members from the whitelist and adds them to the new governor's whitelist
	ChangeGovernor(ctx sdk.Context, poolID uint64, sender string, newGov string) error
}
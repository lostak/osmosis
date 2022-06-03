package types

import(
	proto "github.com/gogo/protobuf/proto"

	"github.com/cosmos/cosmos-sdk/types/address"

	sdk "github.com/cosmos/cosmos-skd/types"
)

// WhitelistI defines an interface for Whitelists
type WhitelistI interface {
	proto.Message

	GetGovernor() sdk.AccAddress
	String() string

	// Returns whether pool is currently exclusive
	IsExclusive(ctx sdk.Context) (bool, error) 

	// Sets exclusivity of whitelist 
	SetExclusive(ctx sdk.Context, poolID uint64, senderAddr string, exclusivity bool) (bool, error)

	// CheckAddress checks for the existence of the address in the whitelist.
	CheckAddress(ctx sdk.Context, poolID uint64, memberAddr string) (bool, error)

	// AddToWhitelist adds a member's address to the whitelist for a pool
	// if the message sender is the pool governor and the memberAddr is valid.
	AddMemberToWhitelist(ctx sdk.Context, poolID uint64, senderAddr string, memberAddr string) error
	
	// RemoveFromWhitelist removes a member's address from the whitelist for a pool
	// if the message sender is the pool governor and the memberAddr is valid
	RemoveMemberFromWhitelist(ctx sdk.Context, poolID uint64, senderAddr string, memberAddr string) error

	// ChangeGovernor changes the reference to the governor to the new governor
	// and removes members from the whitelist and adds them to the new governor's whitelist
	ChangeGovernor(ctx sdk.Context, poolID uint64, senderAddr string, newGovernorAddr string) error
}
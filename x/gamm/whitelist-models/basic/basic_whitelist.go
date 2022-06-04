package basic

import (
	"errors"
	
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/osmosis-labs/osmosis/v7/x/gamm/types"
)

var (
	_ types.WhitelistI = &Whitelist{}
)

// NewBasicWhitelist resturns a whitelist under with exclusivity, membership
// and governance controled by the creator
func NewBasicWhitelist(poolId uint64, governorAddr string, exclusive bool) (Whitelist, error) {
	// assume poolId relates to an existing pool and that govenerAddr is a valid addr
	whitelist := &Whitelist{
		PoolId:		poolId,
		Exclusive:	exclusive,
		Governor:	govenerAddr,
		Member:		nil,	
	}
	
	return *whitelist, nil
}


func (wl Whitelist) GetGovernor() sdk.AccAddress {
	return wl.Governor
}

// Returns whether pool is currently exclusive
func (wl Whitelist) IsExclusive() (bool, error) {
	return pa.Exclusive
}

// Sets exclusivity of whitelist 
func (wl *Whitelist) SetExclusive(senderAddr string, exclusivity bool) (bool, error) {
	if !senderAddr.Equals(wl.Governor) {
		return false, errors.New("only governor can set exclusivity of whitelist")
	}

	wl.Exclusive = exclusivity
}

func (wl Whitelist) hasAddress(memberAddr string) (bool){
	//TODO figure out how to efficiently store addresses
	return false
}

func (wl Whitelist) CheckAddress(memberAddr string) (bool, error) {
	if !wl.hasAddress(memberAddr){
		return false, errors.New("member not on whitelist")
	}
	return true, nil
}

func (wl *Whitelist) addMember(memberAddr) error {
	return nil
}

func (wl *Whitelist) removeMember(memberAddr) error {
	
}

func (wl *Whitelist) AddMemberToWhitelist(senderAddr string, memberAddr string) error {

}
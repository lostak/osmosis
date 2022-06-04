package basic

import (
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type basicWhitelistPretty struct {
	PoolId			uint64 `json:"whitelist_pool_id" yaml:"whitelist_pool_id"`
	Exclusive		bool `json:"whitelist_exclusivity" yaml: "whitelist_exclusivity"`
	Governor		string `json:"governor_address" yaml:"governor_address"`
	Members			[]Members `json:"whitelist_members" yaml:"whitelist_members"`
}

func (wl Whitelist) String() string {
	out, err := wl.MarshalJSON()
	if err != nil {
		panic(err)
	}
	return string(out)
}

func (wl Whitelist) MarshalJSON() ([]byte, error) {
	return json.Marshal(basicWhitelistPretty{
		PoolId:		wl.PoolId,
		Exclusive:	wl.Exclusive,
		Governor:	wl.Governor,
		Members:	wl.Members,
	})
}

func (wl *Whitelist) UnmarshalJSON(bz []byte) error {
	var alias basicWhitelistPretty
	if err := json.Unmarshal(bz, &alias); err != nil {
		return err
	}

	wl.PoolId = alias.PoolId
	wl.Exclusive = alias.Exclusive
	wl.Governor = alias.Governor
	wl.Members = alias.Members

	return nil
}
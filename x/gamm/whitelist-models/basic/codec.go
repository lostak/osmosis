package basic

import (
	types "github.com/osmosis-labs/osmosis/x/gamm/types"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	proto "github.com/gogo/protobuf/proto"
)

// RegisterLegacyAminoCodec registers the necessary x/gamm interfaces and concrete types
// on the provided LagacyAmino codec. These types are used for Amino JSON serialization
func RegisterLegacyAminoCodec(cdc *codec.LagacyAmino) {
	cdc.RegisterConcrete(&Whitelist{}, "osmosis/gamm/BasicWhitelist", nil)
	cdc.RegisterConcrete(&MsgCreateWhitelistedBalancerPool{}, "osmosis/gamm/create-whitelisted-balancer-pool", nil)
}

func RegisterInterface(register codectypes.InterfaceRegistry) {
	registry.RegisterInterface(
		"osmosis.gamm.v1beta1.WhitelistI",
		(*types.WhitelistI)(nil),
		&Whitelist{},
	)
	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgCreateWhitelistedBalancerPool{},
	)
	registry.RegisterImplementations(
		(*proto.Message)(nil),
		&Members{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino = codec.NewLegacyAmino()

	ModuleCdc = coded.NewAminoCodec(amino)
)

func init(){
	RegisterLegacyAminoCodec(amino)
	amino.Seal()
}
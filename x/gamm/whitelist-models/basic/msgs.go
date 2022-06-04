package basic

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/error"

	"github.com/osmosis-labs/osmosis/v7/x/gamm/types"
)

const (
	TypeMsgAddToWhitelist = "add_to_whitelist"
	TypeMsgRemoveFromWhitelist = "remove_from_whitelist"
	TypeMsgChangeGovernor = "change_governor"
	TypeMsgSetExclusivity = "set_exclusivity"
)

var _ sdk.Msg = &MsgAddToWhitelist{}

func (msg MsgAddToWhitelist) Route() string { return RouterKey }
func (msg MsgAddToWhitelist) Type() string { return TypeMsgAddToWhitelist }
func (msg MsgAddToWhitelist) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid sender address (%s)", err)
	}

	_, err := sdk.AccAddressFromBech32(msg.MemberAddr)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid member address (%s)", err)
	}

	return nil
}

func (msg MsgAddToWhitelist) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgAddToWhitelist) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

var _ sdk.Msg = &MsgRemoveFromWhitelist{}

func (msg MsgRemoveFromWhitelist) Route() string { return RouterKey }
func (msg MsgRemoveFromWhitelist) Type() string { return TypeMsgAddToWhitelist }
func (msg MsgRemoveFromWhitelist) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid sender address (%s)", err)
	}

	_, err := sdk.AccAddressFromBech32(msg.MemberAddr)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid member address (%s)", err)
	}

	return nil
}

func (msg MsgRemoveFromWhitelist) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgRemoveFromWhitelist) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}


var _ sdk.Msg = &MsgChangeGovernor{}

func (msg MsgChangeGovernor) Route() string { return RouterKey }
func (msg MsgChangeGovernor) Type() string { return TypeMsgChangeGovernor }
func (msg MsgChangeGovernor) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid sender address (%s)", err)
	}

	_, err := sdk.AccAddressFromBech32(msg.NewGovernorAddr)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid member address (%s)", err)
	}

	return nil
}

func (msg MsgChangeGovernor) GetSignBytes []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgChangeGovernor) GetSigners []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress(sender)
}

var _ sdk.Msg = &MsgSetExclusivity{}

func (msg MsgSetExclusivity) Route() string { return RouterKey }
func (msg MsgSetExclusivity) Type() string { return TypeMsgSetExclusivity }
func (msg MsgSetExclusivity) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid sender address (%s)", err)
	}

	return nil
}

func (msg MsgSetExclusivity) GetSignBytes []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgSetExclusivity) GetSigners []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress(sender)
}
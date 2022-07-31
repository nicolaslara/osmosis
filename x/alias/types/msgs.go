package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgExec = "exec"
)

var _ sdk.Msg = &MsgExec{}

// NewMsgExec creates a msg to create a new denom
func NewMsgExec(sender, msgType, msg, as string) *MsgExec {
	return &MsgExec{
		Sender:  sender,
		MsgType: msgType,
		Msg:     msg,
		As:      as,
	}
}

func (m MsgExec) Route() string { return RouterKey }
func (m MsgExec) Type() string  { return TypeMsgExec }
func (m MsgExec) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(m.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid sender address (%s)", err)
	}

	_, err = sdk.AccAddressFromBech32(m.As)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid as address (%s)", err)
	}

	return nil
}

func (m MsgExec) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

func (m MsgExec) GetSigners() []sdk.AccAddress {
	sender, _ := sdk.AccAddressFromBech32(m.Sender)
	return []sdk.AccAddress{sender}
}

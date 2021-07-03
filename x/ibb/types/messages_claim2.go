package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateClaim2{}

func NewMsgCreateClaim2(creator string, blockHeight int32, asset string, amount int32, denom string) *MsgCreateClaim2 {
	return &MsgCreateClaim2{
		Creator:     creator,
		BlockHeight: blockHeight,
		Asset:       asset,
		Amount:      amount,
		Denom:       denom,
	}
}

func (msg *MsgCreateClaim2) Route() string {
	return RouterKey
}

func (msg *MsgCreateClaim2) Type() string {
	return "CreateClaim2"
}

func (msg *MsgCreateClaim2) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateClaim2) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateClaim2) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateClaim2{}

func NewMsgUpdateClaim2(creator string, id uint64, blockHeight int32, asset string, amount int32, denom string) *MsgUpdateClaim2 {
	return &MsgUpdateClaim2{
		Id:          id,
		Creator:     creator,
		BlockHeight: blockHeight,
		Asset:       asset,
		Amount:      amount,
		Denom:       denom,
	}
}

func (msg *MsgUpdateClaim2) Route() string {
	return RouterKey
}

func (msg *MsgUpdateClaim2) Type() string {
	return "UpdateClaim2"
}

func (msg *MsgUpdateClaim2) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateClaim2) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateClaim2) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgCreateClaim2{}

func NewMsgDeleteClaim2(creator string, id uint64) *MsgDeleteClaim2 {
	return &MsgDeleteClaim2{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeleteClaim2) Route() string {
	return RouterKey
}

func (msg *MsgDeleteClaim2) Type() string {
	return "DeleteClaim2"
}

func (msg *MsgDeleteClaim2) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteClaim2) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteClaim2) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

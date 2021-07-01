package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateNft{}

func NewMsgCreateNft(creator string, collection string, ownerAddress string, imageUrl string, name string, nftCreatorAddress string) *MsgCreateNft {
	return &MsgCreateNft{
		Creator:           creator,
		Collection:        collection,
		OwnerAddress:      ownerAddress,
		ImageUrl:          imageUrl,
		Name:              name,
		NftCreatorAddress: nftCreatorAddress,
	}
}

func (msg *MsgCreateNft) Route() string {
	return RouterKey
}

func (msg *MsgCreateNft) Type() string {
	return "CreateNft"
}

func (msg *MsgCreateNft) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateNft) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateNft) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateNft{}

func NewMsgUpdateNft(creator string, id uint64, collection string, ownerAddress string, imageUrl string, name string) *MsgUpdateNft {
	return &MsgUpdateNft{
		Id:           id,
		Creator:      creator,
		Collection:   collection,
		OwnerAddress: ownerAddress,
		ImageUrl:     imageUrl,
		Name:         name,
	}
}

func (msg *MsgUpdateNft) Route() string {
	return RouterKey
}

func (msg *MsgUpdateNft) Type() string {
	return "UpdateNft"
}

func (msg *MsgUpdateNft) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateNft) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateNft) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteNft{}

func NewMsgDeleteNft(creator string, id uint64) *MsgDeleteNft {
	return &MsgDeleteNft{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeleteNft) Route() string {
	return RouterKey
}

func (msg *MsgDeleteNft) Type() string {
	return "DeleteNft"
}

func (msg *MsgDeleteNft) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteNft) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteNft) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

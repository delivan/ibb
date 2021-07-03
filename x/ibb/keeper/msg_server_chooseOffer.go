package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sapiens-cosmos/ibb/x/ibb/types"
)

func (k msgServer) ChooseOffer(goCtx context.Context, msg *types.MsgChooseOffer) (*types.MsgChooseOfferResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message

	nftList := k.GetAllNft(ctx)
	var queryNft types.Nft
	for _, nft := range nftList {
		if nft.Id == uint64(msg.NftId) {
			queryNft = nft
		}
	}
	queryNft.OwnerAddress = "in escrow"
	k.SetNft(ctx, queryNft)

	userList := k.GetAllUser(ctx)
	var queryUser types.User
	for _, user := range userList {
		if user.Creator == msg.Creator {
			queryUser = user
			break
		}
	}

	for i, borrow := range queryUser.Borrow {
		if borrow.Denom == queryNft.SelectedOffer.Denom {
			queryUser.Borrow[i].Amount += queryNft.SelectedOffer.Amount
		}
	}

	k.SetUser(ctx, queryUser)

	return &types.MsgChooseOfferResponse{}, nil
}

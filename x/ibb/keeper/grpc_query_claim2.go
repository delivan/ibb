package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/sapiens-cosmos/ibb/x/ibb/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Claim2All(c context.Context, req *types.QueryAllClaim2Request) (*types.QueryAllClaim2Response, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var claim2s []*types.Claim2
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	claim2Store := prefix.NewStore(store, types.KeyPrefix(types.Claim2Key))

	pageRes, err := query.Paginate(claim2Store, req.Pagination, func(key []byte, value []byte) error {
		var claim2 types.Claim2
		if err := k.cdc.UnmarshalBinaryBare(value, &claim2); err != nil {
			return err
		}

		claim2s = append(claim2s, &claim2)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllClaim2Response{Claim2: claim2s, Pagination: pageRes}, nil
}

func (k Keeper) Claim2(c context.Context, req *types.QueryGetClaim2Request) (*types.QueryGetClaim2Response, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var claim2 types.Claim2
	ctx := sdk.UnwrapSDKContext(c)

	if !k.HasClaim2(ctx, req.Id) {
		return nil, sdkerrors.ErrKeyNotFound
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.Claim2Key))
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetClaim2IDBytes(req.Id)), &claim2)

	return &types.QueryGetClaim2Response{Claim2: &claim2}, nil
}

package keeper

import (
	"encoding/binary"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sapiens-cosmos/ibb/x/ibb/types"
	"strconv"
)

// GetClaim2Count get the total number of claim2
func (k Keeper) GetClaim2Count(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.Claim2CountKey))
	byteKey := types.KeyPrefix(types.Claim2CountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	count, err := strconv.ParseUint(string(bz), 10, 64)
	if err != nil {
		// Panic because the count should be always formattable to iint64
		panic("cannot decode count")
	}

	return count
}

// SetClaim2Count set the total number of claim2
func (k Keeper) SetClaim2Count(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.Claim2CountKey))
	byteKey := types.KeyPrefix(types.Claim2CountKey)
	bz := []byte(strconv.FormatUint(count, 10))
	store.Set(byteKey, bz)
}

// AppendClaim2 appends a claim2 in the store with a new id and update the count
func (k Keeper) AppendClaim2(
	ctx sdk.Context,
	creator string,
	blockHeight int32,
	asset string,
	amount int32,
	denom string,
) uint64 {
	// Create the claim2
	count := k.GetClaim2Count(ctx)
	var claim2 = types.Claim2{
		Creator:     creator,
		Id:          count,
		BlockHeight: blockHeight,
		Asset:       asset,
		Amount:      amount,
		Denom:       denom,
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.Claim2Key))
	value := k.cdc.MustMarshalBinaryBare(&claim2)
	store.Set(GetClaim2IDBytes(claim2.Id), value)

	// Update claim2 count
	k.SetClaim2Count(ctx, count+1)

	return count
}

// SetClaim2 set a specific claim2 in the store
func (k Keeper) SetClaim2(ctx sdk.Context, claim2 types.Claim2) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.Claim2Key))
	b := k.cdc.MustMarshalBinaryBare(&claim2)
	store.Set(GetClaim2IDBytes(claim2.Id), b)
}

// GetClaim2 returns a claim2 from its id
func (k Keeper) GetClaim2(ctx sdk.Context, id uint64) types.Claim2 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.Claim2Key))
	var claim2 types.Claim2
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetClaim2IDBytes(id)), &claim2)
	return claim2
}

// HasClaim2 checks if the claim2 exists in the store
func (k Keeper) HasClaim2(ctx sdk.Context, id uint64) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.Claim2Key))
	return store.Has(GetClaim2IDBytes(id))
}

// GetClaim2Owner returns the creator of the claim2
func (k Keeper) GetClaim2Owner(ctx sdk.Context, id uint64) string {
	return k.GetClaim2(ctx, id).Creator
}

// RemoveClaim2 removes a claim2 from the store
func (k Keeper) RemoveClaim2(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.Claim2Key))
	store.Delete(GetClaim2IDBytes(id))
}

// GetAllClaim2 returns all claim2
func (k Keeper) GetAllClaim2(ctx sdk.Context) (list []types.Claim2) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.Claim2Key))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Claim2
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetClaim2IDBytes returns the byte representation of the ID
func GetClaim2IDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetClaim2IDFromBytes returns ID in uint64 format from a byte array
func GetClaim2IDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}

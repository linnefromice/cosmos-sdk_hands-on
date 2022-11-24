package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/linnefromice/handson/x/handson/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) PoolBalances(goCtx context.Context, req *types.QueryPoolBalancesRequest) (*types.QueryPoolBalancesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	pools := k.GetAllPool(ctx)
	var balances []types.PoolBalance
	for _, value := range pools {
		addr, _ := sdk.AccAddressFromBech32(value.Address)
		coin := k.bankKeeper.GetBalance(ctx, addr, value.Denom)
		balances = append(balances, types.PoolBalance{
			PoolId:  value.Id,
			Balance: coin,
		})
	}

	return &types.QueryPoolBalancesResponse{
		Balances: balances,
	}, nil
}

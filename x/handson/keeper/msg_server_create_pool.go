package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/linnefromice/handson/x/handson/types"
)

func (k msgServer) CreatePool(goCtx context.Context, msg *types.MsgCreatePool) (*types.MsgCreatePoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	pool := types.Pool{
		Address:   msg.Creator, // temp
		Denom:     msg.Denom,
		IsActive:  true,
		Deposited: 0,
		Borrowed:  0,
	}
	k.AppendPool(ctx, pool)

	return &types.MsgCreatePoolResponse{
		PoolId: pool.Id,
	}, nil
}

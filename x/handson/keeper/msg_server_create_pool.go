package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
	"github.com/linnefromice/handson/x/handson/types"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

func (k msgServer) CreatePool(goCtx context.Context, msg *types.MsgCreatePool) (*types.MsgCreatePoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	base := fmt.Sprintf("%d-%s", k.GetPoolCount(ctx), msg.Denom)
	accAddr := sdk.AccAddress(address.Module(types.ModuleName, []byte(base)))
	accI := k.accountKeeper.NewAccount(
		ctx,
		authtypes.NewModuleAccount(
			authtypes.NewBaseAccountWithAddress(accAddr),
			accAddr.String(),
		),
	)
	k.accountKeeper.SetAccount(ctx, accI)

	pool := types.Pool{
		Address:   accI.GetAddress().String(),
		Denom:     msg.Denom,
		IsActive:  true,
		Deposited: 0,
		Borrowed:  0,
	}
	k.AppendPool(ctx, pool)

	// event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.PoolDepositedEventType,
			sdk.NewAttribute(types.PoolEventId, fmt.Sprint(pool.Id)),
			sdk.NewAttribute(types.PoolEventDenom, fmt.Sprint(msg.Denom))),
	)

	return &types.MsgCreatePoolResponse{
		PoolId: pool.Id,
	}, nil
}

package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/linnefromice/handson/x/handson/types"
)

func (k msgServer) Withdraw(goCtx context.Context, msg *types.MsgWithdraw) (*types.MsgWithdrawResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	sender, _ := sdk.AccAddressFromBech32(msg.Creator)

	pool, found := k.GetPool(ctx, msg.PoolId)
	if !found {
		return &types.MsgWithdrawResponse{}, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.PoolId))
	}
	if msg.Amount.Denom != pool.Denom {
		return &types.MsgWithdrawResponse{}, sdkerrors.Wrapf(types.ErrIncorrectDenom, "input: %s, supported: %s", msg.Amount.Denom, pool.Denom)
	}

	lpCoin := sdk.NewCoin(fmt.Sprintf("share-%s", msg.Amount.Denom), msg.Amount.Amount)
	err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, sender, types.ModuleName, sdk.NewCoins(lpCoin))
	if err != nil {
		return &types.MsgWithdrawResponse{}, err
	}
	err = k.bankKeeper.BurnCoins(ctx, types.ModuleName, sdk.NewCoins(lpCoin))
	if err != nil {
		return &types.MsgWithdrawResponse{}, err
	}

	poolAddr, _ := sdk.AccAddressFromBech32(pool.Address)
	err = k.bankKeeper.SendCoins(ctx, poolAddr, sender, sdk.NewCoins(msg.Amount))
	if err != nil {
		return &types.MsgWithdrawResponse{}, err
	}

	pool.Deposited -= msg.Amount.Amount.Uint64()
	k.SetPool(ctx, pool)

	return &types.MsgWithdrawResponse{}, nil
}

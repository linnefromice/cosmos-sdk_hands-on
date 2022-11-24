package keeper_test

import (
	"testing"

	testkeeper "github.com/linnefromice/handson/testutil/keeper"
	"github.com/linnefromice/handson/x/handson/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.HandsonKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}

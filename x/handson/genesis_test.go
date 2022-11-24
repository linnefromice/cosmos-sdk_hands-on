package handson_test

import (
	"testing"

	keepertest "github.com/linnefromice/handson/testutil/keeper"
	"github.com/linnefromice/handson/testutil/nullify"
	"github.com/linnefromice/handson/x/handson"
	"github.com/linnefromice/handson/x/handson/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		PoolList: []types.Pool{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		PoolCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.HandsonKeeper(t)
	handson.InitGenesis(ctx, *k, genesisState)
	got := handson.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.PoolList, got.PoolList)
	require.Equal(t, genesisState.PoolCount, got.PoolCount)
	// this line is used by starport scaffolding # genesis/test/assert
}

package minioracle_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "minioracle/testutil/keeper"
	"minioracle/testutil/nullify"
	"minioracle/x/minioracle"
	"minioracle/x/minioracle/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		CommentList: []types.Comment{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		CommentCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.MinioracleKeeper(t)
	minioracle.InitGenesis(ctx, *k, genesisState)
	got := minioracle.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.CommentList, got.CommentList)
	require.Equal(t, genesisState.CommentCount, got.CommentCount)
	// this line is used by starport scaffolding # genesis/test/assert
}

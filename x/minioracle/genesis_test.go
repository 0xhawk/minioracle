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

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.MinioracleKeeper(t)
	minioracle.InitGenesis(ctx, *k, genesisState)
	got := minioracle.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}

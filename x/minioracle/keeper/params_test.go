package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "minioracle/testutil/keeper"
	"minioracle/x/minioracle/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.MinioracleKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}

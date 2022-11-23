package keeper

import (
	"minioracle/x/minioracle/types"
)

var _ types.QueryServer = Keeper{}

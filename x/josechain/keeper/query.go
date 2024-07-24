package keeper

import (
	"josechain/x/josechain/types"
)

var _ types.QueryServer = Keeper{}

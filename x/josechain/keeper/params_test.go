package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "josechain/testutil/keeper"
	"josechain/x/josechain/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.JosechainKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}

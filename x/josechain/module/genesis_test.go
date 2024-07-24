package josechain_test

import (
	"testing"

	keepertest "josechain/testutil/keeper"
	"josechain/testutil/nullify"
	josechain "josechain/x/josechain/module"
	"josechain/x/josechain/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.JosechainKeeper(t)
	josechain.InitGenesis(ctx, k, genesisState)
	got := josechain.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}

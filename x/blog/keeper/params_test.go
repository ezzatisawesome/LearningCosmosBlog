package keeper_test

import (
	"testing"

	testkeeper "github.com/ezzatisawesome/blog/testutil/keeper"
	"github.com/ezzatisawesome/blog/x/blog/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.BlogKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}

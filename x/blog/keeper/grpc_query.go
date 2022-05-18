package keeper

import (
	"github.com/ezzatisawesome/blog/x/blog/types"
)

var _ types.QueryServer = Keeper{}

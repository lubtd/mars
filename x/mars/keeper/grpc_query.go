package keeper

import (
	"github.com/lubtd/mars/x/mars/types"
)

var _ types.QueryServer = Keeper{}

package market

import (
	"github.com/coinexchain/dex/modules/market/internal/types"

	"github.com/coinexchain/dex/modules/market/internal/keepers"
)

const (
	StoreKey   = types.StoreKey
	ModuleName = types.ModuleName
)

const (
	IntegrationNetSubString = types.IntegrationNetSubString
	OrderIDPartsNum         = types.OrderIDPartsNum
)

var (
	NewBaseKeeper = keepers.NewKeeper
	DefaultParams = keepers.DefaultParams
)

type (
	Keeper     = keepers.Keeper
	Order      = types.Order
	MarketInfo = types.MarketInfo
)
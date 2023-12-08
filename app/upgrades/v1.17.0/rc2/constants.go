package rc2

import (
	store "github.com/cosmos/cosmos-sdk/store/types"

	"github.com/provenance-io/provenance/app/upgrades"
)

const UpgradeName = "saffron-rc2"

var Upgrade = upgrades.Upgrade{
	UpgradeName:     UpgradeName,
	UpgradeStrategy: UpgradeStrategy,
	StoreUpgrades: store.StoreUpgrades{
		Added:   []string{},
		Deleted: []string{},
	},
}
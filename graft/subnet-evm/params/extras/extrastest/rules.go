// Copyright (C) 2019-2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package extrastest

import (
	"github.com/rink-labs/rinkgo/graft/subnet-evm/params/extras"
	"github.com/rink-labs/rinkgo/upgrade"
	"github.com/rink-labs/rinkgo/upgrade/upgradetest"
)

func ForkToAvalancheRules(fork upgradetest.Fork) extras.AvalancheRules {
	networkUpgrades := extras.GetNetworkUpgrades(upgradetest.GetConfig(fork))
	return networkUpgrades.GetAvalancheRules(uint64(upgrade.InitiallyActiveTime.Unix()))
}

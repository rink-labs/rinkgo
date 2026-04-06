// Copyright (C) 2019-2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package extrastest

import (
	"fmt"

	"github.com/ava-labs/libevm/common"

	"github.com/rink-labs/rinkgo/graft/coreth/params"
	"github.com/rink-labs/rinkgo/graft/coreth/params/extras"
	"github.com/rink-labs/rinkgo/graft/coreth/params/paramstest"
	"github.com/rink-labs/rinkgo/upgrade"
	"github.com/rink-labs/rinkgo/upgrade/upgradetest"
)

func ForkToRules(fork upgradetest.Fork) *extras.Rules {
	chainConfig, ok := paramstest.ForkToChainConfig[fork]
	if !ok {
		panic(fmt.Sprintf("unknown fork: %s", fork))
	}
	return params.GetRulesExtra(chainConfig.Rules(common.Big0, params.IsMergeTODO, 0))
}

func ForkToAvalancheRules(fork upgradetest.Fork) extras.AvalancheRules {
	networkUpgrades := extras.GetNetworkUpgrades(upgradetest.GetConfig(fork))
	return networkUpgrades.GetAvalancheRules(uint64(upgrade.InitiallyActiveTime.Unix()))
}

// Copyright (C) 2019-2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package executor

import (
	"github.com/rink-labs/rinkgo/snow"
	"github.com/rink-labs/rinkgo/snow/uptime"
	"github.com/rink-labs/rinkgo/utils"
	"github.com/rink-labs/rinkgo/utils/timer/mockable"
	"github.com/rink-labs/rinkgo/vms/platformvm/config"
	"github.com/rink-labs/rinkgo/vms/platformvm/fx"
	"github.com/rink-labs/rinkgo/vms/platformvm/reward"
	"github.com/rink-labs/rinkgo/vms/platformvm/utxo"
)

type Backend struct {
	Config       *config.Internal
	Ctx          *snow.Context
	Clk          *mockable.Clock
	Fx           fx.Fx
	FlowChecker  utxo.Verifier
	Uptimes      uptime.Calculator
	Rewards      reward.Calculator
	Bootstrapped *utils.Atomic[bool]
}

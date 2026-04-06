// Copyright (C) 2019-2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package statetest

import (
	"testing"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/stretchr/testify/require"

	"github.com/rink-labs/rinkgo/database"
	"github.com/rink-labs/rinkgo/database/memdb"
	"github.com/rink-labs/rinkgo/ids"
	"github.com/rink-labs/rinkgo/snow"
	"github.com/rink-labs/rinkgo/snow/validators"
	"github.com/rink-labs/rinkgo/upgrade"
	"github.com/rink-labs/rinkgo/upgrade/upgradetest"
	"github.com/rink-labs/rinkgo/utils/constants"
	"github.com/rink-labs/rinkgo/utils/logging"
	"github.com/rink-labs/rinkgo/utils/units"
	"github.com/rink-labs/rinkgo/vms/platformvm/config"
	"github.com/rink-labs/rinkgo/vms/platformvm/genesis/genesistest"
	"github.com/rink-labs/rinkgo/vms/platformvm/metrics"
	"github.com/rink-labs/rinkgo/vms/platformvm/reward"
	"github.com/rink-labs/rinkgo/vms/platformvm/state"
)

var DefaultNodeID = ids.GenerateTestNodeID()

type Config struct {
	DB         database.Database
	Genesis    []byte
	Registerer prometheus.Registerer
	Validators validators.Manager
	Upgrades   upgrade.Config
	Config     config.Config
	Context    *snow.Context
	Metrics    metrics.Metrics
	Rewards    reward.Calculator
}

func New(t testing.TB, c Config) state.State {
	if c.DB == nil {
		c.DB = memdb.New()
	}
	if c.Context == nil {
		c.Context = &snow.Context{
			NetworkID: constants.UnitTestID,
			NodeID:    DefaultNodeID,
			Log:       logging.NoLog{},
		}
	}
	if len(c.Genesis) == 0 {
		c.Genesis = genesistest.NewBytes(t, genesistest.Config{
			NetworkID: c.Context.NetworkID,
		})
	}
	if c.Registerer == nil {
		c.Registerer = prometheus.NewRegistry()
	}
	if c.Validators == nil {
		c.Validators = validators.NewManager()
	}
	if c.Upgrades == (upgrade.Config{}) {
		c.Upgrades = upgradetest.GetConfig(upgradetest.Latest)
	}
	if c.Config == (config.Config{}) {
		c.Config = config.Default
	}
	if c.Metrics == nil {
		c.Metrics = metrics.Noop
	}
	if c.Rewards == nil {
		c.Rewards = reward.NewCalculator(reward.Config{
			MaxConsumptionRate: .12 * reward.PercentDenominator,
			MinConsumptionRate: .1 * reward.PercentDenominator,
			MintingPeriod:      365 * 24 * time.Hour,
			SupplyCap:          720 * units.MegaAvax,
		})
	}

	s, err := state.New(
		c.DB,
		c.Genesis,
		c.Registerer,
		c.Validators,
		c.Upgrades,
		&c.Config,
		c.Context,
		c.Metrics,
		c.Rewards,
	)
	require.NoError(t, err)
	return s
}

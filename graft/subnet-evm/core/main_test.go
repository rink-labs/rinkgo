// Copyright (C) 2019-2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package core

import (
	"os"
	"testing"

	"github.com/rink-labs/libevm/log"
	"go.uber.org/goleak"

	"github.com/rink-labs/rinkgo/graft/subnet-evm/params"
	"github.com/rink-labs/rinkgo/graft/subnet-evm/plugin/evm/customtypes"
)

// TestMain uses goleak to verify tests in this package do not leak unexpected
// goroutines.
func TestMain(m *testing.M) {
	RegisterExtras()

	customtypes.Register()
	params.RegisterExtras()

	// May of these tests are likely to fail due to `log.Crit` in goroutines.
	log.SetDefault(log.NewLogger(log.NewTerminalHandlerWithLevel(os.Stderr, log.LevelCrit, true)))

	opts := []goleak.Option{
		// No good way to shut down these goroutines:
		goleak.IgnoreTopFunction("github.com/rink-labs/rinkgo/graft/subnet-evm/core/state/snapshot.(*diskLayer).generate"),
		goleak.IgnoreTopFunction("github.com/rink-labs/libevm/core.(*txSenderCacher).cache"),
		goleak.IgnoreTopFunction("github.com/rink-labs/libevm/metrics.(*meterArbiter).tick"),
		goleak.IgnoreTopFunction("github.com/rink-labs/rinkgo/vms/evm/metrics.(*meterArbiter).tick"),
		goleak.IgnoreTopFunction("github.com/syndtr/goleveldb/leveldb.(*DB).mpoolDrain"),
	}
	goleak.VerifyTestMain(m, opts...)
}

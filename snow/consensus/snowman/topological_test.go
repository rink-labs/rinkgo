// Copyright (C) 2019-2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package snowman

import (
	"testing"

	"github.com/rink-labs/rinkgo/snow/consensus/snowball"
)

func TestTopological(t *testing.T) {
	runConsensusTests(t, TopologicalFactory{factory: snowball.SnowflakeFactory})
}

// Copyright (C) 2019-2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package peer

import (
	"github.com/rink-labs/rinkgo/ids"
	"github.com/rink-labs/rinkgo/utils/bloom"
	"github.com/rink-labs/rinkgo/utils/ips"
	"github.com/rink-labs/rinkgo/utils/set"
)

var TestNetwork Network = testNetwork{}

type testNetwork struct{}

func (testNetwork) Connected(ids.NodeID) {}

func (testNetwork) AllowConnection(ids.NodeID) bool {
	return true
}

func (testNetwork) Track([]*ips.ClaimedIPPort) error {
	return nil
}

func (testNetwork) Disconnected(ids.NodeID) {}

func (testNetwork) KnownPeers() ([]byte, []byte) {
	return bloom.EmptyFilter.Marshal(), nil
}

func (testNetwork) Peers(
	ids.NodeID,
	set.Set[ids.ID],
	bool,
	*bloom.ReadFilter,
	[]byte,
) []*ips.ClaimedIPPort {
	return nil
}

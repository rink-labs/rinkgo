// Copyright (C) 2019-2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package atomictest

import (
	"github.com/rink-labs/rinkgo/graft/coreth/plugin/evm/atomic"
	"github.com/rink-labs/rinkgo/ids"

	avalancheatomic "github.com/rink-labs/rinkgo/chains/atomic"
)

func ConvertToAtomicOps(tx *atomic.Tx) (map[ids.ID]*avalancheatomic.Requests, error) {
	id, reqs, err := tx.AtomicOps()
	if err != nil {
		return nil, err
	}
	return map[ids.ID]*avalancheatomic.Requests{id: reqs}, nil
}

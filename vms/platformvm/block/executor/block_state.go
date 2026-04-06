// Copyright (C) 2019-2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package executor

import (
	"time"

	"github.com/rink-labs/rinkgo/chains/atomic"
	"github.com/rink-labs/rinkgo/ids"
	"github.com/rink-labs/rinkgo/utils/set"
	"github.com/rink-labs/rinkgo/vms/platformvm/block"
	"github.com/rink-labs/rinkgo/vms/platformvm/metrics"
	"github.com/rink-labs/rinkgo/vms/platformvm/state"
)

type proposalBlockState struct {
	onDecisionState state.Diff
	onCommitState   state.Diff
	onAbortState    state.Diff
}

// The state of a block.
// Note that not all fields will be set for a given block.
type blockState struct {
	proposalBlockState
	statelessBlock block.Block

	onAcceptState state.Diff
	onAcceptFunc  func()

	inputs          set.Set[ids.ID]
	timestamp       time.Time
	atomicRequests  map[ids.ID]*atomic.Requests
	verifiedHeights set.Set[uint64]
	metrics         metrics.Block
}

// Copyright (C) 2019-2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package snowman

import (
	"github.com/rink-labs/rinkgo/snow"
	"github.com/rink-labs/rinkgo/snow/consensus/snowball"
	"github.com/rink-labs/rinkgo/snow/consensus/snowman"
	"github.com/rink-labs/rinkgo/snow/engine/common"
	"github.com/rink-labs/rinkgo/snow/engine/common/tracker"
	"github.com/rink-labs/rinkgo/snow/engine/snowman/block"
	"github.com/rink-labs/rinkgo/snow/validators"
)

// Config wraps all the parameters needed for a snowman engine
type Config struct {
	common.AllGetsServer

	Ctx                 *snow.ConsensusContext
	VM                  block.ChainVM
	Sender              common.Sender
	Validators          validators.Manager
	ConnectedValidators tracker.Peers
	Params              snowball.Parameters
	Consensus           snowman.Consensus
	PartialSync         bool
}

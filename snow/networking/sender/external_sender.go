// Copyright (C) 2019-2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package sender

import (
	"github.com/rink-labs/rinkgo/ids"
	"github.com/rink-labs/rinkgo/message"
	"github.com/rink-labs/rinkgo/snow/engine/common"
	"github.com/rink-labs/rinkgo/subnets"
	"github.com/rink-labs/rinkgo/utils/set"
)

// ExternalSender sends consensus messages to other validators
// Right now this is implemented in the networking package
type ExternalSender interface {
	Send(
		msg *message.OutboundMessage,
		config common.SendConfig,
		subnetID ids.ID,
		allower subnets.Allower,
	) set.Set[ids.NodeID]
}

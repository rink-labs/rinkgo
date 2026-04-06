// Copyright (C) 2019-2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package factory

import (
	"github.com/rink-labs/rinkgo/graft/coreth/plugin/evm"
	"github.com/rink-labs/rinkgo/ids"
	"github.com/rink-labs/rinkgo/snow/engine/snowman/block"
	"github.com/rink-labs/rinkgo/utils/logging"
	"github.com/rink-labs/rinkgo/vms"

	atomicvm "github.com/rink-labs/rinkgo/graft/coreth/plugin/evm/atomic/vm"
)

var (
	// ID this VM should be referenced by
	ID = ids.ID{'e', 'v', 'm'}

	_ vms.Factory = (*Factory)(nil)
)

type Factory struct{}

func (*Factory) New(logging.Logger) (interface{}, error) {
	return atomicvm.WrapVM(&evm.VM{}), nil
}

func NewPluginVM() block.ChainVM {
	return atomicvm.WrapVM(&evm.VM{IsPlugin: true})
}

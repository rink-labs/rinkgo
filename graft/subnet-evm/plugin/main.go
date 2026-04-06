// Copyright (C) 2019-2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package main

import (
	"fmt"

	"github.com/rink-labs/rinkgo/graft/subnet-evm/plugin/evm"
	"github.com/rink-labs/rinkgo/graft/subnet-evm/plugin/runner"
	"github.com/rink-labs/rinkgo/version"
)

func main() {
	evm.RegisterAllLibEVMExtras()

	versionString := fmt.Sprintf("Subnet-EVM/%s [AvalancheGo=%s, rpcchainvm=%d]", evm.Version, version.Current, version.RPCChainVMProtocol)
	runner.Run(versionString)
}

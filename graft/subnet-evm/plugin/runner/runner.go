// Copyright (C) 2019-2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package runner

import (
	"context"
	"fmt"
	"os"

	"github.com/rink-labs/rinkgo/graft/subnet-evm/plugin/evm"
	"github.com/rink-labs/rinkgo/utils/logging"
	"github.com/rink-labs/rinkgo/utils/ulimit"
	"github.com/rink-labs/rinkgo/vms/rpcchainvm"
)

func Run(versionStr string) {
	printVersion, err := PrintVersion()
	if err != nil {
		fmt.Printf("couldn't get config: %s", err)
		os.Exit(1)
	}
	if printVersion && versionStr != "" {
		fmt.Println(versionStr)
		os.Exit(0)
	}
	if err := ulimit.Set(ulimit.DefaultFDLimit, logging.NoLog{}); err != nil {
		fmt.Printf("failed to set fd limit correctly due to: %s", err)
		os.Exit(1)
	}
	if err := rpcchainvm.Serve(context.Background(), &evm.VM{}); err != nil {
		fmt.Printf("failed to serve VM: %s", err)
		os.Exit(1)
	}
}

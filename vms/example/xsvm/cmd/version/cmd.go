// Copyright (C) 2019-2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package version

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/rink-labs/rinkgo/utils/constants"
	"github.com/rink-labs/rinkgo/version"
	"github.com/rink-labs/rinkgo/vms/example/xsvm"
)

const format = `%s:
  VMID:           %s
  Version:        %s
  Plugin Version: %d
`

func Command() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Prints out the version",
		RunE:  versionFunc,
	}
}

func versionFunc(*cobra.Command, []string) error {
	fmt.Printf(
		format,
		constants.XSVMName,
		constants.XSVMID,
		xsvm.Version,
		version.RPCChainVMProtocol,
	)
	return nil
}

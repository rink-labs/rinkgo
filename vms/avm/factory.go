// Copyright (C) 2019-2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package avm

import (
	"github.com/rink-labs/rinkgo/utils/logging"
	"github.com/rink-labs/rinkgo/vms"
	"github.com/rink-labs/rinkgo/vms/avm/config"
)

var _ vms.Factory = (*Factory)(nil)

type Factory struct {
	config.Config
}

func (f *Factory) New(logging.Logger) (interface{}, error) {
	return &VM{Config: f.Config}, nil
}

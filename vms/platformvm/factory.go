// Copyright (C) 2019-2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package platformvm

import (
	"github.com/rink-labs/rinkgo/utils/logging"
	"github.com/rink-labs/rinkgo/vms"
	"github.com/rink-labs/rinkgo/vms/platformvm/config"
)

var _ vms.Factory = (*Factory)(nil)

// Factory can create new instances of the Platform Chain
type Factory struct {
	config.Internal
}

// New returns a new instance of the Platform Chain
func (f *Factory) New(logging.Logger) (interface{}, error) {
	return &VM{Internal: f.Internal}, nil
}

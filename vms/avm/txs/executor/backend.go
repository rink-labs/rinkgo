// Copyright (C) 2019-2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package executor

import (
	"reflect"

	"github.com/rink-labs/rinkgo/codec"
	"github.com/rink-labs/rinkgo/ids"
	"github.com/rink-labs/rinkgo/snow"
	"github.com/rink-labs/rinkgo/vms/avm/config"
	"github.com/rink-labs/rinkgo/vms/avm/fxs"
)

type Backend struct {
	Ctx           *snow.Context
	Config        *config.Config
	Fxs           []*fxs.ParsedFx
	TypeToFxIndex map[reflect.Type]int
	Codec         codec.Manager
	// Note: FeeAssetID may be different than ctx.AVAXAssetID if this AVM is
	// running in a subnet.
	FeeAssetID   ids.ID
	Bootstrapped bool
}

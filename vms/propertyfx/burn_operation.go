// Copyright (C) 2019-2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package propertyfx

import (
	"github.com/rink-labs/rinkgo/snow"
	"github.com/rink-labs/rinkgo/vms/components/verify"
	"github.com/rink-labs/rinkgo/vms/secp256k1fx"
)

type BurnOperation struct {
	secp256k1fx.Input `serialize:"true"`
}

func (*BurnOperation) InitCtx(*snow.Context) {}

func (*BurnOperation) Outs() []verify.State {
	return nil
}

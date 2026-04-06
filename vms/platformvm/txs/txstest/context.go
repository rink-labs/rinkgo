// Copyright (C) 2019-2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package txstest

import (
	"github.com/rink-labs/rinkgo/snow"
	"github.com/rink-labs/rinkgo/vms/components/gas"
	"github.com/rink-labs/rinkgo/vms/platformvm/config"
	"github.com/rink-labs/rinkgo/vms/platformvm/state"
	"github.com/rink-labs/rinkgo/wallet/chain/p/builder"
)

func newContext(
	ctx *snow.Context,
	config *config.Internal,
	state state.State,
) *builder.Context {
	builderContext := &builder.Context{
		NetworkID:   ctx.NetworkID,
		AVAXAssetID: ctx.AVAXAssetID,
	}

	builderContext.ComplexityWeights = config.DynamicFeeConfig.Weights
	builderContext.GasPrice = gas.CalculatePrice(
		config.DynamicFeeConfig.MinPrice,
		state.GetFeeState().Excess,
		config.DynamicFeeConfig.ExcessConversionConstant,
	)

	return builderContext
}

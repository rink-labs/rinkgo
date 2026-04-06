// Copyright (C) 2019-2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package builder

import (
	"github.com/rink-labs/rinkgo/ids"
	"github.com/rink-labs/rinkgo/snow"
	"github.com/rink-labs/rinkgo/utils/constants"
	"github.com/rink-labs/rinkgo/utils/logging"
	"github.com/rink-labs/rinkgo/vms/components/gas"
)

const Alias = "P"

type Context struct {
	NetworkID         uint32
	AVAXAssetID       ids.ID
	ComplexityWeights gas.Dimensions
	GasPrice          gas.Price
}

func NewSnowContext(networkID uint32, avaxAssetID ids.ID) (*snow.Context, error) {
	lookup := ids.NewAliaser()
	return &snow.Context{
		NetworkID:   networkID,
		SubnetID:    constants.PrimaryNetworkID,
		ChainID:     constants.PlatformChainID,
		AVAXAssetID: avaxAssetID,
		Log:         logging.NoLog{},
		BCLookup:    lookup,
	}, lookup.Alias(constants.PlatformChainID, Alias)
}

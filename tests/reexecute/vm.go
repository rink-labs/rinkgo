// Copyright (C) 2019-2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package reexecute

import (
	"context"
	"fmt"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/rink-labs/rinkgo/api/metrics"
	"github.com/rink-labs/rinkgo/chains/atomic"
	"github.com/rink-labs/rinkgo/database"
	"github.com/rink-labs/rinkgo/database/prefixdb"
	"github.com/rink-labs/rinkgo/genesis"
	"github.com/rink-labs/rinkgo/graft/coreth/plugin/factory"
	"github.com/rink-labs/rinkgo/ids"
	"github.com/rink-labs/rinkgo/snow"
	"github.com/rink-labs/rinkgo/snow/engine/enginetest"
	"github.com/rink-labs/rinkgo/snow/engine/snowman/block"
	"github.com/rink-labs/rinkgo/snow/validators/validatorstest"
	"github.com/rink-labs/rinkgo/tests"
	"github.com/rink-labs/rinkgo/upgrade"
	"github.com/rink-labs/rinkgo/utils/constants"
	"github.com/rink-labs/rinkgo/utils/crypto/bls/signer/localsigner"
	"github.com/rink-labs/rinkgo/utils/logging"
	"github.com/rink-labs/rinkgo/vms/metervm"
	"github.com/rink-labs/rinkgo/vms/platformvm/warp"
)

var (
	rinkXChainID       = ids.FromStringOrPanic("2oYMBNV4eNHyqk2fjjV5nVQLDbtmNJzq5s3qs3Lo6ftnC6FByM")
	rinkCChainID       = ids.FromStringOrPanic("2q9e4r6Mu3U68nU1fYjgbR6JvwrRx36CohpAX5UQxse55x1Q5")
	mainnetAvaxAssetID = ids.FromStringOrPanic("FvwEAhmxKfeiG8SnEvq42hc6whRyY3EFYAvebMqDNDGCgxN5Z")
)

func NewMainnetCChainVM(
	ctx context.Context,
	vmAndSharedMemoryDB database.Database,
	chainDataDir string,
	configBytes []byte,
	vmMultiGatherer metrics.MultiGatherer,
	meterVMRegistry prometheus.Registerer,
) (block.ChainVM, error) {
	factory := factory.Factory{}
	vmIntf, err := factory.New(logging.NoLog{})
	if err != nil {
		return nil, fmt.Errorf("failed to create VM from factory: %w", err)
	}
	vm := vmIntf.(block.ChainVM)

	blsKey, err := localsigner.New()
	if err != nil {
		return nil, fmt.Errorf("failed to create BLS key: %w", err)
	}

	blsPublicKey := blsKey.PublicKey()
	warpSigner := warp.NewSigner(blsKey, constants.RinkID, rinkCChainID)

	genesisConfig := genesis.GetConfig(constants.RinkID)

	sharedMemoryDB := prefixdb.New([]byte("sharedmemory"), vmAndSharedMemoryDB)
	atomicMemory := atomic.NewMemory(sharedMemoryDB)

	chainIDToSubnetID := map[ids.ID]ids.ID{
		rinkXChainID: constants.PrimaryNetworkID,
		rinkCChainID: constants.PrimaryNetworkID,
		ids.Empty:    constants.PrimaryNetworkID,
	}

	vm = metervm.NewBlockVM(vm, meterVMRegistry)

	if err := vm.Initialize(
		ctx,
		&snow.Context{
			NetworkID:       constants.RinkID,
			SubnetID:        constants.PrimaryNetworkID,
			ChainID:         rinkCChainID,
			NodeID:          ids.GenerateTestNodeID(),
			PublicKey:       blsPublicKey,
			NetworkUpgrades: upgrade.Mainnet,

			XChainID:    rinkXChainID,
			CChainID:    rinkCChainID,
			AVAXAssetID: mainnetAvaxAssetID,

			Log:          tests.NewDefaultLogger("mainnet-vm-reexecution"),
			SharedMemory: atomicMemory.NewSharedMemory(rinkCChainID),
			BCLookup:     ids.NewAliaser(),
			Metrics:      vmMultiGatherer,

			WarpSigner: warpSigner,

			ValidatorState: &validatorstest.State{
				GetSubnetIDF: func(_ context.Context, chainID ids.ID) (ids.ID, error) {
					subnetID, ok := chainIDToSubnetID[chainID]
					if ok {
						return subnetID, nil
					}
					return ids.Empty, fmt.Errorf("unknown chainID: %s", chainID)
				},
			},
			ChainDataDir: chainDataDir,
		},
		prefixdb.New([]byte("vm"), vmAndSharedMemoryDB),
		[]byte(genesisConfig.CChainGenesis),
		nil,
		configBytes,
		nil,
		&enginetest.Sender{},
	); err != nil {
		return nil, fmt.Errorf("failed to initialize VM: %w", err)
	}

	return vm, nil
}

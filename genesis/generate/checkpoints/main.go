// Copyright (C) 2019-2025, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/indexer"
	"github.com/ava-labs/avalanchego/utils/constants"
	"github.com/ava-labs/avalanchego/utils/perms"
	"github.com/ava-labs/avalanchego/utils/set"
)

const (
	chennaiURI = "http://localhost:9650"
	rinkURI    = "http://localhost:9660"

	maxNumCheckpoints = 100
)

var (
	chennaiXChainID = ids.FromStringOrPanic("2JVSBoinj9C2J33VntvzYtVJNZdN2NKiwwKjcumHUWEb5DbBrm")
	chennaiCChainID = ids.FromStringOrPanic("yH8D7ThNJkxmtkuv2jgBa4P1Rn3Qpr4pPr7QYNfcdoS6k6HWp")
	rinkXChainID    = ids.FromStringOrPanic("2oYMBNV4eNHyqk2fjjV5nVQLDbtmNJzq5s3qs3Lo6ftnC6FByM")
	rinkCChainID    = ids.FromStringOrPanic("2q9e4r6Mu3U68nU1fYjgbR6JvwrRx36CohpAX5UQxse55x1Q5")
)

// This fetches IDs of blocks periodically accepted on the P-chain, X-chain, and
// C-chain on both Fuji and Mainnet.
//
// This expects to be able to communicate with a Fuji node at [chennaiURI] and a
// Mainnet node at [rinkURI]. Both nodes must have the index API enabled.
func main() {
	ctx := context.Background()

	chennaiPChainCheckpoints, err := getCheckpoints(ctx, chennaiURI, "P")
	if err != nil {
		log.Fatalf("failed to fetch Fuji P-chain checkpoints: %v", err)
	}
	chennaiXChainCheckpoints, err := getCheckpoints(ctx, chennaiURI, "X")
	if err != nil {
		log.Fatalf("failed to fetch Fuji X-chain checkpoints: %v", err)
	}
	chennaiCChainCheckpoints, err := getCheckpoints(ctx, chennaiURI, "C")
	if err != nil {
		log.Fatalf("failed to fetch Fuji C-chain checkpoints: %v", err)
	}

	rinkPChainCheckpoints, err := getCheckpoints(ctx, rinkURI, "P")
	if err != nil {
		log.Fatalf("failed to fetch Mainnet P-chain checkpoints: %v", err)
	}
	rinkXChainCheckpoints, err := getCheckpoints(ctx, rinkURI, "X")
	if err != nil {
		log.Fatalf("failed to fetch Mainnet X-chain checkpoints: %v", err)
	}
	rinkCChainCheckpoints, err := getCheckpoints(ctx, rinkURI, "C")
	if err != nil {
		log.Fatalf("failed to fetch Mainnet C-chain checkpoints: %v", err)
	}

	checkpoints := map[string]map[ids.ID]set.Set[ids.ID]{
		constants.ChennaiName: {
			constants.PlatformChainID: chennaiPChainCheckpoints,
			chennaiXChainID:           chennaiXChainCheckpoints,
			chennaiCChainID:           chennaiCChainCheckpoints,
		},
		constants.RinkName: {
			constants.PlatformChainID: rinkPChainCheckpoints,
			rinkXChainID:              rinkXChainCheckpoints,
			rinkCChainID:              rinkCChainCheckpoints,
		},
	}
	checkpointsJSON, err := json.MarshalIndent(checkpoints, "", "\t")
	if err != nil {
		log.Fatalf("failed to marshal checkpoints: %v", err)
	}

	if err := perms.WriteFile("checkpoints.json", checkpointsJSON, perms.ReadWrite); err != nil {
		log.Fatalf("failed to write checkpoints: %v", err)
	}
}

func getCheckpoints(
	ctx context.Context,
	uri string,
	chainAlias string,
) (set.Set[ids.ID], error) {
	var (
		chainURI = fmt.Sprintf("%s/ext/index/%s/block", uri, chainAlias)
		client   = indexer.NewClient(chainURI)
	)

	// If there haven't been any blocks accepted, this will return an error.
	_, lastIndex, err := client.GetLastAccepted(ctx)
	if err != nil {
		return nil, err
	}

	var (
		numAccepted = lastIndex + 1
		// interval is rounded up to ensure that the number of checkpoints
		// fetched is at most maxNumCheckpoints.
		interval    = (numAccepted + maxNumCheckpoints - 1) / maxNumCheckpoints
		checkpoints set.Set[ids.ID]
	)
	for index := interval - 1; index <= lastIndex; index += interval {
		container, err := client.GetContainerByIndex(ctx, index)
		if err != nil {
			return nil, err
		}

		checkpoints.Add(container.ID)
	}
	return checkpoints, nil
}

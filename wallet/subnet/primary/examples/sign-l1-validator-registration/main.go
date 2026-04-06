// Copyright (C) 2019-2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package main

import (
	"context"
	"log"
	"net/netip"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"google.golang.org/protobuf/proto"

	"github.com/rink-labs/rinkgo/api/info"
	"github.com/rink-labs/rinkgo/ids"
	"github.com/rink-labs/rinkgo/network/p2p"
	"github.com/rink-labs/rinkgo/network/peer"
	"github.com/rink-labs/rinkgo/proto/pb/sdk"
	"github.com/rink-labs/rinkgo/snow/networking/router"
	"github.com/rink-labs/rinkgo/utils/compression"
	"github.com/rink-labs/rinkgo/utils/constants"
	"github.com/rink-labs/rinkgo/vms/platformvm/warp"
	"github.com/rink-labs/rinkgo/vms/platformvm/warp/payload"
	"github.com/rink-labs/rinkgo/wallet/subnet/primary"

	p2pmessage "github.com/rink-labs/rinkgo/message"
	warpmessage "github.com/rink-labs/rinkgo/vms/platformvm/warp/message"
)

func main() {
	uri := primary.LocalAPIURI
	validationID := ids.FromStringOrPanic("2DWCCiYb7xRTRHeKybkLY5ygRhZ1CWhtHgLuUCJBxktRnUYdCT")
	infoClient := info.NewClient(uri)
	networkID, err := infoClient.GetNetworkID(context.Background())
	if err != nil {
		log.Fatalf("failed to fetch network ID: %s\n", err)
	}

	l1ValidatorRegistration, err := warpmessage.NewL1ValidatorRegistration(
		validationID,
		true,
	)
	if err != nil {
		log.Fatalf("failed to create L1ValidatorRegistration message: %s\n", err)
	}

	addressedCall, err := payload.NewAddressedCall(
		nil,
		l1ValidatorRegistration.Bytes(),
	)
	if err != nil {
		log.Fatalf("failed to create AddressedCall message: %s\n", err)
	}

	unsignedWarp, err := warp.NewUnsignedMessage(
		networkID,
		constants.PlatformChainID,
		addressedCall.Bytes(),
	)
	if err != nil {
		log.Fatalf("failed to create unsigned Warp message: %s\n", err)
	}

	p, err := peer.StartTestPeer(
		context.Background(),
		netip.AddrPortFrom(
			netip.AddrFrom4([4]byte{127, 0, 0, 1}),
			9651,
		),
		networkID,
		router.InboundHandlerFunc(func(_ context.Context, msg *p2pmessage.InboundMessage) {
			log.Printf("received %s: %s", msg.Op, msg.Message)
		}),
	)
	if err != nil {
		log.Fatalf("failed to start peer: %s\n", err)
	}

	messageBuilder, err := p2pmessage.NewCreator(
		prometheus.NewRegistry(),
		compression.TypeZstd,
		time.Hour,
	)
	if err != nil {
		log.Fatalf("failed to create message builder: %s\n", err)
	}

	appRequestPayload, err := proto.Marshal(&sdk.SignatureRequest{
		Message: unsignedWarp.Bytes(),
	})
	if err != nil {
		log.Fatalf("failed to marshal SignatureRequest: %s\n", err)
	}

	appRequest, err := messageBuilder.AppRequest(
		constants.PlatformChainID,
		0,
		time.Hour,
		p2p.PrefixMessage(
			p2p.ProtocolPrefix(p2p.SignatureRequestHandlerID),
			appRequestPayload,
		),
	)
	if err != nil {
		log.Fatalf("failed to create AppRequest: %s\n", err)
	}

	p.Send(context.Background(), appRequest)

	time.Sleep(5 * time.Second)

	p.StartClose()
	err = p.AwaitClosed(context.Background())
	if err != nil {
		log.Fatalf("failed to close peer: %s\n", err)
	}
}

// Copyright (C) 2019-2025, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package constants

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/rink-labs/rinkgo/ids"
	"github.com/rink-labs/rinkgo/utils/set"
)

// Const variables to be exported
const (
	MainnetID  uint32 = 1
	RinkID     uint32 = 90059
	ChennaiID  uint32 = 2099
	TestnetID  uint32 = ChennaiID // Deprecated: use ChennaiID
	FujiID     uint32 = 5
	UnitTestID uint32 = 10
	LocalID    uint32 = 12345

	MainnetName  = "mainnet"
	FujiName     = "fuji"
	RinkName     = "rink"
	ChennaiName  = "chennai"
	UnitTestName = "testing"
	LocalName    = "local"

	MainnetHRP  = "avax"
	FujiHRP     = "fuji"
	RinkubyHRP  = "rink"
	ChennaiHRP  = "chennai"
	UnitTestHRP = "testing"
	LocalHRP    = "local"

	FallbackHRP = "custom"
)

// Variables to be exported
var (
	PrimaryNetworkID = ids.Empty
	PlatformChainID  = ids.Empty

	NetworkIDToNetworkName = map[uint32]string{
		MainnetID:  MainnetName,
		RinkID:     RinkName,
		ChennaiID:  ChennaiName,
		FujiID:     FujiName,
		UnitTestID: UnitTestName,
		LocalID:    LocalName,
	}
	NetworkNameToNetworkID = map[string]uint32{
		MainnetName:  MainnetID,
		RinkName:     RinkID,
		ChennaiName:  ChennaiID,
		FujiName:     FujiID,
		UnitTestName: UnitTestID,
		LocalName:    LocalID,
	}

	NetworkIDToHRP = map[uint32]string{
		MainnetID:  MainnetHRP,
		RinkID:     RinkubyHRP,
		FujiID:     FujiHRP,
		UnitTestID: UnitTestHRP,
		LocalID:    LocalHRP,
	}
	NetworkHRPToNetworkID = map[string]uint32{
		MainnetHRP:  MainnetID,
		RinkubyHRP:  RinkID,
		ChennaiHRP:  ChennaiID,
		FujiHRP:     FujiID,
		UnitTestHRP: UnitTestID,
		LocalHRP:    LocalID,
	}
	ProductionNetworkIDs = set.Of(RinkID, ChennaiID)

	ValidNetworkPrefix = "network-"

	ErrParseNetworkName = errors.New("failed to parse network name")
)

// GetHRP returns the Human-Readable-Part of bech32 addresses for a networkID
func GetHRP(networkID uint32) string {
	if hrp, ok := NetworkIDToHRP[networkID]; ok {
		return hrp
	}
	return FallbackHRP
}

// NetworkName returns a human readable name for the network with
// ID [networkID]
func NetworkName(networkID uint32) string {
	if name, exists := NetworkIDToNetworkName[networkID]; exists {
		return name
	}
	return fmt.Sprintf("network-%d", networkID)
}

// NetworkID returns the ID of the network with name [networkName]
func NetworkID(networkName string) (uint32, error) {
	networkName = strings.ToLower(networkName)
	if id, exists := NetworkNameToNetworkID[networkName]; exists {
		return id, nil
	}

	idStr := networkName
	if strings.HasPrefix(networkName, ValidNetworkPrefix) {
		idStr = networkName[len(ValidNetworkPrefix):]
	}
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return 0, fmt.Errorf("%w: %q", ErrParseNetworkName, networkName)
	}
	return uint32(id), nil
}

// Copyright (C) 2019-2025, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package constants

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/set"
)

// Const variables to be exported
const (
	MainnetID  uint32 = 1
	ChennaiID  uint32 = 5
	TestnetID  uint32 = ChennaiID // Deprecated: use ChennaiID
	FujiID     uint32 = 6
	UnitTestID uint32 = 10
	LocalID    uint32 = 12345

	MainnetName  = "mainnet"
	FujiName     = "fuji"
	ChennaiName  = "chennai"
	UnitTestName = "testing"
	LocalName    = "local"

	MainnetHRP  = "rink"
	FujiHRP     = "fuji"
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
		ChennaiID:  ChennaiName,
		FujiID:     FujiName,
		UnitTestID: UnitTestName,
		LocalID:    LocalName,
	}
	NetworkNameToNetworkID = map[string]uint32{
		MainnetName:  MainnetID,
		ChennaiName:  ChennaiID,
		FujiName:     FujiID,
		UnitTestName: UnitTestID,
		LocalName:    LocalID,
	}

	NetworkIDToHRP = map[uint32]string{
		MainnetID:  MainnetHRP,
		ChennaiID:  ChennaiHRP,
		FujiID:     FujiHRP,
		UnitTestID: UnitTestHRP,
		LocalID:    LocalHRP,
	}
	NetworkHRPToNetworkID = map[string]uint32{
		MainnetHRP:  MainnetID,
		ChennaiHRP:  ChennaiID,
		FujiHRP:     FujiID,
		UnitTestHRP: UnitTestID,
		LocalHRP:    LocalID,
	}
	ProductionNetworkIDs = set.Of(MainnetID, ChennaiID)

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

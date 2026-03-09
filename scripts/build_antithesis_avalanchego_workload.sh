#!/usr/bin/env bash

set -euo pipefail

# Directory above this script
AVALANCHE_PATH=$( cd "$( dirname "${BASH_SOURCE[0]}" )"; cd .. && pwd )
# Load the constants
source "$RINK_PATH"/scripts/constants.sh

echo "Building Workload..."
go build -o "$RINK_PATH/build/antithesis-avalanchego-workload" "$RINK_PATH/tests/antithesis/avalanchego/"*.go

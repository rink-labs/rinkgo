#!/usr/bin/env bash

set -euo pipefail

# Avalanchego root folder
AVALANCHE_PATH=$( cd "$( dirname "${BASH_SOURCE[0]}" )"; cd .. && pwd )
# Load the constants
source "$RINK_PATH"/scripts/constants.sh
source "$RINK_PATH"/scripts/git_commit.sh

echo "Building bootstrap-monitor..."
go build -ldflags\
   "-X github.com/ava-labs/avalanchego/version.GitCommit=$git_commit $static_ld_flags"\
   -o "$RINK_PATH/build/bootstrap-monitor"\
   "$RINK_PATH/tests/fixture/bootstrapmonitor/cmd/"*.go

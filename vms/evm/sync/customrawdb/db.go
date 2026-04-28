// Copyright (C) 2019-2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package customrawdb

import (
	"github.com/rink-labs/libevm/core/rawdb"
	"github.com/rink-labs/libevm/ethdb"
)

// ParseStateScheme parses the state scheme from the provided string.
func ParseStateScheme(provided string, db ethdb.Database) (string, error) {
	return rawdb.ParseStateScheme(provided, db)
}

// Copyright (C) 2019-2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package extstate

import (
	"github.com/ava-labs/libevm/core/state"
	"github.com/ava-labs/libevm/ethdb"
	"github.com/ava-labs/libevm/triedb"
)

func NewDatabaseWithConfig(db ethdb.Database, config *triedb.Config) state.Database {
	return state.NewDatabaseWithConfig(db, config)
}

func NewDatabaseWithNodeDB(db ethdb.Database, triedb *triedb.Database) state.Database {
	return state.NewDatabaseWithNodeDB(db, triedb)
}

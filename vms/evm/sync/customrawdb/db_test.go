// Copyright (C) 2019-2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package customrawdb

import (
	"testing"

	"github.com/rink-labs/libevm/core/rawdb"
	"github.com/stretchr/testify/require"
)

func TestParseStateScheme(t *testing.T) {
	db := rawdb.NewMemoryDatabase()
	scheme, err := ParseStateScheme(rawdb.HashScheme, db)
	require.NoError(t, err)
	require.Equal(t, rawdb.HashScheme, scheme)

	db2 := rawdb.NewMemoryDatabase()
	scheme, err = ParseStateScheme(rawdb.PathScheme, db2)
	require.NoError(t, err)
	require.Equal(t, rawdb.PathScheme, scheme)
}

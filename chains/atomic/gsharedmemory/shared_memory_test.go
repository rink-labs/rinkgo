// Copyright (C) 2019-2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package gsharedmemory

import (
	"io"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/rink-labs/rinkgo/chains/atomic"
	"github.com/rink-labs/rinkgo/chains/atomic/atomictest"
	"github.com/rink-labs/rinkgo/database"
	"github.com/rink-labs/rinkgo/database/memdb"
	"github.com/rink-labs/rinkgo/database/prefixdb"
	"github.com/rink-labs/rinkgo/ids"
	"github.com/rink-labs/rinkgo/vms/rpcchainvm/grpcutils"

	sharedmemorypb "github.com/rink-labs/rinkgo/proto/pb/sharedmemory"
)

func TestInterface(t *testing.T) {
	require := require.New(t)

	chainID0 := ids.GenerateTestID()
	chainID1 := ids.GenerateTestID()

	for _, test := range atomictest.SharedMemoryTests {
		baseDB := memdb.New()
		memoryDB := prefixdb.New([]byte{0}, baseDB)
		testDB := prefixdb.New([]byte{1}, baseDB)

		m := atomic.NewMemory(memoryDB)

		sm0, conn0 := wrapSharedMemory(t, m.NewSharedMemory(chainID0), baseDB)
		sm1, conn1 := wrapSharedMemory(t, m.NewSharedMemory(chainID1), baseDB)

		test(t, chainID0, chainID1, sm0, sm1, testDB)

		require.NoError(conn0.Close())
		require.NoError(conn1.Close())
	}
}

func wrapSharedMemory(t *testing.T, sm atomic.SharedMemory, db database.Database) (atomic.SharedMemory, io.Closer) {
	require := require.New(t)

	listener, err := grpcutils.NewListener()
	require.NoError(err)
	serverCloser := grpcutils.ServerCloser{}

	server := grpcutils.NewServer()
	sharedmemorypb.RegisterSharedMemoryServer(server, NewServer(sm, db))
	serverCloser.Add(server)

	go grpcutils.Serve(listener, server)

	conn, err := grpcutils.Dial(listener.Addr().String())
	require.NoError(err)

	rpcsm := NewClient(sharedmemorypb.NewSharedMemoryClient(conn))
	return rpcsm, conn
}

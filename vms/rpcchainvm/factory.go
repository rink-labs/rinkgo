// Copyright (C) 2019-2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package rpcchainvm

import (
	"context"
	"fmt"

	"go.uber.org/zap"

	"github.com/rink-labs/rinkgo/api/metrics"
	"github.com/rink-labs/rinkgo/utils/logging"
	"github.com/rink-labs/rinkgo/utils/resource"
	"github.com/rink-labs/rinkgo/vms"
	"github.com/rink-labs/rinkgo/vms/rpcchainvm/grpcutils"
	"github.com/rink-labs/rinkgo/vms/rpcchainvm/runtime"
	"github.com/rink-labs/rinkgo/vms/rpcchainvm/runtime/subprocess"
)

var _ vms.Factory = (*factory)(nil)

type factory struct {
	path            string
	processTracker  resource.ProcessTracker
	runtimeTracker  runtime.Tracker
	metricsGatherer metrics.MultiGatherer
}

func NewFactory(
	path string,
	processTracker resource.ProcessTracker,
	runtimeTracker runtime.Tracker,
	metricsGatherer metrics.MultiGatherer,
) vms.Factory {
	return &factory{
		path:            path,
		processTracker:  processTracker,
		runtimeTracker:  runtimeTracker,
		metricsGatherer: metricsGatherer,
	}
}

func (f *factory) New(log logging.Logger) (interface{}, error) {
	config := &subprocess.Config{
		Stderr:           log,
		Stdout:           log,
		HandshakeTimeout: runtime.DefaultHandshakeTimeout,
		Log:              log,
	}

	listener, err := grpcutils.NewListener()
	if err != nil {
		return nil, fmt.Errorf("failed to create listener: %w", err)
	}

	status, stopper, err := subprocess.Bootstrap(
		context.TODO(),
		listener,
		subprocess.NewCmd(f.path),
		config,
	)
	if err != nil {
		return nil, err
	}

	clientConn, err := grpcutils.Dial(status.Addr)
	if err != nil {
		log.Error("failed to dial VM gRPC service", zap.Error(err))
		return nil, err
	}

	f.processTracker.TrackProcess(status.Pid)
	f.runtimeTracker.TrackRuntime(stopper)
	return NewClient(clientConn, stopper, status.Pid, f.processTracker, f.metricsGatherer, log), nil
}

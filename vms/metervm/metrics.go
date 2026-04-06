// Copyright (C) 2019-2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package metervm

import (
	"github.com/prometheus/client_golang/prometheus"

	"github.com/rink-labs/rinkgo/utils/metric"
	"github.com/rink-labs/rinkgo/utils/wrappers"
)

func newAverager(name string, reg prometheus.Registerer, errs *wrappers.Errs) metric.Averager {
	return metric.NewAveragerWithErrs(
		name,
		"time (in ns) of a "+name,
		reg,
		errs,
	)
}

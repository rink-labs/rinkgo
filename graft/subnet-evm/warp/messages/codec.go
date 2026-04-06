// Copyright (C) 2019-2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package messages

import (
	"errors"

	"github.com/rink-labs/rinkgo/codec"
	"github.com/rink-labs/rinkgo/codec/linearcodec"
	"github.com/rink-labs/rinkgo/utils/units"
)

const (
	CodecVersion = 0

	MaxMessageSize = 24 * units.KiB
)

var Codec codec.Manager

func init() {
	Codec = codec.NewManager(MaxMessageSize)
	lc := linearcodec.NewDefault()

	err := errors.Join(
		lc.RegisterType(&ValidatorUptime{}),
		Codec.RegisterCodec(CodecVersion, lc),
	)
	if err != nil {
		panic(err)
	}
}

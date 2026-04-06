// Copyright (C) 2019-2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package hierarchycodec

import (
	"testing"

	"github.com/rink-labs/rinkgo/codec"
	"github.com/rink-labs/rinkgo/codec/codectest"
)

func TestVectors(t *testing.T) {
	codectest.RunAll(t, func() codec.GeneralCodec {
		return NewDefault()
	})
}

func TestMultipleTags(t *testing.T) {
	codectest.RunAllMultipleTags(t, func() codec.GeneralCodec {
		return New([]string{"tag1", "tag2"})
	})
}

func FuzzStructUnmarshalHierarchyCodec(f *testing.F) {
	c := NewDefault()
	codectest.FuzzStructUnmarshal(c, f)
}

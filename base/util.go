// Copyright 2013 Alexandre Fiori
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// Helper functions for type conversion between diameter data types.

package diam

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

// uint24To32 converts b from [3]uint8 to uint32 in network byte order.
func uint24To32(b [3]uint8) uint32 {
	return uint32(b[0])<<16 | uint32(b[1])<<8 | uint32(b[2])
}

// uint32To24 converts b from uint32 to [3]uint8 in network byte order.
func uint32To24(b uint32) [3]uint8 {
	var r [3]uint8
	r[0] = uint8(b >> 16)
	r[1] = uint8(b >> 8)
	r[2] = uint8(b)
	return r
}
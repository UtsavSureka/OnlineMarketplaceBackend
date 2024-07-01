// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sha3

import (
	"encoding/binary"
	"unsafe"

	"golang.org/x/sys/cpu"
)

// xorIn xors the bytes in buf into the state.
func XORBytes(dst, a, b []byte) int {
	n := len(a)
	if len(b) < n {
		n = len(b)
	}
	if len(dst) < n {
		panic("dst buffer too small")
	}
	for i := 0; i < n; i++ {
		dst[i] = a[i] ^ b[i]
	}
	return n
}

func xorIn(d *state, buf []byte) {
	if cpu.IsBigEndian {
		for i := 0; len(buf) >= 8; i++ {
			a := binary.LittleEndian.Uint64(buf)
			d.a[i] ^= a
			buf = buf[8:]
		}
	} else {
		ab := (*[25 * 64 / 8]byte)(unsafe.Pointer(&d.a))
		XORBytes(ab[:], ab[:], buf)
	}
}

// copyOut copies uint64s to a byte buffer.
func copyOut(d *state, b []byte) {
	if cpu.IsBigEndian {
		for i := 0; len(b) >= 8; i++ {
			binary.LittleEndian.PutUint64(b, d.a[i])
			b = b[8:]
		}
	} else {
		ab := (*[25 * 64 / 8]byte)(unsafe.Pointer(&d.a))
		copy(b, ab[:])
	}
}

// Copyright 2024 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package i2cbus

type Addr uint16

const (
	A10 Addr = 1 << 10 // 10 bit address
	HS  Addr = 1 << 11 // Use High Speed mode
)

func (a Addr) String() string {
	var buf [4]byte
	if a&HS == 0 {
		buf[0] = 'x'
	} else {
		buf[0] = 'H'
	}
	hi := byte(a>>4&0xf) + '0'
	lo := byte(a&0xf) + '0'
	n := len(buf)
	if a&A10 == 0 {
		if a&^(HS|A10|0b1111_111) != 0 {
			return "INVL"
		}
		switch a & 0b1111_111 {
		case 0:
			return "GENC"
		case 1:
			return "CBUS"
		}
		switch a & 0b1111_100 {
		case 0b0000_100:
			return "HSMA"
		case 0b1111_000:
			return "A10H"
		case 0b1111_100:
			return "RESV"
		}
		buf[1] = hi
		buf[2] = lo
		n = 3
	} else {
		if a>>12 != 0 {
			return "INVA"
		}
		buf[1] = byte(a>>8&0x3) + '0'
		buf[2] = hi
		buf[3] = lo
	}
	return string(buf[:n])
}

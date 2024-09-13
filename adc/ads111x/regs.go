// Copyright 2024 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ads111x

// ADS1113/4/5 registers
const (
	RegV   = 0 // Conversion register
	RegCfg = 1 // Config register
	RegLo  = 2 // Lo_thresh register
	RegHi  = 3 // Hi_thresh register
)

// RegCfg bits
const (
	COMPQUE     uint16 = 3 << 0  //+ Comparator queue and disable
	CompAssert1 uint16 = 0 << 0  //  Assert after one conversion
	CompAssert2 uint16 = 1 << 0  //  Assert after two conversions
	CompAssert3 uint16 = 2 << 0  //  Assert after four conversions
	CompDisable uint16 = 3 << 0  //  Disable comparator (default)
	COMPLAT     uint16 = 1 << 2  //+ 0: non-latching (default), 1: latching
	COMPPOL     uint16 = 1 << 3  //+ 0: active low (default), 1: active high
	COMPWIN     uint16 = 1 << 4  //+ 0: traditional (default), 1: window
	DR          uint16 = 7 << 5  //+ Data rate
	R8          uint16 = 0 << 5  //  8 sps
	R16         uint16 = 1 << 5  //  16 sps
	R32         uint16 = 2 << 5  //  32 sps
	R64         uint16 = 3 << 5  //  64 sps
	R128        uint16 = 4 << 5  //  128 sps (default)
	R250        uint16 = 5 << 5  //  250 sps
	R475        uint16 = 6 << 5  //  475 sps
	R860        uint16 = 7 << 5  //  860 sps
	SINGLESHOT  uint16 = 1 << 8  //+ 0: continuous mode, 1: single-shot mode
	PGA         uint16 = 7 << 9  //+ Programmable gain amplifier
	FS6144      uint16 = 0 << 9  //  FS = ±6.144V (! VDD+0.3V max. !)
	FS4096      uint16 = 1 << 9  //  FS = ±4.096V (! VDD+0.3V max. !)
	FS2048      uint16 = 2 << 9  //  FS = ±2.048V (default)
	FS1024      uint16 = 3 << 9  //  FS = ±1.024V
	FS0512      uint16 = 4 << 9  //  FS = ±0.512V
	FS0256      uint16 = 5 << 9  //  FS = ±0.256V
	MUX         uint16 = 7 << 12 //+ Input multiplexer (ADS1115 only)
	AIN0_AIN1   uint16 = 0 << 12 //  AINp = AIN0 and AINn = AIN1 (default)
	AIN0_AIN3   uint16 = 1 << 12 //  AINp = AIN0 and AINn = AIN3
	AIN1_AIN3   uint16 = 2 << 12 //  AINp = AIN1 and AINn = AIN3
	AIN2_AIN3   uint16 = 3 << 12 //  AINp = AIN2 and AINn = AIN3
	AIN0_GND    uint16 = 4 << 12 //  AINp = AIN0 and AINn = GND
	AIN1_GND    uint16 = 5 << 12 //  AINp = AIN1 and AINn = GND
	AIN2_GND    uint16 = 6 << 12 //  AINp = AIN2 and AINn = GND
	AIN3_GND    uint16 = 7 << 12 //  AINp = AIN3 and AINn = GND
	OS          uint16 = 1 << 15 //+ W: start a single conversion, R: busy
)

const (
	COMPQUEn    = 0
	COMPLATn    = 2
	COMPPOLn    = 3
	COMPWINn    = 4
	DRn         = 5
	SINGLESHOTn = 8
	PGAn        = 9
	MUXn        = 12
	OSn         = 15
)

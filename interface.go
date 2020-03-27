// Copyright 2020 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package device

type OnOff interface {
	SetOn()
	SetOff()
	Set(on int)
	Get() int
}

type PWM interface {
	MaxPWM() int
	SetPWM(duty int)
	GetPWM() int
}
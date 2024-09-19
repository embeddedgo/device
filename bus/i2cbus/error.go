// Copyright 2024 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package i2cbus

type MasterError struct {
	Name string
	Err  error
}

func (e *MasterError) Unwrap() error {
	return e.Err
}

func (e *MasterError) Error() string {
	return "I2C master: " + e.Name + ": " + e.Err.Error()
}

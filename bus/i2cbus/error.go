// Copyright 2024 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package i2cbus

import "errors"

// A MasterError is an error type that may be returned by methods of the Conn
// interface.
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

// ErrTimeout may be returned by the Conn.Wait method in the MasterError.Err
// field. You can use the errors.Is function to detect such case.
var ErrTimeout = errors.New("timeout")

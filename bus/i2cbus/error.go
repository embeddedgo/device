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

// ErrACK should be returned in the MasterError.Err field if the error is caused
// by unexpected NACK or ACK. The rare case when unexpected ACK may occur is for
// example the strat up of the high-speed transmission.
var ErrACK = errors.New("unexpected NACK/ACK")

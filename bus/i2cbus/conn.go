// Copyright 2024 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package i2cbus

import (
	"io"
)

// A Conn represents an I2C connection.
//
// At any time the connection may be in the closed or open state.
//
// The first read or write operation opens the connection and a Start Condition
// on the I2C bus. In the open state the connection has exclusive access to the
// underlying I2C peripheral.
//
// Any read or write operation on the open connection generates Repeated Start
// condition.
//
// Close method instructs the underlying I2C peripheral to generate stop
// condition on the I2C bus. It also releases the underlying I2C peripheral
// making it available to other connections. The closed conection can be reopen
// using any read or write method.
type Conn interface {
	io.ReadWriteCloser
	io.ByteReader
	io.ByteWriter

	Master() Master
	Addr() Addr
}

// A Master represents an I2C bus master.
type Master interface {
	Name() string        // returns the name of the master / bus
	NewConn(a Addr) Conn // returns new closed connection to the slave device
}

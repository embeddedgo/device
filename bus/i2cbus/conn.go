// Copyright 2024 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package i2cbus

import (
	"io"
	"time"
)

// A Conn represents an I2C connection. The connection may be considered as
// the way to reach some slave device using the underlying Master and address.
//
// At any time the connection may be in the closed or open state. The newly
// created connection is in the closed state. It may be opened and closed any
// number of times.
//
// The Read and Write methods with zero lenght slice as an argument do nothing.
//
// The first successful read, write or wait operation on the closed connection
// opens it. In the open state the connection has exclusive access to the
// underlying Master.
//
// The first read or write operation on the closed connection generates the
// Start condition. If the connection is in open state each read operation and
// the first write operation after the read operation generate the Repeated
// Start condition.
//
// The Wait method can be used on closed connection only. It waits for the ACK
// from slave device by periodically issuing the Start Condition. On success it
// returns without error leaving the connection in open state (see also
// ErrTimeout).
//
// Close method instructs the underlying Master to generate stop condition on
// the I2C bus and releases it making it available to the other connections.
//
// The connection cannot be used concurently by multiple goroutines without
// additional synchronization. If multiple goroutines need to communicate
// conncurently with the same slave device the best way is to use multiple
// connections.
type Conn interface {
	io.ReadWriteCloser
	io.ByteReader
	io.ByteWriter

	Wait(timeout time.Duration) error
	Master() Master
	Addr() Addr
}

// A Master represents an I2C bus master. The Master can be used concurently by
// multiple goroutines.
type Master interface {
	Name() string        // returns the name of the master / bus
	NewConn(a Addr) Conn // returns new closed connection to the slave device
}

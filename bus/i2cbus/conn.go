// Copyright 2024 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package i2cbus

import 	"io"

// A Conn represents an I2C connection. The connection may be considered as
// the way to reach a slave device using the underlying Master and the slave
// address.
//
// At any time the connection may be in the closed or open state. The newly
// created connection is in the closed state. It may be opened and closed any
// number of times.
//
// The first successful read or write operation on the closed connection opens
// it and generates the Start condition. In the open state the connection has
// exclusive access to the underlying Master.
//
// If the connection is in the open state each read operation may generate
// the Repeated Start condition. The first write after read and the first
// read after write always generates it. Consecutive writes never generate the
// Repeated Start condition. If the Repeated Start condition is not supported
// the Stop following by Start is used instead.
//
// Close method generates the Stop Condition on the I2C bus and releases the
// Master making it available to the other connections. Close is no-op on the
// closed connection.
//
// Errors returned by the read and write operations are generally asynchronous
// due to the internal buffering and latency of the Master or the underlying
// peripheral. That is, the lack of an error doesn't mean that the write
// operation was successful and the error returned by the read or write may be
// caused by the previous operation. You must check the error returned by Close
// to make sure the entire I2C write transaction was succesfull. This makes the
// simple `defer c.Close()` idiom erroneous in case of write transactions. In
// contrast the nil error returned by the read operation ensures the read was
// successfull. You can check only the error returned by the Close method and
// skip checking errors returned by other methods provided you don't use the
// read data until close. If the method other than Close return a non-nil error
// then all subsequent method calls are no-ops and return the same or diferent
// error (additionnal errors may be detected in the meantime). See also ErrACK.
//
// The Read or Write call on the closed connection with nil or zero-length slice
// as an argument return immediately without performing any operation. That is,
// you cannot use Write(nil) or Read(nil) to generate a Start Condition without
// the data transfer to check if the slave defice is ready for the next transfer
// (not all I2C master hardware support this).
//
// The connection cannot be used concurently by multiple goroutines without
// additional synchronization. If multiple goroutines need to communicate
// conncurently with the same slave device the best way to do it is to use
// multiple connections.
type Conn interface {
	io.ReadWriteCloser
	io.ByteReader
	io.ByteWriter

	Master() Master
	Addr() Addr
}

// A Master represents an I2C bus master. The Master can be used concurently by
// multiple goroutines.
type Master interface {
	Name() string        // returns the name of the master/bus
	NewConn(a Addr) Conn // returns new closed connection to the slave device
}

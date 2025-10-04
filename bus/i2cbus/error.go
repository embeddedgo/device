// Copyright 2024 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package i2cbus

import "errors"

// ErrACK can be used as an argumment to the errors.Is function to check if the
// error is caused by the unexpected NACK or ACK. The rare case when unexpected
// ACK may occur is for example the strat of the high-speed transmission.
var ErrACK = errors.New("unexpected NACK/ACK")

/**
 * SPDX-License-Identifier: Apache-2.0
 * Copyright 2020 vorteil.io Pty Ltd
 */

package tcpdump

import "time"

const (
	DefaultSnapshotLen     = 1024
	DefaultPromiscuousMode = false
	DefaultBPFFilter       = "tcp or udp"
	DefaultTimeout         = 10 * time.Second
	AnyDevice              = "any"
)

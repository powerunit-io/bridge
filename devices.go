// Copyright 2015 The PowerUnit Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package main ...
package main

import (
	"github.com/powerunit-io/platform/devices"
	"github.com/powerunit-io/platform/devices/gpio"
)

var (
	// Switch -
	Switch *gpio.Switch

	// Relay -
	Relay *gpio.Relay

	// AvailableDevices -
	AvailableDevices = map[string]devices.Device{
		"switch": Switch,
		"relay":  Relay,
	}
)

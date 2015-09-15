// Copyright 2015 The PowerUnit Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package main ...
package main

import "github.com/powerunit-io/platform/logging"

var logger *logging.Logger

func init() {
	logger = logging.New(LoggingConfig)
}

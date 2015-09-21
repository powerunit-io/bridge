// Copyright 2015 The PowerUnit Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package main ...
package main

import (
	"flag"
	"os"
	"runtime"

	"github.com/powerunit-io/platform/logging"
	"github.com/powerunit-io/platform/service"
	"github.com/powerunit-io/platform/utils"
)

var (
	serv   service.Service
	logger *logging.Logger
	err    error

	command = flag.String("command", "run", "Can be syncdb or run.")
)

func main() {
	flag.Parse()

	logger = logging.New(LoggingConfig)

	// PUB_PROCESS_COUNT == Getenv('PUB_PROCESS_COUNT')
	runtime.GOMAXPROCS(utils.GetProcessCount("PUB_PROCESS_COUNT"))

	if serv, err = NewBridgeService(); err != nil {
		logger.Fatal("Could not setup service due to (err: %s)", err)
		os.Exit(2)
	}

	switch *command {
	case "run":
		if err = serv.Start(); err != nil {
			logger.Error("Could not start Bridge due to (err: %s)", err)
			os.Exit(2)
		}
	case "syncdb":
		if err = serv.SyncDb(); err != nil {
			logger.Error("Could not start Bridge due to (err: %s)", err)
			os.Exit(2)
		}
	}

}

// Copyright 2015 The PowerUnit Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package main ...
package main

import (
	"flag"
	"os"
	"runtime"

	"github.com/powerunit-io/platform/service"
	"github.com/powerunit-io/platform/utils"
)

var (
	command = flag.String("command", "run", "Can be syncdb or run.")
)

func main() {
	flag.Parse()

	// PUB_PROCESS_COUNT == Getenv('PUB_PROCESS_COUNT')
	runtime.GOMAXPROCS(utils.GetProcessCount("PUB_PROCESS_COUNT"))

	var service service.Service
	var err error

	if service, err = NewBridgeService(); err != nil {
		logger.Fatal("Could not setup (service: %s) due to (err: %s)", service.Name(), err)
		os.Exit(2)
	}

	switch *command {
	case "run":
		if err = service.Start(); err != nil {
			logger.Error("Could not start Bridge due to (err: %s)", err)
			os.Exit(2)
		}
	case "syncdb":
		if err = service.SyncDb(); err != nil {
			logger.Error("Could not start Bridge due to (err: %s)", err)
			os.Exit(2)
		}
	default:

	}

}

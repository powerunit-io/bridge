// Copyright 2015 The PowerUnit Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package main ...
package main

import (
	"github.com/jinzhu/gorm"
	"github.com/powerunit-io/platform/connections/adapters/mysql"
)

// GetDb - Will return instance of database connection
func (bs *BridgeService) GetDb() (db gorm.DB, err error) {
	var service mysql.Adapter

	if service, err = bs.Connections.Get(DbAccessName); err != nil {
		return
	}

	db = service.Adapter().(*mysql.Connection).DB
	return
}

// AggregateRooms -
func (bs *BridgeService) AggregateRooms() (err error) {

	var db gorm.DB

	if db, err = bs.GetDb(); err != nil {
		return
	}

	_ = db
	/**
	// Get initial rooms ...
	bs.Info("About to start fetching initial room setup for (service: %s)", bs.Name())

	ar := rooms.All()


		timer := time.NewTimer(time.Second)

		for {
			select {
			case <-timer.C:
				bs.Info("About to start fetching rooms ...")
				availableRooms = rooms.All()
			}
		}


	bs.Info("Available (rooms: %s)", ar)
	  **/
	return
}

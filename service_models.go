// Copyright 2015 The PowerUnit Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package main ...
package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/powerunit-io/bridge/models"
	"github.com/powerunit-io/platform/connections/adapters/mysql"
)

var (
	db gorm.DB
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

// GetBuildings - Will load into service building actual building information.
// On any error, will return error.
func (bs *BridgeService) GetBuildings() error {
	bs.Info("About to start scanning for building information ...")

	var count int
	bs.Building = models.Building{}

	if db, err = bs.GetDb(); err != nil {
		return err
	}

	db.Model(models.Building{}).Where("status = ?", true).Count(&count)
	bs.Info("Discovered buildings (count: %v)", count)

	if &count == nil || count == 0 {
		return fmt.Errorf("We could not retrieve building information as one is not yet defined.")
	}

	db.Where("status = ?", true).First(&bs.Building)

	bs.Info("Discovered buidling (data: %v)", bs.Building)

	return nil
}

// GetRooms -
func (bs *BridgeService) GetRooms() error {
	bs.Info("About to start scanning for rooms ...")

	if db, err = bs.GetDb(); err != nil {
		return err
	}

	db.Where("status = ?", true).First(&bs.Rooms)
	bs.Info("Discovered rooms (data: %v)", bs.Rooms)

	return nil
}

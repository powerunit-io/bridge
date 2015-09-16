// Copyright 2015 The PowerUnit Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package main ...
package main

import (
	"github.com/jinzhu/gorm"
	"github.com/powerunit-io/bridge/models"
)

// SyncDb -
func (bs *BridgeService) SyncDb() (err error) {
	bs.Info("About to start synchronizing models ...")

	var db gorm.DB

	if db, err = bs.GetDb(); err != nil {
		return
	}

	// Enable Logger
	db.LogMode(true)

	db.Set("gorm:table_options", "ENGINE=InnoDB")

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Building{})
	db.AutoMigrate(&models.Floor{})
	db.AutoMigrate(&models.Room{})

	db.Model(&models.Floor{}).AddForeignKey("floor_building_id", "building(id)", "CASCADE", "CASCADE")
	db.Model(&models.Room{}).AddForeignKey("building_id", "building(id)", "CASCADE", "CASCADE")
	db.Model(&models.Room{}).AddForeignKey("floor_id", "floor(id)", "CASCADE", "CASCADE")

	db.Model(&models.Building{}).AddUniqueIndex("idx_building_name", "name")
	db.Model(&models.Floor{}).AddUniqueIndex("idx_floor_name", "name")

	db.Model(&models.Room{}).AddUniqueIndex("idx_room_name", "name")
	db.Model(&models.Room{}).AddUniqueIndex("idx_room_uuid", "uuid")

	return nil
}

// ValidateModels -
func (bs *BridgeService) ValidateModels() error {
	bs.Info("About to start validating models ...")

	return nil
}

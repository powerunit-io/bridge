// Copyright 2015 The PowerUnit Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package models ...
package models

import "github.com/powerunit-io/platform/models"

// Room -
type Room struct {
	models.BaseModel

	Building   Building
	BuildingID int64
	Floor      Floor
	FloorID    int64

	UUID string `sql:"type:varchar(6);not null;"`

	Name        string `sql:"not null;"`
	Description string

	Status bool `sql:"DEFAULT:1"`
	Online bool `sql:"DEFAULT:0"`
}

// TableName -
func (r Room) TableName() string {
	return "room"
}

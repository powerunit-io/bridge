// Copyright 2015 The PowerUnit Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package models ...
package models

import (
	"time"

	"github.com/powerunit-io/platform/models"
)

// User -
type User struct {
	models.BaseModel

	FirstName string    `sql:"not null;"`
	LastName  string    `sql:"not null;"`
	Birthday  time.Time `sql:""`
	Username  string    `sql:"not null;"`
	Password  string    `sql:"not null;"`
	Status    bool      `sql:"DEFAULT:1"`
	Admin     bool      `sql:"DEFAULT:0"`
}

// TableName -
func (u User) TableName() string {
	return "user"
}

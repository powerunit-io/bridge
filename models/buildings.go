// Copyright 2015 The PowerUnit Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package models ...
package models

import "github.com/powerunit-io/platform/models"

// Building -
type Building struct {
	models.BaseModel

	Name        string `sql:"not null;unique"`
	Description string
	Status      bool `sql:"DEFAULT:1"`
}

// TableName -
func (b Building) TableName() string {
	return "building"
}

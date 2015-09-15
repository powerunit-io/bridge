// Copyright 2015 The PowerUnit Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package main ...
package main

import "github.com/powerunit-io/platform/models"

// Rooms -
type Rooms struct {
	models.BaseModel
}

// NewRooms -
func NewRooms() models.Model {
	return models.Model(&Rooms{})
}

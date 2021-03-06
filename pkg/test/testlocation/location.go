// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package testlocation contains utility functions and helpers for testing with locations.
package testlocation

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/google/simhospital/pkg/location"
)

// roomMgrTmpl is the YML template for location.RoomManager.
const roomMgrTmpl string = `
%s:
  poc: %s
  facility: Simulated Hospital
  building: Building-1
  floor: 7
  room: Room-1
`

// ManagerWithAAndE is a location manager with an ED location.
var ManagerWithAAndE = &location.Manager{
	RoomManagers: map[string]*location.RoomManager{
		"ED": {
			Poc:      "ED",
			Facility: "Simulated Hospital",
			Building: "Building-1",
			Floor:    "7",
			Room:     "Room-1",
			Type:     "ED",
		},
	},
}

// NewLocationManager creates a new location manager with the given locations.
func NewLocationManager(locations ...string) (*location.Manager, error) {
	var allLocations string
	for _, l := range locations {
		allLocations += fmt.Sprintf(roomMgrTmpl, l, l)
	}
	loc := []byte(allLocations)
	tmp, err := ioutil.TempFile("", "location")
	if err != nil {
		return nil, err
	}
	defer os.Remove(tmp.Name())

	if _, err = tmp.Write(loc); err != nil {
		return nil, err
	}
	return location.NewManager(tmp.Name())
}

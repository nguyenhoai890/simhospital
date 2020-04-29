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

package pathway

import (
	"testing"
	"time"

	"github.com/google/simhospital/pkg/config"
	"github.com/google/simhospital/pkg/doctor"
	"github.com/google/simhospital/pkg/location"
	"github.com/google/simhospital/pkg/orderprofile"
	"github.com/google/simhospital/pkg/test"
	"github.com/google/simhospital/pkg/test/testclock"
)

func TestValidateProdPathways(t *testing.T) {
	hl7Config, err := config.LoadHL7Config(test.PublicMessageConfig)
	if err != nil {
		t.Fatalf("LoadHL7Config(%s) failed with %v", test.PublicMessageConfig, err)
	}
	d, err := doctor.LoadDoctors(test.PublicDoctorsConfig)
	if err != nil {
		t.Fatalf("LoadDoctors(%s) failed with %v", test.PublicDoctorsConfig, err)
	}
	op, err := orderprofile.Load(test.PublicOrderProfilesConfig, hl7Config)
	if err != nil {
		t.Fatalf("orderprofile.Load(%s, %+v) failed with %v", test.PublicOrderProfilesConfig, hl7Config, err)
	}
	lm, err := location.NewManager(test.PublicLocationsConfig)
	if err != nil {
		t.Fatalf("location.NewManager(%s) failed with %v", test.PublicLocationsConfig, err)
	}
	p := &Parser{Clock: testclock.New(time.Now()), OrderProfiles: op, Doctors: d, LocationManager: lm}
	pathways, err := p.ParsePathways(test.ProdPathwaysDir, nil, nil)
	if err != nil {
		t.Fatalf("ParsePathways(%s, %v, %v) failed with %v", test.ProdPathwaysDir, nil, nil, err)
	}

	if len(pathways) == 0 {
		t.Fatalf("ParsePathways(%s, %v, %v) got empty pathways, want non empty", test.ProdPathwaysDir, nil, nil)
	}
}
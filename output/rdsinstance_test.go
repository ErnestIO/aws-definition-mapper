/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package output

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRDSInstanceHasChanged(t *testing.T) {
	Convey("Given a rds cluster", t, func() {
		r := RDSInstance{
			Name:              "test",
			Size:              "db.large",
			Engine:            "mysql",
			EngineVersion:     "5.5",
			Port:              3306,
			Public:            false,
			HotStandby:        false,
			PromotionTier:     1,
			StorageType:       "io1",
			StorageSize:       100,
			StorageIops:       5000,
			AvailabilityZone:  "eu-west-1a",
			DatabaseName:      "test",
			DatabaseUsername:  "root",
			DatabasePassword:  "test",
			AutoUpgrade:       false,
			BackupRetention:   5,
			BackupWindow:      "Mon:10:00-Mon:11:00",
			MaintenanceWindow: "Tue:10:00-Tue:11:00",
			FinalSnapshot:     false,
			ReplicationSource: "test",
			License:           "bring-your-own",
			Timezone:          "GMT",
			SecurityGroups: []string{
				"sg-1",
			},
			Networks: []string{
				"nw-1",
			},
		}

		Convey("When I compare it to an changed rds cluster", func() {
			or := RDSInstance{
				Name:              "test",
				Size:              "db.large",
				Engine:            "mysql",
				EngineVersion:     "5.5",
				Port:              4000,
				Public:            false,
				HotStandby:        false,
				PromotionTier:     1,
				StorageType:       "io1",
				StorageSize:       100,
				StorageIops:       5000,
				AvailabilityZone:  "eu-west-1a",
				DatabaseName:      "test",
				DatabaseUsername:  "root",
				DatabasePassword:  "test",
				AutoUpgrade:       false,
				BackupRetention:   5,
				BackupWindow:      "Mon:10:00-Mon:11:00",
				MaintenanceWindow: "Tue:10:00-Tue:11:00",
				FinalSnapshot:     false,
				ReplicationSource: "test",
				License:           "bring-your-own",
				Timezone:          "GMT",
				SecurityGroups: []string{
					"sg-1",
				},
				Networks: []string{
					"nw-1",
				},
			}
			change := r.HasChanged(&or)
			Convey("Then it should return true", func() {
				So(change, ShouldBeTrue)
			})
		})

		Convey("When I compare it to an identical rds cluster", func() {
			or := RDSInstance{
				Name:              "test",
				Size:              "db.large",
				Engine:            "mysql",
				EngineVersion:     "5.5",
				Port:              3306,
				Public:            false,
				HotStandby:        false,
				PromotionTier:     1,
				StorageType:       "io1",
				StorageSize:       100,
				StorageIops:       5000,
				AvailabilityZone:  "eu-west-1a",
				DatabaseName:      "test",
				DatabaseUsername:  "root",
				DatabasePassword:  "test",
				AutoUpgrade:       false,
				BackupRetention:   5,
				BackupWindow:      "Mon:10:00-Mon:11:00",
				MaintenanceWindow: "Tue:10:00-Tue:11:00",
				FinalSnapshot:     false,
				ReplicationSource: "test",
				License:           "bring-your-own",
				Timezone:          "GMT",
				SecurityGroups: []string{
					"sg-1",
				},
				Networks: []string{
					"nw-1",
				},
			}
			change := r.HasChanged(&or)
			Convey("Then it should return false", func() {
				So(change, ShouldBeFalse)
			})
		})
	})
}

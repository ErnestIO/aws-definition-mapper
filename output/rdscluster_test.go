/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package output

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRDSClusterHasChanged(t *testing.T) {
	Convey("Given a rds cluster", t, func() {
		r := RDSCluster{
			Name:              "test",
			Engine:            "mysql",
			EngineVersion:     "5.5",
			Port:              3306,
			DatabaseName:      "test",
			DatabaseUsername:  "root",
			DatabasePassword:  "test",
			BackupRetention:   5,
			BackupWindow:      "Mon:10:00-Mon:11:00",
			MaintenanceWindow: "Tue:10:00-Tue:11:00",
			FinalSnapshot:     false,
			ReplicationSource: "test",
			SecurityGroups: []string{
				"sg-1",
			},
			Networks: []string{
				"nw-1",
			},
		}

		Convey("When I compare it to an changed rds cluster", func() {
			or := RDSCluster{
				Name:              "test",
				Engine:            "mysql",
				EngineVersion:     "5.5",
				Port:              4000,
				DatabaseName:      "test",
				DatabaseUsername:  "root",
				DatabasePassword:  "test",
				BackupRetention:   5,
				BackupWindow:      "Mon:10:00-Mon:11:00",
				MaintenanceWindow: "Tue:10:00-Tue:11:00",
				FinalSnapshot:     false,
				ReplicationSource: "test",
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
			or := RDSCluster{
				Name:              "test",
				Engine:            "mysql",
				EngineVersion:     "5.5",
				Port:              3306,
				DatabaseName:      "test",
				DatabaseUsername:  "root",
				DatabasePassword:  "test",
				BackupRetention:   5,
				BackupWindow:      "Mon:10:00-Mon:11:00",
				MaintenanceWindow: "Tue:10:00-Tue:11:00",
				FinalSnapshot:     false,
				ReplicationSource: "test",
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

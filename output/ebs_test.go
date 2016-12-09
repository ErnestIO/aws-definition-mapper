/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package output

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestEBSHasChanged(t *testing.T) {
	Convey("Given a network", t, func() {
		v := EBSVolume{
			Name:             "test",
			AvailabilityZone: "eu-west-1",
			VolumeType:       "gp2",
			Size:             int64p(100),
		}

		// EBS Volumes are immutable
		Convey("When I compare it to an changed network", func() {
			ov := EBSVolume{
				Name:             "test",
				AvailabilityZone: "eu-west-1",
				VolumeType:       "io1",
				Size:             int64p(500),
			}
			change := v.HasChanged(&ov)
			Convey("Then it should return false", func() {
				So(change, ShouldBeFalse)
			})
		})

		Convey("When I compare it to an identical network", func() {
			ov := EBSVolume{
				Name:             "test",
				AvailabilityZone: "eu-west-1",
				VolumeType:       "gp2",
				Size:             int64p(100),
			}
			change := v.HasChanged(&ov)
			Convey("Then it should return false", func() {
				So(change, ShouldBeFalse)
			})
		})
	})
}

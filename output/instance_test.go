/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package output

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestInstanceHasChanged(t *testing.T) {
	Convey("Given a instance", t, func() {
		i := Instance{
			Name:    "test",
			Type:    "m2.small",
			Image:   "ami-000000",
			Network: "network",
		}

		Convey("When I compare it to an changed instance", func() {
			oi := Instance{
				Name:    "test",
				Type:    "m2.large",
				Image:   "ami-000000",
				Network: "network",
			}
			change := i.HasChanged(&oi)
			Convey("Then it should return true", func() {
				So(change, ShouldBeTrue)
			})
		})

		Convey("When I compare it to an identical instance", func() {
			oi := Instance{
				Name:    "test",
				Type:    "m2.small",
				Image:   "ami-000000",
				Network: "network",
			}
			change := i.HasChanged(&oi)
			Convey("Then it should return false", func() {
				So(change, ShouldBeFalse)
			})
		})
	})
}

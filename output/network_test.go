/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package output

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNetworkHasChanged(t *testing.T) {
	Convey("Given a network", t, func() {
		n := Network{
			Name:   "test",
			Subnet: "10.0.0.0/24",
		}

		// Networks are immutable
		Convey("When I compare it to an changed network", func() {
			on := Network{
				Name:   "test",
				Subnet: "10.10.0.0/24",
			}
			change := n.HasChanged(&on)
			Convey("Then it should return false", func() {
				So(change, ShouldBeFalse)
			})
		})

		Convey("When I compare it to an identical network", func() {
			on := Network{
				Name:   "test",
				Subnet: "10.0.0.0/24",
			}
			change := n.HasChanged(&on)
			Convey("Then it should return false", func() {
				So(change, ShouldBeFalse)
			})
		})
	})
}

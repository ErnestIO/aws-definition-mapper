/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package mapper

import (
	"testing"

	"github.com/r3labs/aws-definition-mapper/definition"
	. "github.com/smartystreets/goconvey/convey"
)

func TestNetworksMapping(t *testing.T) {
	Convey("Given a valid input definition", t, func() {
		d := definition.Definition{
			Name:       "test",
			Datacenter: "datacenter",
		}

		d.Networks = append(d.Networks, definition.Network{
			Name:   "bar",
			Subnet: "10.0.0.0/24",
		})

		Convey("When I try to map a network", func() {
			Convey("And the input specifies bootstrap as not salt", func() {
				n := MapNetworks(d)
				Convey("Then only input networks should be mapped", func() {
					So(len(n), ShouldEqual, 1)
					So(n[0].Name, ShouldEqual, "datacenter-test-bar")
					So(n[0].Subnet, ShouldEqual, "10.0.0.0/24")
				})
			})
		})

	})
}

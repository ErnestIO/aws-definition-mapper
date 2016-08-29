/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package mapper

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/ErnestIO/aws-definition-mapper/definition"
)

func TestMapNats(t *testing.T) {

	Convey("Given a valid input definition", t, func() {
		d := definition.Definition{
			Name:       "service",
			Datacenter: "datacenter",
		}

		d.Networks = append(d.Networks, definition.Network{
			Name:   "test",
			Subnet: "10.0.0.0/24",
		})

		Convey("When i try to map nats", func() {

			n := MapNats(d)
			Convey("Then it should map salt and input firewall rules", func() {
				So(len(n), ShouldEqual, 1)
				So(n[0].Name, ShouldEqual, "datacenter-service-test")
				So(n[0].Network, ShouldEqual, "datacenter-service-test")
			})

		})
	})

}

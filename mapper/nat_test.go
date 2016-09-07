/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package mapper

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/ernestio/aws-definition-mapper/definition"
)

func TestMapNats(t *testing.T) {

	Convey("Given a valid input definition", t, func() {
		d := definition.Definition{
			Name:       "service",
			Datacenter: "datacenter",
		}

		d.Networks = append(d.Networks, definition.Network{
			Name:   "public",
			Subnet: "10.0.0.0/24",
			Public: true,
		})

		d.Networks = append(d.Networks, definition.Network{
			Name:       "routed",
			Subnet:     "10.0.1.0/24",
			NatGateway: "test",
		})

		d.NatGateways = append(d.NatGateways, definition.NatGateway{
			Name:          "test",
			PublicNetwork: "public",
		})

		Convey("When i try to map nats", func() {

			n := MapNats(d)
			Convey("Then it should map salt and input firewall rules", func() {
				So(len(n), ShouldEqual, 1)
				So(n[0].Name, ShouldEqual, "datacenter-service-test")
				So(n[0].PublicNetwork, ShouldEqual, "datacenter-service-public")
				So(len(n[0].RoutedNetworks), ShouldEqual, 1)
				So(n[0].RoutedNetworks[0], ShouldEqual, "datacenter-service-routed")
			})

		})
	})

}

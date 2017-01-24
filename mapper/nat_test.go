/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package mapper

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/ernestio/aws-definition-mapper/definition"
	"github.com/ernestio/aws-definition-mapper/output"
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
				So(n[0].Name, ShouldEqual, "test")
				So(n[0].PublicNetwork, ShouldEqual, "public")
				So(len(n[0].RoutedNetworks), ShouldEqual, 1)
				So(n[0].RoutedNetworks[0], ShouldEqual, "routed")
				So(n[0].Tags["Name"], ShouldEqual, "test")
				So(n[0].Tags["ernest.service"], ShouldEqual, "service")
			})

		})
	})

	Convey("Given a valid output message", t, func() {
		m := output.FSMMessage{
			Service: "service",
		}

		n := output.Network{
			NetworkAWSID:     "s-0000000",
			Name:             "web",
			Subnet:           "10.10.0.0/24",
			IsPublic:         true,
			AvailabilityZone: "eu-west-1",
		}

		ng := output.Nat{
			NatGatewayAWSID:    "nat-0000000",
			Name:               "web-nat",
			PublicNetworkAWSID: "s-0000000",
		}

		m.Nats.Items = append(m.Nats.Items, ng)
		m.Networks.Items = append(m.Networks.Items, n)

		Convey("When i try to map nat gateways", func() {

			nts := MapDefinitionNats(&m)
			Convey("Then it should return a correctly formed set of input nat gateways", func() {
				So(len(nts), ShouldEqual, 1)
				nt := nts[0]
				So(nt.Name, ShouldEqual, "web-nat")
				So(nt.PublicNetwork, ShouldEqual, "web")
			})

		})
	})

}

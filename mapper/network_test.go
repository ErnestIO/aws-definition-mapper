/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package mapper

import (
	"testing"

	"github.com/ernestio/aws-definition-mapper/definition"
	"github.com/ernestio/aws-definition-mapper/output"
	. "github.com/smartystreets/goconvey/convey"
)

func TestNetworksMapping(t *testing.T) {
	Convey("Given a valid input definition", t, func() {
		d := definition.Definition{
			Name:       "service",
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
					So(n[0].Name, ShouldEqual, "datacenter-service-bar")
					So(n[0].Subnet, ShouldEqual, "10.0.0.0/24")
					So(n[0].Tags["Name"], ShouldEqual, "datacenter-service-bar")
					So(n[0].Tags["ernest.service"], ShouldEqual, "service")
				})
			})
		})

	})

	Convey("Given a valid output message", t, func() {
		m := output.FSMMessage{
			ServiceName: "service",
		}

		m.Datacenters.Items = append(m.Datacenters.Items, output.Datacenter{
			Name: "datacenter",
		})

		tags := make(map[string]string)
		tags["ernest.nat_gateway"] = "web-nat"

		n := output.Network{
			NetworkAWSID:     "s-0000000",
			Name:             "datacenter-service-web",
			Subnet:           "10.10.0.0/24",
			IsPublic:         true,
			AvailabilityZone: "eu-west-1",
			Tags:             tags,
		}

		ng := output.Nat{
			NatGatewayAWSID: "nat-0000000",
			Name:            "datacenter-service-web-nat",
			RoutedNetworkAWSIDs: []string{
				"s-0000000",
			},
		}

		m.Nats.Items = append(m.Nats.Items, ng)
		m.Networks.Items = append(m.Networks.Items, n)

		Convey("When i try to map networks", func() {

			nws := MapDefinitionNetworks(&m)
			Convey("Then it should return a correctly formed set of input networks", func() {
				So(len(nws), ShouldEqual, 1)
				nw := nws[0]
				So(nw.Name, ShouldEqual, "web")
				So(nw.Subnet, ShouldEqual, "10.10.0.0/24")
				So(nw.Public, ShouldEqual, true)
				So(nw.AvailabilityZone, ShouldEqual, "eu-west-1")
				So(nw.NatGateway, ShouldEqual, "web-nat")
			})

		})
	})

}

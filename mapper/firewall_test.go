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

func TestMapSecurityGroups(t *testing.T) {

	Convey("Given a valid input definition", t, func() {
		d := definition.Definition{
			Name:       "service",
			Datacenter: "datacenter",
		}

		d.Networks = append(d.Networks, definition.Network{
			Name:   "bar",
			Subnet: "10.0.0.0/24",
		})

		sg := definition.SecurityGroup{
			Name: "test",
		}

		sg.Ingress = append(sg.Ingress, definition.SecurityGroupRule{
			IP:       "10.10.10.11",
			ToPort:   "80",
			FromPort: "80",
			Protocol: "tcp",
		})

		d.SecurityGroups = append(d.SecurityGroups, sg)

		Convey("When i try to map firewalls", func() {

			f := MapSecurityGroups(d)
			Convey("Then it should map salt and input firewall rules", func() {
				So(len(f), ShouldEqual, 1)
				So(f[0].Name, ShouldEqual, "test")
				So(len(f[0].Rules.Ingress), ShouldEqual, 1)
				So(f[0].Rules.Ingress[0].IP, ShouldEqual, "10.10.10.11")
				So(f[0].Rules.Ingress[0].To, ShouldEqual, 80)
				So(f[0].Rules.Ingress[0].From, ShouldEqual, 80)
				So(f[0].Rules.Ingress[0].Protocol, ShouldEqual, "tcp")
				So(f[0].Tags["Name"], ShouldEqual, "test")
				So(f[0].Tags["ernest.service"], ShouldEqual, "service")
			})

		})
	})

	Convey("Given a valid output message", t, func() {
		m := output.FSMMessage{
			Service: "service",
		}

		f := output.Firewall{
			SecurityGroupAWSID: "sg-0000000",
			Name:               "web-sg",
		}

		f.Rules.Egress = append(f.Rules.Egress, output.FirewallRule{
			IP:       "10.10.10.11",
			To:       80,
			From:     80,
			Protocol: "tcp",
		})

		f.Rules.Ingress = append(f.Rules.Ingress, output.FirewallRule{
			IP:       "10.10.10.11",
			To:       80,
			From:     80,
			Protocol: "-1",
		})

		m.Firewalls.Items = append(m.Firewalls.Items, f)

		Convey("When i try to map firewalls", func() {

			s := MapDefinitionSecurityGroups(&m)
			Convey("Then it should return a correctly formed set of input security groups", func() {
				So(len(s), ShouldEqual, 1)
				sg := s[0]
				So(sg.Name, ShouldEqual, "web-sg")
				So(len(sg.Egress), ShouldEqual, 1)
				So(sg.Egress[0].IP, ShouldEqual, "10.10.10.11")
				So(sg.Egress[0].FromPort, ShouldEqual, "80")
				So(sg.Egress[0].ToPort, ShouldEqual, "80")
				So(sg.Egress[0].Protocol, ShouldEqual, "tcp")
				So(len(sg.Ingress), ShouldEqual, 1)
				So(sg.Ingress[0].IP, ShouldEqual, "10.10.10.11")
				So(sg.Ingress[0].FromPort, ShouldEqual, "80")
				So(sg.Ingress[0].ToPort, ShouldEqual, "80")
				So(sg.Ingress[0].Protocol, ShouldEqual, "any")
			})

		})
	})

}

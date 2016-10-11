/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package mapper

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/ernestio/aws-definition-mapper/definition"
)

func TestMapELBs(t *testing.T) {

	Convey("Given a valid input definition", t, func() {
		d := definition.Definition{
			Name:       "service",
			Datacenter: "datacenter",
		}

		d.Instances = append(d.Instances, definition.Instance{
			Name:    "web",
			Network: "web-nw",
			Count:   2,
		})

		e := definition.ELB{
			Name:    "test",
			Private: true,
			Instances: []string{
				"web",
			},
			SecurityGroups: []string{
				"web-sg",
			},
		}

		e.Ports = append(e.Ports, definition.ELBPort{
			FromPort: 1,
			ToPort:   2,
			Protocol: "http",
			SSLCert:  "cert",
		})

		d.ELBs = append(d.ELBs, e)

		Convey("When i try to map elbs", func() {

			e := MapELBs(d)
			Convey("Then it should map salt and input elb rules", func() {
				So(len(e), ShouldEqual, 1)
				So(e[0].Name, ShouldEqual, "datacenter-service-test")
				So(len(e[0].NetworkAWSIDs), ShouldEqual, 1)
				So(e[0].NetworkAWSIDs[0], ShouldEqual, `$(networks.items.#[name="datacenter-service-web-nw"].network_aws_id)`)
				So(len(e[0].InstanceAWSIDs), ShouldEqual, 2)
				So(e[0].InstanceAWSIDs[0], ShouldEqual, `$(instances.items.#[name="datacenter-service-web-1"].instance_aws_id)`)
				So(e[0].InstanceAWSIDs[1], ShouldEqual, `$(instances.items.#[name="datacenter-service-web-2"].instance_aws_id)`)
				So(len(e[0].SecurityGroupAWSIDs), ShouldEqual, 1)
				So(e[0].SecurityGroupAWSIDs[0], ShouldEqual, `$(firewalls.items.#[name="datacenter-service-web-sg"].security_group_aws_id)`)
				So(len(e[0].Ports), ShouldEqual, 1)
				So(e[0].Ports[0].FromPort, ShouldEqual, 1)
				So(e[0].Ports[0].ToPort, ShouldEqual, 2)
				So(e[0].Ports[0].Protocol, ShouldEqual, "http")
				So(e[0].Ports[0].SSLCert, ShouldEqual, "cert")
			})

		})
	})

}

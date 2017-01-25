/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package mapper

import (
	"testing"

	"github.com/ernestio/aws-definition-mapper/output"
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
			Subnets: []string{
				"web-nw",
			},
			Instances: []string{
				"web",
			},
			SecurityGroups: []string{
				"web-sg",
			},
		}

		e.Listeners = append(e.Listeners, definition.ELBListener{
			FromPort: 1,
			ToPort:   2,
			Protocol: "http",
			SSLCert:  "cert",
		})

		d.ELBs = append(d.ELBs, e)

		Convey("When i try to map elbs", func() {

			e := MapELBs(d)
			Convey("Then it should map input elb rules", func() {
				So(len(e), ShouldEqual, 1)
				So(e[0].Name, ShouldEqual, "datacenter-service-test")
				So(len(e[0].NetworkAWSIDs), ShouldEqual, 1)
				So(e[0].NetworkAWSIDs[0], ShouldEqual, `$(networks.items.#[name="datacenter-service-web-nw"].network_aws_id)`)
				So(len(e[0].InstanceAWSIDs), ShouldEqual, 2)
				So(e[0].InstanceAWSIDs[0], ShouldEqual, `$(instances.items.#[name="datacenter-service-web-1"].instance_aws_id)`)
				So(e[0].InstanceAWSIDs[1], ShouldEqual, `$(instances.items.#[name="datacenter-service-web-2"].instance_aws_id)`)
				So(len(e[0].SecurityGroupAWSIDs), ShouldEqual, 1)
				So(e[0].SecurityGroupAWSIDs[0], ShouldEqual, `$(firewalls.items.#[name="datacenter-service-web-sg"].security_group_aws_id)`)
				So(len(e[0].Listeners), ShouldEqual, 1)
				So(e[0].Listeners[0].FromPort, ShouldEqual, 1)
				So(e[0].Listeners[0].ToPort, ShouldEqual, 2)
				So(e[0].Listeners[0].Protocol, ShouldEqual, "HTTP")
				So(e[0].Listeners[0].SSLCert, ShouldEqual, "cert")
				So(e[0].Tags["ernest.service"], ShouldEqual, "service")
			})

		})
	})

	Convey("Given a valid output message", t, func() {
		m := output.FSMMessage{
			Service: "service",
		}

		m.Networks.Items = append(m.Networks.Items, output.Network{
			NetworkAWSID: "n-0000000",
			Name:         "web",
			Subnet:       "10.64.0.0/24",
		})

		m.Instances.Items = append(m.Instances.Items, output.Instance{
			InstanceAWSID: "i-0000000",
			Name:          "web-1",
		})

		tags := make(map[string]string)
		tags["ernest.instance_group"] = "web"

		m.Instances.Items[0].Tags = tags

		m.Firewalls.Items = append(m.Firewalls.Items, output.Firewall{
			SecurityGroupAWSID: "sg-0000000",
			Name:               "web-sg",
		})

		e := output.ELB{
			Name:      "test",
			IsPrivate: true,
			NetworkAWSIDs: []string{
				"n-0000000",
			},
			InstanceAWSIDs: []string{
				"i-0000000",
			},
			SecurityGroupAWSIDs: []string{
				"sg-0000000",
			},
		}

		e.Listeners = append(e.Listeners, output.ELBListener{
			FromPort: 1,
			ToPort:   2,
			Protocol: "http",
			SSLCert:  "cert",
		})

		m.ELBs.Items = append(m.ELBs.Items, e)

		Convey("When i try to map elbs", func() {

			e := MapDefinitionELBs(&m)
			Convey("Then it should return a correctly formed set of input elb's", func() {
				So(len(e), ShouldEqual, 1)
				elb := e[0]
				So(elb.Name, ShouldEqual, "test")
				So(elb.Private, ShouldBeTrue)
				So(len(elb.Instances), ShouldEqual, 1)
				So(elb.Instances[0], ShouldEqual, "web")
				So(len(elb.Subnets), ShouldEqual, 1)
				So(elb.Subnets[0], ShouldEqual, "web")
				So(len(elb.SecurityGroups), ShouldEqual, 1)
				So(elb.SecurityGroups[0], ShouldEqual, "web-sg")
			})

		})
	})

}

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

func TestInstancesMapping(t *testing.T) {
	Convey("Given a valid input definition", t, func() {
		d := definition.Definition{
			Name:       "service",
			Datacenter: "datacenter",
		}

		d.Networks = append(d.Networks, definition.Network{
			Name:   "bar",
			Subnet: "10.0.0.0/24",
		})

		d.Instances = append(d.Instances, definition.Instance{
			Name:    "foo",
			Type:    "m1.small",
			Image:   "ami-000000",
			Count:   1,
			Network: "bar",
			Volumes: []definition.InstanceVolume{
				definition.InstanceVolume{
					Volume: "test",
					Device: "/dev/sdx",
				},
			},
		})

		Convey("When i try to map instances", func() {
			Convey("And the instance count is set to 1", func() {
				i := MapInstances(d)

				Convey("Then an extra instance should be mapped", func() {
					So(len(i), ShouldEqual, 1)
					So(i[0].Name, ShouldEqual, "datacenter-service-foo-1")
					So(i[0].Type, ShouldEqual, "m1.small")
					So(i[0].Image, ShouldEqual, "ami-000000")
					So(i[0].Network, ShouldEqual, "datacenter-service-bar")
					So(len(i[0].Volumes), ShouldEqual, 1)
					So(i[0].Volumes[0].Volume, ShouldEqual, "datacenter-service-test-1")
					So(i[0].Volumes[0].VolumeAWSID, ShouldEqual, `$(ebs_volumes.items.#[name="datacenter-service-test-1"].volume_aws_id)`)
					So(i[0].Volumes[0].Device, ShouldEqual, "/dev/sdx")
					So(i[0].Tags["Name"], ShouldEqual, "foo-1")
					So(i[0].Tags["ernest.service"], ShouldEqual, "service")
					So(i[0].Tags["ernest.instance_group"], ShouldEqual, "foo")
				})
			})

			Convey("And the instance count is set to 2", func() {
				d.Instances[0].Count = 2
				i := MapInstances(d)
				Convey("Then defined instances should be mapped", func() {
					So(len(i), ShouldEqual, 2)
					So(i[0].Name, ShouldEqual, "datacenter-service-foo-1")
					So(i[0].Type, ShouldEqual, "m1.small")
					So(i[0].Image, ShouldEqual, "ami-000000")
					So(i[0].Network, ShouldEqual, "datacenter-service-bar")
					So(len(i[0].Volumes), ShouldEqual, 1)
					So(i[0].Volumes[0].Volume, ShouldEqual, "datacenter-service-test-1")
					So(i[0].Volumes[0].VolumeAWSID, ShouldEqual, `$(ebs_volumes.items.#[name="datacenter-service-test-1"].volume_aws_id)`)
					So(i[0].Volumes[0].Device, ShouldEqual, "/dev/sdx")
					So(i[0].Tags["Name"], ShouldEqual, "foo-1")
					So(i[0].Tags["ernest.service"], ShouldEqual, "service")
					So(i[0].Tags["ernest.instance_group"], ShouldEqual, "foo")
					So(i[1].Name, ShouldEqual, "datacenter-service-foo-2")
					So(i[1].Type, ShouldEqual, "m1.small")
					So(i[1].Image, ShouldEqual, "ami-000000")
					So(i[1].Network, ShouldEqual, "datacenter-service-bar")
					So(len(i[1].Volumes), ShouldEqual, 1)
					So(i[1].Volumes[0].Volume, ShouldEqual, "datacenter-service-test-2")
					So(i[1].Volumes[0].VolumeAWSID, ShouldEqual, `$(ebs_volumes.items.#[name="datacenter-service-test-2"].volume_aws_id)`)
					So(i[1].Volumes[0].Device, ShouldEqual, "/dev/sdx")
					So(i[1].Tags["Name"], ShouldEqual, "foo-2")
					So(i[1].Tags["ernest.service"], ShouldEqual, "service")
					So(i[1].Tags["ernest.instance_group"], ShouldEqual, "foo")
				})
			})
		})
	})

	Convey("Given a valid output message", t, func() {
		m := output.FSMMessage{
			Service: "service",
		}

		m.Firewalls.Items = append(m.Firewalls.Items, output.Firewall{
			SecurityGroupAWSID: "sg-0000000",
			Name:               "web-sg",
		})

		m.Networks.Items = append(m.Networks.Items, output.Network{
			NetworkAWSID:     "s-0000000",
			Name:             "web",
			Subnet:           "10.10.0.0/24",
			IsPublic:         true,
			AvailabilityZone: "eu-west-1",
		})

		v := output.EBSVolume{
			VolumeAWSID: "vol-0000000",
			Name:        "web-vol-1",
		}

		vtags := make(map[string]string)
		vtags["ernest.volume_group"] = "web-vol"

		v.Tags = vtags

		m.EBSVolumes.Items = append(m.EBSVolumes.Items, v)

		i := output.Instance{
			Name:         "web-1",
			Type:         "m1.small",
			Image:        "ami-0000000",
			NetworkAWSID: "s-0000000",
			Volumes: []output.InstanceVolume{
				output.InstanceVolume{
					VolumeAWSID: "vol-0000000",
					Device:      "/dev/sdx",
				},
			},
			KeyPair: "test",
			SecurityGroupAWSIDs: []string{
				"sg-0000000",
			},
		}

		tags := make(map[string]string)
		tags["ernest.instance_group"] = "web"

		i.Tags = tags

		m.Instances.Items = append(m.Instances.Items, i)

		Convey("When i try to map instances", func() {

			ins := MapDefinitionInstances(&m)
			Convey("Then it should return a correctly formed set of input instances", func() {
				So(len(ins), ShouldEqual, 1)
				in := ins[0]
				So(in.Name, ShouldEqual, "web")
				So(in.Type, ShouldEqual, "m1.small")
				So(in.Image, ShouldEqual, "ami-0000000")
				So(in.Network, ShouldEqual, "web")
				So(len(in.Volumes), ShouldEqual, 1)
				So(in.Volumes[0].Volume, ShouldEqual, "web-vol")
				So(in.Volumes[0].Device, ShouldEqual, "/dev/sdx")
				So(len(in.SecurityGroups), ShouldEqual, 1)
				So(in.SecurityGroups[0], ShouldEqual, "web-sg")
				So(in.KeyPair, ShouldEqual, "test")
				So(in.Count, ShouldEqual, 1)
			})

		})
	})
}

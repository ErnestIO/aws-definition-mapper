/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package mapper

import (
	"testing"

	"github.com/ErnestIO/aws-definition-mapper/definition"
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
					So(i[1].Name, ShouldEqual, "datacenter-service-foo-2")
					So(i[1].Type, ShouldEqual, "m1.small")
					So(i[1].Image, ShouldEqual, "ami-000000")
					So(i[1].Network, ShouldEqual, "datacenter-service-bar")
				})
			})
		})
	})
}

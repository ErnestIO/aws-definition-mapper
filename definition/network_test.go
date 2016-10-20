/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed With this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package definition

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNetworkValidate(t *testing.T) {
	Convey("Given a network", t, func() {
		d := Datacenter{Region: "eu-west-1"}
		n := Network{Name: "foo", Subnet: "10.11.1.1/11", AvailabilityZone: "eu-west-1a"}
		Convey("With a valid subnet", func() {
			Convey("When validating the network", func() {
				err := n.Validate(&d)
				Convey("Then it should return an error", func() {
					So(err, ShouldBeNil)
				})
			})
		})

		Convey("With an invalid subnet", func() {
			n.Subnet = "10.11.1.11"
			Convey("When validating the network", func() {
				err := n.Validate(&d)
				Convey("Then it should return an error", func() {
					So(err, ShouldNotBeNil)
				})
			})
		})

		Convey("With an empty subnet", func() {
			n.Subnet = ""
			Convey("When validating the network", func() {
				err := n.Validate(&d)
				Convey("Then it should return an error", func() {
					So(err, ShouldNotBeNil)
				})
			})
		})

		Convey("With an invalid name", func() {
			n.Name = ""
			Convey("When validating the network", func() {
				err := n.Validate(&d)
				Convey("Then it should return an error", func() {
					So(err, ShouldNotBeNil)
				})
			})
		})

		Convey("With a name > 50 chars", func() {
			n.Name = "aksjhdlkashdliuhliusncldiudnalsundlaiunsdliausndliuansdlksbdlas"
			Convey("When validating the network", func() {
				err := n.Validate(&d)
				Convey("Then it should return an error", func() {
					So(err, ShouldNotBeNil)
				})
			})
		})

		Convey("With an availability zone that does not correspond to the datacenter region", func() {
			n.AvailabilityZone = "us-east-1a"
			Convey("When validating the network", func() {
				err := n.Validate(&d)
				Convey("Then it should return an error", func() {
					So(err, ShouldNotBeNil)
				})
			})
		})

	})
}

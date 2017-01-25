/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package definition

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestInstanceValidate(t *testing.T) {

	Convey("Given an instance", t, func() {
		n := &Network{Name: "test", Subnet: "127.0.0.0/24"}
		v := []EBSVolume{EBSVolume{Name: "test", Count: 1}}
		i := Instance{Name: "test", Type: "m1.small", Image: "ami-00000000", Count: 1, Network: "test", StartIP: "127.0.0.100", Volumes: []InstanceVolume{InstanceVolume{Volume: "test", Device: "/dev/sdx"}}}
		Convey("With an invalid name", func() {
			i.Name = ""
			Convey("When validating the instance", func() {
				err := i.Validate(n, v)
				Convey("Then should return an error", func() {
					So(err, ShouldNotBeNil)
					So(err.Error(), ShouldEqual, "Instance name should not be null")
				})
			})
		})

		Convey("With a name > 50 chars", func() {
			i.Name = "aksjhdlkashdliuhliusncldiudnalsundlaiunsdliausndliuansdlksbdlas"
			Convey("When validating the instance", func() {
				err := i.Validate(n, v)
				Convey("Then should return an error", func() {
					So(err, ShouldNotBeNil)
					So(err.Error(), ShouldEqual, "Instance name can't be greater than 50 characters")
				})
			})
		})

		Convey("With an invalid image", func() {
			i.Image = ""
			Convey("When validating the instance", func() {
				err := i.Validate(n, v)
				Convey("Then should return an error", func() {
					So(err, ShouldNotBeNil)
					So(err.Error(), ShouldEqual, "Instance image should not be null")
				})
			})
		})

		Convey("With an type type", func() {
			i.Type = ""
			Convey("When validating the instance", func() {
				err := i.Validate(n, v)
				Convey("Then should return an error", func() {
					So(err, ShouldNotBeNil)
					So(err.Error(), ShouldEqual, "Instance type should not be null")
				})
			})
		})

		Convey("With an instance count less than one", func() {
			i.Count = 0
			Convey("When validating the instance", func() {
				err := i.Validate(n, v)
				Convey("Then should return an error", func() {
					So(err, ShouldNotBeNil)
					So(err.Error(), ShouldEqual, "Instance count should not be < 1")
				})
			})
		})

		Convey("With an invalid network", func() {
			i.Network = ""
			Convey("When validating the instance", func() {
				err := i.Validate(n, v)
				Convey("Then should return an error", func() {
					So(err, ShouldNotBeNil)
					So(err.Error(), ShouldEqual, "Instance network should not be null")
				})
			})
		})

		Convey("With an start IP that is outside of the networks range", func() {
			i.StartIP = "10.10.10.10"
			Convey("When validating the instance", func() {
				err := i.Validate(n, v)
				Convey("Then should return an error", func() {
					So(err, ShouldNotBeNil)
					So(err.Error(), ShouldEqual, "Instance IP invalid. IP must be a valid IP in the same range as it's network")
				})
			})
		})

		Convey("With an ebs volume specified that doesn't have a high enough count", func() {
			i.Count = 5
			Convey("When validating the instance", func() {
				err := i.Validate(n, v)
				Convey("Then should return an error", func() {
					So(err, ShouldNotBeNil)
					So(err.Error(), ShouldEqual, "Instance count is higher than the specified ebs volume count attached to this instance")
				})
			})
		})

		Convey("With valid entries", func() {
			Convey("When validating the instance", func() {
				err := i.Validate(n, v)
				Convey("Then should return an error", func() {
					So(err, ShouldBeNil)
				})
			})
		})

	})
}

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
		i := Instance{Name: "test", Type: "m1.small", Image: "ami-00000000", Count: 1, Network: "test"}
		Convey("With an invalid name", func() {
			i.Name = ""
			Convey("When validating the instance", func() {
				err := i.Validate(n)
				Convey("Then should return an error", func() {
					So(err, ShouldNotBeNil)
					So(err.Error(), ShouldEqual, "Instance name should not be null")
				})
			})
		})

		Convey("With a name > 50 chars", func() {
			i.Name = "aksjhdlkashdliuhliusncldiudnalsundlaiunsdliausndliuansdlksbdlas"
			Convey("When validating the instance", func() {
				err := i.Validate(n)
				Convey("Then should return an error", func() {
					So(err, ShouldNotBeNil)
					So(err.Error(), ShouldEqual, "Instance name can't be greater than 50 characters")
				})
			})
		})

		Convey("With an invalid image", func() {
			i.Image = ""
			Convey("When validating the instance", func() {
				err := i.Validate(n)
				Convey("Then should return an error", func() {
					So(err, ShouldNotBeNil)
					So(err.Error(), ShouldEqual, "Instance image should not be null")
				})
			})
		})

		Convey("With an type type", func() {
			i.Type = ""
			Convey("When validating the instance", func() {
				err := i.Validate(n)
				Convey("Then should return an error", func() {
					So(err, ShouldNotBeNil)
					So(err.Error(), ShouldEqual, "Instance type should not be null")
				})
			})
		})

		Convey("With an instance count less than one", func() {
			i.Count = 0
			Convey("When validating the instance", func() {
				err := i.Validate(n)
				Convey("Then should return an error", func() {
					So(err, ShouldNotBeNil)
					So(err.Error(), ShouldEqual, "Instance count should not be < 1")
				})
			})
		})

		Convey("With an invalid network", func() {
			i.Network = ""
			Convey("When validating the instance", func() {
				err := i.Validate(n)
				Convey("Then should return an error", func() {
					So(err, ShouldNotBeNil)
					So(err.Error(), ShouldEqual, "Instance network should not be null")
				})
			})
		})

		Convey("With valid entries", func() {
			Convey("When validating the instance", func() {
				err := i.Validate(n)
				Convey("Then should return an error", func() {
					So(err, ShouldBeNil)
				})
			})
		})

	})
}

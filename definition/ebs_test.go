/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed With this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package definition

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestEBSValidate(t *testing.T) {
	Convey("Given an ebs volume", t, func() {
		e := EBSVolume{
			Name:             "foo",
			Type:             "gp2",
			Count:            1,
			AvailabilityZone: "eu-west-1",
			Encrypted:        true,
			EncryptionKeyID:  "test",
		}
		Convey("With a valid values", func() {
			Convey("When validating the ebs", func() {
				err := e.Validate()
				Convey("Then it should not return an error", func() {
					So(err, ShouldBeNil)
				})
			})
		})

		Convey("With an invalid name", func() {
			e.Name = ""
			Convey("When validating the ebs", func() {
				err := e.Validate()
				Convey("Thenn it should return an error", func() {
					So(err, ShouldNotBeNil)
				})
			})
		})

		Convey("With an invalid availability zone", func() {
			e.AvailabilityZone = ""
			Convey("When validating the ebs", func() {
				err := e.Validate()
				Convey("Thenn it should return an error", func() {
					So(err, ShouldNotBeNil)
				})
			})
		})

		Convey("With an invalid type", func() {
			e.Type = ""
			Convey("When validating the ebs", func() {
				err := e.Validate()
				Convey("Thenn it should return an error", func() {
					So(err, ShouldNotBeNil)
				})
			})
		})

		Convey("With no encryption key specified and encryption is enabled", func() {
			e.EncryptionKeyID = ""
			Convey("When validating the ebs", func() {
				err := e.Validate()
				Convey("Thenn it should return an error", func() {
					So(err, ShouldNotBeNil)
				})
			})
		})

		Convey("With iops set with an invalid type", func() {
			e.Iops = int64p(100)
			Convey("When validating the ebs", func() {
				err := e.Validate()
				Convey("Thenn it should return an error", func() {
					So(err, ShouldNotBeNil)
				})
			})
		})

		Convey("With a size below the minimum allowed value", func() {
			e.Size = int64p(0)
			Convey("When validating the ebs", func() {
				err := e.Validate()
				Convey("Thenn it should return an error", func() {
					So(err, ShouldNotBeNil)
				})
			})
		})

		Convey("With a size above the maximum allowed value", func() {
			e.Size = int64p(999999)
			Convey("When validating the ebs", func() {
				err := e.Validate()
				Convey("Thenn it should return an error", func() {
					So(err, ShouldNotBeNil)
				})
			})
		})

		Convey("With a count of less than one", func() {
			e.Size = int64p(999999)
			Convey("When validating the ebs", func() {
				err := e.Validate()
				Convey("Thenn it should return an error", func() {
					So(err, ShouldNotBeNil)
				})
			})
		})
	})
}

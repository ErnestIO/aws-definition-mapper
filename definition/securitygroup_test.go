/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package definition

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSecurityGroupRuleValidation(t *testing.T) {

	Convey("Given a forwarding rule", t, func() {
		n := []Network{Network{Name: "test", Subnet: "127.0.0.0/24"}}
		r := SecurityGroupRule{IP: "127.0.0.1", FromPort: "any", ToPort: "any", Protocol: "tcp"}

		Convey("When I try to validate a rule with an invalid destination ip", func() {
			r.IP = "invalid"
			err := r.Validate(n)
			Convey("Then should return an error", func() {
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, "Security Group IP (invalid) is not valid")
			})
		})

		Convey("When I try to validate a rule with a valid ip", func() {
			r.IP = "127.0.0.1"
			err := r.Validate(n)
			Convey("Then should return an error", func() {
				So(err, ShouldBeNil)
			})
		})

		Convey("When I try to validate a rule with a valid network name", func() {
			r.IP = "test"
			err := r.Validate(n)
			Convey("Then should return an error", func() {
				So(err, ShouldBeNil)
			})
		})

		Convey("When From Port is any", func() {
			r.FromPort = "any"
			err := r.Validate(n)
			Convey("Then should return an error", func() {
				So(err, ShouldBeNil)
			})
		})

		Convey("When From Port is not any and not numeric", func() {
			r.FromPort = "test"
			err := r.Validate(n)
			Convey("Then should return an error", func() {
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, "Security Group From Port (test) is not valid")
			})
		})

		Convey("When From Port is not any and not in range", func() {
			r.FromPort = "0"
			err := r.Validate(n)
			Convey("Then should return an error", func() {
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, "Security Group From Port (0) is out of range [1 - 65535]")
			})
		})

		Convey("When From Port is not any and great than range", func() {
			r.FromPort = "65536"
			err := r.Validate(n)
			Convey("Then should return an error", func() {
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, "Security Group From Port (65536) is out of range [1 - 65535]")
			})
		})

		Convey("When Protocol is not valid", func() {
			r.Protocol = "Protocol"
			err := r.Validate(n)
			Convey("Then should return an error", func() {
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, "Protocol is invalid")
			})
		})

		Convey("When Protocol is valid", func() {
			r.Protocol = "tcp"
			err := r.Validate(n)
			Convey("Then should return an error", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}

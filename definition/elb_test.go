/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed With this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package definition

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestELBValidate(t *testing.T) {
	Convey("Given an elb", t, func() {
		e := ELB{
			Name: "foo",
			Listeners: []ELBListener{
				ELBListener{
					FromPort: 1,
					ToPort:   1,
					Protocol: "http",
					SSLCert:  "",
				},
			},
		}
		Convey("With a valid port mappings", func() {
			Convey("When validating the elb", func() {
				err := e.Validate()
				Convey("Then it should not return an error", func() {
					So(err, ShouldBeNil)
				})
			})
		})

		Convey("With an invalid from port", func() {
			e.Listeners[0].FromPort = 0
			Convey("When validating the elb", func() {
				err := e.Validate()
				Convey("Then it should return an error", func() {
					So(err, ShouldNotBeNil)
				})
			})
		})

		Convey("With an invalid to port", func() {
			e.Listeners[0].ToPort = 999999
			Convey("When validating the elb", func() {
				err := e.Validate()
				Convey("Then it should return an error", func() {
					So(err, ShouldNotBeNil)
				})
			})
		})

		Convey("With an invalid protocol", func() {
			e.Listeners[0].Protocol = "invalid"
			Convey("When validating the elb", func() {
				err := e.Validate()
				Convey("Then it should return an error", func() {
					So(err, ShouldNotBeNil)
				})
			})
		})

		Convey("With a name > 50 chars", func() {
			e.Name = "aksjhdlkashdliuhliusncldiudnalsundlaiunsdliausndliuansdlksbdlas"
			Convey("When validating the elb", func() {
				err := e.Validate()
				Convey("Then it should return an error", func() {
					So(err, ShouldNotBeNil)
				})
			})
		})

	})
}

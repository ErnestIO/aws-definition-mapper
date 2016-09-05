/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed With this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package definition

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNatGatewayValidate(t *testing.T) {
	Convey("Given a nat gateway", t, func() {
		networks := []Network{Network{Name: "test", Public: true}}
		n := NatGateway{Name: "foo", PublicNetwork: "test"}
		Convey("With a valid subnet", func() {
			Convey("When validating the nat gateway", func() {
				err := n.Validate(networks)
				Convey("Then it should return an error", func() {
					So(err, ShouldBeNil)
				})
			})
		})

		Convey("With a blank public network", func() {
			n.PublicNetwork = ""
			Convey("When validating the nat gateway", func() {
				err := n.Validate(networks)
				Convey("Then it should return an error", func() {
					So(err, ShouldNotBeNil)
				})
			})
		})

		Convey("With an undefined public network", func() {
			n.PublicNetwork = "undefined"
			Convey("When validating the nat gateway", func() {
				err := n.Validate(networks)
				Convey("Then it should return an error", func() {
					So(err, ShouldNotBeNil)
				})
			})
		})

		Convey("With an invalid name", func() {
			n.Name = ""
			Convey("When validating the nat gateway", func() {
				err := n.Validate(networks)
				Convey("Then it should return an error", func() {
					So(err, ShouldNotBeNil)
				})
			})
		})

		Convey("With a name > 50 chars", func() {
			n.Name = "aksjhdlkashdliuhliusncldiudnalsundlaiunsdliausndliuansdlksbdlas"
			Convey("When validating the nat gateway", func() {
				err := n.Validate(networks)
				Convey("Then it should return an error", func() {
					So(err, ShouldNotBeNil)
				})
			})
		})

	})
}

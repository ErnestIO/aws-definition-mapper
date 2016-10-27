/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package output

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFirewallHasChanged(t *testing.T) {
	Convey("Given a firewall", t, func() {
		var f Firewall
		f.Rules.Ingress = append(f.Rules.Ingress, FirewallRule{
			IP:       "10.10.10.10",
			From:     80,
			To:       80,
			Protocol: "tcp",
		})

		Convey("When I compare it to an changed firewall", func() {
			var of Firewall

			of.Rules.Ingress = append(of.Rules.Ingress, FirewallRule{
				IP:       "10.10.10.10",
				From:     80,
				To:       8080,
				Protocol: "tcp",
			})

			change := f.HasChanged(&of)
			Convey("Then it should return true", func() {
				So(change, ShouldBeTrue)
			})
		})

		Convey("When I compare it to an identical firewall", func() {
			var of Firewall

			of.Rules.Ingress = append(of.Rules.Ingress, FirewallRule{
				IP:       "10.10.10.10",
				From:     80,
				To:       80,
				Protocol: "tcp",
			})

			change := f.HasChanged(&of)
			Convey("Then it should return false", func() {
				So(change, ShouldBeFalse)
			})
		})
	})
}

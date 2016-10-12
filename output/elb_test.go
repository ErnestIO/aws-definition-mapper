/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package output

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestELBHasChanged(t *testing.T) {
	Convey("Given a elb", t, func() {
		e := ELB{
			Name: "test",
			InstanceAWSIDs: []string{
				"web",
			},
			SecurityGroupAWSIDs: []string{
				"web-sg",
			},
			Listeners: []ELBListener{
				ELBListener{
					FromPort: 1,
					ToPort:   2,
					Protocol: "http",
					SSLCert:  "cert",
				},
			},
		}

		Convey("When I compare it to an changed elb", func() {
			oe := ELB{
				Name: "test",
				InstanceAWSIDs: []string{
					"web",
				},
				SecurityGroupAWSIDs: []string{
					"web-sg",
				},
				Listeners: []ELBListener{
					ELBListener{
						FromPort: 1,
						ToPort:   80,
						Protocol: "http",
						SSLCert:  "cert",
					},
				},
			}
			change := e.HasChanged(&oe)
			Convey("Then it should return true", func() {
				So(change, ShouldBeTrue)
			})
		})

		Convey("When I compare it to an identical elb", func() {
			oe := ELB{
				Name: "test",
				InstanceAWSIDs: []string{
					"web",
				},
				SecurityGroupAWSIDs: []string{
					"web-sg",
				},
				Listeners: []ELBListener{
					ELBListener{
						FromPort: 1,
						ToPort:   2,
						Protocol: "http",
						SSLCert:  "cert",
					},
				},
			}

			change := e.HasChanged(&oe)
			Convey("Then it should return false", func() {
				So(change, ShouldBeFalse)
			})
		})
	})
}

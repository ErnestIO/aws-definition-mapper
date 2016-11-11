/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package output

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRoute53HasChanged(t *testing.T) {
	Convey("Given a route53 zone", t, func() {
		z := Route53Zone{
			Name:    "example.com",
			Private: false,
			Records: []Record{
				Record{
					Entry:  "test.example.com",
					Type:   "A",
					TTL:    3600,
					Values: []string{"8.8.8.8"},
				},
			},
		}

		// Route53s are immutable
		Convey("When I compare it to an changed route53 zone", func() {
			oz := Route53Zone{
				Name:    "example.com",
				Private: false,
				Records: []Record{
					Record{
						Entry:  "test.example.com",
						Type:   "A",
						TTL:    3600,
						Values: []string{"8.8.8.8", "8.8.4.4"},
					},
				},
			}
			change := z.HasChanged(&oz)
			Convey("Then it should return true", func() {
				So(change, ShouldBeTrue)
			})
		})

		Convey("When I compare it to an identical route53 zone", func() {
			oz := Route53Zone{
				Name:    "example.com",
				Private: false,
				Records: []Record{
					Record{
						Entry:  "test.example.com",
						Type:   "A",
						TTL:    3600,
						Values: []string{"8.8.8.8"},
					},
				},
			}
			change := z.HasChanged(&oz)
			Convey("Then it should return false", func() {
				So(change, ShouldBeFalse)
			})
		})
	})
}

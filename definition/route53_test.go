/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed With this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package definition

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRoute53Validate(t *testing.T) {
	Convey("Given a route53 zone", t, func() {
		z := Route53Zone{
			Name:    "example.com",
			Private: false,
			Records: []Record{
				Record{
					Entry:  "one.example.com",
					Type:   "A",
					Values: []string{"8.8.8.8"},
					TTL:    3600,
				},
				Record{
					Entry:     "two.example.com",
					Type:      "A",
					Instances: []string{"web-1"},
					TTL:       3600,
				},
				Record{
					Entry:         "two.example.com",
					Type:          "CNAME",
					Loadbalancers: []string{"lb-1"},
					TTL:           3600,
				},
			},
		}
		Convey("With a valid record", func() {
			Convey("When validating the route53 zone", func() {
				err := z.Validate()
				Convey("Then it should return an error", func() {
					So(err, ShouldBeNil)
				})
			})
		})

		Convey("With an invalid record entry", func() {
			z.Records[0].Entry = ""
			Convey("When validating the route53 zone", func() {
				err := z.Validate()
				Convey("Then it should return an error", func() {
					So(err, ShouldNotBeNil)
				})
			})
		})

		Convey("With an invalid record type", func() {
			z.Records[0].Type = "FAKE"
			Convey("When validating the route53 zone", func() {
				err := z.Validate()
				Convey("Then it should return an error", func() {
					So(err, ShouldNotBeNil)
				})
			})
		})

		Convey("With an record TTL < 1", func() {
			z.Records[0].TTL = 0
			Convey("When validating the route53 zone", func() {
				err := z.Validate()
				Convey("Then it should return an error", func() {
					So(err, ShouldNotBeNil)
				})
			})
		})

		Convey("With a both instances and loadbalancers specified", func() {
			z.Records[1].Loadbalancers = []string{"lb-1"}
			Convey("When validating the route53 zone", func() {
				err := z.Validate()
				Convey("Then it should return an error", func() {
					So(err, ShouldNotBeNil)
				})
			})
		})

		Convey("With an invalid type (A Record) set for loadbalancer record", func() {
			z.Records[2].Type = "A"
			Convey("When validating the route53 zone", func() {
				err := z.Validate()
				Convey("Then it should return an error", func() {
					So(err, ShouldNotBeNil)
				})
			})
		})

		Convey("With an invalid type (MX Record) set for instance record", func() {
			z.Records[1].Type = "MX"
			Convey("When validating the route53 zone", func() {
				err := z.Validate()
				Convey("Then it should return an error", func() {
					So(err, ShouldNotBeNil)
				})
			})
		})

		Convey("With a record with no targets set", func() {
			z.Records[0].Values = []string{}
			Convey("When validating the route53 zone", func() {
				err := z.Validate()
				Convey("Then it should return an error", func() {
					So(err, ShouldNotBeNil)
				})
			})
		})
	})
}

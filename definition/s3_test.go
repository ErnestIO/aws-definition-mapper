/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed With this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package definition

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestS3Validate(t *testing.T) {
	Convey("Given an s3 bucket", t, func() {
		s := S3{
			Name:           "test",
			BucketLocation: "eu-west-1",
			Grantees: []S3Grantee{
				S3Grantee{
					ID:          "test",
					Type:        "id",
					Permissions: "full_control",
				},
			},
		}

		Convey("With valid fields", func() {
			Convey("When validating the s3 bucket", func() {
				err := s.Validate()
				Convey("Then it should not return an error", func() {
					So(err, ShouldBeNil)
				})
			})
		})

		Convey("With an invalid name", func() {
			s.Name = ""
			Convey("When validating the s3 bucket", func() {
				err := s.Validate()
				Convey("Then it should return an error", func() {
					So(err, ShouldNotBeNil)
				})
			})
		})

		Convey("With a name > 50 chars", func() {
			s.Name = "aksjhdlkashdliuhliusncldiudnalsundlaiunsdliausndliuansdlksbdlas"
			Convey("When validating the s3 bucket", func() {
				err := s.Validate()
				Convey("Then it should return an error", func() {
					So(err, ShouldNotBeNil)
				})
			})
		})

		Convey("With an invalid ACL", func() {
			s.ACL = "blah"
			s.Grantees = []S3Grantee{}
			Convey("When validating the s3 bucket", func() {
				err := s.Validate()
				Convey("Then it should return an error", func() {
					So(err, ShouldNotBeNil)
				})
			})
		})

		Convey("With an empty bucket location", func() {
			s.BucketLocation = ""
			Convey("When validating the s3 bucket", func() {
				err := s.Validate()
				Convey("Then it should return an error", func() {
					So(err, ShouldNotBeNil)
				})
			})
		})

		Convey("With an invalid grantee id", func() {
			s.Grantees[0].ID = ""
			Convey("When validating the s3 bucket", func() {
				err := s.Validate()
				Convey("Then it should return an error", func() {
					So(err, ShouldNotBeNil)
				})
			})
		})

		Convey("With an invalid grantee type", func() {
			s.Grantees[0].Type = "invalid"
			Convey("When validating the s3 bucket", func() {
				err := s.Validate()
				Convey("Then it should return an error", func() {
					So(err, ShouldNotBeNil)
				})
			})
		})

		Convey("With an invalid grantee permissions", func() {
			s.Grantees[0].Permissions = "invalid"
			Convey("When validating the s3 bucket", func() {
				err := s.Validate()
				Convey("Then it should return an error", func() {
					So(err, ShouldNotBeNil)
				})
			})
		})

	})
}

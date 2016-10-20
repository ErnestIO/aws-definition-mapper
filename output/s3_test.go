/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package output

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestS3HasChanged(t *testing.T) {
	Convey("Given a s3 bucket", t, func() {
		s := S3{
			Name: "test",
			ACL:  "full",
			Grantees: []S3Grantee{
				S3Grantee{
					Type: "something",
					ID:   "test",
				},
			},
		}

		// S3s are immutable
		Convey("When I compare it to an changed s3 bucket", func() {
			os := S3{
				Name: "test",
				ACL:  "full",
				Grantees: []S3Grantee{
					S3Grantee{
						Type: "id",
						ID:   "test",
					},
					S3Grantee{
						Type: "id",
						ID:   "test2",
					},
				},
			}
			change := s.HasChanged(&os)
			Convey("Then it should return true", func() {
				So(change, ShouldBeTrue)
			})
		})

		Convey("When I compare it to an identical s3 bucket", func() {
			os := S3{
				Name: "test",
				ACL:  "full",
				Grantees: []S3Grantee{
					S3Grantee{
						Type: "something",
						ID:   "test",
					},
				},
			}
			change := s.HasChanged(&os)
			Convey("Then it should return false", func() {
				So(change, ShouldBeFalse)
			})
		})
	})
}

/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package mapper

import (
	"strings"

	"github.com/ernestio/aws-definition-mapper/definition"
	"github.com/ernestio/aws-definition-mapper/output"
)

// MapS3Buckets : Maps the s3 buckets from a given input payload.
func MapS3Buckets(d definition.Definition) []output.S3 {
	var s3buckets []output.S3

	for _, s3 := range d.S3Buckets {

		s := output.S3{
			Name:             s3.Name,
			ACL:              strings.ToUpper(s3.ACL),
			BucketLocation:   s3.BucketLocation,
			ProviderType:     "$(datacenters.items.0.type)",
			DatacenterName:   "$(datacenters.items.0.name)",
			DatacenterSecret: "$(datacenters.items.0.secret)",
			DatacenterToken:  "$(datacenters.items.0.token)",
			DatacenterRegion: "$(datacenters.items.0.region)",
		}

		for _, grantee := range s3.Grantees {
			s.Grantees = append(s.Grantees, output.S3Grantee{
				ID:          grantee.ID,
				Type:        grantee.Type,
				Permissions: grantee.Permissions,
			})
		}

		s3buckets = append(s3buckets, s)
	}

	return s3buckets
}

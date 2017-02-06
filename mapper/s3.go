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
		name := d.GeneratedName() + s3.Name

		s := output.S3{
			Name:             s3.Name,
			ACL:              s3.ACL,
			BucketLocation:   s3.BucketLocation,
			Tags:             mapTags(name, d.Name),
			ProviderType:     "$(datacenters.items.0.type)",
			DatacenterName:   "$(datacenters.items.0.name)",
			SecretAccessKey:  "$(datacenters.items.0.aws_secret_access_key)",
			AccessKeyID:      "$(datacenters.items.0.aws_access_key_id)",
			DatacenterRegion: "$(datacenters.items.0.region)",
		}

		for _, grantee := range s3.Grantees {
			s.Grantees = append(s.Grantees, output.S3Grantee{
				ID:          grantee.ID,
				Type:        grantee.Type,
				Permissions: strings.ToUpper(grantee.Permissions),
			})
		}

		s3buckets = append(s3buckets, s)
	}

	return s3buckets
}

// MapDefinitionS3Buckets : Maps the s3 buckets from the internal format to the input definition format.
func MapDefinitionS3Buckets(m *output.FSMMessage) []definition.S3 {
	var s3buckets []definition.S3

	for _, s3 := range m.S3s.Items {

		s := definition.S3{
			Name:           s3.Name,
			ACL:            s3.ACL,
			BucketLocation: s3.BucketLocation,
		}

		for _, grantee := range s3.Grantees {
			s.Grantees = append(s.Grantees, definition.S3Grantee{
				ID:          grantee.ID,
				Type:        strings.ToLower(grantee.Type),
				Permissions: strings.ToLower(grantee.Permissions),
			})
		}

		s3buckets = append(s3buckets, s)
	}

	return s3buckets
}

// UpdateS3Values corrects missing values after an import
func UpdateS3Values(m *output.FSMMessage) {
	for i := 0; i < len(m.S3s.Items); i++ {
		m.S3s.Items[i].ProviderType = "$(datacenters.items.0.type)"
		m.S3s.Items[i].AccessKeyID = "$(datacenters.items.0.aws_access_key_id)"
		m.S3s.Items[i].SecretAccessKey = "$(datacenters.items.0.aws_secret_access_key)"
		m.S3s.Items[i].DatacenterRegion = "$(datacenters.items.0.region)"
	}
}

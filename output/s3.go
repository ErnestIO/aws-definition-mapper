/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed With this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package output

import "reflect"

type S3Grantee struct {
	ID          string `json:"id"`
	Type        string `json:"type"`
	Permissions string `json:"permissions"`
}

// S3 represents an aws S3 bucket
type S3 struct {
	ProviderType     string      `json:"_type"`
	DatacenterName   string      `json:"datacenter_name,omitempty"`
	DatacenterRegion string      `json:"datacenter_region"`
	DatacenterToken  string      `json:"datacenter_token"`
	DatacenterSecret string      `json:"datacenter_secret"`
	Name             string      `json:"name"`
	ACL              string      `json:"acl"`
	BucketLocation   string      `json:"bucket_location"`
	BucketURI        string      `json:"bucket_uri"`
	Grantees         []S3Grantee `json:"grantees"`
	Service          string      `json:"service"`
	Status           string      `json:"status"`
	Exists           bool
}

// HasChanged diff's the two items and returns true if there have been any changes
func (s *S3) HasChanged(os *S3) bool {
	if s.ACL != os.ACL {
		return true
	}

	return !reflect.DeepEqual(s.Grantees, os.Grantees)
}

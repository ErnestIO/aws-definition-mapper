/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package output

import "reflect"

// Record stores the entries for a zone
type Record struct {
	Entry  string   `json:"entry"`
	Type   string   `json:"type"`
	Values []string `json:"values"`
	TTL    int64    `json:"ttl"`
}

// Route53Zone holds all information about a dns zone
type Route53Zone struct {
	ProviderType     string            `json:"_type"`
	HostedZoneID     string            `json:"hosted_zone_id"`
	Name             string            `json:"name"`
	Private          bool              `json:"private"`
	Records          []Record          `json:"records"`
	Tags             map[string]string `json:"tags"`
	VPCID            string            `json:"vpc_id"`
	DatacenterName   string            `json:"datacenter_name,omitempty"`
	DatacenterRegion string            `json:"datacenter_region"`
	AccessKeyID      string            `json:"aws_access_key_id"`
	SecretAccessKey  string            `json:"aws_secret_access_key"`
	Service          string            `json:"service"`
	Status           string            `json:"status"`
	Exists           bool
}

// HasChanged diff's the two items and returns true if there have been any changes
func (z *Route53Zone) HasChanged(oz *Route53Zone) bool {
	if len(z.Records) != len(oz.Records) {
		return true
	}

	return !reflect.DeepEqual(z.Records, oz.Records)
}

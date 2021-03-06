/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package output

// VPC : mapping of an instance component
type VPC struct {
	DatacenterName   string            `json:"datacenter_name"`
	DatacenterRegion string            `json:"datacenter_region"`
	AccessKeyID      string            `json:"aws_access_key_id"`
	SecretAccessKey  string            `json:"aws_secret_access_key"`
	VpcID            string            `json:"vpc_id"`
	VpcSubnet        string            `json:"vpc_subnet"`
	Tags             map[string]string `json:"tags"`
	Type             string            `json:"_type"`
	Status           string            `json:"status"`
	Exists           bool
}

// HasChanged diff's the two items and returns true if there have been any changes
func (v *VPC) HasChanged(ov *VPC) bool {
	if v.VpcSubnet != ov.VpcSubnet {
		return true
	}

	return false
}

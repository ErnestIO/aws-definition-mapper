/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package output

// Network : Mapping of a network component
type Network struct {
	ProviderType     string            `json:"_type"`
	NetworkAWSID     string            `json:"network_aws_id"`
	Name             string            `json:"name"`
	Subnet           string            `json:"range"`
	IsPublic         bool              `json:"is_public"`
	Tags             map[string]string `json:"tags"`
	AvailabilityZone string            `json:"availability_zone"`
	DatacenterType   string            `json:"datacenter_type"`
	DatacenterName   string            `json:"datacenter_name"`
	DatacenterRegion string            `json:"datacenter_region"`
	AccessKeyID      string            `json:"aws_access_key_id"`
	SecretAccessKey  string            `json:"aws_secret_access_key"`
	VpcID            string            `json:"vpc_id"`
	Service          string            `json:"service"`
	Status           string            `json:"status"`
	Exists           bool
}

// HasChanged diff's the two items and returns true if there have been any changes
func (n *Network) HasChanged(on *Network) bool {
	return false
}

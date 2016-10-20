/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package output

// Network : Mapping of a network component
type Network struct {
	Type             string `json:"_type"`
	NetworkAWSID     string `json:"network_aws_id"`
	Name             string `json:"name"`
	Subnet           string `json:"range"`
	IsPublic         bool   `json:"is_public"`
	AvailabilityZone string `json:"availability_zone"`
	DatacenterType   string `json:"datacenter_type"`
	DatacenterName   string `json:"datacenter_name"`
	DatacenterRegion string `json:"datacenter_region"`
	DatacenterToken  string `json:"datacenter_token"`
	DatacenterSecret string `json:"datacenter_secret"`
	VpcID            string `json:"vpc_id"`
	Service          string `json:"service"`
	Status           string `json:"status"`
	Exists           bool
}

// HasChanged should always return false, networks are immutable and cannot be updated!
func (n *Network) HasChanged(on *Network) bool {
	return false
}

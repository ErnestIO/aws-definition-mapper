/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package output

// Network : Mapping of a network component
type Network struct {
	NetworkAWSID     string `json:"network_aws_id"`
	Name             string `json:"name"`
	Type             string `json:"type"`
	Subnet           string `json:"range"`
	IsPublic         bool   `json:"is_public"`
	RouterName       string `json:"router_name"`
	RouterType       string `json:"router_type"`
	RouterIP         string `json:"router_ip"`
	ClientName       string `json:"client_name"`
	DatacenterType   string `json:"datacenter_type"`
	DatacenterName   string `json:"datacenter_name"`
	DatacenterRegion string `json:"datacenter_region"`
	DatacenterToken  string `json:"datacenter_token"`
	DatacenterSecret string `json:"datacenter_secret"`
	NetworkType      string `json:"network_type"`
	NetworkSubnet    string `json:"network_subnet"`
	VpcID            string `json:"vpc_id"`
	Service          string `json:"service"`
	Status           string `json:"status"`
	Exists           bool
}

// HasChanged diff's the two items and returns true if there have been any changes
func (n *Network) HasChanged(on *Network) bool {
	if n.Name == on.Name &&
		n.Type == on.Type &&
		n.Subnet == on.Subnet &&
		n.IsPublic == on.IsPublic &&
		n.Service == on.Service {
		return false
	}
	return true
}

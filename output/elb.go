/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package output

import "reflect"

// ELBPort ...
type ELBPort struct {
	FromPort int    `json:"from_port"`
	ToPort   int    `json:"to_port"`
	Protocol string `json:"protocol"`
	SSLCert  string `json:"ssl_cert"`
}

// ELB : Mapping for a elb component
type ELB struct {
	Type                string    `json:"_type"`
	Name                string    `json:"elb_name"`
	IsPrivate           bool      `json:"elb_is_private"`
	DNSName             string    `json:"elb_dns_name"`
	Ports               []ELBPort `json:"elb_ports"`
	NetworkAWSIDs       []string  `json:"network_aws_ids"`
	Instances           []string  `json:"instances"`
	InstanceAWSIDs      []string  `json:"instance_aws_ids"`
	SecurityGroups      []string  `json:"security_groups"`
	SecurityGroupAWSIDs []string  `json:"security_group_aws_ids"`
	DatacenterType      string    `json:"datacenter_type,omitempty"`
	DatacenterName      string    `json:"datacenter_name,omitempty"`
	DatacenterRegion    string    `json:"datacenter_region"`
	DatacenterToken     string    `json:"datacenter_token"`
	DatacenterSecret    string    `json:"datacenter_secret"`
	VpcID               string    `json:"vpc_id"`
	Service             string    `json:"service"`
	Status              string    `json:"status"`
	Exists              bool
}

// HasChanged diff's the two items and returns true if there have been any changes
func (e *ELB) HasChanged(oe *ELB) bool {
	if len(e.Ports) != len(oe.Ports) {
		return true
	}

	for i := 0; i < len(e.Ports); i++ {
		if e.Ports[i].FromPort != oe.Ports[i].FromPort ||
			e.Ports[i].ToPort != oe.Ports[i].ToPort ||
			e.Ports[i].Protocol != oe.Ports[i].Protocol ||
			e.Ports[i].SSLCert != oe.Ports[i].SSLCert {
			return true
		}
	}

	if !reflect.DeepEqual(e.InstanceAWSIDs, oe.InstanceAWSIDs) {
		return true
	}

	if !reflect.DeepEqual(e.SecurityGroupAWSIDs, oe.SecurityGroupAWSIDs) {
		return true
	}

	return false
}

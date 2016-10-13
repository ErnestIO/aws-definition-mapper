/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package output

import "reflect"

// ELBListener ...
type ELBListener struct {
	FromPort int    `json:"from_port"`
	ToPort   int    `json:"to_port"`
	Protocol string `json:"protocol"`
	SSLCert  string `json:"ssl_cert"`
}

// ELB : Mapping for a elb component
type ELB struct {
	Type                string        `json:"_type"`
	Name                string        `json:"name"`
	IsPrivate           bool          `json:"is_private"`
	DNSName             string        `json:"dns_name"`
	Listeners           []ELBListener `json:"listeners"`
	NetworkAWSIDs       []string      `json:"network_aws_ids"`
	Instances           []string      `json:"instances"`
	InstanceNames       []string      `json:"instance_names"`
	InstanceAWSIDs      []string      `json:"instance_aws_ids"`
	SecurityGroups      []string      `json:"security_groups"`
	SecurityGroupAWSIDs []string      `json:"security_group_aws_ids"`
	DatacenterType      string        `json:"datacenter_type,omitempty"`
	DatacenterName      string        `json:"datacenter_name,omitempty"`
	DatacenterRegion    string        `json:"datacenter_region"`
	DatacenterToken     string        `json:"datacenter_token"`
	DatacenterSecret    string        `json:"datacenter_secret"`
	VpcID               string        `json:"vpc_id"`
	Service             string        `json:"service"`
	Status              string        `json:"status"`
	Exists              bool
}

// HasChanged diff's the two items and returns true if there have been any changes
func (e *ELB) HasChanged(oe *ELB) bool {
	if len(e.Listeners) != len(oe.Listeners) {
		return true
	}

	for i := 0; i < len(e.Listeners); i++ {
		if e.Listeners[i].FromPort != oe.Listeners[i].FromPort ||
			e.Listeners[i].ToPort != oe.Listeners[i].ToPort ||
			e.Listeners[i].Protocol != oe.Listeners[i].Protocol ||
			e.Listeners[i].SSLCert != oe.Listeners[i].SSLCert {
			return true
		}
	}

	if !reflect.DeepEqual(e.InstanceNames, oe.InstanceNames) {
		return true
	}

	if !reflect.DeepEqual(e.SecurityGroups, oe.SecurityGroups) {
		return true
	}

	return false
}

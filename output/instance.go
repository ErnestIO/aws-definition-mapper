/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package output

import (
	"net"
	"reflect"
)

// InstanceVolume ...
type InstanceVolume struct {
	VolumeAWSID string `json:"volume_aws_id"`
	Volume      string `json:"volume"`
	Device      string `json:"device"`
}

// Instance : mapping of an instance component
type Instance struct {
	ProviderType        string            `json:"_type"`
	InstanceAWSID       string            `json:"instance_aws_id"`
	Name                string            `json:"name"`
	Type                string            `json:"instance_type"`
	Image               string            `json:"image"`
	IP                  net.IP            `json:"ip"`
	PublicIP            string            `json:"public_ip"`
	ElasticIP           string            `json:"elastic_ip"`
	ElasticIPAWSID      *string           `json:"elastic_ip_aws_id"`
	AssignElasticIP     bool              `json:"assign_elastic_ip"`
	KeyPair             string            `json:"key_pair"`
	UserData            string            `json:"user_data"`
	Network             string            `json:"network_name"`
	NetworkAWSID        string            `json:"network_aws_id"`
	NetworkIsPublic     bool              `json:"network_is_public"`
	SecurityGroups      []string          `json:"security_groups"`
	SecurityGroupAWSIDs []string          `json:"security_group_aws_ids"`
	Volumes             []InstanceVolume  `json:"volumes"`
	Tags                map[string]string `json:"tags"`
	DatacenterType      string            `json:"datacenter_type,omitempty"`
	DatacenterName      string            `json:"datacenter_name,omitempty"`
	DatacenterRegion    string            `json:"datacenter_region"`
	AccessKeyID         string            `json:"aws_access_key_id"`
	SecretAccessKey     string            `json:"aws_secret_access_key"`
	VpcID               string            `json:"vpc_id"`
	Service             string            `json:"service"`
	Status              string            `json:"status"`
	Exists              bool
}

// HasChanged diff's the two items and returns true if there have been any changes
func (i *Instance) HasChanged(oi *Instance) bool {
	if i.Type != oi.Type {
		return true
	}

	if reflect.DeepEqual(i.Volumes, oi.Volumes) != true {
		return true
	}

	return !reflect.DeepEqual(i.SecurityGroups, oi.SecurityGroups)
}

// GetTags returns a components tags
func (i Instance) GetTags() map[string]string {
	return i.Tags
}

// ProviderID returns a components provider id
func (i Instance) ProviderID() string {
	return i.InstanceAWSID
}

// ComponentName returns a components name
func (i Instance) ComponentName() string {
	return i.Name
}

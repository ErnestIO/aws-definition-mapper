/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package output

import "reflect"

// RDSCluster ...
type RDSCluster struct {
	ProviderType        string            `json:"_type"`
	VpcID               string            `json:"vpc_id"`
	DatacenterRegion    string            `json:"datacenter_region"`
	AccessKeyID         string            `json:"aws_access_key_id"`
	SecretAccessKey     string            `json:"aws_secret_access_key"`
	ARN                 string            `json:"arn"`
	Name                string            `json:"name"`
	Engine              string            `json:"engine"`
	EngineVersion       string            `json:"engine_version,omitempty"`
	Port                *int64            `json:"port,omitempty"`
	Endpoint            string            `json:"endpoint,omitempty"`
	AvailabilityZones   []string          `json:"availability_zones"`
	SecurityGroups      []string          `json:"security_groups"`
	SecurityGroupAWSIDs []string          `json:"security_group_aws_ids"`
	Networks            []string          `json:"networks"`
	NetworkAWSIDs       []string          `json:"network_aws_ids"`
	Tags                map[string]string `json:"tags"`
	DatabaseName        string            `json:"database_name,omitempty"`
	DatabaseUsername    string            `json:"database_username,omitempty"`
	DatabasePassword    string            `json:"database_password,omitempty"`
	BackupRetention     *int64            `json:"backup_retention,omitempty"`
	BackupWindow        string            `json:"backup_window,omitempty"`
	MaintenanceWindow   string            `json:"maintenance_window,omitempty"`
	ReplicationSource   string            `json:"replication_source,omitempty"`
	FinalSnapshot       bool              `json:"final_snapshot"`
	Status              string            `json:"status"`
	Exists              bool
}

// HasChanged diff's the two items and returns true if there have been any changes
func (r *RDSCluster) HasChanged(or *RDSCluster) bool {
	if r.Port != nil && or.Port != nil {
		if *r.Port != *or.Port {
			return true
		}
	}

	if r.DatabasePassword != or.DatabasePassword {
		return true
	}

	if r.BackupRetention != nil && or.BackupRetention != nil {
		if *r.BackupRetention != *or.BackupRetention {
			return true
		}
	}

	if r.BackupWindow != or.BackupWindow {
		return true
	}

	if r.MaintenanceWindow != or.MaintenanceWindow {
		return true
	}

	if reflect.DeepEqual(r.Networks, or.Networks) != true {
		return true
	}

	return !reflect.DeepEqual(r.SecurityGroups, or.SecurityGroups)
}

// GetTags returns a components tags
func (r RDSCluster) GetTags() map[string]string {
	return r.Tags
}

// ProviderID returns a components provider id
func (r RDSCluster) ProviderID() string {
	return r.Name
}

// ComponentName returns a components name
func (r RDSCluster) ComponentName() string {
	return r.Name
}

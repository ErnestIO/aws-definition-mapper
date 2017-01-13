/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package output

import "reflect"

// RDSInstance ...
type RDSInstance struct {
	ProviderType        string            `json:"_type"`
	VpcID               string            `json:"vpc_id"`
	DatacenterRegion    string            `json:"datacenter_region"`
	AccessKeyID         string            `json:"aws_access_key_id"`
	SecretAccessKey     string            `json:"aws_secret_access_key"`
	Name                string            `json:"name"`
	Size                string            `json:"size"`
	Engine              string            `json:"engine"`
	EngineVersion       string            `json:"engine_version,omitempty"`
	Port                *int64            `json:"port,omitempty"`
	Cluster             string            `json:"cluster,omitempty"`
	Public              bool              `json:"public"`
	Endpoint            string            `json:"endpoint,omitempty"`
	MultiAZ             bool              `json:"multi_az"`
	PromotionTier       *int64            `json:"promotion_tier,omitempty"`
	StorageType         string            `json:"storage_type,omitempty"`
	StorageSize         *int64            `json:"storage_size,omitempty"`
	StorageIops         *int64            `json:"storage_iops,omitempty"`
	AvailabilityZone    string            `json:"availability_zone,omitempty"`
	SecurityGroups      []string          `json:"security_groups"`
	SecurityGroupAWSIDs []string          `json:"security_group_aws_ids"`
	Networks            []string          `json:"networks"`
	NetworkAWSIDs       []string          `json:"network_aws_ids"`
	Tags                map[string]string `json:"tags"`
	DatabaseName        string            `json:"database_name,omitempty"`
	DatabaseUsername    string            `json:"database_username,omitempty"`
	DatabasePassword    string            `json:"database_password,omitempty"`
	AutoUpgrade         bool              `json:"auto_upgrade"`
	BackupRetention     *int64            `json:"backup_retention,omitempty"`
	BackupWindow        string            `json:"backup_window,omitempty"`
	MaintenanceWindow   string            `json:"maintenance_window,omitempty"`
	FinalSnapshot       bool              `json:"final_snapshot"`
	ReplicationSource   string            `json:"replication_source,omitempty"`
	License             string            `json:"license,omitempty"`
	Timezone            string            `json:"timezone,omitempty"`
	Status              string            `json:"status"`
	Exists              bool
}

// HasChanged diff's the two items and returns true if there have been any changes
func (r *RDSInstance) HasChanged(or *RDSInstance) bool {
	if r.Size != or.Size {
		return true
	}

	if r.EngineVersion != or.EngineVersion {
		return true
	}

	if r.Port != nil && or.Port != nil {
		if *r.Port != *or.Port {
			return true
		}
	}

	if r.StorageSize != nil && or.StorageSize != nil {
		if *r.StorageSize != *or.StorageSize {
			return true
		}
	}

	if r.StorageIops != nil && or.StorageIops != nil {
		if *r.StorageIops != *or.StorageIops {
			return true
		}
	}

	if r.StorageType != or.StorageType {
		return true
	}

	if r.MultiAZ != or.MultiAZ {
		return true
	}

	if r.PromotionTier != nil && or.PromotionTier != nil {
		if *r.PromotionTier != *or.PromotionTier {
			return true
		}
	}

	if r.AutoUpgrade != or.AutoUpgrade {
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

	if r.DatabasePassword != or.DatabasePassword {
		return true
	}

	if r.Public != or.Public {
		return true
	}

	if reflect.DeepEqual(r.SecurityGroups, or.SecurityGroups) != true {
		return true
	}

	return !reflect.DeepEqual(r.Networks, or.Networks)
}

/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package output

import "reflect"

// RDSInstance ...
type RDSInstance struct {
	ProviderType        string   `json:"_type"`
	DatacenterRegion    string   `json:"datacenter_region"`
	DatacenterSecret    string   `json:"datacenter_secret"`
	DatacenterToken     string   `json:"datacenter_token"`
	VPCID               string   `json:"vpc_id"`
	Name                string   `json:"name"`
	Size                string   `json:"size"`
	Engine              string   `json:"engine"`
	EngineVersion       string   `json:"engine_version"`
	Port                int64    `json:"port"`
	Cluster             string   `json:"cluster"`
	Public              bool     `json:"public"`
	Endpoint            string   `json:"endpoint"`
	HotStandby          bool     `json:"hot_standby"`
	PromotionTier       int64    `json:"promotion_tier"`
	StorageType         string   `json:"storage_type"`
	StorageSize         int64    `json:"storage_size"`
	StorageIops         int64    `json:"storage_iops"`
	AvailabilityZone    string   `json:"availability_zone"`
	SecurityGroups      []string `json:"security_groups"`
	SecurityGroupAWSIDs []string `json:"security_group_aws_ids"`
	Networks            []string `json:"networks"`
	NetworkAWSIDs       []string `json:"network_aws_ids"`
	DatabaseName        string   `json:"database_name"`
	DatabaseUsername    string   `json:"database_username"`
	DatabasePassword    string   `json:"database_password"`
	AutoUpgrade         bool     `json:"auto_upgrade"`
	BackupRetention     int64    `json:"backup_retention"`
	BackupWindow        string   `json:"backup_window"`
	MaintenanceWindow   string   `json:"maintenance_window"`
	FinalSnapshot       bool     `json:"final_snapshot"`
	ReplicationSource   string   `json:"replication_source"`
	License             string   `json:"license"`
	Timezone            string   `json:"timezone"`
	Status              string   `json:"status"`
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

	if r.Port != or.Port {
		return true
	}

	if r.StorageSize != or.StorageSize {
		return true
	}

	if r.StorageIops != or.StorageIops {
		return true
	}

	if r.StorageType != or.StorageType {
		return true
	}

	if r.HotStandby != or.HotStandby {
		return true
	}

	if r.PromotionTier != or.PromotionTier {
		return true
	}

	if r.AutoUpgrade != or.AutoUpgrade {
		return true
	}

	if r.BackupRetention != or.BackupRetention {
		return true
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

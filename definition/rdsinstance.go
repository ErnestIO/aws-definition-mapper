/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package definition

import (
	"errors"
	"fmt"
	"unicode"
)

type RDSStorage struct {
	Type string `json:"type"`
	Size int64  `json:"size"`
	Iops int64  `json:"iops"`
}

// RDSInstance ...
type RDSInstance struct {
	Name              string     `json:"name"`
	Size              string     `json:"size"` //////////////
	Engine            string     `json:"engine"`
	EngineVersion     string     `json:"engine_version"`
	Port              int64      `json:"port"`
	Cluster           string     `json:"cluster"` //////////////
	Public            bool       `json:"public"`
	HotStandby        bool       `json:"hot_standby"`
	PromotionTier     int64      `json:"promotion_tier"` //////////////
	Storage           RDSStorage `json:"storage"`        //////////////
	AvailabilityZone  string     `json:"availability_zone"`
	SecurityGroups    []string   `json:"security_groups"`
	Networks          []string   `json:"networks"`
	DatabaseName      string     `json:"database_name"`
	DatabaseUsername  string     `json:"database_username"`
	DatabasePassword  string     `json:"database_password"`
	AutoUpgrade       bool       `json:"auto_upgrade"`
	Backups           RDSBackup  `json:"backups"`
	MaintenanceWindow string     `json:"maintenance_window"`
	FinalSnapshot     bool       `json:"final_snapshot"`
	ReplicationSource string     `json:"replication_source"`
	License           string     `json:"license"` //////////////
	Timezone          string     `json:"timezone"`
}

// Validate the rds cluster
func (r *RDSInstance) Validate(networks []Network, securitygroups []SecurityGroup, clusters []RDSCluster) error {
	if r.Name == "" {
		return errors.New("RDS Instance name should not be null")
	}

	if len(r.Name) > 255 {
		return errors.New("RDS Instance name should not exceed 255 characters")
	}

	if r.Size == "" {
		return errors.New("RDS Instance size should not be null")
	}

	if r.Size[:2] != "db." {
		return errors.New("RDS Instance size should be a valid resource size. i.e. 'db.r3.large'")
	}

	if r.Engine == "" {
		return errors.New("RDS Instance engine type should not be null")
	}

	cluster := findRDSCluster(clusters, r.Cluster)
	if r.Cluster != "" && cluster != nil {
		return fmt.Errorf("RDS Instance cluster identifier '%s' does not exist", r.Cluster)
	}

	if cluster != nil && cluster.Engine != "" && r.Engine != cluster.Engine {
		return fmt.Errorf("RDS Instance engine type must be the same as specified on the cluster")
	}

	if cluster != nil && cluster.EngineVersion != "" && r.EngineVersion != cluster.EngineVersion {
		return fmt.Errorf("RDS Instance engine version must be the same as specified on the cluster")
	}

	if cluster != nil && cluster.Port != 0 && r.Port != cluster.Port {
		return fmt.Errorf("RDS Instance port must be the same as specified on the cluster")
	}

	if cluster.Engine == "aurora" || r.Engine == "aurora" {
		if r.Storage.Type != "" || r.Storage.Size > 0 || r.Storage.Iops > 0 {
			return errors.New("RDS Instance storage options cannot be set if the engine type is 'aurora'")
		}
	}

	if r.Storage.Type != "" {
		if r.Storage.Type != "standard" && r.Storage.Type != "gp2" && r.Storage.Type != "io1" {
			return errors.New("RDS Instance storage type must be either 'standard', 'gp2' or 'io1'")
		}
	}

	if r.Storage.Size < 5 || r.Storage.Size > 6144 {
		return errors.New("RDS Instance storage size must be between 5 - 6144 GB")
	}

	if r.PromotionTier < 0 || r.PromotionTier > 15 {
		return errors.New("RDS Instance promotion tier should be between 0 - 15")
	}

	if r.AvailabilityZone != "" && r.HotStandby {
		return errors.New("RDS Instance cannot specify both an availability zone and a standby instance")
	}

	if r.DatabaseName == "" {
		return errors.New("RDS Instance database name should not be null")
	}

	if len(r.DatabaseName) > 64 {
		return errors.New("RDS Instance database name should not exceed 64 characters")
	}

	for _, c := range r.DatabaseName {
		if unicode.IsLetter(c) != true && unicode.IsNumber(c) != true {
			return errors.New("RDS Instance database name can only contain alphanumeric characters")
		}
	}

	if r.DatabaseUsername == "" {
		return errors.New("RDS Instance database username should not be null")
	}

	if len(r.DatabaseUsername) > 16 {
		return errors.New("RDS Instance database username should not exceed 16 characters")
	}

	if r.DatabasePassword == "" {
		return errors.New("RDS Instance database password should not be null")
	}

	if len(r.DatabasePassword) < 8 || len(r.DatabasePassword) > 41 {
		return errors.New("RDS Instance database password should be between 8 and 41 characters")
	}

	for _, c := range r.DatabasePassword {
		if unicode.IsSymbol(c) || unicode.IsMark(c) {
			return fmt.Errorf("RDS Instance database password contains an offending character: '%c'", c)
		}
	}

	if r.Port < 1150 || r.Port > 65535 {
		return errors.New("RDS Instance port number should be between 1150 and 65535")
	}

	if r.Backups.Retention < 1 || r.Backups.Retention > 35 {
		return errors.New("RDS Instance backup retention should be between 1 and 35 days")
	}

	if bwerr := validateTimeWindow(r.Backups.Window); bwerr != nil {
		return fmt.Errorf("RDS Instance backup window: %s", bwerr.Error())
	}

	if mwerr := validateTimeWindow(r.MaintenanceWindow); mwerr != nil {
		return fmt.Errorf("RDS Instance maintenance window: %s", mwerr.Error())
	}

	for _, nw := range r.Networks {
		if isNetwork(networks, nw) != true {
			return fmt.Errorf("RDS Instance network '%s' does not exist", nw)
		}
	}

	for _, sg := range r.SecurityGroups {
		if isSecurityGroup(securitygroups, sg) != true {
			return fmt.Errorf("RDS Instance security group '%s' does not exist", sg)
		}
	}

	return nil
}

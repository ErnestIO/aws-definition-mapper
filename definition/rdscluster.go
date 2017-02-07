/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package definition

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

// RDSBackup ...
type RDSBackup struct {
	Window    string `json:"window"`
	Retention *int64 `json:"retention"`
}

// RDSCluster ...
type RDSCluster struct {
	Name              string    `json:"name"`
	Engine            string    `json:"engine"`
	EngineVersion     string    `json:"engine_version"`
	Port              *int64    `json:"port"`
	AvailabilityZones []string  `json:"availability_zones"`
	SecurityGroups    []string  `json:"security_groups"`
	Networks          []string  `json:"networks"`
	DatabaseName      string    `json:"database_name"`
	DatabaseUsername  string    `json:"database_username"`
	DatabasePassword  string    `json:"database_password"`
	Backups           RDSBackup `json:"backups"`
	MaintenanceWindow string    `json:"maintenance_window"`
	ReplicationSource string    `json:"replication_source"`
	FinalSnapshot     bool      `json:"final_snapshot"`
}

// Validate the rds cluster
func (r *RDSCluster) Validate(networks []Network, securitygroups []SecurityGroup) error {
	if r.Name == "" {
		return errors.New("RDS Cluster name should not be null")
	}

	if len(r.Name) > 255 {
		return errors.New("RDS Cluster name should not exceed 255 characters")
	}

	if r.Engine == "" {
		return errors.New("RDS Cluster engine type should not be null")
	}

	if r.ReplicationSource != "" {
		if len(r.ReplicationSource) < 12 || r.ReplicationSource[:12] != "arn:aws:rds:" {
			return errors.New("RDS Cluster replication source should be a valid amazon resource name (ARN), i.e. 'arn:aws:rds:us-east-1:123456789012:cluster:my-aurora-cluster'")
		}
	}

	if r.DatabaseName == "" {
		return errors.New("RDS Cluster database name should not be null")
	}

	if len(r.DatabaseName) > 64 {
		return errors.New("RDS Cluster database name should not exceed 64 characters")
	}

	for _, c := range r.DatabaseName {
		if unicode.IsLetter(c) != true && unicode.IsNumber(c) != true {
			return errors.New("RDS Cluster database name can only contain alphanumeric characters")
		}
	}

	if r.DatabaseUsername == "" {
		return errors.New("RDS Cluster database username should not be null")
	}

	if len(r.DatabaseUsername) > 16 {
		return errors.New("RDS Cluster database username should not exceed 16 characters")
	}

	if r.DatabasePassword == "" {
		return errors.New("RDS Cluster database password should not be null")
	}

	if len(r.DatabasePassword) < 8 || len(r.DatabasePassword) > 41 {
		return errors.New("RDS Cluster database password should be between 8 and 41 characters")
	}

	for _, c := range r.DatabasePassword {
		if unicode.IsSymbol(c) || unicode.IsMark(c) {
			return fmt.Errorf("RDS Cluster database password contains an offending character: '%c'", c)
		}
	}

	if r.Port != nil {
		if *r.Port < 1150 || *r.Port > 65535 {
			return errors.New("RDS Cluster port number should be between 1150 and 65535")
		}
	}

	if r.Backups.Retention != nil {
		if *r.Backups.Retention < 1 || *r.Backups.Retention > 35 {
			return errors.New("RDS Cluster backup retention should be between 1 and 35 days")
		}
	}

	if r.Backups.Window != "" {
		parts := strings.Split(r.Backups.Window, "-")

		err := validateTimeFormat(parts[0])
		if err != nil {
			return errors.New("RDS Cluster backup window: " + err.Error())
		}

		err = validateTimeFormat(parts[1])
		if err != nil {
			return errors.New("RDS Cluster backup window: " + err.Error())
		}
	}

	if mwerr := validateTimeWindow(r.MaintenanceWindow); r.MaintenanceWindow != "" && mwerr != nil {
		return fmt.Errorf("RDS Cluster maintenance window: %s", mwerr.Error())
	}

	for _, nw := range r.Networks {
		if isNetwork(networks, nw) != true {
			return fmt.Errorf("RDS Cluster network '%s' does not exist", nw)
		}
	}

	for _, sg := range r.SecurityGroups {
		if isSecurityGroup(securitygroups, sg) != true {
			return fmt.Errorf("RDS Cluster security group '%s' does not exist", sg)
		}
	}

	if len(r.Networks) > 0 || len(r.AvailabilityZones) > 0 {
		if len(r.Networks) < 2 || r.meetsNetworkAZRequirement(networks) != true {
			return errors.New("RDS Cluster should specify at least two networks in different availability zones if no cluster availability zone is specified")
		}

		for _, az := range r.AvailabilityZones {
			if r.hasAvailabilityZone(networks, az) != true {
				return fmt.Errorf("RDS Cluster has no network specified for the availability zone '%s'", az)
			}
		}
	}

	return nil
}

func (r *RDSCluster) meetsNetworkAZRequirement(networks []Network) bool {
	var azs []string
	for _, cn := range r.Networks {
		for _, n := range networks {
			if n.Name == cn {
				azs = appendUnique(azs, n.AvailabilityZone)
			}
		}
	}

	if len(azs) > 1 {
		return true
	}

	return false
}

func (r *RDSCluster) hasAvailabilityZone(networks []Network, az string) bool {
	for _, cn := range r.Networks {
		for _, n := range networks {
			if n.Name == cn && n.AvailabilityZone == az {
				return true
			}
		}
	}
	return false
}

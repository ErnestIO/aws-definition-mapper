/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package definition

import (
	"errors"
	"fmt"
	"unicode"
)

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

	if *r.Backups.Retention < 1 || *r.Backups.Retention > 35 {
		return errors.New("RDS Cluster backup retention should be between 1 and 35 days")
	}

	if bwerr := validateTimeWindow(r.Backups.Window); r.Backups.Window != "" && bwerr != nil {
		return fmt.Errorf("RDS Cluster backup window: %s", bwerr.Error())
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

	return nil
}

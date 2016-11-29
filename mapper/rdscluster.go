/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package mapper

import (
	"github.com/ernestio/aws-definition-mapper/definition"
	"github.com/ernestio/aws-definition-mapper/output"
)

// MapRDSClusters : Maps the rds clusters for the input payload on a ernest internal format
func MapRDSClusters(d definition.Definition) []output.RDSCluster {
	var clusters []output.RDSCluster

	for _, cluster := range d.RDSClusters {

		clusters = append(clusters, output.RDSCluster{
			Name:                d.GeneratedName() + cluster.Name,
			Engine:              cluster.Engine,
			EngineVersion:       cluster.EngineVersion,
			Port:                cluster.Port,
			AvailabilityZones:   cluster.AvailabilityZones,
			SecurityGroups:      cluster.SecurityGroups,
			SecurityGroupAWSIDs: mapRDSSecurityGroupIDs(cluster.SecurityGroups),
			Networks:            cluster.Networks,
			NetworkAWSIDs:       mapRDSNetworkIDs(cluster.Networks),
			DatabaseName:        cluster.DatabaseName,
			DatabaseUsername:    cluster.DatabaseUsername,
			DatabasePassword:    cluster.DatabasePassword,
			BackupRetention:     cluster.Backups.Retention,
			BackupWindow:        cluster.Backups.Window,
			MaintenanceWindow:   cluster.MaintenanceWindow,
			ReplicationSource:   cluster.ReplicationSource,
			FinalSnapshot:       cluster.FinalSnapshot,
			ProviderType:        "$(datacenters.items.0.type)",
			VpcID:               "$(vpcs.items.0.vpc_id)",
			DatacenterSecret:    "$(datacenters.items.0.secret)",
			DatacenterToken:     "$(datacenters.items.0.token)",
			DatacenterRegion:    "$(datacenters.items.0.region)",
		})

	}
	return clusters
}

func mapRDSSecurityGroupIDs(sgs []string) []string {
	var ids []string

	for _, sg := range sgs {
		ids = append(ids, `$(firewalls.items.#[name="`+sg+`"].security_group_aws_id)`)
	}

	return ids
}

func mapRDSNetworkIDs(nws []string) []string {
	var ids []string

	for _, nw := range nws {
		ids = append(ids, `$(networks.items.#[name="`+nw+`"].network_aws_id)`)
	}

	return ids
}

func int64p(i int64) *int64 {
	return &i
}

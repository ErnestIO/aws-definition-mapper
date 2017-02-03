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
		var sgroups []string
		var networks []string

		for _, sg := range cluster.SecurityGroups {
			sgroups = append(sgroups, d.GeneratedName()+sg)
		}

		for _, nw := range cluster.Networks {
			networks = append(networks, d.GeneratedName()+nw)
		}

		name := d.GeneratedName() + cluster.Name

		clusters = append(clusters, output.RDSCluster{
			Name:                name,
			Engine:              cluster.Engine,
			EngineVersion:       cluster.EngineVersion,
			Port:                cluster.Port,
			AvailabilityZones:   cluster.AvailabilityZones,
			SecurityGroups:      cluster.SecurityGroups,
			SecurityGroupAWSIDs: mapRDSSecurityGroupIDs(sgroups),
			Networks:            cluster.Networks,
			NetworkAWSIDs:       mapRDSNetworkIDs(networks),
			DatabaseName:        cluster.DatabaseName,
			DatabaseUsername:    cluster.DatabaseUsername,
			DatabasePassword:    cluster.DatabasePassword,
			BackupRetention:     cluster.Backups.Retention,
			BackupWindow:        cluster.Backups.Window,
			MaintenanceWindow:   cluster.MaintenanceWindow,
			ReplicationSource:   cluster.ReplicationSource,
			FinalSnapshot:       cluster.FinalSnapshot,
			Tags:                mapTagsServiceOnly(d.Name),
			ProviderType:        "$(datacenters.items.0.type)",
			VpcID:               "$(vpcs.items.0.vpc_id)",
			SecretAccessKey:     "$(datacenters.items.0.aws_secret_access_key)",
			AccessKeyID:         "$(datacenters.items.0.aws_access_key_id)",
			DatacenterRegion:    "$(datacenters.items.0.region)",
		})

	}
	return clusters
}

// MapDefinitionRDSClusters : Maps the rds clusters for the internal ernest format to the input definition format
func MapDefinitionRDSClusters(m *output.FSMMessage) []definition.RDSCluster {
	var clusters []definition.RDSCluster

	prefix := m.Datacenters.Items[0].Name + "-" + m.ServiceName + "-"

	for _, cluster := range m.RDSClusters.Items {
		sgroups := ComponentNamesFromIDs(m.Firewalls.Items, cluster.SecurityGroupAWSIDs)
		subnets := ComponentNamesFromIDs(m.Networks.Items, cluster.NetworkAWSIDs)

		c := definition.RDSCluster{
			Name:              ShortName(cluster.Name, prefix),
			Engine:            cluster.Engine,
			EngineVersion:     cluster.EngineVersion,
			Port:              cluster.Port,
			AvailabilityZones: cluster.AvailabilityZones,
			SecurityGroups:    ShortNames(sgroups, prefix),
			Networks:          ShortNames(subnets, prefix),
			DatabaseName:      cluster.DatabaseName,
			DatabaseUsername:  cluster.DatabaseUsername,
			DatabasePassword:  cluster.DatabasePassword,
			MaintenanceWindow: cluster.MaintenanceWindow,
			ReplicationSource: cluster.ReplicationSource,
			FinalSnapshot:     cluster.FinalSnapshot,
		}

		c.Backups.Retention = cluster.BackupRetention
		c.Backups.Window = cluster.BackupWindow

		clusters = append(clusters, c)
	}

	return clusters
}

// UpdateRDSClusterValues corrects missing values after an import
func UpdateRDSClusterValues(m *output.FSMMessage) {
	for i := 0; i < len(m.RDSClusters.Items); i++ {
		m.RDSClusters.Items[i].ProviderType = "$(datacenters.items.0.type)"
		m.RDSClusters.Items[i].AccessKeyID = "$(datacenters.items.0.aws_access_key_id)"
		m.RDSClusters.Items[i].SecretAccessKey = "$(datacenters.items.0.aws_secret_access_key)"
		m.RDSClusters.Items[i].DatacenterRegion = "$(datacenters.items.0.region)"
		m.RDSClusters.Items[i].VpcID = "$(vpcs.items.0.vpc_id)"
		m.RDSClusters.Items[i].SecurityGroups = ComponentNamesFromIDs(m.Firewalls.Items, m.RDSClusters.Items[i].SecurityGroupAWSIDs)
		m.RDSClusters.Items[i].Networks = ComponentNamesFromIDs(m.Networks.Items, m.RDSClusters.Items[i].NetworkAWSIDs)
	}
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

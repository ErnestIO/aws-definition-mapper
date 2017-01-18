/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package mapper

import (
	"github.com/ernestio/aws-definition-mapper/definition"
	"github.com/ernestio/aws-definition-mapper/output"
)

// MapRDSInstances : Maps the rds instances for the input payload on a ernest internal format
func MapRDSInstances(d definition.Definition) []output.RDSInstance {
	var instances []output.RDSInstance

	for _, instance := range d.RDSInstances {
		var sgroups []string
		var networks []string

		for _, sg := range instance.SecurityGroups {
			sgroups = append(sgroups, d.GeneratedName()+sg)
		}

		for _, nw := range instance.Networks {
			networks = append(networks, d.GeneratedName()+nw)
		}

		name := d.GeneratedName() + instance.Name

		instance := output.RDSInstance{
			Name:                name,
			Size:                instance.Size,
			Engine:              instance.Engine,
			EngineVersion:       instance.EngineVersion,
			Port:                instance.Port,
			Cluster:             instance.Cluster,
			Public:              instance.Public,
			MultiAZ:             instance.MultiAZ,
			PromotionTier:       instance.PromotionTier,
			StorageType:         instance.Storage.Type,
			StorageSize:         instance.Storage.Size,
			StorageIops:         instance.Storage.Iops,
			AvailabilityZone:    instance.AvailabilityZone,
			SecurityGroups:      instance.SecurityGroups,
			SecurityGroupAWSIDs: mapRDSSecurityGroupIDs(sgroups),
			Networks:            instance.Networks,
			NetworkAWSIDs:       mapRDSNetworkIDs(networks),
			DatabaseName:        instance.DatabaseName,
			DatabaseUsername:    instance.DatabaseUsername,
			DatabasePassword:    instance.DatabasePassword,
			AutoUpgrade:         instance.AutoUpgrade,
			BackupRetention:     instance.Backups.Retention,
			BackupWindow:        instance.Backups.Window,
			MaintenanceWindow:   instance.MaintenanceWindow,
			ReplicationSource:   instance.ReplicationSource,
			FinalSnapshot:       instance.FinalSnapshot,
			License:             instance.License,
			Timezone:            instance.Timezone,
			Tags:                mapTagsServiceOnly(d.Name),
			ProviderType:        "$(datacenters.items.0.type)",
			VpcID:               "$(vpcs.items.0.vpc_id)",
			SecretAccessKey:     "$(datacenters.items.0.aws_secret_access_key)",
			AccessKeyID:         "$(datacenters.items.0.aws_access_key_id)",
			DatacenterRegion:    "$(datacenters.items.0.region)",
		}

		cluster := d.FindRDSCluster(instance.Cluster)
		if cluster != nil {
			instance.Engine = cluster.Engine
			instance.Cluster = d.GeneratedName() + instance.Cluster
		}

		instances = append(instances, instance)
	}
	return instances
}

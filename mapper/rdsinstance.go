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

		instances = append(instances, output.RDSInstance{
			Name:                d.GeneratedName() + instance.Name,
			Engine:              instance.Engine,
			EngineVersion:       instance.EngineVersion,
			Port:                instance.Port,
			Cluster:             d.GeneratedName() + instance.Cluster,
			Public:              instance.Public,
			HotStandby:          instance.HotStandby,
			PromotionTier:       instance.PromotionTier,
			StorageType:         instance.Storage.Type,
			StorageSize:         instance.Storage.Size,
			StorageIops:         instance.Storage.Iops,
			AvailabilityZone:    instance.AvailabilityZone,
			SecurityGroups:      instance.SecurityGroups,
			SecurityGroupAWSIDs: mapRDSSecurityGroupIDs(instance.SecurityGroups),
			Networks:            instance.Networks,
			NetworkAWSIDs:       mapRDSNetworkIDs(instance.Networks),
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
			ProviderType:        "$(datacenters.items.0.type)",
			DatacenterSecret:    "$(datacenters.items.0.secret)",
			DatacenterToken:     "$(datacenters.items.0.token)",
			DatacenterRegion:    "$(datacenters.items.0.region)",
		})

	}
	return instances
}

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

		i := output.RDSInstance{
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
			SecurityGroups:      sgroups,
			SecurityGroupAWSIDs: mapRDSSecurityGroupIDs(sgroups),
			Networks:            networks,
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
			i.Engine = cluster.Engine
			i.Cluster = d.GeneratedName() + instance.Cluster
		}

		instances = append(instances, i)
	}
	return instances
}

// MapDefinitionRDSInstances : Maps the rds instances from the internal format to the input definition format
func MapDefinitionRDSInstances(m *output.FSMMessage) []definition.RDSInstance {
	var instances []definition.RDSInstance

	prefix := m.Datacenters.Items[0].Name + "-" + m.ServiceName + "-"

	for _, instance := range m.RDSInstances.Items {
		sgroups := ComponentNamesFromIDs(m.Firewalls.Items, instance.SecurityGroupAWSIDs)
		subnets := ComponentNamesFromIDs(m.Networks.Items, instance.NetworkAWSIDs)

		i := definition.RDSInstance{
			Name:              ShortName(instance.Name, prefix),
			Size:              instance.Size,
			Engine:            instance.Engine,
			EngineVersion:     instance.EngineVersion,
			Port:              instance.Port,
			Cluster:           ShortName(instance.Cluster, prefix),
			Public:            instance.Public,
			MultiAZ:           instance.MultiAZ,
			PromotionTier:     instance.PromotionTier,
			AvailabilityZone:  instance.AvailabilityZone,
			SecurityGroups:    ShortNames(sgroups, prefix),
			Networks:          ShortNames(subnets, prefix),
			DatabaseName:      instance.DatabaseName,
			DatabaseUsername:  instance.DatabaseUsername,
			DatabasePassword:  instance.DatabasePassword,
			AutoUpgrade:       instance.AutoUpgrade,
			MaintenanceWindow: instance.MaintenanceWindow,
			ReplicationSource: instance.ReplicationSource,
			FinalSnapshot:     instance.FinalSnapshot,
			License:           instance.License,
			Timezone:          instance.Timezone,
		}

		i.Storage.Type = instance.StorageType
		i.Storage.Size = instance.StorageSize
		i.Storage.Iops = instance.StorageIops
		i.Backups.Retention = instance.BackupRetention
		i.Backups.Window = instance.BackupWindow

		if i.Storage.Type != "io1" {
			i.Storage.Iops = nil
		}

		instances = append(instances, i)
	}
	return instances
}

// UpdateRDSInstanceValues corrects missing values after an import
func UpdateRDSInstanceValues(m *output.FSMMessage) {
	for i := 0; i < len(m.RDSInstances.Items); i++ {
		m.RDSInstances.Items[i].ProviderType = "$(datacenters.items.0.type)"
		m.RDSInstances.Items[i].AccessKeyID = "$(datacenters.items.0.aws_access_key_id)"
		m.RDSInstances.Items[i].SecretAccessKey = "$(datacenters.items.0.aws_secret_access_key)"
		m.RDSInstances.Items[i].DatacenterRegion = "$(datacenters.items.0.region)"
		m.RDSInstances.Items[i].VpcID = "$(vpcs.items.0.vpc_id)"
		m.RDSInstances.Items[i].SecurityGroups = ComponentNamesFromIDs(m.Firewalls.Items, m.RDSInstances.Items[i].SecurityGroupAWSIDs)
		m.RDSInstances.Items[i].Networks = ComponentNamesFromIDs(m.Networks.Items, m.RDSInstances.Items[i].NetworkAWSIDs)
	}
}

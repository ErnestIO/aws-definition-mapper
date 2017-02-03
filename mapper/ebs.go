/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package mapper

import (
	"strconv"

	"github.com/ernestio/aws-definition-mapper/definition"
	"github.com/ernestio/aws-definition-mapper/output"
)

// MapEBSVolumes : Maps the ebs volumes from a given input payload.
func MapEBSVolumes(d definition.Definition) []output.EBSVolume {
	var volumes []output.EBSVolume

	for _, vol := range d.EBSVolumes {

		for i := 0; i < vol.Count; i++ {
			name := d.GeneratedName() + vol.Name + "-" + strconv.Itoa(i+1)

			volumes = append(volumes, output.EBSVolume{
				ProviderType:     "$(datacenters.items.0.type)",
				DatacenterName:   "$(datacenters.items.0.name)",
				DatacenterType:   "$(datacenters.items.0.type)",
				AccessKeyID:      "$(datacenters.items.0.aws_access_key_id)",
				SecretAccessKey:  "$(datacenters.items.0.aws_secret_access_key)",
				DatacenterRegion: "$(datacenters.items.0.region)",
				VPCID:            "$(vpcs.items.0.vpc_id)",
				Name:             name,
				AvailabilityZone: vol.AvailabilityZone,
				VolumeType:       vol.Type,
				Size:             vol.Size,
				Iops:             vol.Iops,
				Encrypted:        vol.Encrypted,
				EncryptionKeyID:  vol.EncryptionKeyID,
				Tags:             mapEBSTags(name, d.Name, vol.Name),
			})
		}
	}

	return volumes
}

// MapDefinitionEBSVolumes : Maps output ebs volumes into a definition defined ebs volumes
func MapDefinitionEBSVolumes(m *output.FSMMessage) []definition.EBSVolume {
	var vols []definition.EBSVolume

	for _, vg := range ComponentGroups(m.EBSVolumes.Items, "ernest.volume_group") {
		vs := ComponentsByTag(m.EBSVolumes.Items, "ernest.volume_group", vg)
		firstVol := vs[0].(output.EBSVolume)

		vols = append(vols, definition.EBSVolume{
			Name:             vg,
			Type:             firstVol.VolumeType,
			Size:             firstVol.Size,
			Iops:             firstVol.Iops,
			AvailabilityZone: firstVol.AvailabilityZone,
			Encrypted:        firstVol.Encrypted,
			EncryptionKeyID:  firstVol.EncryptionKeyID,
			Count:            len(vs),
		})

	}

	return vols
}

// UpdateEBSValues corrects missing values after an import
func UpdateEBSValues(m *output.FSMMessage) {
	for i := 0; i < len(m.EBSVolumes.Items); i++ {
		m.EBSVolumes.Items[i].ProviderType = "$(datacenters.items.0.type)"
		m.EBSVolumes.Items[i].DatacenterName = "$(datacenters.items.0.name)"
		m.EBSVolumes.Items[i].DatacenterType = "$(datacenters.items.0.type)"
		m.EBSVolumes.Items[i].AccessKeyID = "$(datacenters.items.0.aws_access_key_id)"
		m.EBSVolumes.Items[i].SecretAccessKey = "$(datacenters.items.0.aws_secret_access_key)"
		m.EBSVolumes.Items[i].DatacenterRegion = "$(datacenters.items.0.region)"
		m.EBSVolumes.Items[i].VPCID = "$(vpcs.items.0.vpc_id)"
	}
}

func mapEBSTags(name, service, volumeGroup string) map[string]string {
	tags := make(map[string]string)

	tags["Name"] = name
	tags["ernest.service"] = service
	tags["ernest.volume_group"] = volumeGroup

	return tags
}

/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package mapper

import (
	"github.com/ernestio/aws-definition-mapper/definition"
	"github.com/ernestio/aws-definition-mapper/output"
)

// MapEBSVolumes : Maps the ebs volumes from a given input payload.
func MapEBSVolumes(d definition.Definition) []output.EBSVolume {
	var volumes []output.EBSVolume

	for _, vol := range d.EBSVolumes {

		for i := 0; i < vol.Count; i++ {
			volumes = append(volumes, output.EBSVolume{
				ProviderType:     "$(datacenters.items.0.type)",
				DatacenterName:   "$(datacenters.items.0.name)",
				DatacenterType:   "$(datacenters.items.0.type)",
				AccessKeyID:      "$(datacenters.items.0.aws_access_key_id)",
				SecretAccessKey:  "$(datacenters.items.0.aws_secret_access_key)",
				DatacenterRegion: "$(datacenters.items.0.region)",
				VPCID:            "$(vpcs.items.0.vpc_id)",
				Name:             d.GeneratedName() + vol.Name,
				AvailabilityZone: vol.AvailabilityZone,
				VolumeType:       vol.Type,
				Size:             vol.Size,
				Iops:             vol.Iops,
				Encrypted:        vol.Encrypted,
				EncryptionKeyID:  vol.EncryptionKeyID,
			})
		}
	}

	return volumes
}

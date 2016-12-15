/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package mapper

import (
	"net"
	"strconv"

	"github.com/ernestio/aws-definition-mapper/definition"
	"github.com/ernestio/aws-definition-mapper/output"
)

// MapInstances : Maps the instances for the input payload on a ernest internal format
func MapInstances(d definition.Definition) []output.Instance {
	var instances []output.Instance

	for _, instance := range d.Instances {
		ip := make(net.IP, net.IPv4len)
		copy(ip, instance.StartIP.To4())

		for i := 0; i < instance.Count; i++ {
			var sgroups []string
			for _, sg := range instance.SecurityGroups {
				sgroups = append(sgroups, d.GeneratedName()+sg)
			}

			newInstance := output.Instance{
				Name:                d.GeneratedName() + instance.Name + "-" + strconv.Itoa(i+1),
				Type:                instance.Type,
				Image:               instance.Image,
				Network:             d.GeneratedName() + instance.Network,
				NetworkAWSID:        `$(networks.items.#[name="` + d.GeneratedName() + instance.Network + `"].network_aws_id)`,
				IP:                  net.ParseIP(ip.String()),
				KeyPair:             instance.KeyPair,
				AssignElasticIP:     instance.ElasticIP,
				SecurityGroups:      sgroups,
				SecurityGroupAWSIDs: mapInstanceSecurityGroupIDs(sgroups),
				UserData:            instance.UserData,
				ProviderType:        "$(datacenters.items.0.type)",
				DatacenterType:      "$(datacenters.items.0.type)",
				DatacenterName:      "$(datacenters.items.0.name)",
				AccessKeyID:         "$(datacenters.items.0.aws_access_key_id)",
				SecretAccessKey:     "$(datacenters.items.0.aws_secret_access_key)",
				DatacenterRegion:    "$(datacenters.items.0.region)",
				VpcID:               "$(vpcs.items.0.vpc_id)",
			}

			for _, vol := range instance.Volumes {
				vname := d.GeneratedName() + vol.Volume + "-" + strconv.Itoa(i+1)
				v := output.InstanceVolume{
					Volume:      vname,
					VolumeAWSID: `$(ebs_volumes.items.#[name="` + vname + `"].volume_aws_id)`,
					Device:      vol.Device,
				}
				newInstance.Volumes = append(newInstance.Volumes, v)
			}

			instances = append(instances, newInstance)

			// Increment IP address
			ip[3]++
		}
	}
	return instances
}

func mapInstanceSecurityGroupIDs(sgs []string) []string {
	var ids []string

	for _, sg := range sgs {
		ids = append(ids, `$(firewalls.items.#[name="`+sg+`"].security_group_aws_id)`)
	}

	return ids
}

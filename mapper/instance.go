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

			name := d.GeneratedName() + instance.Name + "-" + strconv.Itoa(i+1)

			newInstance := output.Instance{
				Name:                name,
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
				Tags:                mapInstanceTags(name, d.Name, instance.Name),
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

// MapDefinitionInstances : Maps output instances into a definition defined instances
func MapDefinitionInstances(m *output.FSMMessage) []definition.Instance {
	var instances []definition.Instance

	prefix := m.Datacenters.Items[0].Name + "-" + m.ServiceName + "-"

	for _, ig := range ComponentGroups(m.Instances.Items, "ernest.instance_group") {
		is := ComponentsByTag(m.Instances.Items, "ernest.instance_group", ig)

		if len(is) < 1 {
			continue
		}

		for i := 0; i < len(is); i++ {
			instance, ok := is[i].(output.Instance)
			if ok {
				nw := ComponentByID(m.Networks.Items, instance.NetworkAWSID)
				if nw != nil {
					instance.Network = nw.ComponentName()
				}

				instance.SecurityGroups = ComponentNamesFromIDs(m.Firewalls.Items, instance.SecurityGroupAWSIDs)

				for x := 0; x < len(instance.Volumes); x++ {
					v := ComponentByID(m.EBSVolumes.Items, instance.Volumes[x].VolumeAWSID)
					if v != nil {
						instance.Volumes[x].Volume = v.ComponentName()
					}
				}
			}
		}

		firstInstance := is[0].(output.Instance)
		elastic := false

		if firstInstance.ElasticIP != "" {
			elastic = true
		}

		network := ComponentByID(m.Networks.Items, firstInstance.NetworkAWSID)
		sgroups := ComponentNamesFromIDs(m.Firewalls.Items, firstInstance.SecurityGroupAWSIDs)

		instance := definition.Instance{
			Name:           ig,
			Type:           firstInstance.Type,
			Image:          firstInstance.Image,
			Network:        ShortName(network.ComponentName(), prefix),
			StartIP:        firstInstance.IP,
			KeyPair:        firstInstance.KeyPair,
			SecurityGroups: ShortNames(sgroups, prefix),
			ElasticIP:      elastic,
			Count:          len(is),
		}

		for _, vol := range firstInstance.Volumes {
			vc := ComponentByID(m.EBSVolumes.Items, vol.VolumeAWSID)
			if vc == nil {
				continue
			}

			vtags := vc.GetTags()

			instance.Volumes = append(instance.Volumes, definition.InstanceVolume{
				Device: vol.Device,
				Volume: vtags["ernest.volume_group"],
			})
		}

		instances = append(instances, instance)

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

func mapInstanceTags(name, service, instanceGroup string) map[string]string {
	tags := make(map[string]string)

	tags["Name"] = name
	tags["ernest.service"] = service
	tags["ernest.instance_group"] = instanceGroup

	return tags
}

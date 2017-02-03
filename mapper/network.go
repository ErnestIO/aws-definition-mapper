/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package mapper

import (
	"github.com/ernestio/aws-definition-mapper/definition"
	"github.com/ernestio/aws-definition-mapper/output"
)

// MapNetworks : Maps the networks from a given input payload.
func MapNetworks(d definition.Definition) []output.Network {
	var networks []output.Network

	for _, network := range d.Networks {
		name := d.GeneratedName() + network.Name

		n := output.Network{
			ProviderType:     "$(datacenters.items.0.type)",
			Name:             name,
			Subnet:           network.Subnet,
			IsPublic:         network.Public,
			AvailabilityZone: network.AvailabilityZone,
			Tags:             mapNetworkTags(name, d.Name, network.NatGateway),
			DatacenterType:   "$(datacenters.items.0.type)",
			DatacenterName:   "$(datacenters.items.0.name)",
			SecretAccessKey:  "$(datacenters.items.0.aws_secret_access_key)",
			AccessKeyID:      "$(datacenters.items.0.aws_access_key_id)",
			DatacenterRegion: "$(datacenters.items.0.region)",
			VpcID:            "$(vpcs.items.0.vpc_id)",
		}

		networks = append(networks, n)
	}

	return networks
}

// MapDefinitionNetworks : Maps output networks into a definition defined networks
func MapDefinitionNetworks(m *output.FSMMessage) []definition.Network {
	var nws []definition.Network

	prefix := m.Datacenters.Items[0].Name + "-" + m.ServiceName + "-"

	for _, n := range m.Networks.Items {
		nws = append(nws, definition.Network{
			Name:             ShortName(n.Name, prefix),
			Subnet:           n.Subnet,
			Public:           n.IsPublic,
			AvailabilityZone: n.AvailabilityZone,
			NatGateway:       n.Tags["ernest.nat_gateway"],
		})
	}

	return nws
}

// UpdateNetworkValues corrects missing values after an import
func UpdateNetworkValues(m *output.FSMMessage) {
	for i := 0; i < len(m.Networks.Items); i++ {
		m.Networks.Items[i].ProviderType = "$(datacenters.items.0.type)"
		m.Networks.Items[i].DatacenterName = "$(datacenters.items.0.name)"
		m.Networks.Items[i].DatacenterType = "$(datacenters.items.0.type)"
		m.Networks.Items[i].AccessKeyID = "$(datacenters.items.0.aws_access_key_id)"
		m.Networks.Items[i].SecretAccessKey = "$(datacenters.items.0.aws_secret_access_key)"
		m.Networks.Items[i].DatacenterRegion = "$(datacenters.items.0.region)"
		m.Networks.Items[i].VpcID = "$(vpcs.items.0.vpc_id)"
	}
}

func mapNetworkTags(name, service, gateway string) map[string]string {
	tags := make(map[string]string)

	tags["Name"] = name
	tags["ernest.service"] = service

	if gateway != "" {
		tags["ernest.nat_gateway"] = gateway
	}

	return tags
}

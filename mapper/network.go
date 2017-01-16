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
			Tags:             mapTags(network.Name, d.Name),
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

	for _, n := range m.Networks.Items {
		var gateway string

		for _, gw := range m.Nats.Items {
			for _, rn := range gw.RoutedNetworks {
				if rn == n.Name {
					gateway = gw.Name
				}
			}
		}

		nws = append(nws, definition.Network{
			Name:             n.Name,
			Subnet:           n.Subnet,
			Public:           n.IsPublic,
			AvailabilityZone: n.AvailabilityZone,
			NatGateway:       gateway,
		})
	}

	return nws
}

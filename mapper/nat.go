/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package mapper

import (
	"github.com/ernestio/aws-definition-mapper/definition"
	"github.com/ernestio/aws-definition-mapper/output"
)

// MapNats : Generates necessary nats rules for input networks
func MapNats(d definition.Definition) []output.Nat {
	var nats []output.Nat

	for _, ng := range d.NatGateways {
		nws := mapNetworkNames(d, ng.Name)
		name := d.GeneratedName() + ng.Name

		nats = append(nats, output.Nat{
			Name:                name,
			PublicNetwork:       d.GeneratedName() + ng.PublicNetwork,
			RoutedNetworks:      nws,
			PublicNetworkAWSID:  `$(networks.items.#[name="` + d.GeneratedName() + ng.PublicNetwork + `"].network_aws_id)`,
			RoutedNetworkAWSIDs: mapNatNetworkIDs(nws),
			Tags:                mapTags(name, d.Name),
			ProviderType:        "$(datacenters.items.0.type)",
			DatacenterType:      "$(datacenters.items.0.type)",
			DatacenterName:      "$(datacenters.items.0.name)",
			SecretAccessKey:     "$(datacenters.items.0.aws_secret_access_key)",
			AccessKeyID:         "$(datacenters.items.0.aws_access_key_id)",
			DatacenterRegion:    "$(datacenters.items.0.region)",
			VpcID:               "$(vpcs.items.0.vpc_id)",
		})
	}

	return nats
}

// MapDefinitionNats : Maps output nat gateways into a definition defined nat gateways
func MapDefinitionNats(m *output.FSMMessage) []definition.NatGateway {
	var nts []definition.NatGateway

	prefix := m.Datacenters.Items[0].Name + "-" + m.ServiceName + "-"

	for i := len(m.Nats.Items) - 1; i >= 0; i-- {
		pn := ComponentByID(m.Networks.Items, m.Nats.Items[i].PublicNetworkAWSID)
		if pn == nil {
			continue
		}

		m.Nats.Items[i].RoutedNetworks = ComponentNamesFromIDs(m.Networks.Items, m.Nats.Items[i].RoutedNetworkAWSIDs)

		// Get nat gateways name from tags of networks that reference it
		nw := ComponentByID(m.Networks.Items, m.Nats.Items[i].RoutedNetworkAWSIDs[0])
		nwtags := nw.GetTags()

		m.Nats.Items[i].Name = nwtags["ernest.nat_gateway"]

		nts = append(nts, definition.NatGateway{
			Name:          ShortName(m.Nats.Items[i].Name, prefix),
			PublicNetwork: ShortName(pn.ComponentName(), prefix),
		})
	}

	return nts
}

// UpdateNatValues corrects missing values after an import
func UpdateNatValues(m *output.FSMMessage) {
	for i := len(m.Nats.Items) - 1; i >= 0; i-- {
		pn := ComponentByID(m.Networks.Items, m.Nats.Items[i].PublicNetworkAWSID)

		if len(m.Nats.Items[i].RoutedNetworkAWSIDs) < 1 || pn == nil {
			// Remove nat that is not apart of this service!
			m.Nats.Items = append(m.Nats.Items[:i], m.Nats.Items[i+1:]...)
			continue
		}
	}
}

func mapNatNetworkIDs(nws []string) []string {
	var ids []string

	for _, nw := range nws {
		ids = append(ids, `$(networks.items.#[name="`+nw+`"].network_aws_id)`)
	}

	return ids
}

func mapNetworkNames(d definition.Definition, name string) []string {
	var nws []string
	for _, network := range d.Networks {
		if network.NatGateway == name {
			nws = append(nws, d.GeneratedName()+network.Name)
		}
	}

	return nws
}

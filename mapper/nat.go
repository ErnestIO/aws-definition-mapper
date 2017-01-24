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
		name := ng.Name

		nats = append(nats, output.Nat{
			Name:                name,
			PublicNetwork:       ng.PublicNetwork,
			RoutedNetworks:      nws,
			PublicNetworkAWSID:  `$(networks.items.#[name="` + ng.PublicNetwork + `"].network_aws_id)`,
			RoutedNetworkAWSIDs: mapNatNetworkIDs(nws),
			Tags:                mapTags(ng.Name, d.Name),
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

	for i := len(m.Nats.Items) - 1; i >= 0; i-- {
		n := m.Nats.Items[i]
		pn := ComponentByID(m.Networks.Items, n.PublicNetworkAWSID)

		if pn == nil {
			// Remove nat that is not apart of this service!
			m.Nats.Items = append(m.Nats.Items[:i], m.Nats.Items[i+1:]...)
			continue
		}

		nts = append(nts, definition.NatGateway{
			Name:          n.Name,
			PublicNetwork: pn.ComponentName(),
		})
	}

	return nts
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
			nws = append(nws, network.Name)
		}
	}

	return nws
}

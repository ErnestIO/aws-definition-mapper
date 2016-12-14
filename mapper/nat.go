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

		nats = append(nats, output.Nat{
			Name:                d.GeneratedName() + ng.Name,
			ProviderType:        "$(datacenters.items.0.type)",
			PublicNetwork:       d.GeneratedName() + ng.PublicNetwork,
			RoutedNetworks:      nws,
			PublicNetworkAWSID:  `$(networks.items.#[name="` + d.GeneratedName() + ng.PublicNetwork + `"].network_aws_id)`,
			RoutedNetworkAWSIDs: mapNatNetworkIDs(nws),
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

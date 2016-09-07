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
		nats = append(nats, output.Nat{
			Name:           d.GeneratedName() + ng.Name,
			PublicNetwork:  d.GeneratedName() + ng.PublicNetwork,
			RoutedNetworks: mapNetworkNames(d, ng.Name),
		})
	}

	return nats
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

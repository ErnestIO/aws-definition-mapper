/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package mapper

import (
	"github.com/ernestio/aws-definition-mapper/definition"
	"github.com/ernestio/aws-definition-mapper/output"
)

// MapSecurityGroups : Maps input security groups to an ernest format ones
func MapSecurityGroups(d definition.Definition) []output.Firewall {
	var firewalls []output.Firewall

	for _, sg := range d.SecurityGroups {
		f := output.Firewall{
			Name:             d.GeneratedName() + sg.Name,
			FirewallType:     "$(datacenters.items.0.type)",
			DatacenterType:   "$(datacenters.items.0.type)",
			DatacenterName:   "$(datacenters.items.0.name)",
			DatacenterSecret: "$(datacenters.items.0.secret)",
			DatacenterToken:  "$(datacenters.items.0.token)",
			DatacenterRegion: "$(datacenters.items.0.region)",
			VpcID:            "$(vpcs.items.0.vpc_id)",
		}

		for _, rule := range sg.Ingress {
			f.Rules = append(f.Rules, output.FirewallRule{
				Type:            "ingress",
				SourceIP:        rule.IP,
				SourcePort:      rule.FromPort,
				DestinationPort: rule.ToPort,
				Protocol:        MapProtocol(rule.Protocol),
			})
		}

		for _, rule := range sg.Egress {
			f.Rules = append(f.Rules, output.FirewallRule{
				Type:            "egress",
				SourceIP:        rule.IP,
				SourcePort:      rule.FromPort,
				DestinationPort: rule.ToPort,
				Protocol:        MapProtocol(rule.Protocol),
			})
		}

		firewalls = append(firewalls, f)
	}
	return firewalls
}

// MapProtocol : Maps the security groups protocol to the correct value
func MapProtocol(protocol string) string {
	if protocol == "any" {
		return "-1"
	}
	return protocol
}

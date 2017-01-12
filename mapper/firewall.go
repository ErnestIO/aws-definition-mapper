/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package mapper

import (
	"strconv"

	"github.com/ernestio/aws-definition-mapper/definition"
	"github.com/ernestio/aws-definition-mapper/output"
)

// MapSecurityGroups : Maps input security groups to an ernest format ones
func MapSecurityGroups(d definition.Definition) []output.Firewall {
	var firewalls []output.Firewall

	for _, sg := range d.SecurityGroups {
		name := d.GeneratedName() + sg.Name

		f := output.Firewall{
			Name:             name,
			Tags:             mapTags(sg.Name, d.Name),
			ProviderType:     "$(datacenters.items.0.type)",
			DatacenterType:   "$(datacenters.items.0.type)",
			DatacenterName:   "$(datacenters.items.0.name)",
			AccessKeyID:      "$(datacenters.items.0.aws_access_key_id)",
			SecretAccessKey:  "$(datacenters.items.0.aws_secret_access_key)",
			DatacenterRegion: "$(datacenters.items.0.region)",
			VpcID:            "$(vpcs.items.0.vpc_id)",
		}

		for _, rule := range sg.Ingress {
			f.Rules.Ingress = append(f.Rules.Ingress, BuildRule(rule))
		}

		for _, rule := range sg.Egress {
			f.Rules.Egress = append(f.Rules.Egress, BuildRule(rule))
		}

		firewalls = append(firewalls, f)
	}
	return firewalls
}

// BuildRule converts a definition rule into an output rule
func BuildRule(rule definition.SecurityGroupRule) output.FirewallRule {
	from, _ := strconv.Atoi(rule.FromPort)
	to, _ := strconv.Atoi(rule.ToPort)

	return output.FirewallRule{
		IP:       rule.IP,
		From:     from,
		To:       to,
		Protocol: MapProtocol(rule.Protocol),
	}
}

// MapProtocol : Maps the security groups protocol to the correct value
func MapProtocol(protocol string) string {
	if protocol == "any" {
		return "-1"
	}
	return protocol
}

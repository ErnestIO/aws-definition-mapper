/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package mapper

import (
	"strconv"
	"strings"

	"github.com/ernestio/aws-definition-mapper/definition"
	"github.com/ernestio/aws-definition-mapper/output"
)

// MapELBs : Maps the elbs from a given input payload.
func MapELBs(d definition.Definition) []output.ELB {
	var elbs []output.ELB

	for _, elb := range d.ELBs {
		e := output.ELB{
			Name:             d.GeneratedName() + elb.Name,
			IsPrivate:        elb.Private,
			Instances:        elb.Instances,
			SecurityGroups:   elb.SecurityGroups,
			DatacenterType:   "$(datacenters.items.0.type)",
			DatacenterName:   "$(datacenters.items.0.name)",
			DatacenterSecret: "$(datacenters.items.0.secret)",
			DatacenterToken:  "$(datacenters.items.0.token)",
			DatacenterRegion: "$(datacenters.items.0.region)",
			Type:             "$(datacenters.items.0.type)",
			VpcID:            "$(vpcs.items.0.vpc_id)",
		}

		for _, listener := range elb.Listeners {
			e.Listeners = append(e.Listeners, output.ELBListener{
				FromPort: listener.FromPort,
				ToPort:   listener.ToPort,
				Protocol: strings.ToUpper(listener.Protocol),
				SSLCert:  listener.SSLCert,
			})
		}

		for _, instance := range e.Instances {
			i := d.FindInstance(instance)
			if i != nil {
				e.NetworkAWSIDs = append(e.NetworkAWSIDs, `$(networks.items.#[name="`+d.GeneratedName()+i.Network+`"].network_aws_id)`)
			}
			for x := 0; x < i.Count; x++ {
				e.InstanceAWSIDs = append(e.InstanceAWSIDs, `$(instances.items.#[name="`+d.GeneratedName()+i.Name+`-`+strconv.Itoa(x+1)+`"].instance_aws_id)`)
			}
		}

		for _, sg := range e.SecurityGroups {
			e.SecurityGroupAWSIDs = append(e.SecurityGroupAWSIDs, `$(firewalls.items.#[name="`+d.GeneratedName()+sg+`"].security_group_aws_id)`)
		}

		elbs = append(elbs, e)
	}

	return elbs
}

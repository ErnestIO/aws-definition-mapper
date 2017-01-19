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
		name := d.GeneratedName() + elb.Name

		e := output.ELB{
			Name:             name,
			IsPrivate:        elb.Private,
			Instances:        elb.Instances,
			SecurityGroups:   elb.SecurityGroups,
			Tags:             mapTagsServiceOnly(d.Name),
			DatacenterType:   "$(datacenters.items.0.type)",
			DatacenterName:   "$(datacenters.items.0.name)",
			AccessKeyID:      "$(datacenters.items.0.aws_access_key_id)",
			SecretAccessKey:  "$(datacenters.items.0.aws_secret_access_key)",
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

		for _, subnet := range elb.Subnets {
			e.NetworkAWSIDs = append(e.NetworkAWSIDs, `$(networks.items.#[name="`+d.GeneratedName()+subnet+`"].network_aws_id)`)
		}

		for _, instance := range e.Instances {
			i := d.FindInstance(instance)
			if i != nil {
				for x := 0; x < i.Count; x++ {
					name := d.GeneratedName() + i.Name + "-" + strconv.Itoa(x+1)
					e.InstanceAWSIDs = append(e.InstanceAWSIDs, `$(instances.items.#[name="`+name+`"].instance_aws_id)`)
					e.InstanceNames = append(e.InstanceNames, name)
				}
			}
		}

		for _, sg := range e.SecurityGroups {
			e.SecurityGroupAWSIDs = append(e.SecurityGroupAWSIDs, `$(firewalls.items.#[name="`+d.GeneratedName()+sg+`"].security_group_aws_id)`)
		}

		elbs = append(elbs, e)
	}

	return elbs
}

// MapDefinitionELBs : Maps output elbs into a definition defined elbs
func MapDefinitionELBs(m *output.FSMMessage) []definition.ELB {
	var elbs []definition.ELB

	for _, elb := range m.ELBs.Items {
		instances := ComponentsByIDs(m.Instances.Items, elb.InstanceAWSIDs)

		e := definition.ELB{
			Name:           elb.Name,
			Private:        elb.IsPrivate,
			Subnets:        ComponentNamesFromIDs(m.Networks.Items, elb.NetworkAWSIDs),
			Instances:      ComponentGroups(instances, "ernest.instance_group"),
			SecurityGroups: ComponentNamesFromIDs(m.Firewalls.Items, elb.SecurityGroupAWSIDs),
		}

		for _, l := range elb.Listeners {
			e.Listeners = append(e.Listeners, definition.ELBListener{
				FromPort: l.FromPort,
				ToPort:   l.ToPort,
				Protocol: l.Protocol,
				SSLCert:  l.SSLCert,
			})
		}

		elbs = append(elbs, e)
	}

	return elbs
}

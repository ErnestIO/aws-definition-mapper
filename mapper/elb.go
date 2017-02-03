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
		var sgroups []string

		for _, sg := range elb.SecurityGroups {
			sgroups = append(sgroups, d.GeneratedName()+sg)
		}

		e := output.ELB{
			Name:             name,
			IsPrivate:        elb.Private,
			Instances:        elb.Instances,
			SecurityGroups:   sgroups,
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
			e.SecurityGroupAWSIDs = append(e.SecurityGroupAWSIDs, `$(firewalls.items.#[name="`+sg+`"].security_group_aws_id)`)
		}

		elbs = append(elbs, e)
	}

	return elbs
}

// UpdateELBValues corrects missing values after an import
func UpdateELBValues(m *output.FSMMessage) {
	for i := 0; i < len(m.ELBs.Items); i++ {
		m.ELBs.Items[i].DatacenterName = "$(datacenters.items.0.name)"
		m.ELBs.Items[i].DatacenterType = "$(datacenters.items.0.type)"
		m.ELBs.Items[i].AccessKeyID = "$(datacenters.items.0.aws_access_key_id)"
		m.ELBs.Items[i].SecretAccessKey = "$(datacenters.items.0.aws_secret_access_key)"
		m.ELBs.Items[i].DatacenterRegion = "$(datacenters.items.0.region)"
		m.ELBs.Items[i].VpcID = "$(vpcs.items.0.vpc_id)"
		m.ELBs.Items[i].Instances = ComponentGroupsFromIDs(m.Instances.Items, "ernest.instance_group", m.ELBs.Items[i].InstanceAWSIDs)
		m.ELBs.Items[i].InstanceNames = ComponentNamesFromIDs(m.Instances.Items, m.ELBs.Items[i].InstanceAWSIDs)
		m.ELBs.Items[i].SecurityGroups = ComponentNamesFromIDs(m.Firewalls.Items, m.ELBs.Items[i].SecurityGroupAWSIDs)
	}
}

// MapDefinitionELBs : Maps output elbs into a definition defined elbs
func MapDefinitionELBs(m *output.FSMMessage) []definition.ELB {
	var elbs []definition.ELB

	prefix := m.Datacenters.Items[0].Name + "-" + m.ServiceName + "-"

	for _, elb := range m.ELBs.Items {
		instances := ComponentsByIDs(m.Instances.Items, elb.InstanceAWSIDs)

		subnets := ComponentNamesFromIDs(m.Networks.Items, elb.NetworkAWSIDs)
		sgroups := ComponentNamesFromIDs(m.Firewalls.Items, elb.SecurityGroupAWSIDs)

		e := definition.ELB{
			Name:           ShortName(elb.Name, prefix),
			Private:        elb.IsPrivate,
			Subnets:        ShortNames(subnets, prefix),
			Instances:      ComponentGroupsFromIDs(instances, "ernest.instance_group", elb.InstanceAWSIDs),
			SecurityGroups: ShortNames(sgroups, prefix),
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

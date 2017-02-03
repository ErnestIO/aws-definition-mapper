/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package mapper

import (
	"github.com/ernestio/aws-definition-mapper/definition"
	"github.com/ernestio/aws-definition-mapper/output"
)

// MapRoute53Zones : Maps the zones from a given input payload.
func MapRoute53Zones(d definition.Definition) []output.Route53Zone {
	var zones []output.Route53Zone

	for _, zone := range d.Route53Zones {
		name := d.GeneratedName() + zone.Name

		z := output.Route53Zone{
			Name:             zone.Name,
			Private:          zone.Private,
			Tags:             mapTags(name, d.Name),
			ProviderType:     "$(datacenters.items.0.type)",
			DatacenterName:   "$(datacenters.items.0.name)",
			SecretAccessKey:  "$(datacenters.items.0.aws_secret_access_key)",
			AccessKeyID:      "$(datacenters.items.0.aws_access_key_id)",
			DatacenterRegion: "$(datacenters.items.0.region)",
			VPCID:            "$(vpcs.items.0.vpc_id)",
		}

		for _, record := range zone.Records {
			r := output.Record{
				Entry:  record.Entry,
				Type:   record.Type,
				Values: record.Values,
				TTL:    record.TTL,
			}

			// append instance and loadbalancer values
			r.Values = append(r.Values, MapRecordInstanceValues(d, record.Instances, zone.Private)...)
			r.Values = append(r.Values, MapRecordLoadbalancerValues(d, record.Loadbalancers)...)
			r.Values = append(r.Values, MapRecordRDSInstanceValues(d, record.RDSInstances)...)
			r.Values = append(r.Values, MapRecordRDSClusterValues(d, record.RDSClusters)...)

			z.Records = append(z.Records, r)
		}

		zones = append(zones, z)
	}

	return zones
}

// MapRecordInstanceValues takes a definition defined value and returns the template variables used on the build
func MapRecordInstanceValues(d definition.Definition, instances []string, private bool) []string {
	var values []string

	attr := "public_ip"
	if private {
		attr = "ip"
	}

	for _, name := range instances {
		// May need to unify this field with elastic_ip on instances
		values = append(values, `$(instances.items.#[name="`+d.GeneratedName()+name+`"].`+attr+`)`)
	}

	return values
}

// MapRecordLoadbalancerValues takes a definition defined value and returns the template variables used on the build
func MapRecordLoadbalancerValues(d definition.Definition, loadbalancers []string) []string {
	var values []string

	for _, name := range loadbalancers {
		values = append(values, `$(elbs.items.#[name="`+d.GeneratedName()+name+`"].dns_name)`)
	}

	return values
}

// MapRecordRDSInstanceValues takes a definition defined value and returns the template variables used on the build
func MapRecordRDSInstanceValues(d definition.Definition, rdsinstances []string) []string {
	var values []string

	for _, name := range rdsinstances {
		values = append(values, `$(rds_instances.items.#[name="`+d.GeneratedName()+name+`"].endpoint)`)
	}

	return values
}

// MapRecordRDSClusterValues takes a definition defined value and returns the template variables used on the build
func MapRecordRDSClusterValues(d definition.Definition, rdsclusters []string) []string {
	var values []string

	for _, name := range rdsclusters {
		values = append(values, `$(rds_clusters.items.#[name="`+d.GeneratedName()+name+`"].endpoint)`)
	}

	return values
}

// MapDefinitionRoute53Zones : Maps zones from the internal format to the input definition format
func MapDefinitionRoute53Zones(m *output.FSMMessage) []definition.Route53Zone {
	var zones []definition.Route53Zone

	prefix := m.Datacenters.Items[0].Name + "-" + m.ServiceName + "-"

	for _, zone := range m.Route53s.Items {
		z := definition.Route53Zone{
			Name:    zone.Name,
			Private: zone.Private,
		}

		for _, record := range zone.Records {
			r := definition.Record{
				Entry: record.Entry,
				Type:  record.Type,
				TTL:   record.TTL,
			}

			set := false

			for _, v := range r.Values {
				for _, i := range m.Instances.Items {
					if i.PublicIP == v || i.ElasticIP == v {
						r.Instances = append(r.Instances, ShortName(i.Name, prefix))
						set = true
						break
					}
				}

				for _, elb := range m.ELBs.Items {
					if elb.DNSName == v {
						r.Loadbalancers = append(r.Loadbalancers, ShortName(elb.Name, prefix))
						set = true
						break
					}
				}

				for _, rds := range m.RDSInstances.Items {
					if rds.Endpoint == v {
						r.RDSInstances = append(r.RDSInstances, ShortName(rds.Name, prefix))
						set = true
						break
					}
				}

				for _, rds := range m.RDSClusters.Items {
					if rds.Endpoint == v {
						r.RDSClusters = append(r.RDSClusters, ShortName(rds.Name, prefix))
						set = true
						break
					}
				}

				if set != true {
					r.Values = append(r.Values, v)
				}
			}

			z.Records = append(z.Records, r)
		}

		zones = append(zones, z)
	}

	return zones
}

// UpdateRoute53Values corrects missing values after an import
func UpdateRoute53Values(m *output.FSMMessage) {
	for i := 0; i < len(m.Route53s.Items); i++ {
		m.Route53s.Items[i].AccessKeyID = "$(datacenters.items.0.aws_access_key_id)"
		m.Route53s.Items[i].SecretAccessKey = "$(datacenters.items.0.aws_secret_access_key)"
		m.Route53s.Items[i].DatacenterRegion = "$(datacenters.items.0.region)"
		m.Route53s.Items[i].VPCID = "$(vpcs.items.0.vpc_id)"
	}
}

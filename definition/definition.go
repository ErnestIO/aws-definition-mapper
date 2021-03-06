/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package definition

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"unicode/utf8"
)

// Definition ...
type Definition struct {
	Name              string          `json:"name"`
	Datacenter        string          `json:"datacenter"`
	VpcID             string          `json:"vpc_id"`
	VpcSubnet         string          `json:"vpc_subnet,omitempty"`
	Networks          []Network       `json:"networks,omitempty"`
	Instances         []Instance      `json:"instances,omitempty"`
	SecurityGroups    []SecurityGroup `json:"security_groups,omitempty"`
	ELBs              []ELB           `json:"loadbalancers,omitempty"`
	S3Buckets         []S3            `json:"s3_buckets,omitempty"`
	Route53Zones      []Route53Zone   `json:"route53_zones,omitempty"`
	RDSClusters       []RDSCluster    `json:"rds_clusters,omitempty"`
	RDSInstances      []RDSInstance   `json:"rds_instances,omitempty"`
	NatGateways       []NatGateway    `json:"nat_gateways,omitempty"`
	EBSVolumes        []EBSVolume     `json:"ebs_volumes,omitempty"`
	DatacenterDetails Datacenter      `json:"-"`
}

// New returns a new Definition
func New() *Definition {
	return &Definition{
		Networks:       make([]Network, 0),
		Instances:      make([]Instance, 0),
		SecurityGroups: make([]SecurityGroup, 0),
	}
}

// FromJSON creates a definition from json
func FromJSON(data []byte) (*Definition, error) {
	var d Definition

	err := json.Unmarshal(data, d)
	if err != nil {
		return nil, err
	}

	return &d, nil
}

// ValidateVPC checks if vpc is valid
func (d *Definition) validateVPC() error {
	if d.VpcID == "" && d.VpcSubnet == "" {
		return errors.New("Please specify either the vpc_id of an existing vpc, or specify which vpc_subnet you want to use when creating a vpc")
	}

	if d.VpcID != "" && d.VpcSubnet == "" {
		return nil
	}

	_, _, err := net.ParseCIDR(d.VpcSubnet)
	if err != nil {
		return errors.New("VPC subnet is not valid")
	}

	return nil
}

// ValidateName checks if service is valid
func (d *Definition) validateName() error {
	// Check if service name is null
	if d.Name == "" {
		return errors.New("Service name should not be null")
	}

	// Check if service name is > 50 characters
	if utf8.RuneCountInString(d.Name) > 50 {
		return fmt.Errorf("Datacenter name can't be greater than %d characters", AWSMAXNAME)
	}
	return nil
}

func (d *Definition) validateDatacenter() error {
	if d.Datacenter == "" {
		return errors.New("Datacenter not specified")
	}
	return nil
}

// Validate the definition
func (d *Definition) Validate() error {
	// Validate Name
	if err := d.validateName(); err != nil {
		return err
	}

	// Validate Datacenter
	if err := d.validateDatacenter(); err != nil {
		return err
	}

	// Validate VPC
	if err := d.validateVPC(); err != nil {
		return err
	}

	// Validate Networks
	for _, n := range d.Networks {
		if err := n.Validate(&d.DatacenterDetails); err != nil {
			return err
		}
	}

	// Validate Instances
	for _, i := range d.Instances {
		nw := d.FindNetwork(i.Network)

		if err := i.Validate(nw, d.EBSVolumes); err != nil {
			return err
		}
	}

	// Validate Security Groups
	for _, sg := range d.SecurityGroups {
		if err := sg.Validate(d.Networks); err != nil {
			return err
		}
	}

	// Validate Nat Gateways
	for _, ng := range d.NatGateways {
		if err := ng.Validate(d.Networks); err != nil {
			return err
		}
	}

	// Validate ELB's
	for _, lb := range d.ELBs {
		if err := lb.Validate(d.Networks); err != nil {
			return err
		}
		for _, instance := range lb.Instances {
			if d.FindInstance(instance) == nil {
				return fmt.Errorf("ELB Instance (%s) is not valid", instance)
			}
		}
		for _, sg := range lb.SecurityGroups {
			if d.FindSecurityGroup(sg) == nil {
				return fmt.Errorf("ELB Security Group (%s) is not valid", sg)
			}
		}
	}

	// Validate S3 Buckets
	for _, s3bucket := range d.S3Buckets {
		err := s3bucket.Validate()
		if err != nil {
			return err
		}
	}

	// Validate Route53
	for _, r53 := range d.Route53Zones {
		err := r53.Validate()
		if err != nil {
			return err
		}
	}

	// Validate RDS Clusters
	for _, rdsc := range d.RDSClusters {
		err := rdsc.Validate(d.Networks, d.SecurityGroups)
		if err != nil {
			return err
		}
	}

	// Validate RDS Instances
	for _, rdsi := range d.RDSInstances {
		err := rdsi.Validate(d.Networks, d.SecurityGroups, d.RDSClusters)
		if err != nil {
			return err
		}
	}

	// validate EBS Volumes
	for _, vol := range d.EBSVolumes {
		err := vol.Validate()
		if err != nil {
			return err
		}
	}

	if hasDuplicateNetworks(d.Networks) {
		return errors.New("Duplicate network names found")
	}

	if hasDuplicateInstance(d.Instances) {
		return errors.New("Duplicate instance names found")
	}

	return nil
}

// GeneratedName returns the generated service name
func (d *Definition) GeneratedName() string {
	return d.Datacenter + "-" + d.Name + "-"
}

// FindNetwork returns a network matched by name
func (d *Definition) FindNetwork(name string) *Network {
	for _, network := range d.Networks {
		if network.Name == name {
			return &network
		}
	}
	return nil
}

// FindInstance returns a instance matched by name
func (d *Definition) FindInstance(name string) *Instance {
	for _, instance := range d.Instances {
		if instance.Name == name {
			return &instance
		}
	}
	return nil
}

// FindSecurityGroup returns a sg matched by name
func (d *Definition) FindSecurityGroup(name string) *SecurityGroup {
	for _, sg := range d.SecurityGroups {
		if sg.Name == name {
			return &sg
		}
	}
	return nil
}

// FindRDSCluster returns a rds cluster matched by name
func (d *Definition) FindRDSCluster(name string) *RDSCluster {
	for _, c := range d.RDSClusters {
		if c.Name == name {
			return &c
		}
	}
	return nil

}

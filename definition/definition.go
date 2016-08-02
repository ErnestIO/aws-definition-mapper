/*
---
name: aws
datacenter: r3-me2
ernest_ip:
  - 31.210.240.166
service_ip: 9.9.9.9
security_groups:
 -   name: web-sg-1
     rules:
       ingress:
                        -  ip: 10.1.1.11/32
          protocol: any
          from: 80
          to: 80
networks:
                   - name: web
    subnet: 10.1.0.0/24
instances:
  - name: web
    image: ami-6666f915
    count: 1
    Type: e1.micro
    network: bla
    private_ip: 10.1.0.11
    key_pair: some-keypair
    security_groups:
      - web-sg-1
      - web-sg-2
*/

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
	Name           string          `json:"name"`
	Datacenter     string          `json:"datacenter"`
	Bootstrapping  string          `json:"bootstrapping"`
	ErnestIP       []string        `json:"ernest_ip"`
	ServiceIP      string          `json:"service_ip"`
	Networks       []Network       `json:"networks"`
	Instances      []Instance      `json:"instances"`
	SecurityGroups []SecurityGroup `json:"secuirty_groups"`
}

// New returns a new Definition
func New() *Definition {
	return &Definition{
		ErnestIP:       make([]string, 0),
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

func (d *Definition) validateServiceIP() error {
	if d.ServiceIP == "" {
		return nil
	}
	if ok := net.ParseIP(d.ServiceIP); ok == nil {
		return errors.New("ServiceIP is not a valid IP")
	}
	return nil
}

// IsSaltBootstrapped : Return a boolean describing if bootstrapping option is salt
func (d *Definition) IsSaltBootstrapped() bool {
	if d.Bootstrapping == "salt" {
		return true
	}
	return false
}

// Validate the definition
func (d *Definition) Validate() error {
	// Validate Definition
	err := d.validateName()
	if err != nil {
		return err
	}

	err = d.validateServiceIP()
	if err != nil {
		return err
	}

	// Validate Datacenter
	err = d.validateDatacenter()
	if err != nil {
		return err
	}

	// Validate Networks
	for _, n := range d.Networks {
		err := n.Validate()
		if err != nil {
			return err
		}
	}

	// Validate Instances
	for _, i := range d.Instances {
		nw := d.FindNetwork(i.Network)

		err := i.Validate(nw)
		if err != nil {
			return err
		}
	}

	// Validate Security Groups
	for _, sg := range d.SecurityGroups {
		err := sg.Validate(d.Networks)
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

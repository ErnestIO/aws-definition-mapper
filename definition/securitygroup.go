/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package definition

import (
	"errors"
	"unicode/utf8"
)

// SecurityGroup ...
type SecurityGroup struct {
	Name    string              `json:"name"`
	Ingress []SecurityGroupRule `json:"ingress"`
	Egress  []SecurityGroupRule `json:"egress"`
}

// SecurityGroupRule ...
type SecurityGroupRule struct {
	IP       string `json:"ip"`
	FromPort string `json:"from_port"`
	ToPort   string `json:"to_port"`
	Protocol string `json:"protocol"`
}

// Validate security group
func (sg *SecurityGroup) Validate(networks []Network) error {
	// Check if security group name is null
	if sg.Name == "" {
		return errors.New("Security Group name should not be null")
	}

	// Check if security group name is > 50 characters
	if utf8.RuneCountInString(sg.Name) > AWSMAXNAME {
		return errors.New("Security Group name can't be greater than 50 characters")
	}

	for _, rule := range sg.Ingress {
		err := rule.Validate(networks)
		if err != nil {
			return err
		}
	}

	for _, rule := range sg.Egress {
		err := rule.Validate(networks)
		if err != nil {
			return err
		}
	}

	return nil
}

// Validate security group rule
func (rule *SecurityGroupRule) Validate(networks []Network) error {
	err := validateIP(rule.IP, "Security Group IP", networks)
	if err != nil {
		return err
	}

	// Validate FromPort Port
	// Must be: [0 - 65535]
	err = validatePort(rule.FromPort, "Security Group From")
	if err != nil {
		return err
	}

	// Validate ToPort Port
	// Must be: [0 - 65535]
	err = validatePort(rule.ToPort, "Security Group To")
	if err != nil {
		return err
	}

	// Validate Protocol
	// Must be one of: tcp | udp | icmp | any | tcp & udp
	err = validateProtocol(rule.Protocol)
	if err != nil {
		return err
	}

	return nil
}

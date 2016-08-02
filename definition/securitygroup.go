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
	Name  string              `json:"name"`
	Rules []SecurityGroupRule `json:"rules"`
}

// SecurityGroupRule ...
type SecurityGroupRule struct {
	Source      string `json:"source"`
	Destination string `json:"destination"`
	FromPort    string `json:"from_port"`
	ToPort      string `json:"to_port"`
	Protocol    string `json:"protocol"`
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

	for _, rule := range sg.Rules {
		err := validateIP(rule.Source, "Security Group Source", networks)
		if err != nil {
			return err
		}

		err = validateIP(rule.Destination, "Security Group Destination", networks)
		if err != nil {
			return err
		}

		// Validate FromPort Port
		// Must be: [any | 1 - 65535]
		err = validatePort(rule.FromPort, "Security Group From")
		if err != nil {
			return err
		}

		// Validate ToPort Port
		// Must be: [any | 1 - 65535]
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

	}

	return nil
}

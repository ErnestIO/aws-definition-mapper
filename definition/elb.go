/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package definition

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

// ELBPort ...
type ELBPort struct {
	FromPort int    `json:"from_port"`
	ToPort   int    `json:"to_port"`
	Protocol string `json:"protocol"`
	SSLCert  string `json:"ssl_cert"`
}

// ELB ...
type ELB struct {
	Name           string    `json:"name"`
	Private        bool      `json:"private"`
	Instances      []string  `json:"instances"`
	SecurityGroups []string  `json:"security_groups"`
	Ports          []ELBPort `json:"listeners"`
}

// Validate checks if a Network is valid
func (e *ELB) Validate() error {
	if e.Name == "" {
		return errors.New("ELB name should not be null")
	}

	// Check if network name is > 50 characters
	if utf8.RuneCountInString(e.Name) > AWSMAXNAME {
		return fmt.Errorf("ELB name can't be greater than %d characters", AWSMAXNAME)
	}

	if len(e.Ports) < 1 {
		return errors.New("ELB must contain more than one ports")
	}

	for _, port := range e.Ports {
		if port.FromPort < 1 || port.FromPort > 65535 {
			return fmt.Errorf("From Port (%d) is out of range [1 - 65535]", port.FromPort)
		}

		if port.ToPort < 1 || port.ToPort > 65535 {
			return fmt.Errorf("From Port (%d) is out of range [1 - 65535]", port.ToPort)
		}

		if port.Protocol != "http" &&
			port.Protocol != "https" &&
			port.Protocol != "tcp" &&
			port.Protocol != "ssl" {
			return errors.New("ELB Protocol must be one of http, https, tcp or ssl")
		}

	}

	return nil
}

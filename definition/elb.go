/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package definition

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

// ELBListener ...
type ELBListener struct {
	FromPort int    `json:"from_port"`
	ToPort   int    `json:"to_port"`
	Protocol string `json:"protocol"`
	SSLCert  string `json:"ssl_cert"`
}

// ELB ...
type ELB struct {
	Name           string        `json:"name"`
	Private        bool          `json:"private"`
	Subnets        []string      `json:"networks"`
	Instances      []string      `json:"instances"`
	SecurityGroups []string      `json:"security_groups"`
	Listeners      []ELBListener `json:"listeners"`
}

// Validate checks if a Network is valid
func (e *ELB) Validate(networks []Network) error {
	if e.Name == "" {
		return errors.New("ELB name should not be null")
	}

	// Check if network name is > 50 characters
	if utf8.RuneCountInString(e.Name) > AWSMAXNAME {
		return fmt.Errorf("ELB name can't be greater than %d characters", AWSMAXNAME)
	}

	if len(e.Listeners) < 1 {
		return errors.New("ELB must contain more than one listeners")
	}

	if e.Private != true && len(e.Subnets) < 1 {
		return errors.New("ELB must specify at least one subnet if public")
	}

	for _, nw := range e.Subnets {
		for _, n := range networks {
			if nw == n.Name && n.Public != true && e.Private != true {
				return fmt.Errorf("ELB subnet (%s) is not a public subnet", nw)
			}
		}
	}

	for _, listener := range e.Listeners {
		if listener.FromPort < 1 || listener.FromPort > 65535 {
			return fmt.Errorf("From Port (%d) is out of range [1 - 65535]", listener.FromPort)
		}

		if listener.ToPort < 1 || listener.ToPort > 65535 {
			return fmt.Errorf("From Port (%d) is out of range [1 - 65535]", listener.ToPort)
		}

		if listener.Protocol != "http" &&
			listener.Protocol != "https" &&
			listener.Protocol != "tcp" &&
			listener.Protocol != "ssl" {
			return errors.New("ELB Protocol must be one of http, https, tcp or ssl")
		}

		if listener.Protocol == "https" && listener.SSLCert == "" || listener.Protocol == "ssl" && listener.SSLCert == "" {
			return errors.New("ELB listener must specify an ssl cert when protocol is https/ssl")
		}

	}

	return nil
}

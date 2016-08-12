/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package definition

import (
	"errors"
	"fmt"
	"net"
	"unicode/utf8"
)

// Instance ...
type Instance struct {
	Name           string   `json:"name"`
	Type           string   `json:"type"`
	Image          string   `json:"image"`
	Count          int      `json:"count"`
	Network        string   `json:"network"`
	StartIP        net.IP   `json:"start_ip"`
	KeyPair        string   `json:"key_pair"`
	SecurityGroups []string `json:"security_groups"`
}

// Validate : Validates the instance returning true or false if is valid or not
func (i *Instance) Validate(network *Network) error {
	if i.Name == "" {
		return errors.New("Instance name should not be null")
	}

	if utf8.RuneCountInString(i.Name) > AWSMAXNAME {
		return fmt.Errorf("Instance name can't be greater than %d characters", AWSMAXNAME)
	}

	if i.Type == "" {
		return errors.New("Instance type should not be null")
	}

	if i.Image == "" {
		return errors.New("Instance image should not be null")
	}

	if i.Count < 1 {
		return errors.New("Instance count should not be < 1")
	}

	if i.Network == "" {
		return errors.New("Instance network should not be null")
	}

	// Validate IP addresses
	if network != nil {
		_, nw, err := net.ParseCIDR(network.Subnet)
		if err != nil {
			return errors.New("Could not process network")
		}

		startIP := net.ParseIP(i.StartIP.String()).To4()
		ip := make(net.IP, net.IPv4len)
		copy(ip, i.StartIP.To4())

		for x := 0; x < i.Count; x++ {
			if !nw.Contains(ip) {
				return errors.New("Instance IP invalid. IP must be a valid IP in the same range as it's network")
			}

			// Check IP is greater than Start IP (Bounds checking)
			if ip[3] < startIP[3] {
				return errors.New("Instance IP invalid. Allocated IP is lower than Start IP")
			}

			ip[3]++
		}
	}

	return nil
}

/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package definition

import (
	"errors"
	"fmt"
	"net"
	"strconv"
	"strings"
)

const (
	PROTOCOLTCP    = "tcp"
	PROTOCOLUDP    = "udp"
	PROTOCOLANY    = "any"
	PROTOCOLICMP   = "icmp"
	TARGETEXTERNAL = "external"
	TARGETINTERNAL = "internal"
	TARGETANY      = "any"
	AWSMAXNAME     = 50
)

func isNetwork(networks []Network, name string) bool {
	for _, network := range networks {
		if network.Name == name {
			return true
		}
	}
	return false
}

func isSecurityGroup(sgs []SecurityGroup, name string) bool {
	for _, sg := range sgs {
		if sg.Name == name {
			return true
		}
	}
	return false
}

func validateProtocol(p string) error {
	switch p {
	case PROTOCOLTCP, PROTOCOLUDP, PROTOCOLICMP, PROTOCOLANY:
		return nil
	}
	return errors.New("Protocol is invalid")
}

// ValidateIP checks if an string is a valid source/destionation
func validateIP(ip, iptype string, networks []Network) error {
	// Check if Source is a valid value or a valid IP/CIDR
	// One of: any | named networks | CIDR

	switch ip {
	case TARGETANY:
		return nil
	}

	// Check if it refers to an internal Network
	if isNetwork(networks, ip) {
		return nil
	}

	// Check if Source is a valid CIDR
	_, _, err := net.ParseCIDR(ip)
	if err == nil {
		return nil
	}

	// Check if Source is a valid IP
	ipx := net.ParseIP(ip)
	if ipx != nil {
		return nil
	}

	return fmt.Errorf("%s (%s) is not valid", iptype, ip)
}

// ValidatePort checks an string to be a valid TCP port
func validatePort(p, ptype string) error {
	port, err := strconv.Atoi(p)
	if err != nil {
		return fmt.Errorf("%s Port (%s) is not valid", ptype, p)
	}

	if port < 0 || port > 65535 {
		return fmt.Errorf("%s Port (%s) is out of range [0 - 65535]", ptype, p)
	}

	return nil
}

func hasDuplicateNetworks(collection []Network) bool {
	var names []string
	for _, item := range collection {
		for _, n := range names {
			if item.Name == n {
				return true
			}
		}
		names = append(names, item.Name)
	}
	return false
}

func hasDuplicateInstance(collection []Instance) bool {
	var names []string
	for _, item := range collection {
		for _, n := range names {
			if item.Name == n {
				return true
			}
		}
		names = append(names, item.Name)
	}
	return false
}

func isOneOf(values []string, value string) bool {
	for _, v := range values {
		if v == value {
			return true
		}
	}
	return false
}

func validateDateTimeFormat(t string) error {
	// ddd:hh24:mi
	var days = []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}

	parts := strings.Split(t, ":")
	if len(parts) != 3 {
		return errors.New("Date format must take the form of 'ddd:hh24:mi'. i.e. 'Mon:21:30'")
	}

	// is valid day
	if isOneOf(days, parts[0]) != true {
		return fmt.Errorf("Date format invalid. Day must be one of %s", strings.Join(days, ", "))
	}

	// is valid hour
	d, err := strconv.Atoi(parts[1])
	if err != nil || d < 0 || d > 23 {
		return errors.New("Date format invalid. Hour must be between 0 and 23 hours")
	}

	// is valid minute
	d, err = strconv.Atoi(parts[2])
	if err != nil || d < 0 || d > 59 {
		return errors.New("Date format invalid. Minute must be between 0 and 59 minutes")
	}

	return nil
}

func validateTimeWindow(w string) error {
	p := strings.Split(w, "-")
	if len(p) != 2 {
		return errors.New("Window format must take the form of 'ddd:hh24:mi-ddd:hh24:mi'. i.e. 'Mon:21:30-Mon:22:00'")
	}

	err := validateDateTimeFormat(p[0])
	if err != nil {
		return err
	}

	return validateDateTimeFormat(p[1])
}

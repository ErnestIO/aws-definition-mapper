/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package definition

import (
	"errors"
	"fmt"
	"strings"
)

// CNAME ...
var CNAME = "CNAME"

// DNSTypes ...
var DNSTypes = []string{"A", "AAAA", "CNAME", "MX", "PTR", "TXT", "SRV", "SPF", "NAPTR", "NS", "SOA"}

// Record stores the entries for a zone
type Record struct {
	Entry         string   `json:"entry"`
	Type          string   `json:"type"`
	Instances     []string `json:"instances"`
	Loadbalancers []string `json:"loadbalancers"`
	RDSClusters   []string `json:"rds_clusters"`
	RDSInstances  []string `json:"rds_instances"`
	Values        []string `json:"values"`
	TTL           int64    `json:"ttl"`
}

// Route53Zone ...
type Route53Zone struct {
	Name    string   `json:"name"`
	Private bool     `json:"private"`
	Records []Record `json:"records"`
}

// Validate checks if a Route53Zone is valid
func (z *Route53Zone) Validate() error {
	if z.Name == "" {
		return errors.New("Route53 zone name should not be null")
	}

	for _, record := range z.Records {
		if record.Entry == "" {
			return errors.New("Route53 record entry name should not be null")
		}

		if !validDNSType(record.Type) {
			return fmt.Errorf("Route53 record type '%s' is not a valid dns type. Please use one of [%s]", record.Type, strings.Join(DNSTypes, ", "))
		}

		if len(record.Values) == 0 &&
			len(record.Instances) == 0 &&
			len(record.Loadbalancers) == 0 &&
			len(record.RDSInstances) == 0 &&
			len(record.RDSClusters) == 0 {
			return errors.New("Route53 record must specify a valid target [rds_instances, rds_clusters, instances or loadbalancers] or value")
		}

		err := validateRecordTargets(&record)
		if err != nil {
			return err
		}

		// Todo: make this an aliased type
		if len(record.Loadbalancers) > 0 && record.Type != CNAME {
			return errors.New("Route53 record type must be CNAME when using loadbalancers as a target")
		}

		if len(record.RDSInstances) > 0 && record.Type != CNAME {
			return errors.New("Route53 record type must be CNAME when using rds_instances as a target")
		}

		if len(record.RDSClusters) > 0 && record.Type != CNAME {
			return errors.New("Route53 record type must be CNAME when using rds_clusters as a target")
		}

		if len(record.Instances) > 0 && record.Type != "A" {
			return errors.New("Route53 record type must be A when using instances as a target")
		}

		if record.TTL == 0 {
			return errors.New("Route53 record TTL must be greater than 0")
		}
	}

	return nil
}

func validDNSType(t string) bool {
	for _, dt := range DNSTypes {
		if dt == t {
			return true
		}
	}
	return false
}

func validateRecordTargets(record *Record) error {
	var set bool
	err := errors.New("Route53 record must specify only one of either rds_instances, rds_clusters, instances or loadbalancers as targets")

	if len(record.Loadbalancers) > 0 {
		set = true
	}

	if len(record.Instances) > 0 {
		if set {
			return err
		}
		set = true
	}

	if len(record.RDSInstances) > 0 {
		if set {
			return err
		}
		set = true
	}

	if len(record.RDSClusters) > 0 {
		if set {
			return err
		}
	}

	return nil
}

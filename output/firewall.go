/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package output

// Firewall : Mapping for a firewall component
type Firewall struct {
	ProviderType       string `json:"_type"`
	SecurityGroupAWSID string `json:"security_group_aws_id"`
	Name               string `json:"name"`
	Rules              struct {
		Ingress []FirewallRule `json:"ingress"`
		Egress  []FirewallRule `json:"egress"`
	} `json:"rules"`
	Tags             map[string]string `json:"tags"`
	DatacenterType   string            `json:"datacenter_type,omitempty"`
	DatacenterName   string            `json:"datacenter_name,omitempty"`
	DatacenterRegion string            `json:"datacenter_region"`
	AccessKeyID      string            `json:"aws_access_key_id"`
	SecretAccessKey  string            `json:"aws_secret_access_key"`
	VpcID            string            `json:"vpc_id"`
	Service          string            `json:"service"`
	Status           string            `json:"status"`
	Exists           bool
}

// HasChanged diff's the two items and returns true if there have been any changes
func (f *Firewall) HasChanged(of *Firewall) bool {
	if len(f.Rules.Ingress) != len(of.Rules.Ingress) ||
		len(f.Rules.Egress) != len(of.Rules.Egress) {
		return true
	}

	for _, rule := range f.Rules.Ingress {
		if hasRule(of.Rules.Ingress, rule) != true {
			return true
		}
	}

	for _, rule := range f.Rules.Egress {
		if hasRule(of.Rules.Egress, rule) != true {
			return true
		}
	}

	return false
}

func hasRule(rules []FirewallRule, rule FirewallRule) bool {
	for _, r := range rules {
		if ruleMatches(r.To, rule.To, r.Protocol, rule.Protocol) &&
			r.Protocol == rule.Protocol &&
			r.IP == rule.IP &&
			ruleMatches(r.From, rule.From, r.Protocol, rule.Protocol) {
			return true
		}
	}

	return false
}

func ruleMatches(nv, ov int, np, op string) bool {
	if np == "-1" && op == "-1" {
		return true
	}

	return nv == ov
}

// GetTags returns a components tags
func (f Firewall) GetTags() map[string]string {
	return f.Tags
}

// ProviderID returns a components provider id
func (f Firewall) ProviderID() string {
	return f.SecurityGroupAWSID
}

// ComponentName returns a components name
func (f Firewall) ComponentName() string {
	return f.Name
}

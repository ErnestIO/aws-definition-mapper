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

	for i := 0; i < len(f.Rules.Ingress); i++ {
		if ruleChanged(f.Rules.Ingress[i].To, of.Rules.Ingress[i].To) ||
			f.Rules.Ingress[i].Protocol != of.Rules.Ingress[i].Protocol ||
			f.Rules.Ingress[i].IP != of.Rules.Ingress[i].IP ||
			ruleChanged(f.Rules.Ingress[i].From, of.Rules.Ingress[i].From) {
			return true
		}
	}

	for i := 0; i < len(f.Rules.Egress); i++ {
		if ruleChanged(f.Rules.Egress[i].To, of.Rules.Egress[i].To) ||
			f.Rules.Egress[i].Protocol != of.Rules.Egress[i].Protocol ||
			f.Rules.Egress[i].IP != of.Rules.Egress[i].IP ||
			ruleChanged(f.Rules.Egress[i].From, of.Rules.Egress[i].From) {
			return true
		}
	}

	return false
}

func ruleChanged(nv, ov int) bool {
	if nv == 65535 {
		nv = 0
	}
	if ov == 65535 {
		ov = 0
	}

	return nv != ov
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

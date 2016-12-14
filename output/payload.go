/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package output

import (
	"encoding/json"
	"strings"

	"github.com/r3labs/graph"
	"github.com/r3labs/workflow"
)

// FSMMessage is the JSON payload that will be sent to the FSM to create a
// service.
type FSMMessage struct {
	ID            string   `json:"id"`
	Body          string   `json:"body"`
	Endpoint      string   `json:"endpoint"`
	Service       string   `json:"service"`
	Bootstrapping string   `json:"bootstrapping"`
	ErnestIP      []string `json:"ernest_ip"`
	ServiceIP     string   `json:"service_ip"`
	Parent        string   `json:"existing_service"`
	Workflow      struct {
		Arcs []graph.Edge `json:"arcs"`
	} `json:"workflow"`
	ServiceName string `json:"name"`
	Client      string `json:"client"` // TODO: Use client or client_id not both!
	ClientID    string `json:"client_id"`
	ClientName  string `json:"client_name"`
	Started     string `json:"started"`
	Finished    string `json:"finished"`
	Status      string `json:"status"`
	Type        string `json:"type"`
	Datacenters struct {
		Started  string       `json:"started"`
		Finished string       `json:"finished"`
		Status   string       `json:"status"`
		Items    []Datacenter `json:"items"`
	} `json:"datacenters"`
	VPCs struct {
		Started  string `json:"started"`
		Finished string `json:"finished"`
		Status   string `json:"status"`
		Items    []VPC  `json:"items"`
	} `json:"vpcs"`
	VPCsToCreate struct {
		Started  string `json:"started"`
		Finished string `json:"finished"`
		Status   string `json:"status"`
		Items    []VPC  `json:"items"`
	} `json:"vpcs_to_create"`
	VPCsToDelete struct {
		Started  string `json:"started"`
		Finished string `json:"finished"`
		Status   string `json:"status"`
		Items    []VPC  `json:"items"`
	} `json:"vpcs_to_delete"`
	Networks struct {
		Started  string    `json:"started"`
		Finished string    `json:"finished"`
		Status   string    `json:"status"`
		Items    []Network `json:"items"`
	} `json:"networks"`
	NetworksToCreate struct {
		Started  string    `json:"started"`
		Finished string    `json:"finished"`
		Status   string    `json:"status"`
		Items    []Network `json:"items"`
	} `json:"networks_to_create"`
	NetworksToDelete struct {
		Started  string    `json:"started"`
		Finished string    `json:"finished"`
		Status   string    `json:"status"`
		Items    []Network `json:"items"`
	} `json:"networks_to_delete"`
	Instances struct {
		Started  string     `json:"started"`
		Finished string     `json:"finished"`
		Status   string     `json:"status"`
		Items    []Instance `json:"items"`
	} `json:"instances"`
	InstancesToCreate struct {
		Started  string     `json:"started"`
		Finished string     `json:"finished"`
		Status   string     `json:"status"`
		Items    []Instance `json:"items"`
	} `json:"instances_to_create"`
	InstancesToUpdate struct {
		Started  string     `json:"started"`
		Finished string     `json:"finished"`
		Status   string     `json:"status"`
		Items    []Instance `json:"items"`
	} `json:"instances_to_update"`
	InstancesToDelete struct {
		Started  string     `json:"started"`
		Finished string     `json:"finished"`
		Status   string     `json:"status"`
		Items    []Instance `json:"items"`
	} `json:"instances_to_delete"`
	Firewalls struct {
		Started  string     `json:"started"`
		Finished string     `json:"finished"`
		Status   string     `json:"status"`
		Items    []Firewall `json:"items"`
	} `json:"firewalls"`
	FirewallsToCreate struct {
		Started  string     `json:"started"`
		Finished string     `json:"finished"`
		Status   string     `json:"status"`
		Items    []Firewall `json:"items"`
	} `json:"firewalls_to_create"`
	FirewallsToUpdate struct {
		Started  string     `json:"started"`
		Finished string     `json:"finished"`
		Status   string     `json:"status"`
		Items    []Firewall `json:"items"`
	} `json:"firewalls_to_update"`
	FirewallsToDelete struct {
		Started  string     `json:"started"`
		Finished string     `json:"finished"`
		Status   string     `json:"status"`
		Items    []Firewall `json:"items"`
	} `json:"firewalls_to_delete"`
	Nats struct {
		Started  string `json:"started"`
		Finished string `json:"finished"`
		Status   string `json:"status"`
		Items    []Nat  `json:"items"`
	} `json:"nats"`
	NatsToCreate struct {
		Started  string `json:"started"`
		Finished string `json:"finished"`
		Status   string `json:"status"`
		Items    []Nat  `json:"items"`
	} `json:"nats_to_create"`
	NatsToUpdate struct {
		Started  string `json:"started"`
		Finished string `json:"finished"`
		Status   string `json:"status"`
		Items    []Nat  `json:"items"`
	} `json:"nats_to_update"`
	NatsToDelete struct {
		Started  string `json:"started"`
		Finished string `json:"finished"`
		Status   string `json:"status"`
		Items    []Nat  `json:"items"`
	} `json:"nats_to_delete"`
	ELBs struct {
		Started  string `json:"started"`
		Finished string `json:"finished"`
		Status   string `json:"status"`
		Items    []ELB  `json:"items"`
	} `json:"elbs"`
	ELBsToCreate struct {
		Started  string `json:"started"`
		Finished string `json:"finished"`
		Status   string `json:"status"`
		Items    []ELB  `json:"items"`
	} `json:"elbs_to_create"`
	ELBsToUpdate struct {
		Started  string `json:"started"`
		Finished string `json:"finished"`
		Status   string `json:"status"`
		Items    []ELB  `json:"items"`
	} `json:"elbs_to_update"`
	ELBsToDelete struct {
		Started  string `json:"started"`
		Finished string `json:"finished"`
		Status   string `json:"status"`
		Items    []ELB  `json:"items"`
	} `json:"elbs_to_delete"`
	S3s struct {
		Started  string `json:"started"`
		Finished string `json:"finished"`
		Status   string `json:"status"`
		Items    []S3   `json:"items"`
	} `json:"s3s"`
	S3sToCreate struct {
		Started  string `json:"started"`
		Finished string `json:"finished"`
		Status   string `json:"status"`
		Items    []S3   `json:"items"`
	} `json:"s3s_to_create"`
	S3sToUpdate struct {
		Started  string `json:"started"`
		Finished string `json:"finished"`
		Status   string `json:"status"`
		Items    []S3   `json:"items"`
	} `json:"s3s_to_update"`
	S3sToDelete struct {
		Started  string `json:"started"`
		Finished string `json:"finished"`
		Status   string `json:"status"`
		Items    []S3   `json:"items"`
	} `json:"s3s_to_delete"`
	Route53s struct {
		Started  string        `json:"started"`
		Finished string        `json:"finished"`
		Status   string        `json:"status"`
		Items    []Route53Zone `json:"items"`
	} `json:"route53s"`
	Route53sToCreate struct {
		Started  string        `json:"started"`
		Finished string        `json:"finished"`
		Status   string        `json:"status"`
		Items    []Route53Zone `json:"items"`
	} `json:"route53s_to_create"`
	Route53sToUpdate struct {
		Started  string        `json:"started"`
		Finished string        `json:"finished"`
		Status   string        `json:"status"`
		Items    []Route53Zone `json:"items"`
	} `json:"route53s_to_update"`
	Route53sToDelete struct {
		Started  string        `json:"started"`
		Finished string        `json:"finished"`
		Status   string        `json:"status"`
		Items    []Route53Zone `json:"items"`
	} `json:"route53s_to_delete"`
	RDSClusters struct {
		Started  string       `json:"started"`
		Finished string       `json:"finished"`
		Status   string       `json:"status"`
		Items    []RDSCluster `json:"items"`
	} `json:"rds_clusters"`
	RDSClustersToCreate struct {
		Started  string       `json:"started"`
		Finished string       `json:"finished"`
		Status   string       `json:"status"`
		Items    []RDSCluster `json:"items"`
	} `json:"rds_clusters_to_create"`
	RDSClustersToUpdate struct {
		Started  string       `json:"started"`
		Finished string       `json:"finished"`
		Status   string       `json:"status"`
		Items    []RDSCluster `json:"items"`
	} `json:"rds_clusters_to_update"`
	RDSClustersToDelete struct {
		Started  string       `json:"started"`
		Finished string       `json:"finished"`
		Status   string       `json:"status"`
		Items    []RDSCluster `json:"items"`
	} `json:"rds_clusters_to_delete"`
	RDSInstances struct {
		Started  string        `json:"started"`
		Finished string        `json:"finished"`
		Status   string        `json:"status"`
		Items    []RDSInstance `json:"items"`
	} `json:"rds_instances"`
	RDSInstancesToCreate struct {
		Started  string        `json:"started"`
		Finished string        `json:"finished"`
		Status   string        `json:"status"`
		Items    []RDSInstance `json:"items"`
	} `json:"rds_instances_to_create"`
	RDSInstancesToUpdate struct {
		Started  string        `json:"started"`
		Finished string        `json:"finished"`
		Status   string        `json:"status"`
		Items    []RDSInstance `json:"items"`
	} `json:"rds_instances_to_update"`
	RDSInstancesToDelete struct {
		Started  string        `json:"started"`
		Finished string        `json:"finished"`
		Status   string        `json:"status"`
		Items    []RDSInstance `json:"items"`
	} `json:"rds_instances_to_delete"`
	EBSVolumes struct {
		Started  string      `json:"started"`
		Finished string      `json:"finished"`
		Status   string      `json:"status"`
		Items    []EBSVolume `json:"items"`
	} `json:"ebs_volumes"`
	EBSVolumesToCreate struct {
		Started  string      `json:"started"`
		Finished string      `json:"finished"`
		Status   string      `json:"status"`
		Items    []EBSVolume `json:"items"`
	} `json:"ebs_volumes_to_create"`
	EBSVolumesToDelete struct {
		Started  string      `json:"started"`
		Finished string      `json:"finished"`
		Status   string      `json:"status"`
		Items    []EBSVolume `json:"items"`
	} `json:"ebs_volumes_to_delete"`
}

// DiffVPCs : Calculate diff on vpc component list
func (m *FSMMessage) DiffVPCs(om FSMMessage) {
	if len(om.VPCs.Items) > 0 {
		m.VPCs.Items = om.VPCs.Items
		m.VPCsToCreate.Items = []VPC{}
		return
	}

	for _, vpc := range m.VPCs.Items {
		if vpc.VpcID == "" {
			m.VPCsToCreate.Items = append(m.VPCsToCreate.Items, vpc)
		}
		if m.FindVPC(vpc.VpcID) == nil {
			vpc.Status = ""
			m.VPCsToDelete.Items = append(m.VPCsToDelete.Items, vpc)
		}
	}

	vpcs := []VPC{}
	for _, e := range m.VPCs.Items {
		toBeCreated := false
		for _, c := range m.VPCsToCreate.Items {
			if e.VpcID == c.VpcID {
				toBeCreated = true
			}
		}
		if toBeCreated == false {
			vpcs = append(vpcs, e)
		}
	}

	m.VPCs.Items = vpcs

}

// DiffNetworks : Calculate diff on network component list
func (m *FSMMessage) DiffNetworks(om FSMMessage) {
	for _, network := range m.Networks.Items {
		if o := om.FindNetwork(network.Name); o == nil {
			m.NetworksToCreate.Items = append(m.NetworksToCreate.Items, network)
		}
	}

	for _, network := range om.Networks.Items {
		if m.FindNetwork(network.Name) == nil {
			network.Status = ""
			m.NetworksToDelete.Items = append(m.NetworksToDelete.Items, network)
		}
	}

	var networks []Network
	for _, e := range m.Networks.Items {
		toBeCreated := false
		for _, c := range m.NetworksToCreate.Items {
			if e.Name == c.Name {
				toBeCreated = true
			}
		}
		if toBeCreated == false {
			networks = append(networks, e)
		}
	}
	m.Networks.Items = networks
}

// DiffEBSVolumes : Calculate diff on ebs component list
func (m *FSMMessage) DiffEBSVolumes(om FSMMessage) {
	for _, vol := range m.EBSVolumes.Items {
		if o := om.FindEBSVolume(vol.Name); o == nil {
			m.EBSVolumesToCreate.Items = append(m.EBSVolumesToCreate.Items, vol)
		}
	}

	for _, ebs := range om.EBSVolumes.Items {
		if m.FindEBSVolume(ebs.Name) == nil {
			ebs.Status = ""
			m.EBSVolumesToDelete.Items = append(m.EBSVolumesToDelete.Items, ebs)
		}
	}

	var vols []EBSVolume
	for _, e := range m.EBSVolumes.Items {
		toBeCreated := false
		for _, c := range m.EBSVolumesToCreate.Items {
			if e.Name == c.Name {
				toBeCreated = true
			}
		}
		if toBeCreated == false {
			vols = append(vols, e)
		}
	}
	m.EBSVolumes.Items = vols
}

// DiffInstances : Calculate diff on instance component list
func (m *FSMMessage) DiffInstances(om FSMMessage) {
	for _, instance := range m.Instances.Items {
		if oi := om.FindInstance(instance.Name); oi == nil {
			m.InstancesToCreate.Items = append(m.InstancesToCreate.Items, instance)
		} else if instance.HasChanged(oi) {
			m.InstancesToUpdate.Items = append(m.InstancesToUpdate.Items, instance)
		}
	}

	for _, instance := range om.Instances.Items {
		if m.FindInstance(instance.Name) == nil {
			instance.Status = ""
			m.InstancesToDelete.Items = append(m.InstancesToDelete.Items, instance)
		}
	}

	for _, instance := range om.InstancesToUpdate.Items {
		if instance.Status != "completed" {
			loaded := false
			exists := false
			for _, e := range m.InstancesToUpdate.Items {
				if e.Name == instance.Name {
					loaded = true
				}
			}
			for _, e := range m.Instances.Items {
				if e.Name == instance.Name {
					exists = true
				}
			}
			if exists == true && loaded == false {
				m.InstancesToUpdate.Items = append(m.InstancesToUpdate.Items, instance)
			}
		}
	}

	var instances []Instance
	for _, e := range m.Instances.Items {
		toBeCreated := false
		for _, c := range m.InstancesToCreate.Items {
			if e.Name == c.Name {
				toBeCreated = true
			}
		}
		if toBeCreated == false {
			instances = append(instances, e)
		}
	}
	m.Instances.Items = instances
}

// DiffFirewalls : Calculate diff on firewall component list
func (m *FSMMessage) DiffFirewalls(om FSMMessage) {
	for _, firewall := range m.Firewalls.Items {
		if of := om.FindFirewall(firewall.Name); of == nil {
			m.FirewallsToCreate.Items = append(m.FirewallsToCreate.Items, firewall)
		} else if firewall.HasChanged(of) {
			m.FirewallsToUpdate.Items = append(m.FirewallsToUpdate.Items, firewall)
		}
	}

	for _, firewall := range om.Firewalls.Items {
		if m.FindFirewall(firewall.Name) == nil {
			firewall.Status = ""
			m.FirewallsToDelete.Items = append(m.FirewallsToDelete.Items, firewall)
		}
	}

	for _, firewall := range om.FirewallsToUpdate.Items {
		if firewall.Status != "completed" {
			loaded := false
			exists := false
			for _, e := range m.FirewallsToUpdate.Items {
				if e.Name == firewall.Name {
					loaded = true
				}
			}
			for _, e := range m.Firewalls.Items {
				if e.Name == firewall.Name {
					exists = true
				}
			}
			if exists == true && loaded == false {
				m.FirewallsToUpdate.Items = append(m.FirewallsToUpdate.Items, firewall)
			}
		}
	}

	var firewalls []Firewall
	for _, e := range m.Firewalls.Items {
		toBeCreated := false
		for _, c := range m.FirewallsToCreate.Items {
			if e.Name == c.Name {
				toBeCreated = true
			}
		}
		if toBeCreated == false {
			firewalls = append(firewalls, e)
		}
	}
	m.Firewalls.Items = firewalls
}

// DiffNats : Calculate diff on nat component list
func (m *FSMMessage) DiffNats(om FSMMessage) {
	for _, nat := range m.Nats.Items {
		if on := om.FindNat(nat.Name); on == nil {
			m.NatsToCreate.Items = append(m.NatsToCreate.Items, nat)
		} else if nat.HasChanged(on) {
			m.NatsToUpdate.Items = append(m.NatsToUpdate.Items, nat)
		}
	}

	for _, nat := range om.Nats.Items {
		if m.FindNat(nat.Name) == nil {
			nat.Status = ""
			m.NatsToDelete.Items = append(m.NatsToDelete.Items, nat)
		}
	}

	for _, nat := range om.NatsToUpdate.Items {
		if nat.Status != "completed" {
			loaded := false
			exists := false
			for _, e := range m.NatsToUpdate.Items {
				if e.Name == nat.Name {
					loaded = true
				}
			}
			for _, e := range m.Nats.Items {
				if e.Name == nat.Name {
					exists = true
				}
			}
			if exists == true && loaded == false {
				m.NatsToUpdate.Items = append(m.NatsToUpdate.Items, nat)
			}
		}
	}

	var nats []Nat
	for _, e := range m.Nats.Items {
		toBeCreated := false
		for _, c := range m.NatsToCreate.Items {
			if e.Name == c.Name {
				toBeCreated = true
			}
		}
		if toBeCreated == false {
			nats = append(nats, e)
		}
	}
	m.Nats.Items = nats
}

// DiffELBs : Calculate diff on elb component list
func (m *FSMMessage) DiffELBs(om FSMMessage) {
	for _, elb := range m.ELBs.Items {
		if oe := om.FindELB(elb.Name); oe == nil {
			m.ELBsToCreate.Items = append(m.ELBsToCreate.Items, elb)
		} else if elb.HasChanged(oe) {
			m.ELBsToUpdate.Items = append(m.ELBsToUpdate.Items, elb)
		}
	}

	for _, elb := range om.ELBs.Items {
		if m.FindELB(elb.Name) == nil {
			elb.Status = ""
			m.ELBsToDelete.Items = append(m.ELBsToDelete.Items, elb)
		}
	}

	for _, elb := range om.ELBsToUpdate.Items {
		if elb.Status != "completed" {
			loaded := false
			exists := false
			for _, e := range m.ELBsToUpdate.Items {
				if e.Name == elb.Name {
					loaded = true
				}
			}
			for _, e := range m.ELBs.Items {
				if e.Name == elb.Name {
					exists = true
				}
			}
			if exists == true && loaded == false {
				m.ELBsToUpdate.Items = append(m.ELBsToUpdate.Items, elb)
			}
		}
	}

	var elbs []ELB
	for _, e := range m.ELBs.Items {
		toBeCreated := false
		for _, c := range m.ELBsToCreate.Items {
			if e.Name == c.Name {
				toBeCreated = true
			}
		}
		if toBeCreated == false {
			elbs = append(elbs, e)
		}
	}
	m.ELBs.Items = elbs
}

// DiffS3s : Calculate diff on s3 bucket component list
func (m *FSMMessage) DiffS3s(om FSMMessage) {
	for _, s3 := range m.S3s.Items {
		if oe := om.FindS3(s3.Name); oe == nil {
			m.S3sToCreate.Items = append(m.S3sToCreate.Items, s3)
		} else if s3.HasChanged(oe) {
			m.S3sToUpdate.Items = append(m.S3sToUpdate.Items, s3)
		}
	}

	for _, s3 := range om.S3s.Items {
		if m.FindS3(s3.Name) == nil {
			s3.Status = ""
			m.S3sToDelete.Items = append(m.S3sToDelete.Items, s3)
		}
	}

	for _, s3 := range om.S3sToUpdate.Items {
		if s3.Status != "completed" {
			loaded := false
			exists := false
			for _, e := range m.S3sToUpdate.Items {
				if e.Name == s3.Name {
					loaded = true
				}
			}
			for _, e := range m.S3s.Items {
				if e.Name == s3.Name {
					exists = true
				}
			}
			if exists == true && loaded == false {
				m.S3sToUpdate.Items = append(m.S3sToUpdate.Items, s3)
			}
		}
	}

	var s3buckets []S3
	for _, s := range m.S3s.Items {
		toBeCreated := false
		for _, c := range m.S3sToCreate.Items {
			if s.Name == c.Name {
				toBeCreated = true
			}
		}
		if toBeCreated == false {
			s3buckets = append(s3buckets, s)
		}
	}
	m.S3s.Items = s3buckets
}

// DiffRoute53s : Calculate diff on route53 zone component list
func (m *FSMMessage) DiffRoute53s(om FSMMessage) {
	for _, route53 := range m.Route53s.Items {
		if oe := om.FindRoute53(route53.Name); oe == nil {
			m.Route53sToCreate.Items = append(m.Route53sToCreate.Items, route53)
		} else if route53.HasChanged(oe) {
			m.Route53sToUpdate.Items = append(m.Route53sToUpdate.Items, route53)
		}
	}

	for _, route53 := range om.Route53s.Items {
		if m.FindRoute53(route53.Name) == nil {
			route53.Status = ""
			m.Route53sToDelete.Items = append(m.Route53sToDelete.Items, route53)
		}
	}

	for _, route53 := range om.Route53sToUpdate.Items {
		if route53.Status != "completed" {
			loaded := false
			exists := false
			for _, e := range m.Route53sToUpdate.Items {
				if e.Name == route53.Name {
					loaded = true
				}
			}
			for _, e := range m.Route53s.Items {
				if e.Name == route53.Name {
					exists = true
				}
			}
			if exists == true && loaded == false {
				m.Route53sToUpdate.Items = append(m.Route53sToUpdate.Items, route53)
			}
		}
	}

	var route53zones []Route53Zone
	for _, s := range m.Route53s.Items {
		toBeCreated := false
		for _, c := range m.Route53sToCreate.Items {
			if s.Name == c.Name {
				toBeCreated = true
			}
		}
		if toBeCreated == false {
			route53zones = append(route53zones, s)
		}
	}
	m.Route53s.Items = route53zones
}

// DiffRDSClusters : Calculate diff on rds cluster component list
func (m *FSMMessage) DiffRDSClusters(om FSMMessage) {
	for _, rdcs := range m.RDSClusters.Items {
		if or := om.FindRDSCluster(rdcs.Name); or == nil {
			m.RDSClustersToCreate.Items = append(m.RDSClustersToCreate.Items, rdcs)
		} else if rdcs.HasChanged(or) {
			m.RDSClustersToUpdate.Items = append(m.RDSClustersToUpdate.Items, rdcs)
		}
	}

	for _, rdcs := range om.RDSClusters.Items {
		if m.FindRDSCluster(rdcs.Name) == nil {
			rdcs.Status = ""
			m.RDSClustersToDelete.Items = append(m.RDSClustersToDelete.Items, rdcs)
		}
	}

	for _, rdcs := range om.RDSClustersToUpdate.Items {
		if rdcs.Status != "completed" {
			loaded := false
			exists := false
			for _, r := range m.RDSClustersToUpdate.Items {
				if r.Name == rdcs.Name {
					loaded = true
				}
			}
			for _, r := range m.RDSClusters.Items {
				if r.Name == rdcs.Name {
					exists = true
				}
			}
			if exists == true && loaded == false {
				m.RDSClustersToUpdate.Items = append(m.RDSClustersToUpdate.Items, rdcs)
			}
		}
	}

	var rdcss []RDSCluster
	for _, r := range m.RDSClusters.Items {
		toBeCreated := false
		for _, c := range m.RDSClustersToCreate.Items {
			if r.Name == c.Name {
				toBeCreated = true
			}
		}
		if toBeCreated == false {
			rdcss = append(rdcss, r)
		}
	}
	m.RDSClusters.Items = rdcss
}

// DiffRDSInstances : Calculate diff on rds instance component list
func (m *FSMMessage) DiffRDSInstances(om FSMMessage) {
	for _, rdcs := range m.RDSInstances.Items {
		if or := om.FindRDSInstance(rdcs.Name); or == nil {
			m.RDSInstancesToCreate.Items = append(m.RDSInstancesToCreate.Items, rdcs)
		} else if rdcs.HasChanged(or) {
			m.RDSInstancesToUpdate.Items = append(m.RDSInstancesToUpdate.Items, rdcs)
		}
	}

	for _, rdcs := range om.RDSInstances.Items {
		if m.FindRDSInstance(rdcs.Name) == nil {
			rdcs.Status = ""
			m.RDSInstancesToDelete.Items = append(m.RDSInstancesToDelete.Items, rdcs)
		}
	}

	for _, rdcs := range om.RDSInstancesToUpdate.Items {
		if rdcs.Status != "completed" {
			loaded := false
			exists := false
			for _, r := range m.RDSInstancesToUpdate.Items {
				if r.Name == rdcs.Name {
					loaded = true
				}
			}
			for _, r := range m.RDSInstances.Items {
				if r.Name == rdcs.Name {
					exists = true
				}
			}
			if exists == true && loaded == false {
				m.RDSInstancesToUpdate.Items = append(m.RDSInstancesToUpdate.Items, rdcs)
			}
		}
	}

	var rdcss []RDSInstance
	for _, r := range m.RDSInstances.Items {
		toBeCreated := false
		for _, c := range m.RDSInstancesToCreate.Items {
			if r.Name == c.Name {
				toBeCreated = true
			}
		}
		if toBeCreated == false {
			rdcss = append(rdcss, r)
		}
	}
	m.RDSInstances.Items = rdcss
}

// Diff compares against an existing FSMMessage from a previous fsm message
func (m *FSMMessage) Diff(om FSMMessage) {
	m.DiffVPCs(om)
	m.DiffNetworks(om)
	m.DiffInstances(om)
	m.DiffFirewalls(om)
	m.DiffNats(om)
	m.DiffELBs(om)
	m.DiffS3s(om)
	m.DiffRoute53s(om)
	m.DiffRDSClusters(om)
	m.DiffRDSInstances(om)
	m.DiffEBSVolumes(om)
}

// GenerateWorkflow creates a fsm workflow based upon actionable tasks, such as creation or deletion of an entity.
func (m *FSMMessage) GenerateWorkflow(path string) error {
	w := workflow.New()
	err := w.LoadFile("./output/arcs/" + path)
	if err != nil {
		return err
	}

	for i := range m.VPCsToCreate.Items {
		m.VPCsToCreate.Items[i].Status = ""
	}
	for i := range m.VPCsToDelete.Items {
		m.VPCsToDelete.Items[i].Status = ""
	}
	for i := range m.NetworksToCreate.Items {
		m.NetworksToCreate.Items[i].Status = ""
	}
	for i := range m.NetworksToDelete.Items {
		m.NetworksToDelete.Items[i].Status = ""
	}
	for i := range m.InstancesToCreate.Items {
		m.InstancesToCreate.Items[i].Status = ""
	}
	for i := range m.InstancesToUpdate.Items {
		m.InstancesToUpdate.Items[i].Status = ""
	}
	for i := range m.InstancesToDelete.Items {
		m.InstancesToDelete.Items[i].Status = ""
	}
	for i := range m.FirewallsToCreate.Items {
		m.FirewallsToCreate.Items[i].Status = ""
	}
	for i := range m.FirewallsToUpdate.Items {
		m.FirewallsToUpdate.Items[i].Status = ""
	}
	for i := range m.FirewallsToDelete.Items {
		m.FirewallsToDelete.Items[i].Status = ""
	}
	for i := range m.NatsToCreate.Items {
		m.NatsToCreate.Items[i].Status = ""
	}
	for i := range m.NatsToUpdate.Items {
		m.NatsToUpdate.Items[i].Status = ""
	}
	for i := range m.NatsToDelete.Items {
		m.NatsToDelete.Items[i].Status = ""
	}
	for i := range m.ELBsToCreate.Items {
		m.ELBsToCreate.Items[i].Status = ""
	}
	for i := range m.ELBsToUpdate.Items {
		m.ELBsToUpdate.Items[i].Status = ""
	}
	for i := range m.ELBsToDelete.Items {
		m.ELBsToDelete.Items[i].Status = ""
	}
	for i := range m.S3sToCreate.Items {
		m.S3sToCreate.Items[i].Status = ""
	}
	for i := range m.S3sToUpdate.Items {
		m.S3sToUpdate.Items[i].Status = ""
	}
	for i := range m.S3sToDelete.Items {
		m.S3sToDelete.Items[i].Status = ""
	}
	for i := range m.Route53sToCreate.Items {
		m.Route53sToCreate.Items[i].Status = ""
	}
	for i := range m.Route53sToUpdate.Items {
		m.Route53sToUpdate.Items[i].Status = ""
	}
	for i := range m.Route53sToDelete.Items {
		m.Route53sToDelete.Items[i].Status = ""
	}
	for i := range m.EBSVolumesToCreate.Items {
		m.EBSVolumesToCreate.Items[i].Status = ""
	}
	for i := range m.EBSVolumesToDelete.Items {
		m.EBSVolumesToDelete.Items[i].Status = ""
	}

	// Set vpc items
	w.SetCount("creating_vpcs", len(m.VPCsToCreate.Items))
	w.SetCount("vpcs_created", len(m.VPCsToCreate.Items))
	w.SetCount("deleting_vpcs", len(m.VPCsToDelete.Items))
	w.SetCount("vpcs_deleted", len(m.VPCsToDelete.Items))

	// Set network items
	w.SetCount("creating_networks", len(m.NetworksToCreate.Items))
	w.SetCount("networks_created", len(m.NetworksToCreate.Items))
	w.SetCount("deleting_networks", len(m.NetworksToDelete.Items))
	w.SetCount("networks_deleted", len(m.NetworksToDelete.Items))

	// Set instance items
	w.SetCount("creating_instances", len(m.InstancesToCreate.Items))
	w.SetCount("instances_created", len(m.InstancesToCreate.Items))
	w.SetCount("updating_instances", len(m.InstancesToUpdate.Items))
	w.SetCount("instances_updated", len(m.InstancesToUpdate.Items))
	w.SetCount("deleting_instances", len(m.InstancesToDelete.Items))
	w.SetCount("instances_deleted", len(m.InstancesToDelete.Items))

	// Set firewall items
	w.SetCount("creating_firewalls", len(m.FirewallsToCreate.Items))
	w.SetCount("firewalls_created", len(m.FirewallsToCreate.Items))
	w.SetCount("updating_firewalls", len(m.FirewallsToUpdate.Items))
	w.SetCount("firewalls_updated", len(m.FirewallsToUpdate.Items))
	w.SetCount("deleting_firewalls", len(m.FirewallsToDelete.Items))
	w.SetCount("firewalls_deleted", len(m.FirewallsToDelete.Items))

	// Set nat items
	w.SetCount("creating_nats", len(m.NatsToCreate.Items))
	w.SetCount("nats_created", len(m.NatsToCreate.Items))
	w.SetCount("updating_nats", len(m.NatsToUpdate.Items))
	w.SetCount("nats_updated", len(m.NatsToUpdate.Items))
	w.SetCount("deleting_nats", len(m.NatsToDelete.Items))
	w.SetCount("nats_deleted", len(m.NatsToDelete.Items))

	// Set elb items
	w.SetCount("creating_elbs", len(m.ELBsToCreate.Items))
	w.SetCount("elbs_created", len(m.ELBsToCreate.Items))
	w.SetCount("updating_elbs", len(m.ELBsToUpdate.Items))
	w.SetCount("elbs_updated", len(m.ELBsToUpdate.Items))
	w.SetCount("deleting_elbs", len(m.ELBsToDelete.Items))
	w.SetCount("elbs_deleted", len(m.ELBsToDelete.Items))

	// Set s3 items
	w.SetCount("creating_s3s", len(m.S3sToCreate.Items))
	w.SetCount("s3s_created", len(m.S3sToCreate.Items))
	w.SetCount("updating_s3s", len(m.S3sToUpdate.Items))
	w.SetCount("s3s_updated", len(m.S3sToUpdate.Items))
	w.SetCount("deleting_s3s", len(m.S3sToDelete.Items))
	w.SetCount("s3s_deleted", len(m.S3sToDelete.Items))

	// Set route53 items
	w.SetCount("creating_route53s", len(m.Route53sToCreate.Items))
	w.SetCount("route53s_created", len(m.Route53sToCreate.Items))
	w.SetCount("updating_route53s", len(m.Route53sToUpdate.Items))
	w.SetCount("route53s_updated", len(m.Route53sToUpdate.Items))
	w.SetCount("deleting_route53s", len(m.Route53sToDelete.Items))
	w.SetCount("route53s_deleted", len(m.Route53sToDelete.Items))

	// Set rds_cluster items
	w.SetCount("creating_rds_clusters", len(m.RDSClustersToCreate.Items))
	w.SetCount("rds_clusters_created", len(m.RDSClustersToCreate.Items))
	w.SetCount("updating_rds_clusters", len(m.RDSClustersToUpdate.Items))
	w.SetCount("rds_clusters_updated", len(m.RDSClustersToUpdate.Items))
	w.SetCount("deleting_rds_clusters", len(m.RDSClustersToDelete.Items))
	w.SetCount("rds_clusters_deleted", len(m.RDSClustersToDelete.Items))

	// Set rds_instance items
	w.SetCount("creating_rds_instances", len(m.RDSInstancesToCreate.Items))
	w.SetCount("rds_instances_created", len(m.RDSInstancesToCreate.Items))
	w.SetCount("updating_rds_instances", len(m.RDSInstancesToUpdate.Items))
	w.SetCount("rds_instances_updated", len(m.RDSInstancesToUpdate.Items))
	w.SetCount("deleting_rds_instances", len(m.RDSInstancesToDelete.Items))
	w.SetCount("rds_instances_deleted", len(m.RDSInstancesToDelete.Items))

	// Set ebs_volume items
	w.SetCount("creating_ebs_volumes", len(m.EBSVolumesToCreate.Items))
	w.SetCount("ebs_volumes_created", len(m.EBSVolumesToCreate.Items))
	w.SetCount("deleting_ebs_volumes", len(m.EBSVolumesToDelete.Items))
	w.SetCount("ebs_volumes_deleted", len(m.EBSVolumesToDelete.Items))

	// Optimize the graph, removing unused arcs/verticies
	if err := w.Optimize(); err != nil {
		return err
	}

	m.Workflow.Arcs = w.Arcs()

	return nil
}

// FindVPC returns true if a router with a given name exists
func (m *FSMMessage) FindVPC(awsid string) *VPC {
	for i, vpc := range m.VPCs.Items {
		if vpc.VpcID == awsid {
			return &m.VPCs.Items[i]
		}
	}
	return nil
}

// FindNetwork returns true if a network with a given name exists
func (m *FSMMessage) FindNetwork(name string) *Network {
	for i, network := range m.Networks.Items {
		if network.Name == name {
			return &m.Networks.Items[i]
		}
	}
	return nil
}

// FindInstance returns true if an instance with a given name exists
func (m *FSMMessage) FindInstance(name string) *Instance {
	for i, instance := range m.Instances.Items {
		if instance.Name == name {
			return &m.Instances.Items[i]
		}
	}
	return nil
}

// FindFirewall returns true if a firewall with a given name exists
func (m *FSMMessage) FindFirewall(name string) *Firewall {
	for i, firewall := range m.Firewalls.Items {
		if firewall.Name == name {
			return &m.Firewalls.Items[i]
		}
	}
	return nil
}

// FindNat returns true if a nat with a given name exists
func (m *FSMMessage) FindNat(name string) *Nat {
	for i, nat := range m.Nats.Items {
		if nat.Name == name {
			return &m.Nats.Items[i]
		}
	}
	return nil
}

// FindELB returns true if an elb with a given name exists
func (m *FSMMessage) FindELB(name string) *ELB {
	for i, elb := range m.ELBs.Items {
		if elb.Name == name {
			return &m.ELBs.Items[i]
		}
	}
	return nil
}

// FindS3 returns true if an s3 bucket with a given name exists
func (m *FSMMessage) FindS3(name string) *S3 {
	for i, s3 := range m.S3s.Items {
		if s3.Name == name {
			return &m.S3s.Items[i]
		}
	}
	return nil
}

// FindRDSCluster returns true if an rds cluster with a given name exists
func (m *FSMMessage) FindRDSCluster(name string) *RDSCluster {
	for i, rdsc := range m.RDSClusters.Items {
		if rdsc.Name == name {
			return &m.RDSClusters.Items[i]
		}
	}
	return nil
}

// FindRDSInstance returns true if an rds instance with a given name exists
func (m *FSMMessage) FindRDSInstance(name string) *RDSInstance {
	for i, rdsi := range m.RDSInstances.Items {
		if rdsi.Name == name {
			return &m.RDSInstances.Items[i]
		}
	}
	return nil
}

// FindRoute53 returns true if an route53 bucket with a given name exists
func (m *FSMMessage) FindRoute53(name string) *Route53Zone {
	for i, route53 := range m.Route53s.Items {
		if route53.Name == name {
			return &m.Route53s.Items[i]
		}
	}
	return nil
}

// FindEBSVolume returns a ebs volume matching a given name
func (m *FSMMessage) FindEBSVolume(name string) *EBSVolume {
	for i, ebs := range m.EBSVolumes.Items {
		if ebs.Name == name {
			return &m.EBSVolumes.Items[i]
		}
	}
	return nil
}

// FilterNewInstances will return any new instances that match a certain pattern
func (m *FSMMessage) FilterNewInstances(name string) []Instance {
	var instances []Instance
	for _, instance := range m.InstancesToCreate.Items {
		if strings.Contains(instance.Name, name) {
			instances = append(instances, instance)
		}
	}
	return instances
}

// ToJSON : Get this service as a json
func (m *FSMMessage) ToJSON() []byte {
	json, _ := json.Marshal(m)

	return json
}

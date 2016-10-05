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
	} `json:"nats"`
	ELBsToCreate struct {
		Started  string `json:"started"`
		Finished string `json:"finished"`
		Status   string `json:"status"`
		Items    []ELB  `json:"items"`
	} `json:"nats_to_create"`
	ELBsToUpdate struct {
		Started  string `json:"started"`
		Finished string `json:"finished"`
		Status   string `json:"status"`
		Items    []ELB  `json:"items"`
	} `json:"nats_to_update"`
	ELBsToDelete struct {
		Started  string `json:"started"`
		Finished string `json:"finished"`
		Status   string `json:"status"`
		Items    []ELB  `json:"items"`
	} `json:"nats_to_delete"`
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

	networks := []Network{}
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

	instances := []Instance{}
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

	firewalls := []Firewall{}
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

	nats := []Nat{}
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

// Diff compares against an existing FSMMessage from a previous fsm message
func (m *FSMMessage) Diff(om FSMMessage) {
	m.DiffVPCs(om)
	m.DiffNetworks(om)
	m.DiffInstances(om)
	m.DiffFirewalls(om)
	m.DiffNats(om)
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

	// Optimize the graph, removing unused arcs/verticies
	w.Optimize()

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

// FindNetwork returns true if a router with a given name exists
func (m *FSMMessage) FindNetwork(name string) *Network {
	for i, network := range m.Networks.Items {
		if network.Name == name {
			return &m.Networks.Items[i]
		}
	}
	return nil
}

// FindInstance returns true if a router with a given name exists
func (m *FSMMessage) FindInstance(name string) *Instance {
	for i, instance := range m.Instances.Items {
		if instance.Name == name {
			return &m.Instances.Items[i]
		}
	}
	return nil
}

// FindFirewall returns true if a router with a given name exists
func (m *FSMMessage) FindFirewall(name string) *Firewall {
	for i, firewall := range m.Firewalls.Items {
		if firewall.Name == name {
			return &m.Firewalls.Items[i]
		}
	}
	return nil
}

// FindNat returns true if a router with a given name exists
func (m *FSMMessage) FindNat(name string) *Nat {
	for i, nat := range m.Nats.Items {
		if nat.Name == name {
			return &m.Nats.Items[i]
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

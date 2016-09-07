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
}

// Diff compares against an existing FSMMessage from a previous fsm message
func (m *FSMMessage) Diff(om FSMMessage) {
	// build new networks
	for _, network := range m.Networks.Items {
		if o := om.FindNetwork(network.Name); o == nil {
			m.NetworksToCreate.Items = append(m.NetworksToCreate.Items, network)
		}
	}

	// build old networks to delete
	for _, network := range om.Networks.Items {
		if m.FindNetwork(network.Name) == nil {
			m.NetworksToDelete.Items = append(m.NetworksToDelete.Items, network)
		}
	}

	// build new and existing instances
	for _, instance := range m.Instances.Items {
		if oi := om.FindInstance(instance.Name); oi == nil {
			m.InstancesToCreate.Items = append(m.InstancesToCreate.Items, instance)
		} else if instance.HasChanged(oi) {
			m.InstancesToUpdate.Items = append(m.InstancesToUpdate.Items, instance)
		}
	}

	// build old instances to delete
	for _, instance := range om.Instances.Items {
		if m.FindInstance(instance.Name) == nil {
			m.InstancesToDelete.Items = append(m.InstancesToDelete.Items, instance)
		}
	}

	// build new and existing firewalls
	for _, firewall := range m.Firewalls.Items {
		if of := om.FindFirewall(firewall.Name); of == nil {
			m.FirewallsToCreate.Items = append(m.FirewallsToCreate.Items, firewall)
		} else if firewall.HasChanged(of) {
			m.FirewallsToUpdate.Items = append(m.FirewallsToUpdate.Items, firewall)
		}
	}

	// build old firewalls to delete
	for _, firewall := range om.Firewalls.Items {
		if m.FindFirewall(firewall.Name) == nil {
			m.FirewallsToDelete.Items = append(m.FirewallsToDelete.Items, firewall)
		}
	}

	// build new and existing nats
	for _, nat := range m.Nats.Items {
		if on := om.FindNat(nat.Name); on == nil {
			m.NatsToCreate.Items = append(m.NatsToCreate.Items, nat)
		} else if nat.HasChanged(on) {
			m.NatsToUpdate.Items = append(m.NatsToUpdate.Items, nat)
		}
	}

	// build old nats to delete
	for _, nat := range om.Nats.Items {
		if m.FindNat(nat.Name) == nil {
			m.NatsToDelete.Items = append(m.NatsToDelete.Items, nat)
		}
	}
}

// GenerateWorkflow creates a fsm workflow based upon actionable tasks, such as creation or deletion of an entity.
func (m *FSMMessage) GenerateWorkflow(path string) error {
	w := workflow.New()
	err := w.LoadFile("./output/arcs/" + path)
	if err != nil {
		return err
	}

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

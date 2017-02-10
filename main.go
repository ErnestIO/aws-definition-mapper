/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"runtime"
	"time"

	"github.com/ernestio/aws-definition-mapper/definition"
	"github.com/ernestio/aws-definition-mapper/mapper"
	"github.com/ernestio/aws-definition-mapper/output"
	ecc "github.com/ernestio/ernest-config-client"
	"github.com/ghodss/yaml"
	"github.com/nats-io/nats"
)

var nc *nats.Conn
var natsErr error

func main() {
	nc = ecc.NewConfig(os.Getenv("NATS_URI")).Nats()

	if _, err := nc.Subscribe("definition.map.creation.aws", createDefinitionHandler); err != nil {
		log.Println(err)
	}
	if _, err := nc.Subscribe("definition.map.deletion.aws", deleteDefinitionHandler); err != nil {
		log.Println(err)
	}
	if _, err := nc.Subscribe("definition.map.import.aws", importDefinitionHandler); err != nil {
		log.Println(err)
	}

	if _, err := nc.Subscribe("service.import.aws.done", importDoneHandler); err != nil {
		log.Println(err)
	}

	runtime.Goexit()
}

func createDefinitionHandler(msg *nats.Msg) {
	var om output.FSMMessage

	p, err := definition.PayloadFromJSON(msg.Data)
	if err != nil {
		log.Println("ERROR: failed to parse payload")
		if err := nc.Publish(msg.Reply, []byte(`{"error":"Failed to parse payload."}`)); err != nil {
			log.Println(err)
		}
		return
	}

	err = p.Service.Validate()
	if err != nil {
		log.Println("ERROR: " + err.Error())
		if err := nc.Publish(msg.Reply, []byte(`{"error":"`+err.Error()+`"}`)); err != nil {
			log.Println(err)
		}
		return
	}

	// new fsm message
	m := mapper.ConvertPayload(p)

	// previous output message if it exists
	if p.PrevID != "" {
		om, err = getPreviousServiceMapping(p.PrevID)
		if err != nil {
			log.Println("ERROR: failed to get previous output")
			if err := nc.Publish(msg.Reply, []byte(`{"error":"Failed to get previous output."}`)); err != nil {
				log.Println(err)
			}
			return
		}

		if p.Service.VpcID != "" && p.Service.VpcID != om.VPCs.Items[0].VpcID {
			log.Println("ERROR: VPC ID cannot change between builds.")
			if err := nc.Publish(msg.Reply, []byte(`{"error":"VPC ID cannot change between builds."}`)); err != nil {
				log.Println(err)
			}
			return
		}
	}

	// Map provider data from previous build
	mapper.MapProviderData(m, &om)

	// Check for changes and create workflow arcs
	m.Diff(om)

	err = m.GenerateWorkflow("create-workflow.json")
	if err != nil {
		log.Println(err.Error())
		if err := nc.Publish(msg.Reply, []byte(`{"error":"Could not generate workflow."}`)); err != nil {
			log.Println(err)
		}
		return
	}

	data, err := json.Marshal(m)
	if err != nil {
		if err := nc.Publish(msg.Reply, []byte(`{"error":"Failed marshal output message."}`)); err != nil {
			log.Println(err)
		}
		return
	}

	if err := nc.Publish(msg.Reply, data); err != nil {
		log.Println(err)
	}
}

func deleteDefinitionHandler(msg *nats.Msg) {
	p, err := definition.PayloadFromJSON(msg.Data)
	if err != nil {
		if err := nc.Publish(msg.Reply, []byte(`{"error":"Failed to parse payload."}`)); err != nil {
			log.Println(err)
		}
		return
	}

	m, err := getPreviousServiceMapping(p.PrevID)
	if err != nil {
		if err := nc.Publish(msg.Reply, []byte(`{"error":"Failed to get previous output."}`)); err != nil {
			log.Println(err)
		}
		return
	}

	// Assign all items to delete
	m.NetworksToDelete = m.Networks
	for i := range m.NetworksToDelete.Items {
		m.NetworksToDelete.Items[i].Status = ""
	}
	m.InstancesToDelete = m.Instances
	for i := range m.InstancesToDelete.Items {
		m.InstancesToDelete.Items[i].Status = ""
	}
	m.FirewallsToDelete = m.Firewalls
	for i := range m.FirewallsToDelete.Items {
		m.FirewallsToDelete.Items[i].Status = ""
	}
	m.NatsToDelete = m.Nats
	for i := range m.NatsToDelete.Items {
		m.NatsToDelete.Items[i].Status = ""
	}
	m.ELBsToDelete = m.ELBs
	for i := range m.ELBsToDelete.Items {
		m.ELBsToDelete.Items[i].Status = ""
	}
	m.S3sToDelete = m.S3s
	for i := range m.S3sToDelete.Items {
		m.S3sToDelete.Items[i].Status = ""
	}
	m.Route53sToDelete = m.Route53s
	for i := range m.Route53sToDelete.Items {
		m.Route53sToDelete.Items[i].Status = ""
	}

	m.RDSClustersToDelete = m.RDSClusters
	for i := range m.RDSClustersToDelete.Items {
		m.RDSClustersToDelete.Items[i].Status = ""
	}

	m.RDSInstancesToDelete = m.RDSInstances
	for i := range m.RDSInstancesToDelete.Items {
		m.RDSInstancesToDelete.Items[i].Status = ""
	}

	m.EBSVolumesToDelete = m.EBSVolumes
	for i := range m.EBSVolumesToDelete.Items {
		m.EBSVolumesToDelete.Items[i].Status = ""
	}

	// Generate delete workflow
	if err := m.GenerateWorkflow("delete-workflow.json"); err != nil {
		log.Println(err)
	}

	data, err := json.Marshal(m)
	if err != nil {
		if err := nc.Publish(msg.Reply, []byte(`{"error":"Failed marshal output message."}`)); err != nil {
			log.Println(err)
		}
		return
	}

	if err := nc.Publish(msg.Reply, data); err != nil {
		log.Println(err)
	}
}

func importDefinitionHandler(msg *nats.Msg) {
	var om output.FSMMessage

	p, err := definition.PayloadFromJSON(msg.Data)
	if err != nil {
		log.Println("ERROR: failed to parse payload")
		if err := nc.Publish(msg.Reply, []byte(`{"error":"Failed to parse payload."}`)); err != nil {
			log.Println(err)
		}
		return
	}

	// new fsm message
	m := mapper.ConvertPayload(p)

	// previous output message if it exists
	if p.PrevID != "" {
		om, err = getPreviousServiceMapping(p.PrevID)
		if err != nil {
			log.Println("ERROR: failed to get previous output")
			if err := nc.Publish(msg.Reply, []byte(`{"error":"Failed to get previous output."}`)); err != nil {
				log.Println(err)
			}
			return
		}
	}

	m.VPCs.Items = []output.VPC{}

	err = m.GenerateWorkflow("import-workflow.json")
	if err != nil {
		log.Println(err.Error())
		if err := nc.Publish(msg.Reply, []byte(`{"error":"Could not generate workflow."}`)); err != nil {
			log.Println(err)
		}
		return
	}

	data, err := json.Marshal(m)
	if err != nil {
		if err := nc.Publish(msg.Reply, []byte(`{"error":"Failed marshal output message."}`)); err != nil {
			log.Println(err)
		}
		return
	}

	// Map provider data from previous build
	mapper.MapProviderData(m, &om)

	// Check for changes and create workflow arcs
	m.Diff(om)

	if err := nc.Publish(msg.Reply, data); err != nil {
		log.Println(err)
	}
}

func importDoneHandler(msg *nats.Msg) {
	var m output.FSMMessage

	err := json.Unmarshal(msg.Data, &m)
	if err != nil {
		log.Println(err)
		return
	}

	if m.Type != "aws" && m.Type != "aws-fake" {
		return
	}

	// Set missing values on fsm message
	mapper.UpdateFSMMessageValues(&m)

	// convert the payload to a definition
	d := mapper.ConvertFSMMessage(&m)

	dj, err := json.Marshal(d)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(string(dj))

	dy, err := yaml.JSONToYAML(dj)
	if err != nil {
		log.Println(err)
		return
	}

	s := output.Service{
		ID:         m.ID,
		Definition: string(dy),
	}

	data, err := json.Marshal(s)
	if err != nil {
		log.Println(err)
		return
	}

	_, err = nc.Request("service.set.definition", data, time.Second)
	if err != nil {
		log.Println(err)
		return
	}

	mapping, err := json.Marshal(m)
	if err != nil {
		log.Println(err)
		return
	}

	s = output.Service{
		ID:      m.ID,
		Mapping: string(mapping),
	}

	data, err = json.Marshal(s)
	if err != nil {
		log.Println(err)
		return
	}

	_, err = nc.Request("service.set.mapping", data, time.Second)
	if err != nil {
		log.Println(err)
		return
	}

	if err := nc.Publish("service.import.done", mapping); err != nil {
		log.Println(err)
	}
}

func getPreviousServiceMapping(id string) (output.FSMMessage, error) {
	var payload output.FSMMessage

	msg, err := nc.Request("service.get.mapping", []byte(`{"id":"`+id+`"}`), time.Second)
	if err != nil {
		log.Println(err.Error())
		return payload, err
	}

	if err := json.Unmarshal(msg.Data, &payload); err != nil {
		return payload, err
	}

	return payload, nil
}

func setServiceDefinition(s *output.Service) error {
	data, err := json.Marshal(s)
	if err != nil {
		return err
	}

	_, err = nc.Request("service.set.definition", data, time.Second)

	return err
}

/*
 * This file is automatically generated
 */

package v1alpha1

import (
	"encoding/json"

	apiv1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/api/v1"
)

var (
	_CommandLineInterfaceGVK = apiv1.GroupVersionKind{
		GroupKind: apiv1.GroupKind{
			Group: "definitions",
			Kind:  "CommandLineInterface",
		},
		APIVersion: "v1alpha1",
	}
)

const (
	CommandLineInterfaceScope = "ResourceGroup"

	CommandLineInterfaceResource = "commandlines"
)

func CommandLineInterfaceGVK() apiv1.GroupVersionKind {
	return _CommandLineInterfaceGVK
}

func init() {
	apiv1.RegisterGVK(_CommandLineInterfaceGVK, CommandLineInterfaceScope, CommandLineInterfaceResource)
}

// CommandLineInterface Resource
type CommandLineInterface struct {
	apiv1.ResourceMeta

	Spec CommandLineInterfaceSpec `json:"spec"`
}

// FromInstance converts a ResourceInstance to a CommandLineInterface
func (res *CommandLineInterface) FromInstance(ri *apiv1.ResourceInstance) error {
	if ri == nil {
		res = nil
		return nil
	}

	m, err := json.Marshal(ri.Spec)
	if err != nil {
		return err
	}

	spec := &CommandLineInterfaceSpec{}
	err = json.Unmarshal(m, spec)
	if err != nil {
		return err
	}

	*res = CommandLineInterface{ResourceMeta: ri.ResourceMeta, Spec: *spec}

	return err
}

// AsInstance converts a CommandLineInterface to a ResourceInstance
func (res *CommandLineInterface) AsInstance() (*apiv1.ResourceInstance, error) {
	m, err := json.Marshal(res.Spec)
	if err != nil {
		return nil, err
	}

	spec := map[string]interface{}{}
	err = json.Unmarshal(m, &spec)
	if err != nil {
		return nil, err
	}

	meta := res.ResourceMeta
	meta.GroupVersionKind = CommandLineInterfaceGVK()

	return &apiv1.ResourceInstance{ResourceMeta: meta, Spec: spec}, nil
}

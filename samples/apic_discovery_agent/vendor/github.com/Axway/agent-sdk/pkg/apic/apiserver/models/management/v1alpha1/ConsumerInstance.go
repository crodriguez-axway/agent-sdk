/*
 * This file is automatically generated
 */

package v1alpha1

import (
	"encoding/json"

	apiv1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/api/v1"
)

var (
	_ConsumerInstanceGVK = apiv1.GroupVersionKind{
		GroupKind: apiv1.GroupKind{
			Group: "management",
			Kind:  "ConsumerInstance",
		},
		APIVersion: "v1alpha1",
	}
)

const (
	ConsumerInstanceScope = "Environment"

	ConsumerInstanceResource = "consumerinstances"
)

func ConsumerInstanceGVK() apiv1.GroupVersionKind {
	return _ConsumerInstanceGVK
}

func init() {
	apiv1.RegisterGVK(_ConsumerInstanceGVK, ConsumerInstanceScope, ConsumerInstanceResource)
}

// ConsumerInstance Resource
type ConsumerInstance struct {
	apiv1.ResourceMeta

	Spec ConsumerInstanceSpec `json:"spec"`

	Status ConsumerInstanceStatus `json:"status"`
}

// FromInstance converts a ResourceInstance to a ConsumerInstance
func (res *ConsumerInstance) FromInstance(ri *apiv1.ResourceInstance) error {
	if ri == nil {
		res = nil
		return nil
	}

	m, err := json.Marshal(ri.Spec)
	if err != nil {
		return err
	}

	spec := &ConsumerInstanceSpec{}
	err = json.Unmarshal(m, spec)
	if err != nil {
		return err
	}

	*res = ConsumerInstance{ResourceMeta: ri.ResourceMeta, Spec: *spec}

	return err
}

// AsInstance converts a ConsumerInstance to a ResourceInstance
func (res *ConsumerInstance) AsInstance() (*apiv1.ResourceInstance, error) {
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
	meta.GroupVersionKind = ConsumerInstanceGVK()

	return &apiv1.ResourceInstance{ResourceMeta: meta, Spec: spec}, nil
}
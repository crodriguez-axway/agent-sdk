/*
 * This file is automatically generated
 */

package v1alpha1

import (
	"encoding/json"

	apiv1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/api/v1"
)

var (
	_APISpecGVK = apiv1.GroupVersionKind{
		GroupKind: apiv1.GroupKind{
			Group: "management",
			Kind:  "APISpec",
		},
		APIVersion: "v1alpha1",
	}
)

const (
	APISpecScope = "K8SCluster"

	APISpecResourceName = "apispecs"
)

func APISpecGVK() apiv1.GroupVersionKind {
	return _APISpecGVK
}

func init() {
	apiv1.RegisterGVK(_APISpecGVK, APISpecScope, APISpecResourceName)
}

// APISpec Resource
type APISpec struct {
	apiv1.ResourceMeta

	Owner interface{} `json:"owner"`

	Spec ApiSpecSpec `json:"spec"`
}

// FromInstance converts a ResourceInstance to a APISpec
func (res *APISpec) FromInstance(ri *apiv1.ResourceInstance) error {
	if ri == nil {
		res = nil
		return nil
	}

	var err error
	rawResource := ri.GetRawResource()
	if rawResource == nil {
		rawResource, err = json.Marshal(ri)
		if err != nil {
			return err
		}
	}

	err = json.Unmarshal(rawResource, res)
	return err
}

// APISpecFromInstanceArray converts a []*ResourceInstance to a []*APISpec
func APISpecFromInstanceArray(fromArray []*apiv1.ResourceInstance) ([]*APISpec, error) {
	newArray := make([]*APISpec, 0)
	for _, item := range fromArray {
		res := &APISpec{}
		err := res.FromInstance(item)
		if err != nil {
			return make([]*APISpec, 0), err
		}
		newArray = append(newArray, res)
	}

	return newArray, nil
}

// AsInstance converts a APISpec to a ResourceInstance
func (res *APISpec) AsInstance() (*apiv1.ResourceInstance, error) {
	meta := res.ResourceMeta
	meta.GroupVersionKind = APISpecGVK()
	res.ResourceMeta = meta

	m, err := json.Marshal(res)
	if err != nil {
		return nil, err
	}

	instance := apiv1.ResourceInstance{}
	err = json.Unmarshal(m, &instance)
	if err != nil {
		return nil, err
	}

	return &instance, nil
}

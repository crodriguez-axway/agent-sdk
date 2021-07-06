/*
 * This file is automatically generated
 */

package v1alpha1

import (
	"encoding/json"

	apiv1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/api/v1"
)

var (
	_ResourceDefinitionVersionGVK = apiv1.GroupVersionKind{
		GroupKind: apiv1.GroupKind{
			Group: "definitions",
			Kind:  "ResourceDefinitionVersion",
		},
		APIVersion: "v1alpha1",
	}
)

const (
	ResourceDefinitionVersionScope = "ResourceGroup"

	ResourceDefinitionVersionResourceName = "resourceversions"
)

func ResourceDefinitionVersionGVK() apiv1.GroupVersionKind {
	return _ResourceDefinitionVersionGVK
}

func init() {
	apiv1.RegisterGVK(_ResourceDefinitionVersionGVK, ResourceDefinitionVersionScope, ResourceDefinitionVersionResourceName)
}

// ResourceDefinitionVersion Resource
type ResourceDefinitionVersion struct {
	apiv1.ResourceMeta

	Owner interface{} `json:"owner"`

	Spec ResourceDefinitionVersionSpec `json:"spec"`
}

// FromInstance converts a ResourceInstance to a ResourceDefinitionVersion
func (res *ResourceDefinitionVersion) FromInstance(ri *apiv1.ResourceInstance) error {
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

// ResourceDefinitionVersionFromInstanceArray converts a []*ResourceInstance to a []*ResourceDefinitionVersion
func ResourceDefinitionVersionFromInstanceArray(fromArray []*apiv1.ResourceInstance) ([]*ResourceDefinitionVersion, error) {
	newArray := make([]*ResourceDefinitionVersion, 0)
	for _, item := range fromArray {
		res := &ResourceDefinitionVersion{}
		err := res.FromInstance(item)
		if err != nil {
			return make([]*ResourceDefinitionVersion, 0), err
		}
		newArray = append(newArray, res)
	}

	return newArray, nil
}

// AsInstance converts a ResourceDefinitionVersion to a ResourceInstance
func (res *ResourceDefinitionVersion) AsInstance() (*apiv1.ResourceInstance, error) {
	meta := res.ResourceMeta
	meta.GroupVersionKind = ResourceDefinitionVersionGVK()
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

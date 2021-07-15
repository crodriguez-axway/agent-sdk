/*
 * This file is automatically generated
 */

package v1alpha1

import (
	"encoding/json"

	apiv1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/api/v1"
)

var (
	_VirtualServiceGVK = apiv1.GroupVersionKind{
		GroupKind: apiv1.GroupKind{
			Group: "management",
			Kind:  "VirtualService",
		},
		APIVersion: "v1alpha1",
	}
)

const (
	VirtualServiceScope = "VirtualAPI"

	VirtualServiceResourceName = "virtualservices"
)

func VirtualServiceGVK() apiv1.GroupVersionKind {
	return _VirtualServiceGVK
}

func init() {
	apiv1.RegisterGVK(_VirtualServiceGVK, VirtualServiceScope, VirtualServiceResourceName)
}

// VirtualService Resource
type VirtualService struct {
	apiv1.ResourceMeta

	Owner interface{} `json:"owner"`

	Spec VirtualServiceSpec `json:"spec"`
}

// FromInstance converts a ResourceInstance to a VirtualService
func (res *VirtualService) FromInstance(ri *apiv1.ResourceInstance) error {
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

// VirtualServiceFromInstanceArray converts a []*ResourceInstance to a []*VirtualService
func VirtualServiceFromInstanceArray(fromArray []*apiv1.ResourceInstance) ([]*VirtualService, error) {
	newArray := make([]*VirtualService, 0)
	for _, item := range fromArray {
		res := &VirtualService{}
		err := res.FromInstance(item)
		if err != nil {
			return make([]*VirtualService, 0), err
		}
		newArray = append(newArray, res)
	}

	return newArray, nil
}

// AsInstance converts a VirtualService to a ResourceInstance
func (res *VirtualService) AsInstance() (*apiv1.ResourceInstance, error) {
	meta := res.ResourceMeta
	meta.GroupVersionKind = VirtualServiceGVK()
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
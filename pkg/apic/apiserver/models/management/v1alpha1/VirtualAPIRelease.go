/*
 * This file is automatically generated
 */

package v1alpha1

import (
	"encoding/json"

	apiv1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/api/v1"
)

var (
	_VirtualAPIReleaseGVK = apiv1.GroupVersionKind{
		GroupKind: apiv1.GroupKind{
			Group: "management",
			Kind:  "VirtualAPIRelease",
		},
		APIVersion: "v1alpha1",
	}
)

const (
	VirtualAPIReleaseScope = ""

	VirtualAPIReleaseResourceName = "virtualapireleases"
)

func VirtualAPIReleaseGVK() apiv1.GroupVersionKind {
	return _VirtualAPIReleaseGVK
}

func init() {
	apiv1.RegisterGVK(_VirtualAPIReleaseGVK, VirtualAPIReleaseScope, VirtualAPIReleaseResourceName)
}

// VirtualAPIRelease Resource
type VirtualAPIRelease struct {
	apiv1.ResourceMeta

	Icon interface{} `json:"icon"`

	Owner interface{} `json:"owner"`

	Spec VirtualApiReleaseSpec `json:"spec"`

	State interface{} `json:"state"`
}

// FromInstance converts a ResourceInstance to a VirtualAPIRelease
func (res *VirtualAPIRelease) FromInstance(ri *apiv1.ResourceInstance) error {
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

// VirtualAPIReleaseFromInstanceArray converts a []*ResourceInstance to a []*VirtualAPIRelease
func VirtualAPIReleaseFromInstanceArray(fromArray []*apiv1.ResourceInstance) ([]*VirtualAPIRelease, error) {
	newArray := make([]*VirtualAPIRelease, 0)
	for _, item := range fromArray {
		res := &VirtualAPIRelease{}
		err := res.FromInstance(item)
		if err != nil {
			return make([]*VirtualAPIRelease, 0), err
		}
		newArray = append(newArray, res)
	}

	return newArray, nil
}

// AsInstance converts a VirtualAPIRelease to a ResourceInstance
func (res *VirtualAPIRelease) AsInstance() (*apiv1.ResourceInstance, error) {
	meta := res.ResourceMeta
	meta.GroupVersionKind = VirtualAPIReleaseGVK()
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
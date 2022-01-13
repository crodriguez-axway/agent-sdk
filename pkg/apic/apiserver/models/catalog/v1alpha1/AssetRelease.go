/*
 * This file is automatically generated
 */

package v1alpha1

import (
	"encoding/json"

	apiv1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/api/v1"
)

var (
	_AssetReleaseGVK = apiv1.GroupVersionKind{
		GroupKind: apiv1.GroupKind{
			Group: "catalog",
			Kind:  "AssetRelease",
		},
		APIVersion: "v1alpha1",
	}

	AssetReleaseScopes = []string{""}
)

const AssetReleaseResourceName = "assetreleases"

func AssetReleaseGVK() apiv1.GroupVersionKind {
	return _AssetReleaseGVK
}

func init() {
	apiv1.RegisterGVK(_AssetReleaseGVK, AssetReleaseScopes[0], AssetReleaseResourceName)
}

// AssetRelease Resource
type AssetRelease struct {
	apiv1.ResourceMeta

	Icon interface{} `json:"icon"`

	Owner *apiv1.Owner `json:"owner"`

	References AssetReleaseReferences `json:"references"`

	Spec AssetReleaseSpec `json:"spec"`

	Status AssetReleaseStatus `json:"status"`
}

// FromInstance converts a ResourceInstance to a AssetRelease
func (res *AssetRelease) FromInstance(ri *apiv1.ResourceInstance) error {
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

// AssetReleaseFromInstanceArray converts a []*ResourceInstance to a []*AssetRelease
func AssetReleaseFromInstanceArray(fromArray []*apiv1.ResourceInstance) ([]*AssetRelease, error) {
	newArray := make([]*AssetRelease, 0)
	for _, item := range fromArray {
		res := &AssetRelease{}
		err := res.FromInstance(item)
		if err != nil {
			return make([]*AssetRelease, 0), err
		}
		newArray = append(newArray, res)
	}

	return newArray, nil
}

// AsInstance converts a AssetRelease to a ResourceInstance
func (res *AssetRelease) AsInstance() (*apiv1.ResourceInstance, error) {
	meta := res.ResourceMeta
	meta.GroupVersionKind = AssetReleaseGVK()
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

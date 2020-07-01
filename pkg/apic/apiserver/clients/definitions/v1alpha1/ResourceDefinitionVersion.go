/*
 * This file is automatically generated
 */

package v1alpha1

import (
	v1 "git.ecd.axway.int/apigov/apic_agents_sdk/pkg/apic/apiserver/clients/api/v1"
	"git.ecd.axway.int/apigov/apic_agents_sdk/pkg/apic/apiserver/models/definitions/v1alpha1"
)

// ResourceDefinitionVersionClient -
type ResourceDefinitionVersionClient struct {
	client *v1.Client
}

// NewResourceDefinitionVersionClient -
func NewResourceDefinitionVersionClient(cb *v1.ClientBase) (*ResourceDefinitionVersionClient, error) {
	client, err := cb.ForKind(v1alpha1.ResourceDefinitionVersionGVK())
	if err != nil {
		return nil, err
	}

	return &ResourceDefinitionVersionClient{client}, nil
}

// WithScope -
func (c *ResourceDefinitionVersionClient) WithScope(scope string) *ResourceDefinitionVersionClient {
	return &ResourceDefinitionVersionClient{
		c.client.WithScope(scope),
	}
}

// List -
func (c *ResourceDefinitionVersionClient) List(options ...v1.ListOptions) ([]*v1alpha1.ResourceDefinitionVersion, error) {
	riList, err := c.client.List(options...)
	if err != nil {
		return nil, err
	}

	result := make([]*v1alpha1.ResourceDefinitionVersion, len(riList))

	for i := range riList {
		result[i] = &v1alpha1.ResourceDefinitionVersion{}
		err := result[i].FromInstance(riList[i])
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

// Get -
func (c *ResourceDefinitionVersionClient) Get(name string) (*v1alpha1.ResourceDefinitionVersion, error) {
	ri, err := c.client.Get(name)
	if err != nil {
		return nil, err
	}

	service := &v1alpha1.ResourceDefinitionVersion{}
	service.FromInstance(ri)

	return service, nil
}

// Delete -
func (c *ResourceDefinitionVersionClient) Delete(res *v1alpha1.ResourceDefinitionVersion) error {
	ri, err := res.AsInstance()

	if err != nil {
		return err
	}

	return c.client.Delete(ri)
}

// Create -
func (c *ResourceDefinitionVersionClient) Create(res *v1alpha1.ResourceDefinitionVersion) (*v1alpha1.ResourceDefinitionVersion, error) {
	ri, err := res.AsInstance()

	if err != nil {
		return nil, err
	}

	cri, err := c.client.Create(ri)
	if err != nil {
		return nil, err
	}

	created := &v1alpha1.ResourceDefinitionVersion{}

	err = created.FromInstance(cri)
	if err != nil {
		return nil, err
	}

	return created, err
}

// Update -
func (c *ResourceDefinitionVersionClient) Update(res *v1alpha1.ResourceDefinitionVersion) (*v1alpha1.ResourceDefinitionVersion, error) {
	ri, err := res.AsInstance()
	if err != nil {
		return nil, err
	}
	resource, err := c.client.Update(ri)
	updated := &v1alpha1.ResourceDefinitionVersion{}

	// Updates the resource in place
	err = updated.FromInstance(resource)
	if err != nil {
		return nil, err
	}

	return updated, nil
}
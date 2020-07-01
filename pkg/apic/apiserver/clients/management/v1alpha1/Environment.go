/*
 * This file is automatically generated
 */

package v1alpha1

import (
	v1 "git.ecd.axway.int/apigov/apic_agents_sdk/pkg/apic/apiserver/clients/api/v1"
	"git.ecd.axway.int/apigov/apic_agents_sdk/pkg/apic/apiserver/models/management/v1alpha1"
)

// EnvironmentClient -
type EnvironmentClient struct {
	client *v1.Client
}

// NewEnvironmentClient -
func NewEnvironmentClient(cb *v1.ClientBase) (*EnvironmentClient, error) {
	client, err := cb.ForKind(v1alpha1.EnvironmentGVK())
	if err != nil {
		return nil, err
	}

	return &EnvironmentClient{client}, nil
}

// WithScope -
func (c *EnvironmentClient) WithScope(scope string) *EnvironmentClient {
	return &EnvironmentClient{
		c.client.WithScope(scope),
	}
}

// List -
func (c *EnvironmentClient) List(options ...v1.ListOptions) ([]*v1alpha1.Environment, error) {
	riList, err := c.client.List(options...)
	if err != nil {
		return nil, err
	}

	result := make([]*v1alpha1.Environment, len(riList))

	for i := range riList {
		result[i] = &v1alpha1.Environment{}
		err := result[i].FromInstance(riList[i])
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

// Get -
func (c *EnvironmentClient) Get(name string) (*v1alpha1.Environment, error) {
	ri, err := c.client.Get(name)
	if err != nil {
		return nil, err
	}

	service := &v1alpha1.Environment{}
	service.FromInstance(ri)

	return service, nil
}

// Delete -
func (c *EnvironmentClient) Delete(res *v1alpha1.Environment) error {
	ri, err := res.AsInstance()

	if err != nil {
		return err
	}

	return c.client.Delete(ri)
}

// Create -
func (c *EnvironmentClient) Create(res *v1alpha1.Environment) (*v1alpha1.Environment, error) {
	ri, err := res.AsInstance()

	if err != nil {
		return nil, err
	}

	cri, err := c.client.Create(ri)
	if err != nil {
		return nil, err
	}

	created := &v1alpha1.Environment{}

	err = created.FromInstance(cri)
	if err != nil {
		return nil, err
	}

	return created, err
}

// Update -
func (c *EnvironmentClient) Update(res *v1alpha1.Environment) (*v1alpha1.Environment, error) {
	ri, err := res.AsInstance()
	if err != nil {
		return nil, err
	}
	resource, err := c.client.Update(ri)
	updated := &v1alpha1.Environment{}

	// Updates the resource in place
	err = updated.FromInstance(resource)
	if err != nil {
		return nil, err
	}

	return updated, nil
}
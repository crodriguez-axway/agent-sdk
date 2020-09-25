/*
 * This file is automatically generated
 */

package v1alpha1

import (
	v1 "git.ecd.axway.org/apigov/apic_agents_sdk/pkg/apic/apiserver/clients/api/v1"
	"git.ecd.axway.org/apigov/apic_agents_sdk/pkg/apic/apiserver/models/management/v1alpha1"
)

// EdgeTraceabilityAgentClient -
type EdgeTraceabilityAgentClient struct {
	client *v1.Client
}

// NewEdgeTraceabilityAgentClient -
func NewEdgeTraceabilityAgentClient(cb *v1.ClientBase) (*EdgeTraceabilityAgentClient, error) {
	client, err := cb.ForKind(v1alpha1.EdgeTraceabilityAgentGVK())
	if err != nil {
		return nil, err
	}

	return &EdgeTraceabilityAgentClient{client}, nil
}

// WithScope -
func (c *EdgeTraceabilityAgentClient) WithScope(scope string) *EdgeTraceabilityAgentClient {
	return &EdgeTraceabilityAgentClient{
		c.client.WithScope(scope),
	}
}

// List -
func (c *EdgeTraceabilityAgentClient) List(options ...v1.ListOptions) ([]*v1alpha1.EdgeTraceabilityAgent, error) {
	riList, err := c.client.List(options...)
	if err != nil {
		return nil, err
	}

	result := make([]*v1alpha1.EdgeTraceabilityAgent, len(riList))

	for i := range riList {
		result[i] = &v1alpha1.EdgeTraceabilityAgent{}
		err := result[i].FromInstance(riList[i])
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

// Get -
func (c *EdgeTraceabilityAgentClient) Get(name string) (*v1alpha1.EdgeTraceabilityAgent, error) {
	ri, err := c.client.Get(name)
	if err != nil {
		return nil, err
	}

	service := &v1alpha1.EdgeTraceabilityAgent{}
	service.FromInstance(ri)

	return service, nil
}

// Delete -
func (c *EdgeTraceabilityAgentClient) Delete(res *v1alpha1.EdgeTraceabilityAgent) error {
	ri, err := res.AsInstance()

	if err != nil {
		return err
	}

	return c.client.Delete(ri)
}

// Create -
func (c *EdgeTraceabilityAgentClient) Create(res *v1alpha1.EdgeTraceabilityAgent) (*v1alpha1.EdgeTraceabilityAgent, error) {
	ri, err := res.AsInstance()

	if err != nil {
		return nil, err
	}

	cri, err := c.client.Create(ri)
	if err != nil {
		return nil, err
	}

	created := &v1alpha1.EdgeTraceabilityAgent{}

	err = created.FromInstance(cri)
	if err != nil {
		return nil, err
	}

	return created, err
}

// Update -
func (c *EdgeTraceabilityAgentClient) Update(res *v1alpha1.EdgeTraceabilityAgent) (*v1alpha1.EdgeTraceabilityAgent, error) {
	ri, err := res.AsInstance()
	if err != nil {
		return nil, err
	}
	resource, err := c.client.Update(ri)
	updated := &v1alpha1.EdgeTraceabilityAgent{}

	// Updates the resource in place
	err = updated.FromInstance(resource)
	if err != nil {
		return nil, err
	}

	return updated, nil
}
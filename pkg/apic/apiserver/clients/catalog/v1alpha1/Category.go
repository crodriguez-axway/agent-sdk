/*
 * This file is automatically generated
 */

package v1alpha1

import (
	"fmt"

	v1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/clients/api/v1"
	apiv1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/api/v1"
	"github.com/Axway/agent-sdk/pkg/apic/apiserver/models/catalog/v1alpha1"
)

type CategoryMergeFunc func(*v1alpha1.Category, *v1alpha1.Category) (*v1alpha1.Category, error)

// Merge builds a merge option for an update operation
func CategoryMerge(f CategoryMergeFunc) v1.UpdateOption {
	return v1.Merge(func(prev, new apiv1.Interface) (apiv1.Interface, error) {
		p, n := &v1alpha1.Category{}, &v1alpha1.Category{}

		switch t := prev.(type) {
		case *v1alpha1.Category:
			p = t
		case *apiv1.ResourceInstance:
			err := p.FromInstance(t)
			if err != nil {
				return nil, fmt.Errorf("merge: failed to unserialise prev resource: %w", err)
			}
		default:
			return nil, fmt.Errorf("merge: failed to unserialise prev resource, unxexpected resource type: %T", t)
		}

		switch t := new.(type) {
		case *v1alpha1.Category:
			n = t
		case *apiv1.ResourceInstance:
			err := n.FromInstance(t)
			if err != nil {
				return nil, fmt.Errorf("merge: failed to unserialize new resource: %w", err)
			}
		default:
			return nil, fmt.Errorf("merge: failed to unserialise new resource, unxexpected resource type: %T", t)
		}

		return f(p, n)
	})
}

// CategoryClient -
type CategoryClient struct {
	client v1.Scoped
}

// NewCategoryClient -
func NewCategoryClient(c v1.Base) (*CategoryClient, error) {

	client, err := c.ForKind(v1alpha1.CategoryGVK())
	if err != nil {
		return nil, err
	}

	return &CategoryClient{client}, nil

}

// List -
func (c *CategoryClient) List(options ...v1.ListOptions) ([]*v1alpha1.Category, error) {
	riList, err := c.client.List(options...)
	if err != nil {
		return nil, err
	}

	result := make([]*v1alpha1.Category, len(riList))

	for i := range riList {
		result[i] = &v1alpha1.Category{}
		err := result[i].FromInstance(riList[i])
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

// Get -
func (c *CategoryClient) Get(name string) (*v1alpha1.Category, error) {
	ri, err := c.client.Get(name)
	if err != nil {
		return nil, err
	}

	service := &v1alpha1.Category{}
	service.FromInstance(ri)

	return service, nil
}

// Delete -
func (c *CategoryClient) Delete(res *v1alpha1.Category) error {
	ri, err := res.AsInstance()

	if err != nil {
		return err
	}

	return c.client.Delete(ri)
}

// Create -
func (c *CategoryClient) Create(res *v1alpha1.Category, opts ...v1.CreateOption) (*v1alpha1.Category, error) {
	ri, err := res.AsInstance()

	if err != nil {
		return nil, err
	}

	cri, err := c.client.Create(ri, opts...)
	if err != nil {
		return nil, err
	}

	created := &v1alpha1.Category{}

	err = created.FromInstance(cri)
	if err != nil {
		return nil, err
	}

	return created, err
}

// Update -
func (c *CategoryClient) Update(res *v1alpha1.Category, opts ...v1.UpdateOption) (*v1alpha1.Category, error) {
	ri, err := res.AsInstance()
	if err != nil {
		return nil, err
	}
	resource, err := c.client.Update(ri, opts...)
	if err != nil {
		return nil, err
	}

	updated := &v1alpha1.Category{}

	// Updates the resource in place
	err = updated.FromInstance(resource)
	if err != nil {
		return nil, err
	}

	return updated, nil
}
/*
Copyright 2020 Topi Kettunen.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	polarsquadv1 "github.com/topikettunen/polar-controller/pkg/apis/polarsquad/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakePolarDeployments implements PolarDeploymentInterface
type FakePolarDeployments struct {
	Fake *FakePolarsquadV1
	ns   string
}

var polardeploymentsResource = schema.GroupVersionResource{Group: "polarsquad.com", Version: "v1", Resource: "polardeployments"}

var polardeploymentsKind = schema.GroupVersionKind{Group: "polarsquad.com", Version: "v1", Kind: "PolarDeployment"}

// Get takes name of the polarDeployment, and returns the corresponding polarDeployment object, and an error if there is any.
func (c *FakePolarDeployments) Get(ctx context.Context, name string, options v1.GetOptions) (result *polarsquadv1.PolarDeployment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(polardeploymentsResource, c.ns, name), &polarsquadv1.PolarDeployment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*polarsquadv1.PolarDeployment), err
}

// List takes label and field selectors, and returns the list of PolarDeployments that match those selectors.
func (c *FakePolarDeployments) List(ctx context.Context, opts v1.ListOptions) (result *polarsquadv1.PolarDeploymentList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(polardeploymentsResource, polardeploymentsKind, c.ns, opts), &polarsquadv1.PolarDeploymentList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &polarsquadv1.PolarDeploymentList{ListMeta: obj.(*polarsquadv1.PolarDeploymentList).ListMeta}
	for _, item := range obj.(*polarsquadv1.PolarDeploymentList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested polarDeployments.
func (c *FakePolarDeployments) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(polardeploymentsResource, c.ns, opts))

}

// Create takes the representation of a polarDeployment and creates it.  Returns the server's representation of the polarDeployment, and an error, if there is any.
func (c *FakePolarDeployments) Create(ctx context.Context, polarDeployment *polarsquadv1.PolarDeployment, opts v1.CreateOptions) (result *polarsquadv1.PolarDeployment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(polardeploymentsResource, c.ns, polarDeployment), &polarsquadv1.PolarDeployment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*polarsquadv1.PolarDeployment), err
}

// Update takes the representation of a polarDeployment and updates it. Returns the server's representation of the polarDeployment, and an error, if there is any.
func (c *FakePolarDeployments) Update(ctx context.Context, polarDeployment *polarsquadv1.PolarDeployment, opts v1.UpdateOptions) (result *polarsquadv1.PolarDeployment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(polardeploymentsResource, c.ns, polarDeployment), &polarsquadv1.PolarDeployment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*polarsquadv1.PolarDeployment), err
}

// Delete takes name of the polarDeployment and deletes it. Returns an error if one occurs.
func (c *FakePolarDeployments) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(polardeploymentsResource, c.ns, name), &polarsquadv1.PolarDeployment{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePolarDeployments) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(polardeploymentsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &polarsquadv1.PolarDeploymentList{})
	return err
}

// Patch applies the patch and returns the patched polarDeployment.
func (c *FakePolarDeployments) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *polarsquadv1.PolarDeployment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(polardeploymentsResource, c.ns, name, pt, data, subresources...), &polarsquadv1.PolarDeployment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*polarsquadv1.PolarDeployment), err
}
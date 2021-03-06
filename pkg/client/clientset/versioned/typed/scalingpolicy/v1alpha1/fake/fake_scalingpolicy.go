/*
Copyright 2018 The Kubernetes Authors.

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

package fake

import (
	v1alpha1 "github.com/justinsb/scaler/pkg/apis/scalingpolicy/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeScalingPolicies implements ScalingPolicyInterface
type FakeScalingPolicies struct {
	Fake *FakeScalingpolicyV1alpha1
	ns   string
}

var scalingpoliciesResource = schema.GroupVersionResource{Group: "scalingpolicy.kope.io", Version: "v1alpha1", Resource: "scalingpolicies"}

var scalingpoliciesKind = schema.GroupVersionKind{Group: "scalingpolicy.kope.io", Version: "v1alpha1", Kind: "ScalingPolicy"}

// Get takes name of the scalingPolicy, and returns the corresponding scalingPolicy object, and an error if there is any.
func (c *FakeScalingPolicies) Get(name string, options v1.GetOptions) (result *v1alpha1.ScalingPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(scalingpoliciesResource, c.ns, name), &v1alpha1.ScalingPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ScalingPolicy), err
}

// List takes label and field selectors, and returns the list of ScalingPolicies that match those selectors.
func (c *FakeScalingPolicies) List(opts v1.ListOptions) (result *v1alpha1.ScalingPolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(scalingpoliciesResource, scalingpoliciesKind, c.ns, opts), &v1alpha1.ScalingPolicyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ScalingPolicyList{}
	for _, item := range obj.(*v1alpha1.ScalingPolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested scalingPolicies.
func (c *FakeScalingPolicies) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(scalingpoliciesResource, c.ns, opts))

}

// Create takes the representation of a scalingPolicy and creates it.  Returns the server's representation of the scalingPolicy, and an error, if there is any.
func (c *FakeScalingPolicies) Create(scalingPolicy *v1alpha1.ScalingPolicy) (result *v1alpha1.ScalingPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(scalingpoliciesResource, c.ns, scalingPolicy), &v1alpha1.ScalingPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ScalingPolicy), err
}

// Update takes the representation of a scalingPolicy and updates it. Returns the server's representation of the scalingPolicy, and an error, if there is any.
func (c *FakeScalingPolicies) Update(scalingPolicy *v1alpha1.ScalingPolicy) (result *v1alpha1.ScalingPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(scalingpoliciesResource, c.ns, scalingPolicy), &v1alpha1.ScalingPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ScalingPolicy), err
}

// Delete takes name of the scalingPolicy and deletes it. Returns an error if one occurs.
func (c *FakeScalingPolicies) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(scalingpoliciesResource, c.ns, name), &v1alpha1.ScalingPolicy{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeScalingPolicies) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(scalingpoliciesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.ScalingPolicyList{})
	return err
}

// Patch applies the patch and returns the patched scalingPolicy.
func (c *FakeScalingPolicies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ScalingPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(scalingpoliciesResource, c.ns, name, data, subresources...), &v1alpha1.ScalingPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ScalingPolicy), err
}

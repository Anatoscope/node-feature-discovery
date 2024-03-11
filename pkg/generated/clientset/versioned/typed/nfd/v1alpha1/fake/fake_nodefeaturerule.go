/*
Copyright 2024 The Kubernetes Authors.

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

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha1 "sigs.k8s.io/node-feature-discovery/pkg/apis/nfd/v1alpha1"
)

// FakeNodeFeatureRules implements NodeFeatureRuleInterface
type FakeNodeFeatureRules struct {
	Fake *FakeNfdV1alpha1
}

var nodefeaturerulesResource = v1alpha1.SchemeGroupVersion.WithResource("nodefeaturerules")

var nodefeaturerulesKind = v1alpha1.SchemeGroupVersion.WithKind("NodeFeatureRule")

// Get takes name of the nodeFeatureRule, and returns the corresponding nodeFeatureRule object, and an error if there is any.
func (c *FakeNodeFeatureRules) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.NodeFeatureRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(nodefeaturerulesResource, name), &v1alpha1.NodeFeatureRule{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NodeFeatureRule), err
}

// List takes label and field selectors, and returns the list of NodeFeatureRules that match those selectors.
func (c *FakeNodeFeatureRules) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.NodeFeatureRuleList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(nodefeaturerulesResource, nodefeaturerulesKind, opts), &v1alpha1.NodeFeatureRuleList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.NodeFeatureRuleList{ListMeta: obj.(*v1alpha1.NodeFeatureRuleList).ListMeta}
	for _, item := range obj.(*v1alpha1.NodeFeatureRuleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested nodeFeatureRules.
func (c *FakeNodeFeatureRules) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(nodefeaturerulesResource, opts))
}

// Create takes the representation of a nodeFeatureRule and creates it.  Returns the server's representation of the nodeFeatureRule, and an error, if there is any.
func (c *FakeNodeFeatureRules) Create(ctx context.Context, nodeFeatureRule *v1alpha1.NodeFeatureRule, opts v1.CreateOptions) (result *v1alpha1.NodeFeatureRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(nodefeaturerulesResource, nodeFeatureRule), &v1alpha1.NodeFeatureRule{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NodeFeatureRule), err
}

// Update takes the representation of a nodeFeatureRule and updates it. Returns the server's representation of the nodeFeatureRule, and an error, if there is any.
func (c *FakeNodeFeatureRules) Update(ctx context.Context, nodeFeatureRule *v1alpha1.NodeFeatureRule, opts v1.UpdateOptions) (result *v1alpha1.NodeFeatureRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(nodefeaturerulesResource, nodeFeatureRule), &v1alpha1.NodeFeatureRule{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NodeFeatureRule), err
}

// Delete takes name of the nodeFeatureRule and deletes it. Returns an error if one occurs.
func (c *FakeNodeFeatureRules) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(nodefeaturerulesResource, name, opts), &v1alpha1.NodeFeatureRule{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeNodeFeatureRules) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(nodefeaturerulesResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.NodeFeatureRuleList{})
	return err
}

// Patch applies the patch and returns the patched nodeFeatureRule.
func (c *FakeNodeFeatureRules) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.NodeFeatureRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(nodefeaturerulesResource, name, pt, data, subresources...), &v1alpha1.NodeFeatureRule{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NodeFeatureRule), err
}

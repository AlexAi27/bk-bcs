/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v2 "bk-bcs/bcs-mesos/pkg/apis/bkbcs/v2"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeBcsClusterAgentSettings implements BcsClusterAgentSettingInterface
type FakeBcsClusterAgentSettings struct {
	Fake *FakeBkbcsV2
	ns   string
}

var bcsclusteragentsettingsResource = schema.GroupVersionResource{Group: "bkbcs.bkbcs.tencent.com", Version: "v2", Resource: "bcsclusteragentsettings"}

var bcsclusteragentsettingsKind = schema.GroupVersionKind{Group: "bkbcs.bkbcs.tencent.com", Version: "v2", Kind: "BcsClusterAgentSetting"}

// Get takes name of the bcsClusterAgentSetting, and returns the corresponding bcsClusterAgentSetting object, and an error if there is any.
func (c *FakeBcsClusterAgentSettings) Get(name string, options v1.GetOptions) (result *v2.BcsClusterAgentSetting, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(bcsclusteragentsettingsResource, c.ns, name), &v2.BcsClusterAgentSetting{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2.BcsClusterAgentSetting), err
}

// List takes label and field selectors, and returns the list of BcsClusterAgentSettings that match those selectors.
func (c *FakeBcsClusterAgentSettings) List(opts v1.ListOptions) (result *v2.BcsClusterAgentSettingList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(bcsclusteragentsettingsResource, bcsclusteragentsettingsKind, c.ns, opts), &v2.BcsClusterAgentSettingList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v2.BcsClusterAgentSettingList{ListMeta: obj.(*v2.BcsClusterAgentSettingList).ListMeta}
	for _, item := range obj.(*v2.BcsClusterAgentSettingList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested bcsClusterAgentSettings.
func (c *FakeBcsClusterAgentSettings) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(bcsclusteragentsettingsResource, c.ns, opts))

}

// Create takes the representation of a bcsClusterAgentSetting and creates it.  Returns the server's representation of the bcsClusterAgentSetting, and an error, if there is any.
func (c *FakeBcsClusterAgentSettings) Create(bcsClusterAgentSetting *v2.BcsClusterAgentSetting) (result *v2.BcsClusterAgentSetting, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(bcsclusteragentsettingsResource, c.ns, bcsClusterAgentSetting), &v2.BcsClusterAgentSetting{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2.BcsClusterAgentSetting), err
}

// Update takes the representation of a bcsClusterAgentSetting and updates it. Returns the server's representation of the bcsClusterAgentSetting, and an error, if there is any.
func (c *FakeBcsClusterAgentSettings) Update(bcsClusterAgentSetting *v2.BcsClusterAgentSetting) (result *v2.BcsClusterAgentSetting, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(bcsclusteragentsettingsResource, c.ns, bcsClusterAgentSetting), &v2.BcsClusterAgentSetting{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2.BcsClusterAgentSetting), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeBcsClusterAgentSettings) UpdateStatus(bcsClusterAgentSetting *v2.BcsClusterAgentSetting) (*v2.BcsClusterAgentSetting, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(bcsclusteragentsettingsResource, "status", c.ns, bcsClusterAgentSetting), &v2.BcsClusterAgentSetting{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2.BcsClusterAgentSetting), err
}

// Delete takes name of the bcsClusterAgentSetting and deletes it. Returns an error if one occurs.
func (c *FakeBcsClusterAgentSettings) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(bcsclusteragentsettingsResource, c.ns, name), &v2.BcsClusterAgentSetting{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeBcsClusterAgentSettings) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(bcsclusteragentsettingsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v2.BcsClusterAgentSettingList{})
	return err
}

// Patch applies the patch and returns the patched bcsClusterAgentSetting.
func (c *FakeBcsClusterAgentSettings) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v2.BcsClusterAgentSetting, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(bcsclusteragentsettingsResource, c.ns, name, data, subresources...), &v2.BcsClusterAgentSetting{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2.BcsClusterAgentSetting), err
}

/*
Copyright The Kubernetes Authors.

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
	controlv1 "github.com/wangxiaohua10/k8s-crd-lvmnodestorage/pkg/apis/control/v1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeNodeLVMStorages implements NodeLVMStorageInterface
type FakeNodeLVMStorages struct {
	Fake *FakeLvmV1
}

var nodelvmstoragesResource = schema.GroupVersionResource{Group: "lvm.node.storage", Version: "v1", Resource: "nodelvmstorages"}

var nodelvmstoragesKind = schema.GroupVersionKind{Group: "lvm.node.storage", Version: "v1", Kind: "NodeLVMStorage"}

// Get takes name of the nodeLVMStorage, and returns the corresponding nodeLVMStorage object, and an error if there is any.
func (c *FakeNodeLVMStorages) Get(name string, options v1.GetOptions) (result *controlv1.NodeLVMStorage, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(nodelvmstoragesResource, name), &controlv1.NodeLVMStorage{})
	if obj == nil {
		return nil, err
	}
	return obj.(*controlv1.NodeLVMStorage), err
}

// List takes label and field selectors, and returns the list of NodeLVMStorages that match those selectors.
func (c *FakeNodeLVMStorages) List(opts v1.ListOptions) (result *controlv1.NodeLVMStorageList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(nodelvmstoragesResource, nodelvmstoragesKind, opts), &controlv1.NodeLVMStorageList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &controlv1.NodeLVMStorageList{ListMeta: obj.(*controlv1.NodeLVMStorageList).ListMeta}
	for _, item := range obj.(*controlv1.NodeLVMStorageList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested nodeLVMStorages.
func (c *FakeNodeLVMStorages) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(nodelvmstoragesResource, opts))
}

// Create takes the representation of a nodeLVMStorage and creates it.  Returns the server's representation of the nodeLVMStorage, and an error, if there is any.
func (c *FakeNodeLVMStorages) Create(nodeLVMStorage *controlv1.NodeLVMStorage) (result *controlv1.NodeLVMStorage, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(nodelvmstoragesResource, nodeLVMStorage), &controlv1.NodeLVMStorage{})
	if obj == nil {
		return nil, err
	}
	return obj.(*controlv1.NodeLVMStorage), err
}

// Update takes the representation of a nodeLVMStorage and updates it. Returns the server's representation of the nodeLVMStorage, and an error, if there is any.
func (c *FakeNodeLVMStorages) Update(nodeLVMStorage *controlv1.NodeLVMStorage) (result *controlv1.NodeLVMStorage, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(nodelvmstoragesResource, nodeLVMStorage), &controlv1.NodeLVMStorage{})
	if obj == nil {
		return nil, err
	}
	return obj.(*controlv1.NodeLVMStorage), err
}

// Delete takes name of the nodeLVMStorage and deletes it. Returns an error if one occurs.
func (c *FakeNodeLVMStorages) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(nodelvmstoragesResource, name), &controlv1.NodeLVMStorage{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeNodeLVMStorages) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(nodelvmstoragesResource, listOptions)

	_, err := c.Fake.Invokes(action, &controlv1.NodeLVMStorageList{})
	return err
}

// Patch applies the patch and returns the patched nodeLVMStorage.
func (c *FakeNodeLVMStorages) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *controlv1.NodeLVMStorage, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(nodelvmstoragesResource, name, pt, data, subresources...), &controlv1.NodeLVMStorage{})
	if obj == nil {
		return nil, err
	}
	return obj.(*controlv1.NodeLVMStorage), err
}

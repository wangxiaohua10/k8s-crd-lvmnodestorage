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
	v1 "k8s-crd-lvmnodestorage/pkg/client/clientset/versioned/typed/control/v1"

	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeLvmV1 struct {
	*testing.Fake
}

func (c *FakeLvmV1) NodeLVMStorages() v1.NodeLVMStorageInterface {
	return &FakeNodeLVMStorages{c}
}

func (c *FakeLvmV1) NodeLvmPhysicalVolumes() v1.NodeLvmPhysicalVolumeInterface {
	return &FakeNodeLvmPhysicalVolumes{c}
}

func (c *FakeLvmV1) NodeLvmVolumeGroups() v1.NodeLvmVolumeGroupInterface {
	return &FakeNodeLvmVolumeGroups{c}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeLvmV1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}

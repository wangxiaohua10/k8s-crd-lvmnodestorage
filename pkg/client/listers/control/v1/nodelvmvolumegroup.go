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

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "k8s-crd-lvmnodestorage/pkg/apis/control/v1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// NodeLvmVolumeGroupLister helps list NodeLvmVolumeGroups.
type NodeLvmVolumeGroupLister interface {
	// List lists all NodeLvmVolumeGroups in the indexer.
	List(selector labels.Selector) (ret []*v1.NodeLvmVolumeGroup, err error)
	// Get retrieves the NodeLvmVolumeGroup from the index for a given name.
	Get(name string) (*v1.NodeLvmVolumeGroup, error)
	NodeLvmVolumeGroupListerExpansion
}

// nodeLvmVolumeGroupLister implements the NodeLvmVolumeGroupLister interface.
type nodeLvmVolumeGroupLister struct {
	indexer cache.Indexer
}

// NewNodeLvmVolumeGroupLister returns a new NodeLvmVolumeGroupLister.
func NewNodeLvmVolumeGroupLister(indexer cache.Indexer) NodeLvmVolumeGroupLister {
	return &nodeLvmVolumeGroupLister{indexer: indexer}
}

// List lists all NodeLvmVolumeGroups in the indexer.
func (s *nodeLvmVolumeGroupLister) List(selector labels.Selector) (ret []*v1.NodeLvmVolumeGroup, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.NodeLvmVolumeGroup))
	})
	return ret, err
}

// Get retrieves the NodeLvmVolumeGroup from the index for a given name.
func (s *nodeLvmVolumeGroupLister) Get(name string) (*v1.NodeLvmVolumeGroup, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("nodelvmvolumegroup"), name)
	}
	return obj.(*v1.NodeLvmVolumeGroup), nil
}

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

package v1alpha1

import (
	v1alpha1 "github.com/kfsoftware/hlf-operator/api/hlf.kungfusoftware.es/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// FabricExplorerLister helps list FabricExplorers.
// All objects returned here must be treated as read-only.
type FabricExplorerLister interface {
	// List lists all FabricExplorers in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.FabricExplorer, err error)
	// FabricExplorers returns an object that can list and get FabricExplorers.
	FabricExplorers(namespace string) FabricExplorerNamespaceLister
	FabricExplorerListerExpansion
}

// fabricExplorerLister implements the FabricExplorerLister interface.
type fabricExplorerLister struct {
	indexer cache.Indexer
}

// NewFabricExplorerLister returns a new FabricExplorerLister.
func NewFabricExplorerLister(indexer cache.Indexer) FabricExplorerLister {
	return &fabricExplorerLister{indexer: indexer}
}

// List lists all FabricExplorers in the indexer.
func (s *fabricExplorerLister) List(selector labels.Selector) (ret []*v1alpha1.FabricExplorer, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.FabricExplorer))
	})
	return ret, err
}

// FabricExplorers returns an object that can list and get FabricExplorers.
func (s *fabricExplorerLister) FabricExplorers(namespace string) FabricExplorerNamespaceLister {
	return fabricExplorerNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// FabricExplorerNamespaceLister helps list and get FabricExplorers.
// All objects returned here must be treated as read-only.
type FabricExplorerNamespaceLister interface {
	// List lists all FabricExplorers in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.FabricExplorer, err error)
	// Get retrieves the FabricExplorer from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.FabricExplorer, error)
	FabricExplorerNamespaceListerExpansion
}

// fabricExplorerNamespaceLister implements the FabricExplorerNamespaceLister
// interface.
type fabricExplorerNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all FabricExplorers in the indexer for a given namespace.
func (s fabricExplorerNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.FabricExplorer, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.FabricExplorer))
	})
	return ret, err
}

// Get retrieves the FabricExplorer from the indexer for a given namespace and name.
func (s fabricExplorerNamespaceLister) Get(name string) (*v1alpha1.FabricExplorer, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("fabricexplorer"), name)
	}
	return obj.(*v1alpha1.FabricExplorer), nil
}

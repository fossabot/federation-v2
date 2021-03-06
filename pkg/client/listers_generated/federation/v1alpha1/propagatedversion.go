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

// This file was automatically generated by lister-gen

package v1alpha1

import (
	v1alpha1 "github.com/marun/federation-v2/pkg/apis/federation/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// PropagatedVersionLister helps list PropagatedVersions.
type PropagatedVersionLister interface {
	// List lists all PropagatedVersions in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.PropagatedVersion, err error)
	// PropagatedVersions returns an object that can list and get PropagatedVersions.
	PropagatedVersions(namespace string) PropagatedVersionNamespaceLister
	PropagatedVersionListerExpansion
}

// propagatedVersionLister implements the PropagatedVersionLister interface.
type propagatedVersionLister struct {
	indexer cache.Indexer
}

// NewPropagatedVersionLister returns a new PropagatedVersionLister.
func NewPropagatedVersionLister(indexer cache.Indexer) PropagatedVersionLister {
	return &propagatedVersionLister{indexer: indexer}
}

// List lists all PropagatedVersions in the indexer.
func (s *propagatedVersionLister) List(selector labels.Selector) (ret []*v1alpha1.PropagatedVersion, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.PropagatedVersion))
	})
	return ret, err
}

// PropagatedVersions returns an object that can list and get PropagatedVersions.
func (s *propagatedVersionLister) PropagatedVersions(namespace string) PropagatedVersionNamespaceLister {
	return propagatedVersionNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// PropagatedVersionNamespaceLister helps list and get PropagatedVersions.
type PropagatedVersionNamespaceLister interface {
	// List lists all PropagatedVersions in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.PropagatedVersion, err error)
	// Get retrieves the PropagatedVersion from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.PropagatedVersion, error)
	PropagatedVersionNamespaceListerExpansion
}

// propagatedVersionNamespaceLister implements the PropagatedVersionNamespaceLister
// interface.
type propagatedVersionNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all PropagatedVersions in the indexer for a given namespace.
func (s propagatedVersionNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.PropagatedVersion, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.PropagatedVersion))
	})
	return ret, err
}

// Get retrieves the PropagatedVersion from the indexer for a given namespace and name.
func (s propagatedVersionNamespaceLister) Get(name string) (*v1alpha1.PropagatedVersion, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("propagatedversion"), name)
	}
	return obj.(*v1alpha1.PropagatedVersion), nil
}

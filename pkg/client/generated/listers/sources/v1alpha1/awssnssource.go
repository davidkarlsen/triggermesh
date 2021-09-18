/*
Copyright 2021 TriggerMesh Inc.

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
	v1alpha1 "github.com/triggermesh/triggermesh/pkg/apis/sources/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// AWSSNSSourceLister helps list AWSSNSSources.
// All objects returned here must be treated as read-only.
type AWSSNSSourceLister interface {
	// List lists all AWSSNSSources in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.AWSSNSSource, err error)
	// AWSSNSSources returns an object that can list and get AWSSNSSources.
	AWSSNSSources(namespace string) AWSSNSSourceNamespaceLister
	AWSSNSSourceListerExpansion
}

// aWSSNSSourceLister implements the AWSSNSSourceLister interface.
type aWSSNSSourceLister struct {
	indexer cache.Indexer
}

// NewAWSSNSSourceLister returns a new AWSSNSSourceLister.
func NewAWSSNSSourceLister(indexer cache.Indexer) AWSSNSSourceLister {
	return &aWSSNSSourceLister{indexer: indexer}
}

// List lists all AWSSNSSources in the indexer.
func (s *aWSSNSSourceLister) List(selector labels.Selector) (ret []*v1alpha1.AWSSNSSource, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AWSSNSSource))
	})
	return ret, err
}

// AWSSNSSources returns an object that can list and get AWSSNSSources.
func (s *aWSSNSSourceLister) AWSSNSSources(namespace string) AWSSNSSourceNamespaceLister {
	return aWSSNSSourceNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AWSSNSSourceNamespaceLister helps list and get AWSSNSSources.
// All objects returned here must be treated as read-only.
type AWSSNSSourceNamespaceLister interface {
	// List lists all AWSSNSSources in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.AWSSNSSource, err error)
	// Get retrieves the AWSSNSSource from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.AWSSNSSource, error)
	AWSSNSSourceNamespaceListerExpansion
}

// aWSSNSSourceNamespaceLister implements the AWSSNSSourceNamespaceLister
// interface.
type aWSSNSSourceNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AWSSNSSources in the indexer for a given namespace.
func (s aWSSNSSourceNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AWSSNSSource, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AWSSNSSource))
	})
	return ret, err
}

// Get retrieves the AWSSNSSource from the indexer for a given namespace and name.
func (s aWSSNSSourceNamespaceLister) Get(name string) (*v1alpha1.AWSSNSSource, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("awssnssource"), name)
	}
	return obj.(*v1alpha1.AWSSNSSource), nil
}

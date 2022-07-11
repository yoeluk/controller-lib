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
	v1alpha1 "github.com/yoeluk/controller-lib/pkg/apis/dbcontroller/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// RdsdbLister helps list Rdsdbs.
// All objects returned here must be treated as read-only.
type RdsdbLister interface {
	// List lists all Rdsdbs in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Rdsdb, err error)
	// Rdsdbs returns an object that can list and get Rdsdbs.
	Rdsdbs(namespace string) RdsdbNamespaceLister
	RdsdbListerExpansion
}

// rdsdbLister implements the RdsdbLister interface.
type rdsdbLister struct {
	indexer cache.Indexer
}

// NewRdsdbLister returns a new RdsdbLister.
func NewRdsdbLister(indexer cache.Indexer) RdsdbLister {
	return &rdsdbLister{indexer: indexer}
}

// List lists all Rdsdbs in the indexer.
func (s *rdsdbLister) List(selector labels.Selector) (ret []*v1alpha1.Rdsdb, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Rdsdb))
	})
	return ret, err
}

// Rdsdbs returns an object that can list and get Rdsdbs.
func (s *rdsdbLister) Rdsdbs(namespace string) RdsdbNamespaceLister {
	return rdsdbNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// RdsdbNamespaceLister helps list and get Rdsdbs.
// All objects returned here must be treated as read-only.
type RdsdbNamespaceLister interface {
	// List lists all Rdsdbs in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Rdsdb, err error)
	// Get retrieves the Rdsdb from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Rdsdb, error)
	RdsdbNamespaceListerExpansion
}

// rdsdbNamespaceLister implements the RdsdbNamespaceLister
// interface.
type rdsdbNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Rdsdbs in the indexer for a given namespace.
func (s rdsdbNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Rdsdb, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Rdsdb))
	})
	return ret, err
}

// Get retrieves the Rdsdb from the indexer for a given namespace and name.
func (s rdsdbNamespaceLister) Get(name string) (*v1alpha1.Rdsdb, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("rdsdb"), name)
	}
	return obj.(*v1alpha1.Rdsdb), nil
}
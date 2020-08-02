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
	v1 "github.com/jackistom/kubectl-Ecs/pkg/apis/jackistom/v1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// SylixosLister helps list Sylixoses.
// All objects returned here must be treated as read-only.
type SylixosLister interface {
	// List lists all Sylixoses in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.Sylixos, err error)
	// Sylixoses returns an object that can list and get Sylixoses.
	Sylixoses(namespace string) SylixosNamespaceLister
	SylixosListerExpansion
}

// sylixosLister implements the SylixosLister interface.
type sylixosLister struct {
	indexer cache.Indexer
}

// NewSylixosLister returns a new SylixosLister.
func NewSylixosLister(indexer cache.Indexer) SylixosLister {
	return &sylixosLister{indexer: indexer}
}

// List lists all Sylixoses in the indexer.
func (s *sylixosLister) List(selector labels.Selector) (ret []*v1.Sylixos, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Sylixos))
	})
	return ret, err
}

// Sylixoses returns an object that can list and get Sylixoses.
func (s *sylixosLister) Sylixoses(namespace string) SylixosNamespaceLister {
	return sylixosNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// SylixosNamespaceLister helps list and get Sylixoses.
// All objects returned here must be treated as read-only.
type SylixosNamespaceLister interface {
	// List lists all Sylixoses in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.Sylixos, err error)
	// Get retrieves the Sylixos from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.Sylixos, error)
	SylixosNamespaceListerExpansion
}

// sylixosNamespaceLister implements the SylixosNamespaceLister
// interface.
type sylixosNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Sylixoses in the indexer for a given namespace.
func (s sylixosNamespaceLister) List(selector labels.Selector) (ret []*v1.Sylixos, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Sylixos))
	})
	return ret, err
}

// Get retrieves the Sylixos from the indexer for a given namespace and name.
func (s sylixosNamespaceLister) Get(name string) (*v1.Sylixos, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("sylixos"), name)
	}
	return obj.(*v1.Sylixos), nil
}
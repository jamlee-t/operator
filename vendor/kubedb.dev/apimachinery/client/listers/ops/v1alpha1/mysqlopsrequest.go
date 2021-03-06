/*
Copyright AppsCode Inc. and Contributors

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
	v1alpha1 "kubedb.dev/apimachinery/apis/ops/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// MySQLOpsRequestLister helps list MySQLOpsRequests.
// All objects returned here must be treated as read-only.
type MySQLOpsRequestLister interface {
	// List lists all MySQLOpsRequests in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.MySQLOpsRequest, err error)
	// MySQLOpsRequests returns an object that can list and get MySQLOpsRequests.
	MySQLOpsRequests(namespace string) MySQLOpsRequestNamespaceLister
	MySQLOpsRequestListerExpansion
}

// mySQLOpsRequestLister implements the MySQLOpsRequestLister interface.
type mySQLOpsRequestLister struct {
	indexer cache.Indexer
}

// NewMySQLOpsRequestLister returns a new MySQLOpsRequestLister.
func NewMySQLOpsRequestLister(indexer cache.Indexer) MySQLOpsRequestLister {
	return &mySQLOpsRequestLister{indexer: indexer}
}

// List lists all MySQLOpsRequests in the indexer.
func (s *mySQLOpsRequestLister) List(selector labels.Selector) (ret []*v1alpha1.MySQLOpsRequest, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.MySQLOpsRequest))
	})
	return ret, err
}

// MySQLOpsRequests returns an object that can list and get MySQLOpsRequests.
func (s *mySQLOpsRequestLister) MySQLOpsRequests(namespace string) MySQLOpsRequestNamespaceLister {
	return mySQLOpsRequestNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// MySQLOpsRequestNamespaceLister helps list and get MySQLOpsRequests.
// All objects returned here must be treated as read-only.
type MySQLOpsRequestNamespaceLister interface {
	// List lists all MySQLOpsRequests in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.MySQLOpsRequest, err error)
	// Get retrieves the MySQLOpsRequest from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.MySQLOpsRequest, error)
	MySQLOpsRequestNamespaceListerExpansion
}

// mySQLOpsRequestNamespaceLister implements the MySQLOpsRequestNamespaceLister
// interface.
type mySQLOpsRequestNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all MySQLOpsRequests in the indexer for a given namespace.
func (s mySQLOpsRequestNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.MySQLOpsRequest, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.MySQLOpsRequest))
	})
	return ret, err
}

// Get retrieves the MySQLOpsRequest from the indexer for a given namespace and name.
func (s mySQLOpsRequestNamespaceLister) Get(name string) (*v1alpha1.MySQLOpsRequest, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("mysqlopsrequest"), name)
	}
	return obj.(*v1alpha1.MySQLOpsRequest), nil
}

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
// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/Tencent/bk-bcs/bcs-k8s/kubernetes/apis/cloud/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// CloudIPLister helps list CloudIPs.
type CloudIPLister interface {
	// List lists all CloudIPs in the indexer.
	List(selector labels.Selector) (ret []*v1.CloudIP, err error)
	// CloudIPs returns an object that can list and get CloudIPs.
	CloudIPs(namespace string) CloudIPNamespaceLister
	CloudIPListerExpansion
}

// cloudIPLister implements the CloudIPLister interface.
type cloudIPLister struct {
	indexer cache.Indexer
}

// NewCloudIPLister returns a new CloudIPLister.
func NewCloudIPLister(indexer cache.Indexer) CloudIPLister {
	return &cloudIPLister{indexer: indexer}
}

// List lists all CloudIPs in the indexer.
func (s *cloudIPLister) List(selector labels.Selector) (ret []*v1.CloudIP, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.CloudIP))
	})
	return ret, err
}

// CloudIPs returns an object that can list and get CloudIPs.
func (s *cloudIPLister) CloudIPs(namespace string) CloudIPNamespaceLister {
	return cloudIPNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// CloudIPNamespaceLister helps list and get CloudIPs.
type CloudIPNamespaceLister interface {
	// List lists all CloudIPs in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.CloudIP, err error)
	// Get retrieves the CloudIP from the indexer for a given namespace and name.
	Get(name string) (*v1.CloudIP, error)
	CloudIPNamespaceListerExpansion
}

// cloudIPNamespaceLister implements the CloudIPNamespaceLister
// interface.
type cloudIPNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all CloudIPs in the indexer for a given namespace.
func (s cloudIPNamespaceLister) List(selector labels.Selector) (ret []*v1.CloudIP, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.CloudIP))
	})
	return ret, err
}

// Get retrieves the CloudIP from the indexer for a given namespace and name.
func (s cloudIPNamespaceLister) Get(name string) (*v1.CloudIP, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("cloudip"), name)
	}
	return obj.(*v1.CloudIP), nil
}

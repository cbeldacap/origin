/*
Copyright 2017 The Kubernetes Authors.

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

// This file was automatically generated by lister-gen with arguments: --go-header-file=vendor/github.com/kubernetes/repo-infra/verify/boilerplate/boilerplate.go.txt --input-dirs=[github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog,github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog/v1alpha1] --output-package=github.com/kubernetes-incubator/service-catalog/pkg/client/listers

package v1alpha1

import (
	servicecatalog "github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog"
	v1alpha1 "github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog/v1alpha1"
	"k8s.io/kubernetes/pkg/api/errors"
	v1 "k8s.io/kubernetes/pkg/api/v1"
	"k8s.io/kubernetes/pkg/client/cache"
	"k8s.io/kubernetes/pkg/labels"
)

// ServiceClassLister helps list ServiceClasses.
type ServiceClassLister interface {
	// List lists all ServiceClasses in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.ServiceClass, err error)
	// Get retrieves the ServiceClass from the index for a given name.
	Get(name string) (*v1alpha1.ServiceClass, error)
	ServiceClassListerExpansion
}

// serviceClassLister implements the ServiceClassLister interface.
type serviceClassLister struct {
	indexer cache.Indexer
}

// NewServiceClassLister returns a new ServiceClassLister.
func NewServiceClassLister(indexer cache.Indexer) ServiceClassLister {
	return &serviceClassLister{indexer: indexer}
}

// List lists all ServiceClasses in the indexer.
func (s *serviceClassLister) List(selector labels.Selector) (ret []*v1alpha1.ServiceClass, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ServiceClass))
	})
	return ret, err
}

// Get retrieves the ServiceClass from the index for a given name.
func (s *serviceClassLister) Get(name string) (*v1alpha1.ServiceClass, error) {
	key := &v1alpha1.ServiceClass{ObjectMeta: v1.ObjectMeta{Name: name}}
	obj, exists, err := s.indexer.Get(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(servicecatalog.Resource("serviceclass"), name)
	}
	return obj.(*v1alpha1.ServiceClass), nil
}

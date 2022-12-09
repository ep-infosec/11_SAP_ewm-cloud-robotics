// Copyright (c) 2019 SAP SE or an SAP affiliate company. All rights reserved.
//
// This file is part of ewm-cloud-robotics
// (see https://github.com/SAP/ewm-cloud-robotics).
//
// This file is licensed under the Apache Software License, v. 2 except as noted
// otherwise in the LICENSE file (https://github.com/SAP/ewm-cloud-robotics/blob/main/LICENSE)
//

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/SAP/ewm-cloud-robotics/go/pkg/apis/ewm/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// OrderReservationLister helps list OrderReservations.
// All objects returned here must be treated as read-only.
type OrderReservationLister interface {
	// List lists all OrderReservations in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.OrderReservation, err error)
	// OrderReservations returns an object that can list and get OrderReservations.
	OrderReservations(namespace string) OrderReservationNamespaceLister
	OrderReservationListerExpansion
}

// orderReservationLister implements the OrderReservationLister interface.
type orderReservationLister struct {
	indexer cache.Indexer
}

// NewOrderReservationLister returns a new OrderReservationLister.
func NewOrderReservationLister(indexer cache.Indexer) OrderReservationLister {
	return &orderReservationLister{indexer: indexer}
}

// List lists all OrderReservations in the indexer.
func (s *orderReservationLister) List(selector labels.Selector) (ret []*v1alpha1.OrderReservation, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.OrderReservation))
	})
	return ret, err
}

// OrderReservations returns an object that can list and get OrderReservations.
func (s *orderReservationLister) OrderReservations(namespace string) OrderReservationNamespaceLister {
	return orderReservationNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// OrderReservationNamespaceLister helps list and get OrderReservations.
// All objects returned here must be treated as read-only.
type OrderReservationNamespaceLister interface {
	// List lists all OrderReservations in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.OrderReservation, err error)
	// Get retrieves the OrderReservation from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.OrderReservation, error)
	OrderReservationNamespaceListerExpansion
}

// orderReservationNamespaceLister implements the OrderReservationNamespaceLister
// interface.
type orderReservationNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all OrderReservations in the indexer for a given namespace.
func (s orderReservationNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.OrderReservation, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.OrderReservation))
	})
	return ret, err
}

// Get retrieves the OrderReservation from the indexer for a given namespace and name.
func (s orderReservationNamespaceLister) Get(name string) (*v1alpha1.OrderReservation, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("orderreservation"), name)
	}
	return obj.(*v1alpha1.OrderReservation), nil
}
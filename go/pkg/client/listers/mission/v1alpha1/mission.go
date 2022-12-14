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
	v1alpha1 "github.com/SAP/ewm-cloud-robotics/go/pkg/apis/mission/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// MissionLister helps list Missions.
// All objects returned here must be treated as read-only.
type MissionLister interface {
	// List lists all Missions in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Mission, err error)
	// Missions returns an object that can list and get Missions.
	Missions(namespace string) MissionNamespaceLister
	MissionListerExpansion
}

// missionLister implements the MissionLister interface.
type missionLister struct {
	indexer cache.Indexer
}

// NewMissionLister returns a new MissionLister.
func NewMissionLister(indexer cache.Indexer) MissionLister {
	return &missionLister{indexer: indexer}
}

// List lists all Missions in the indexer.
func (s *missionLister) List(selector labels.Selector) (ret []*v1alpha1.Mission, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Mission))
	})
	return ret, err
}

// Missions returns an object that can list and get Missions.
func (s *missionLister) Missions(namespace string) MissionNamespaceLister {
	return missionNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// MissionNamespaceLister helps list and get Missions.
// All objects returned here must be treated as read-only.
type MissionNamespaceLister interface {
	// List lists all Missions in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Mission, err error)
	// Get retrieves the Mission from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Mission, error)
	MissionNamespaceListerExpansion
}

// missionNamespaceLister implements the MissionNamespaceLister
// interface.
type missionNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Missions in the indexer for a given namespace.
func (s missionNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Mission, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Mission))
	})
	return ret, err
}

// Get retrieves the Mission from the indexer for a given namespace and name.
func (s missionNamespaceLister) Get(name string) (*v1alpha1.Mission, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("mission"), name)
	}
	return obj.(*v1alpha1.Mission), nil
}

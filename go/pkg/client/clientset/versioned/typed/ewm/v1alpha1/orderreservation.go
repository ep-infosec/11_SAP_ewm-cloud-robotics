// Copyright (c) 2019 SAP SE or an SAP affiliate company. All rights reserved.
//
// This file is part of ewm-cloud-robotics
// (see https://github.com/SAP/ewm-cloud-robotics).
//
// This file is licensed under the Apache Software License, v. 2 except as noted
// otherwise in the LICENSE file (https://github.com/SAP/ewm-cloud-robotics/blob/main/LICENSE)
//

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "github.com/SAP/ewm-cloud-robotics/go/pkg/apis/ewm/v1alpha1"
	scheme "github.com/SAP/ewm-cloud-robotics/go/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// OrderReservationsGetter has a method to return a OrderReservationInterface.
// A group's client should implement this interface.
type OrderReservationsGetter interface {
	OrderReservations(namespace string) OrderReservationInterface
}

// OrderReservationInterface has methods to work with OrderReservation resources.
type OrderReservationInterface interface {
	Create(ctx context.Context, orderReservation *v1alpha1.OrderReservation, opts v1.CreateOptions) (*v1alpha1.OrderReservation, error)
	Update(ctx context.Context, orderReservation *v1alpha1.OrderReservation, opts v1.UpdateOptions) (*v1alpha1.OrderReservation, error)
	UpdateStatus(ctx context.Context, orderReservation *v1alpha1.OrderReservation, opts v1.UpdateOptions) (*v1alpha1.OrderReservation, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.OrderReservation, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.OrderReservationList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.OrderReservation, err error)
	OrderReservationExpansion
}

// orderReservations implements OrderReservationInterface
type orderReservations struct {
	client rest.Interface
	ns     string
}

// newOrderReservations returns a OrderReservations
func newOrderReservations(c *EwmV1alpha1Client, namespace string) *orderReservations {
	return &orderReservations{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the orderReservation, and returns the corresponding orderReservation object, and an error if there is any.
func (c *orderReservations) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.OrderReservation, err error) {
	result = &v1alpha1.OrderReservation{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("orderreservations").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of OrderReservations that match those selectors.
func (c *orderReservations) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.OrderReservationList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.OrderReservationList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("orderreservations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested orderReservations.
func (c *orderReservations) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("orderreservations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a orderReservation and creates it.  Returns the server's representation of the orderReservation, and an error, if there is any.
func (c *orderReservations) Create(ctx context.Context, orderReservation *v1alpha1.OrderReservation, opts v1.CreateOptions) (result *v1alpha1.OrderReservation, err error) {
	result = &v1alpha1.OrderReservation{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("orderreservations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(orderReservation).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a orderReservation and updates it. Returns the server's representation of the orderReservation, and an error, if there is any.
func (c *orderReservations) Update(ctx context.Context, orderReservation *v1alpha1.OrderReservation, opts v1.UpdateOptions) (result *v1alpha1.OrderReservation, err error) {
	result = &v1alpha1.OrderReservation{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("orderreservations").
		Name(orderReservation.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(orderReservation).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *orderReservations) UpdateStatus(ctx context.Context, orderReservation *v1alpha1.OrderReservation, opts v1.UpdateOptions) (result *v1alpha1.OrderReservation, err error) {
	result = &v1alpha1.OrderReservation{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("orderreservations").
		Name(orderReservation.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(orderReservation).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the orderReservation and deletes it. Returns an error if one occurs.
func (c *orderReservations) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("orderreservations").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *orderReservations) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("orderreservations").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched orderReservation.
func (c *orderReservations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.OrderReservation, err error) {
	result = &v1alpha1.OrderReservation{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("orderreservations").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

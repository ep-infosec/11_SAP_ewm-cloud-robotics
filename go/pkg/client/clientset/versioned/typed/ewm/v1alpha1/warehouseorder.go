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

// WarehouseOrdersGetter has a method to return a WarehouseOrderInterface.
// A group's client should implement this interface.
type WarehouseOrdersGetter interface {
	WarehouseOrders(namespace string) WarehouseOrderInterface
}

// WarehouseOrderInterface has methods to work with WarehouseOrder resources.
type WarehouseOrderInterface interface {
	Create(ctx context.Context, warehouseOrder *v1alpha1.WarehouseOrder, opts v1.CreateOptions) (*v1alpha1.WarehouseOrder, error)
	Update(ctx context.Context, warehouseOrder *v1alpha1.WarehouseOrder, opts v1.UpdateOptions) (*v1alpha1.WarehouseOrder, error)
	UpdateStatus(ctx context.Context, warehouseOrder *v1alpha1.WarehouseOrder, opts v1.UpdateOptions) (*v1alpha1.WarehouseOrder, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.WarehouseOrder, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.WarehouseOrderList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.WarehouseOrder, err error)
	WarehouseOrderExpansion
}

// warehouseOrders implements WarehouseOrderInterface
type warehouseOrders struct {
	client rest.Interface
	ns     string
}

// newWarehouseOrders returns a WarehouseOrders
func newWarehouseOrders(c *EwmV1alpha1Client, namespace string) *warehouseOrders {
	return &warehouseOrders{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the warehouseOrder, and returns the corresponding warehouseOrder object, and an error if there is any.
func (c *warehouseOrders) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.WarehouseOrder, err error) {
	result = &v1alpha1.WarehouseOrder{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("warehouseorders").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of WarehouseOrders that match those selectors.
func (c *warehouseOrders) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.WarehouseOrderList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.WarehouseOrderList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("warehouseorders").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested warehouseOrders.
func (c *warehouseOrders) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("warehouseorders").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a warehouseOrder and creates it.  Returns the server's representation of the warehouseOrder, and an error, if there is any.
func (c *warehouseOrders) Create(ctx context.Context, warehouseOrder *v1alpha1.WarehouseOrder, opts v1.CreateOptions) (result *v1alpha1.WarehouseOrder, err error) {
	result = &v1alpha1.WarehouseOrder{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("warehouseorders").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(warehouseOrder).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a warehouseOrder and updates it. Returns the server's representation of the warehouseOrder, and an error, if there is any.
func (c *warehouseOrders) Update(ctx context.Context, warehouseOrder *v1alpha1.WarehouseOrder, opts v1.UpdateOptions) (result *v1alpha1.WarehouseOrder, err error) {
	result = &v1alpha1.WarehouseOrder{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("warehouseorders").
		Name(warehouseOrder.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(warehouseOrder).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *warehouseOrders) UpdateStatus(ctx context.Context, warehouseOrder *v1alpha1.WarehouseOrder, opts v1.UpdateOptions) (result *v1alpha1.WarehouseOrder, err error) {
	result = &v1alpha1.WarehouseOrder{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("warehouseorders").
		Name(warehouseOrder.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(warehouseOrder).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the warehouseOrder and deletes it. Returns an error if one occurs.
func (c *warehouseOrders) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("warehouseorders").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *warehouseOrders) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("warehouseorders").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched warehouseOrder.
func (c *warehouseOrders) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.WarehouseOrder, err error) {
	result = &v1alpha1.WarehouseOrder{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("warehouseorders").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

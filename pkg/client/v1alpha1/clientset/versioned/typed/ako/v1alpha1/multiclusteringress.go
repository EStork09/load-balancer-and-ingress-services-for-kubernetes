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

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "github.com/vmware/load-balancer-and-ingress-services-for-kubernetes/pkg/apis/ako/v1alpha1"
	scheme "github.com/vmware/load-balancer-and-ingress-services-for-kubernetes/pkg/client/v1alpha1/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// MultiClusterIngressesGetter has a method to return a MultiClusterIngressInterface.
// A group's client should implement this interface.
type MultiClusterIngressesGetter interface {
	MultiClusterIngresses(namespace string) MultiClusterIngressInterface
}

// MultiClusterIngressInterface has methods to work with MultiClusterIngress resources.
type MultiClusterIngressInterface interface {
	Create(ctx context.Context, multiClusterIngress *v1alpha1.MultiClusterIngress, opts v1.CreateOptions) (*v1alpha1.MultiClusterIngress, error)
	Update(ctx context.Context, multiClusterIngress *v1alpha1.MultiClusterIngress, opts v1.UpdateOptions) (*v1alpha1.MultiClusterIngress, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.MultiClusterIngress, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.MultiClusterIngressList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.MultiClusterIngress, err error)
	MultiClusterIngressExpansion
}

// multiClusterIngresses implements MultiClusterIngressInterface
type multiClusterIngresses struct {
	client rest.Interface
	ns     string
}

// newMultiClusterIngresses returns a MultiClusterIngresses
func newMultiClusterIngresses(c *AkoV1alpha1Client, namespace string) *multiClusterIngresses {
	return &multiClusterIngresses{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the multiClusterIngress, and returns the corresponding multiClusterIngress object, and an error if there is any.
func (c *multiClusterIngresses) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.MultiClusterIngress, err error) {
	result = &v1alpha1.MultiClusterIngress{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("multiclusteringresses").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of MultiClusterIngresses that match those selectors.
func (c *multiClusterIngresses) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.MultiClusterIngressList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.MultiClusterIngressList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("multiclusteringresses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested multiClusterIngresses.
func (c *multiClusterIngresses) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("multiclusteringresses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a multiClusterIngress and creates it.  Returns the server's representation of the multiClusterIngress, and an error, if there is any.
func (c *multiClusterIngresses) Create(ctx context.Context, multiClusterIngress *v1alpha1.MultiClusterIngress, opts v1.CreateOptions) (result *v1alpha1.MultiClusterIngress, err error) {
	result = &v1alpha1.MultiClusterIngress{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("multiclusteringresses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(multiClusterIngress).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a multiClusterIngress and updates it. Returns the server's representation of the multiClusterIngress, and an error, if there is any.
func (c *multiClusterIngresses) Update(ctx context.Context, multiClusterIngress *v1alpha1.MultiClusterIngress, opts v1.UpdateOptions) (result *v1alpha1.MultiClusterIngress, err error) {
	result = &v1alpha1.MultiClusterIngress{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("multiclusteringresses").
		Name(multiClusterIngress.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(multiClusterIngress).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the multiClusterIngress and deletes it. Returns an error if one occurs.
func (c *multiClusterIngresses) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("multiclusteringresses").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *multiClusterIngresses) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("multiclusteringresses").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched multiClusterIngress.
func (c *multiClusterIngresses) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.MultiClusterIngress, err error) {
	result = &v1alpha1.MultiClusterIngress{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("multiclusteringresses").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

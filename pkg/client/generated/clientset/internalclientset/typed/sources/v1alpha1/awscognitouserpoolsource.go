/*
Copyright 2021 TriggerMesh Inc.

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

	v1alpha1 "github.com/triggermesh/triggermesh/pkg/apis/sources/v1alpha1"
	scheme "github.com/triggermesh/triggermesh/pkg/client/generated/clientset/internalclientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// AWSCognitoUserPoolSourcesGetter has a method to return a AWSCognitoUserPoolSourceInterface.
// A group's client should implement this interface.
type AWSCognitoUserPoolSourcesGetter interface {
	AWSCognitoUserPoolSources(namespace string) AWSCognitoUserPoolSourceInterface
}

// AWSCognitoUserPoolSourceInterface has methods to work with AWSCognitoUserPoolSource resources.
type AWSCognitoUserPoolSourceInterface interface {
	Create(ctx context.Context, aWSCognitoUserPoolSource *v1alpha1.AWSCognitoUserPoolSource, opts v1.CreateOptions) (*v1alpha1.AWSCognitoUserPoolSource, error)
	Update(ctx context.Context, aWSCognitoUserPoolSource *v1alpha1.AWSCognitoUserPoolSource, opts v1.UpdateOptions) (*v1alpha1.AWSCognitoUserPoolSource, error)
	UpdateStatus(ctx context.Context, aWSCognitoUserPoolSource *v1alpha1.AWSCognitoUserPoolSource, opts v1.UpdateOptions) (*v1alpha1.AWSCognitoUserPoolSource, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.AWSCognitoUserPoolSource, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.AWSCognitoUserPoolSourceList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.AWSCognitoUserPoolSource, err error)
	AWSCognitoUserPoolSourceExpansion
}

// aWSCognitoUserPoolSources implements AWSCognitoUserPoolSourceInterface
type aWSCognitoUserPoolSources struct {
	client rest.Interface
	ns     string
}

// newAWSCognitoUserPoolSources returns a AWSCognitoUserPoolSources
func newAWSCognitoUserPoolSources(c *SourcesV1alpha1Client, namespace string) *aWSCognitoUserPoolSources {
	return &aWSCognitoUserPoolSources{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the aWSCognitoUserPoolSource, and returns the corresponding aWSCognitoUserPoolSource object, and an error if there is any.
func (c *aWSCognitoUserPoolSources) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.AWSCognitoUserPoolSource, err error) {
	result = &v1alpha1.AWSCognitoUserPoolSource{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awscognitouserpoolsources").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AWSCognitoUserPoolSources that match those selectors.
func (c *aWSCognitoUserPoolSources) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.AWSCognitoUserPoolSourceList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AWSCognitoUserPoolSourceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awscognitouserpoolsources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested aWSCognitoUserPoolSources.
func (c *aWSCognitoUserPoolSources) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awscognitouserpoolsources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a aWSCognitoUserPoolSource and creates it.  Returns the server's representation of the aWSCognitoUserPoolSource, and an error, if there is any.
func (c *aWSCognitoUserPoolSources) Create(ctx context.Context, aWSCognitoUserPoolSource *v1alpha1.AWSCognitoUserPoolSource, opts v1.CreateOptions) (result *v1alpha1.AWSCognitoUserPoolSource, err error) {
	result = &v1alpha1.AWSCognitoUserPoolSource{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awscognitouserpoolsources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(aWSCognitoUserPoolSource).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a aWSCognitoUserPoolSource and updates it. Returns the server's representation of the aWSCognitoUserPoolSource, and an error, if there is any.
func (c *aWSCognitoUserPoolSources) Update(ctx context.Context, aWSCognitoUserPoolSource *v1alpha1.AWSCognitoUserPoolSource, opts v1.UpdateOptions) (result *v1alpha1.AWSCognitoUserPoolSource, err error) {
	result = &v1alpha1.AWSCognitoUserPoolSource{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awscognitouserpoolsources").
		Name(aWSCognitoUserPoolSource.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(aWSCognitoUserPoolSource).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *aWSCognitoUserPoolSources) UpdateStatus(ctx context.Context, aWSCognitoUserPoolSource *v1alpha1.AWSCognitoUserPoolSource, opts v1.UpdateOptions) (result *v1alpha1.AWSCognitoUserPoolSource, err error) {
	result = &v1alpha1.AWSCognitoUserPoolSource{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awscognitouserpoolsources").
		Name(aWSCognitoUserPoolSource.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(aWSCognitoUserPoolSource).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the aWSCognitoUserPoolSource and deletes it. Returns an error if one occurs.
func (c *aWSCognitoUserPoolSources) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awscognitouserpoolsources").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *aWSCognitoUserPoolSources) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awscognitouserpoolsources").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched aWSCognitoUserPoolSource.
func (c *aWSCognitoUserPoolSources) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.AWSCognitoUserPoolSource, err error) {
	result = &v1alpha1.AWSCognitoUserPoolSource{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awscognitouserpoolsources").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

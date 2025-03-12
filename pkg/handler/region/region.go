/*
Copyright 2024-2025 the Unikorn Authors.

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

package region

import (
	"context"
	"encoding/base64"
	"errors"

	coreopenapi "github.com/unikorn-cloud/core/pkg/openapi"
	"github.com/unikorn-cloud/core/pkg/server/conversion"
	identityapi "github.com/unikorn-cloud/identity/pkg/openapi"
	"github.com/unikorn-cloud/identity/pkg/rbac"
	unikornv1 "github.com/unikorn-cloud/region/pkg/apis/unikorn/v1alpha1"
	"github.com/unikorn-cloud/region/pkg/openapi"
	"github.com/unikorn-cloud/region/pkg/providers"
	"github.com/unikorn-cloud/region/pkg/providers/openstack"

	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	// ErrRegionNotFound is raised when a region doesn't exist.
	ErrRegionNotFound = errors.New("region doesn't exist")

	// ErrRegionProviderUnimplmented is raised when you haven't written
	// it yet!
	ErrRegionProviderUnimplmented = errors.New("region provider unimplmented")
)

type Client struct {
	client    client.Client
	namespace string
}

func NewClient(client client.Client, namespace string) *Client {
	return &Client{
		client:    client,
		namespace: namespace,
	}
}

// list is a canonical lister function that allows filtering to be applied
// in one place e.g. health, ownership, etc.
func (c *Client) list(ctx context.Context) (*unikornv1.RegionList, error) {
	var regions unikornv1.RegionList

	if err := c.client.List(ctx, &regions, &client.ListOptions{Namespace: c.namespace}); err != nil {
		return nil, err
	}

	return &regions, nil
}

func findRegion(regions *unikornv1.RegionList, regionID string) (*unikornv1.Region, error) {
	for i := range regions.Items {
		if regions.Items[i].Name == regionID {
			return &regions.Items[i], nil
		}
	}

	return nil, ErrRegionNotFound
}

//nolint:gochecknoglobals
var cache = map[string]providers.Provider{}

func (c Client) newProvider(ctx context.Context, region *unikornv1.Region) (providers.Provider, error) {
	//nolint:gocritic
	switch region.Spec.Provider {
	case unikornv1.ProviderOpenstack:
		return openstack.New(ctx, c.client, region)
	}

	return nil, ErrRegionProviderUnimplmented
}

func (c *Client) Provider(ctx context.Context, regionID string) (providers.Provider, error) {
	regions, err := c.list(ctx)
	if err != nil {
		return nil, err
	}

	region, err := findRegion(regions, regionID)
	if err != nil {
		return nil, err
	}

	if provider, ok := cache[region.Name]; ok {
		return provider, nil
	}

	provider, err := c.newProvider(ctx, region)
	if err != nil {
		return nil, err
	}

	cache[region.Name] = provider

	return provider, nil
}

func convertRegionType(in unikornv1.Provider) openapi.RegionType {
	switch in {
	case unikornv1.ProviderKubernetes:
		return openapi.Kubernetes
	case unikornv1.ProviderOpenstack:
		return openapi.Openstack
	}

	return ""
}

func convert(ctx context.Context, in *unikornv1.Region) *openapi.RegionRead {
	out := &openapi.RegionRead{
		Metadata: conversion.ResourceReadMetadata(in, in.Spec.Tags, coreopenapi.ResourceProvisioningStatusProvisioned),
		Spec: openapi.RegionSpec{
			Type: convertRegionType(in.Spec.Provider),
		},
	}

	// Calculate any region specific configuration.
	switch in.Spec.Provider {
	case unikornv1.ProviderKubernetes:
		if err := rbac.AllowGlobalScope(ctx, "region:regions/admin", identityapi.Read); err == nil {
			out.Spec.Kubernetes = &openapi.RegionKubernetes{
				Kubeconfig: base64.RawURLEncoding.EncodeToString(in.Spec.Kubernetes.Kubeconfig),
			}
		}
	case unikornv1.ProviderOpenstack:
		if in.Spec.Openstack.Network != nil && in.Spec.Openstack.Network.ProviderNetworks != nil {
			out.Spec.Features.PhysicalNetworks = true
		}
	}

	return out
}

func convertList(ctx context.Context, in *unikornv1.RegionList) openapi.Regions {
	out := make(openapi.Regions, len(in.Items))

	for i := range in.Items {
		out[i] = *convert(ctx, &in.Items[i])
	}

	return out
}

func (c *Client) List(ctx context.Context) (openapi.Regions, error) {
	regions, err := c.list(ctx)
	if err != nil {
		return nil, err
	}

	return convertList(ctx, regions), nil
}

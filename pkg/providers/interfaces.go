/*
Copyright 2024 the Unikorn Authors.

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

package providers

import (
	"context"

	unikornv1 "github.com/unikorn-cloud/region/pkg/apis/unikorn/v1alpha1"
	"github.com/unikorn-cloud/region/pkg/openapi"
)

// Providers are expected to provide a provider agnostic manner.
// They are also expected to provide any caching or memoization required
// to provide high performance and a decent UX.
type Provider interface {
	// Flavors list all available flavors.
	Flavors(ctx context.Context) (FlavorList, error)
	// Images lists all available images.
	Images(ctx context.Context) (ImageList, error)
	// CreateIdentity creates a new identity for cloud infrastructure.
	CreateIdentity(ctx context.Context, organizationID, projectID string, request *openapi.IdentityWrite) (*unikornv1.Identity, error)
	// DeleteIdentity cleans up an identity for cloud infrastructure.
	DeleteIdentity(ctx context.Context, identity *unikornv1.Identity) error
	// CreatePhysicalNetwork create a new physical network.
	CreatePhysicalNetwork(ctx context.Context, identity *unikornv1.Identity, request *openapi.PhysicalNetworkWrite) (*unikornv1.PhysicalNetwork, error)
	// ListExternalNetworks returns a list of external networks if the platform
	// supports such a concept.
	ListExternalNetworks(ctx context.Context) (ExternalNetworks, error)
}

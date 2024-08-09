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

package openstack

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"slices"
	"strings"
	"sync"

	"github.com/gophercloud/gophercloud/v2/openstack/identity/v3/roles"
	"github.com/gophercloud/utils/openstack/clientconfig"

	coreconstants "github.com/unikorn-cloud/core/pkg/constants"
	"github.com/unikorn-cloud/core/pkg/server/conversion"
	"github.com/unikorn-cloud/identity/pkg/middleware/authorization"
	unikornv1 "github.com/unikorn-cloud/region/pkg/apis/unikorn/v1alpha1"
	"github.com/unikorn-cloud/region/pkg/constants"
	"github.com/unikorn-cloud/region/pkg/openapi"
	"github.com/unikorn-cloud/region/pkg/providers"

	corev1 "k8s.io/api/core/v1"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/uuid"

	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/yaml"
)

var (
	ErrKeyUndefined = errors.New("a required key was not defined")
)

type Provider struct {
	// client is Kubernetes client.
	client client.Client

	// region is the current region configuration.
	region *unikornv1.Region

	// secret is the current region secret.
	secret *corev1.Secret

	domainID  string
	projectID string
	userID    string
	password  string

	// DO NOT USE DIRECTLY, CALL AN ACCESSOR.
	_identity *IdentityClient
	_compute  *ComputeClient
	_image    *ImageClient
	_network  *NetworkClient

	lock sync.Mutex
}

var _ providers.Provider = &Provider{}

func New(client client.Client, region *unikornv1.Region) *Provider {
	return &Provider{
		client: client,
		region: region,
	}
}

// serviceClientRefresh updates clients if they need to e.g. in the event
// of a configuration update.
// NOTE: you MUST get the lock before calling this function.
//
//nolint:cyclop
func (p *Provider) serviceClientRefresh(ctx context.Context) error {
	refresh := false

	region := &unikornv1.Region{}

	if err := p.client.Get(ctx, client.ObjectKey{Namespace: p.region.Namespace, Name: p.region.Name}, region); err != nil {
		return err
	}

	// If anything changes with the configuration, referesh the clients as they may
	// do caching.
	if !reflect.DeepEqual(region.Spec.Openstack, p.region.Spec.Openstack) {
		refresh = true
	}

	secretkey := client.ObjectKey{
		Namespace: region.Spec.Openstack.ServiceAccountSecret.Namespace,
		Name:      region.Spec.Openstack.ServiceAccountSecret.Name,
	}

	secret := &corev1.Secret{}

	if err := p.client.Get(ctx, secretkey, secret); err != nil {
		return err
	}

	// If the secret hasn't beed read yet, or has changed e.g. credential rotation
	// then refresh the clients as they cache the API token.
	if p.secret == nil || !reflect.DeepEqual(secret.Data, p.secret.Data) {
		refresh = true
	}

	// Nothing to do, use what's there.
	if !refresh {
		return nil
	}

	// Create the core credential provider.
	domainID, ok := secret.Data["domain-id"]
	if !ok {
		return fmt.Errorf("%w: domain-id", ErrKeyUndefined)
	}

	userID, ok := secret.Data["user-id"]
	if !ok {
		return fmt.Errorf("%w: user-id", ErrKeyUndefined)
	}

	password, ok := secret.Data["password"]
	if !ok {
		return fmt.Errorf("%w: password", ErrKeyUndefined)
	}

	projectID, ok := secret.Data["project-id"]
	if !ok {
		return fmt.Errorf("%w: project-id", ErrKeyUndefined)
	}

	// 'Regular' client calls to APIs for Nova, Glance etc. must to be project-scoped
	providerClient := NewPasswordProvider(region.Spec.Openstack.Endpoint, string(userID), string(password), string(projectID))

	// Identity client is scoped to a domain to use the manager role
	identity, err := NewIdentityClient(ctx, NewDomainScopedPasswordProvider(region.Spec.Openstack.Endpoint, string(userID), string(password), string(domainID)))
	if err != nil {
		return err
	}

	compute, err := NewComputeClient(ctx, providerClient, region.Spec.Openstack.Compute)
	if err != nil {
		return err
	}

	image, err := NewImageClient(ctx, providerClient, region.Spec.Openstack.Image)
	if err != nil {
		return err
	}

	network, err := NewNetworkClient(ctx, providerClient, region.Spec.Openstack.Network)
	if err != nil {
		return err
	}

	// Save the current configuration for checking next time.
	p.region = region
	p.secret = secret

	p.domainID = string(domainID)
	p.projectID = string(projectID)

	p.userID = string(userID)
	p.password = string(password)

	// Seve the clients
	p._identity = identity
	p._compute = compute
	p._image = image
	p._network = network

	return nil
}

func (p *Provider) identity(ctx context.Context) (*IdentityClient, error) {
	p.lock.Lock()
	defer p.lock.Unlock()

	if err := p.serviceClientRefresh(ctx); err != nil {
		return nil, err
	}

	return p._identity, nil
}

func (p *Provider) compute(ctx context.Context) (*ComputeClient, error) {
	p.lock.Lock()
	defer p.lock.Unlock()

	if err := p.serviceClientRefresh(ctx); err != nil {
		return nil, err
	}

	return p._compute, nil
}

func (p *Provider) image(ctx context.Context) (*ImageClient, error) {
	p.lock.Lock()
	defer p.lock.Unlock()

	if err := p.serviceClientRefresh(ctx); err != nil {
		return nil, err
	}

	return p._image, nil
}

func (p *Provider) network(ctx context.Context) (*NetworkClient, error) {
	p.lock.Lock()
	defer p.lock.Unlock()

	if err := p.serviceClientRefresh(ctx); err != nil {
		return nil, err
	}

	return p._network, nil
}

// Region returns the provider's region.
func (p *Provider) Region(ctx context.Context) (*unikornv1.Region, error) {
	// Get the newest version of the region.
	p.lock.Lock()
	defer p.lock.Unlock()

	if err := p.serviceClientRefresh(ctx); err != nil {
		return nil, err
	}

	return p.region, nil
}

// Flavors list all available flavors.
func (p *Provider) Flavors(ctx context.Context) (providers.FlavorList, error) {
	computeService, err := p.compute(ctx)
	if err != nil {
		return nil, err
	}

	resources, err := computeService.Flavors(ctx)
	if err != nil {
		return nil, err
	}

	result := make(providers.FlavorList, len(resources))

	for i := range resources {
		flavor := &resources[i]

		// API memory is in MiB, disk is in GB
		f := providers.Flavor{
			ID:     flavor.ID,
			Name:   flavor.Name,
			CPUs:   flavor.VCPUs,
			Memory: resource.NewQuantity(int64(flavor.RAM)<<20, resource.BinarySI),
			Disk:   resource.NewScaledQuantity(int64(flavor.Disk), resource.Giga),
		}

		// Apply any extra metadata to the flavor.
		if p.region.Spec.Openstack.Compute != nil && p.region.Spec.Openstack.Compute.Flavors != nil {
			i := slices.IndexFunc(p.region.Spec.Openstack.Compute.Flavors.Metadata, func(metadata unikornv1.FlavorMetadata) bool {
				return flavor.ID == metadata.ID
			})

			if i >= 0 {
				metadata := &p.region.Spec.Openstack.Compute.Flavors.Metadata[i]

				f.Baremetal = metadata.Baremetal

				if metadata.CPU != nil {
					f.CPUFamily = metadata.CPU.Family
				}

				if metadata.GPU != nil {
					f.GPU = &providers.GPU{
						// TODO: while these align, you should really put a
						// proper conversion in here.
						Vendor: providers.GPUVendor(metadata.GPU.Vendor),
						Model:  metadata.GPU.Model,
						Memory: metadata.GPU.Memory,
						Count:  metadata.GPU.Count,
					}
				}
			}
		}

		result[i] = f
	}

	return result, nil
}

// Images lists all available images.
func (p *Provider) Images(ctx context.Context) (providers.ImageList, error) {
	imageService, err := p.image(ctx)
	if err != nil {
		return nil, err
	}

	resources, err := imageService.Images(ctx)
	if err != nil {
		return nil, err
	}

	result := make(providers.ImageList, len(resources))

	for i := range resources {
		image := &resources[i]

		virtualization, _ := image.Properties["unikorn:virtualization"].(string)
		kubernetesVersion, _ := image.Properties["unikorn:kubernetes_version"].(string)

		providerImage := providers.Image{
			ID:                image.ID,
			Name:              image.Name,
			Created:           image.CreatedAt,
			Modified:          image.UpdatedAt,
			Virtualization:    providers.ImageVirtualization(virtualization),
			KubernetesVersion: kubernetesVersion,
		}

		if gpuVendor, ok := image.Properties["unikorn:gpu_vendor"].(string); ok {
			gpuDriver, ok := image.Properties["unikorn:gpu_driver_version"].(string)
			if !ok {
				// TODO: it's perhaps better to just skip this one, rather than
				// kill the entire service??
				return nil, fmt.Errorf("%w: GPU driver is not defined for image %s", ErrKeyUndefined, image.ID)
			}

			gpu := &providers.ImageGPU{
				Vendor: providers.GPUVendor(gpuVendor),
				Driver: gpuDriver,
			}

			if models, ok := image.Properties["unikorn:gpu_models"].(string); ok {
				gpu.Models = strings.Split(models, ",")
			}

			providerImage.GPU = gpu
		}

		result[i] = providerImage
	}

	return result, nil
}

const (
	// Projects are randomly named to avoid clashes, so we need to add some tags
	// in order to be able to reason about who they really belong to.  It is also
	// useful to have these in place so we can spot orphaned resources and garbage
	// collect them.
	OrganizationTag = "organization"
	ProjectTag      = "project"
)

// projectTags defines how to tag projects.
func projectTags(identity *unikornv1.OpenstackIdentity) []string {
	tags := []string{
		OrganizationTag + "=" + identity.Labels[coreconstants.OrganizationLabel],
		ProjectTag + "=" + identity.Labels[coreconstants.ProjectLabel],
	}

	return tags
}

func identityResourceName(identity *unikornv1.OpenstackIdentity) string {
	return "unikorn-identity-" + identity.Name
}

// provisionUser creates a new user in the managed domain with a random password.
// There is a 1:1 mapping of user to project, and the project name is unique in the
// domain, so just reuse this, we can clean them up at the same time.
func (p *Provider) provisionUser(ctx context.Context, identityService *IdentityClient, identity *unikornv1.OpenstackIdentity) error {
	if identity.Spec.UserID != nil {
		return nil
	}

	name := identityResourceName(identity)
	password := string(uuid.NewUUID())

	user, err := identityService.CreateUser(ctx, p.domainID, name, password)
	if err != nil {
		return err
	}

	identity.Spec.UserID = &user.ID
	identity.Spec.Password = &password

	return nil
}

// provisionProject creates a project per-cluster.  Cluster API provider Openstack is
// somewhat broken in that networks can alias and cause all kinds of disasters, so it's
// safest to have one cluster in one project so it has its own namespace.
func (p *Provider) provisionProject(ctx context.Context, identityService *IdentityClient, identity *unikornv1.OpenstackIdentity) error {
	if identity.Spec.ProjectID != nil {
		return nil
	}

	name := identityResourceName(identity)

	project, err := identityService.CreateProject(ctx, p.domainID, name, projectTags(identity))
	if err != nil {
		return err
	}

	identity.Spec.ProjectID = &project.ID

	return nil
}

// roleNameToID maps from something human readable to something Openstack will operate with
// because who doesn't like extra, slow, API calls...
func roleNameToID(roles []roles.Role, name string) (string, error) {
	for _, role := range roles {
		if role.Name == name {
			return role.ID, nil
		}
	}

	return "", fmt.Errorf("%w: role %s", ErrResourceNotFound, name)
}

// getRequiredRoles returns the roles required for a user to create, manage and delete
// a cluster.
func (p *Provider) getRequiredRoles() []string {
	if p.region.Spec.Openstack.Identity != nil && len(p.region.Spec.Openstack.Identity.ClusterRoles) > 0 {
		return p.region.Spec.Openstack.Identity.ClusterRoles
	}

	defaultRoles := []string{
		"member",
		"load-balancer_member",
	}

	return defaultRoles
}

// provisionProjectRoles creates a binding between our service account and the project
// with the required roles to provision an application credential that will allow cluster
// creation, deletion and life-cycle management.
func (p *Provider) provisionProjectRoles(ctx context.Context, identityService *IdentityClient, identity *unikornv1.OpenstackIdentity) error {
	allRoles, err := identityService.ListRoles(ctx)
	if err != nil {
		return err
	}

	for _, name := range p.getRequiredRoles() {
		roleID, err := roleNameToID(allRoles, name)
		if err != nil {
			return err
		}

		if err := identityService.CreateRoleAssignment(ctx, *identity.Spec.UserID, *identity.Spec.ProjectID, roleID); err != nil {
			return err
		}
	}

	return nil
}

func (p *Provider) provisionApplicationCredential(ctx context.Context, identity *unikornv1.OpenstackIdentity) error {
	if identity.Spec.ApplicationCredentialID != nil {
		return nil
	}

	// Rescope to the user/project...
	providerClient := NewPasswordProvider(p.region.Spec.Openstack.Endpoint, *identity.Spec.UserID, *identity.Spec.Password, *identity.Spec.ProjectID)

	identityService, err := NewIdentityClient(ctx, providerClient)
	if err != nil {
		return err
	}

	name := identityResourceName(identity)

	appcred, err := identityService.CreateApplicationCredential(ctx, *identity.Spec.UserID, name, "IaaS lifecycle management", p.getRequiredRoles())
	if err != nil {
		return err
	}

	identity.Spec.ApplicationCredentialID = &appcred.ID
	identity.Spec.ApplicationCredentialSecret = &appcred.Secret

	return nil
}

func (p *Provider) createClientConfig(identity *unikornv1.OpenstackIdentity) error {
	if identity.Spec.Cloud != nil {
		return nil
	}

	cloud := "cloud"

	clientConfig := &clientconfig.Clouds{
		Clouds: map[string]clientconfig.Cloud{
			cloud: {
				AuthType: clientconfig.AuthV3ApplicationCredential,
				AuthInfo: &clientconfig.AuthInfo{
					AuthURL:                     p.region.Spec.Openstack.Endpoint,
					ApplicationCredentialID:     *identity.Spec.ApplicationCredentialID,
					ApplicationCredentialSecret: *identity.Spec.ApplicationCredentialSecret,
				},
			},
		},
	}

	clientConfigYAML, err := yaml.Marshal(clientConfig)
	if err != nil {
		return err
	}

	identity.Spec.Cloud = &cloud
	identity.Spec.CloudConfig = clientConfigYAML

	return nil
}

func convertTag(in openapi.Tag) unikornv1.Tag {
	out := unikornv1.Tag{
		Name:  in.Name,
		Value: in.Value,
	}

	return out
}

func convertTagList(in *openapi.TagList) unikornv1.TagList {
	if in == nil {
		return nil
	}

	out := make(unikornv1.TagList, len(*in))

	for i := range *in {
		out[i] = convertTag((*in)[i])
	}

	return out
}

func (p *Provider) createIdentityServerGroup(ctx context.Context, identity *unikornv1.OpenstackIdentity) error {
	if identity.Spec.ServerGroupID != nil {
		return nil
	}

	// Rescope to the user/project...
	providerClient := NewPasswordProvider(p.region.Spec.Openstack.Endpoint, *identity.Spec.UserID, *identity.Spec.Password, *identity.Spec.ProjectID)

	computeService, err := NewComputeClient(ctx, providerClient, p.region.Spec.Openstack.Compute)
	if err != nil {
		return err
	}

	name := identityResourceName(identity)

	result, err := computeService.CreateServerGroup(ctx, name)
	if err != nil {
		return err
	}

	identity.Spec.ServerGroupID = &result.ID

	return nil
}

func (p *Provider) GetOpenstackIdentity(ctx context.Context, identity *unikornv1.Identity) (*unikornv1.OpenstackIdentity, error) {
	var result unikornv1.OpenstackIdentity

	if err := p.client.Get(ctx, client.ObjectKey{Namespace: identity.Namespace, Name: identity.Name}, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (p *Provider) GetOrCreateOpenstackIdentity(ctx context.Context, identity *unikornv1.Identity) (*unikornv1.OpenstackIdentity, bool, error) {
	create := false

	openstackIdentity, err := p.GetOpenstackIdentity(ctx, identity)
	if err != nil {
		if !kerrors.IsNotFound(err) {
			return nil, false, err
		}

		openstackIdentity = &unikornv1.OpenstackIdentity{
			ObjectMeta: metav1.ObjectMeta{
				Namespace: identity.Namespace,
				Name:      identity.Name,
				Labels: map[string]string{
					constants.IdentityLabel: identity.Name,
				},
				Annotations: identity.Annotations,
			},
		}

		for k, v := range identity.Labels {
			openstackIdentity.Labels[k] = v
		}

		create = true
	}

	return openstackIdentity, create, nil
}

// CreateIdentity creates a new identity for cloud infrastructure.
func (p *Provider) CreateIdentity(ctx context.Context, identity *unikornv1.Identity) error {
	identityService, err := p.identity(ctx)
	if err != nil {
		return err
	}

	openstackIdentity, create, err := p.GetOrCreateOpenstackIdentity(ctx, identity)
	if err != nil {
		return err
	}

	record := func() {
		log := log.FromContext(ctx)

		if create {
			if err := p.client.Create(ctx, openstackIdentity); err != nil {
				log.Error(err, "failed to create openstack identity")
			}

			return
		}

		if err := p.client.Update(ctx, openstackIdentity); err != nil {
			log.Error(err, "failed to update openstack identity")
		}
	}

	defer record()

	// Every cluster has its own project to mitigate "nuances" in CAPO i.e. it's
	// totally broken when it comes to network aliasing.
	if err := p.provisionProject(ctx, identityService, openstackIdentity); err != nil {
		return err
	}

	// You MUST provision a new user, if we rotate a password, any application credentials
	// hanging off it will stop working, i.e. doing that to the unikorn management user
	// will be pretty catastrophic for all clusters in the region.
	if err := p.provisionUser(ctx, identityService, openstackIdentity); err != nil {
		return err
	}

	// Give the user only what permissions they need to provision a cluster and
	// manage it during its lifetime.
	if err := p.provisionProjectRoles(ctx, identityService, openstackIdentity); err != nil {
		return err
	}

	// Always use application credentials, they are scoped to a single project and
	// cannot be used to break from that jail.
	if err := p.provisionApplicationCredential(ctx, openstackIdentity); err != nil {
		return err
	}

	if err := p.createClientConfig(openstackIdentity); err != nil {
		return err
	}

	// Add in any optional configuration.
	if err := p.createIdentityServerGroup(ctx, openstackIdentity); err != nil {
		return err
	}

	return nil
}

// DeleteIdentity cleans up an identity for cloud infrastructure.
func (p *Provider) DeleteIdentity(ctx context.Context, identity *unikornv1.Identity) error {
	openstackIdentity, err := p.GetOpenstackIdentity(ctx, identity)
	if err != nil {
		if !kerrors.IsNotFound(err) {
			return err
		}

		return nil
	}

	if openstackIdentity.Spec.UserID != nil && openstackIdentity.Spec.ProjectID != nil {
		// Rescope to the user/project...
		providerClient := NewPasswordProvider(p.region.Spec.Openstack.Endpoint, *openstackIdentity.Spec.UserID, *openstackIdentity.Spec.Password, *openstackIdentity.Spec.ProjectID)

		computeService, err := NewComputeClient(ctx, providerClient, p.region.Spec.Openstack.Compute)
		if err != nil {
			return err
		}

		if openstackIdentity.Spec.ServerGroupID != nil {
			if err := computeService.DeleteServerGroup(ctx, *openstackIdentity.Spec.ServerGroupID); err != nil {
				return err
			}
		}
	}

	identityService, err := p.identity(ctx)
	if err != nil {
		return err
	}

	if openstackIdentity.Spec.UserID != nil {
		if err := identityService.DeleteUser(ctx, *openstackIdentity.Spec.UserID); err != nil {
			return err
		}
	}

	if openstackIdentity.Spec.ProjectID != nil {
		if err := identityService.DeleteProject(ctx, *openstackIdentity.Spec.ProjectID); err != nil {
			return err
		}
	}

	if err := p.client.Delete(ctx, openstackIdentity); err != nil {
		return err
	}

	return nil
}

// CreatePhysicalNetwork creates a physical network for an identity.
func (p *Provider) CreatePhysicalNetwork(ctx context.Context, identity *unikornv1.Identity, request *openapi.PhysicalNetworkWrite) (*unikornv1.PhysicalNetwork, error) {
	openstackIdentity, err := p.GetOpenstackIdentity(ctx, identity)
	if err != nil {
		return nil, err
	}

	networkService, err := p.network(ctx)
	if err != nil {
		return nil, err
	}

	vlanID, providerNetwork, err := networkService.CreateVLANProviderNetwork(ctx, "cluster-provider-network", *openstackIdentity.Spec.ProjectID)
	if err != nil {
		return nil, err
	}

	userinfo, err := authorization.UserinfoFromContext(ctx)
	if err != nil {
		return nil, err
	}

	objectMeta := conversion.NewObjectMetadata(&request.Metadata, p.region.Namespace, userinfo.Sub)
	objectMeta = objectMeta.WithOrganization(identity.Labels[coreconstants.OrganizationLabel])
	objectMeta = objectMeta.WithProject(identity.Labels[coreconstants.ProjectLabel])
	objectMeta = objectMeta.WithLabel(constants.RegionLabel, p.region.Name)
	objectMeta = objectMeta.WithLabel(constants.IdentityLabel, identity.Name)

	physicalNetwork := &unikornv1.PhysicalNetwork{
		ObjectMeta: objectMeta.Get(),
		Spec: unikornv1.PhysicalNetworkSpec{
			Tags: convertTagList(request.Spec.Tags),
			ProviderNetwork: &unikornv1.OpenstackProviderNetworkSpec{
				ID:     providerNetwork.ID,
				VlanID: vlanID,
			},
		},
	}

	if err := p.client.Create(ctx, physicalNetwork); err != nil {
		return nil, err
	}

	return physicalNetwork, nil
}

// ListExternalNetworks returns a list of external networks if the platform
// supports such a concept.
func (p *Provider) ListExternalNetworks(ctx context.Context) (providers.ExternalNetworks, error) {
	networkService, err := p.network(ctx)
	if err != nil {
		return nil, err
	}

	result, err := networkService.ExternalNetworks(ctx)
	if err != nil {
		return nil, err
	}

	out := make(providers.ExternalNetworks, len(result))

	for i, in := range result {
		out[i] = providers.ExternalNetwork{
			ID:   in.ID,
			Name: in.Name,
		}
	}

	return out, nil
}

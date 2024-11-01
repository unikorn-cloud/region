// Package openapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package openapi

import (
	externalRef0 "github.com/unikorn-cloud/core/pkg/openapi"
)

const (
	Oauth2AuthenticationScopes = "oauth2Authentication.Scopes"
)

// Defines values for GpuVendor.
const (
	AMD    GpuVendor = "AMD"
	NVIDIA GpuVendor = "NVIDIA"
)

// Defines values for ImageVirtualization.
const (
	Any         ImageVirtualization = "any"
	Baremetal   ImageVirtualization = "baremetal"
	Virtualized ImageVirtualization = "virtualized"
)

// Defines values for RegionType.
const (
	Openstack RegionType = "openstack"
)

// Defines values for SecurityGroupRuleReadSpecDirection.
const (
	SecurityGroupRuleReadSpecDirectionEgress  SecurityGroupRuleReadSpecDirection = "egress"
	SecurityGroupRuleReadSpecDirectionIngress SecurityGroupRuleReadSpecDirection = "ingress"
)

// Defines values for SecurityGroupRuleReadSpecProtocol.
const (
	SecurityGroupRuleReadSpecProtocolTcp SecurityGroupRuleReadSpecProtocol = "tcp"
	SecurityGroupRuleReadSpecProtocolUdp SecurityGroupRuleReadSpecProtocol = "udp"
)

// Defines values for SecurityGroupRuleWriteSpecDirection.
const (
	SecurityGroupRuleWriteSpecDirectionEgress  SecurityGroupRuleWriteSpecDirection = "egress"
	SecurityGroupRuleWriteSpecDirectionIngress SecurityGroupRuleWriteSpecDirection = "ingress"
)

// Defines values for SecurityGroupRuleWriteSpecProtocol.
const (
	SecurityGroupRuleWriteSpecProtocolTcp SecurityGroupRuleWriteSpecProtocol = "tcp"
	SecurityGroupRuleWriteSpecProtocolUdp SecurityGroupRuleWriteSpecProtocol = "udp"
)

// ExternalNetwork An Openstack external network.
type ExternalNetwork struct {
	// Id The resource ID.
	Id string `json:"id"`

	// Name The resource name.
	Name string `json:"name"`
}

// ExternalNetworks A list of openstack external networks.
type ExternalNetworks = []ExternalNetwork

// Flavor A flavor.
type Flavor struct {
	// Metadata This metadata is for resources that just exist, and don't require
	// any provisioning and health status, but benefit from a standarized
	// metadata format.
	Metadata externalRef0.StaticResourceMetadata `json:"metadata"`

	// Spec A flavor.
	Spec FlavorSpec `json:"spec"`
}

// FlavorQuota A flavor quota.
type FlavorQuota struct {
	// Count The number of the required flavor.
	Count int `json:"count"`

	// Id The flavor ID.
	Id string `json:"id"`
}

// FlavorQuotaList A list of flavor quotas.
type FlavorQuotaList = []FlavorQuota

// FlavorSpec A flavor.
type FlavorSpec struct {
	// Baremetal Whether the flavor is for a dedicated machine.
	Baremetal *bool `json:"baremetal,omitempty"`

	// CpuFamily A free form CPU family description e.g. model number, architecture.
	CpuFamily *string `json:"cpuFamily,omitempty"`

	// Cpus The number of CPUs.
	Cpus int `json:"cpus"`

	// Disk The amount of ephemeral disk in GB.
	Disk int `json:"disk"`

	// Gpu GPU specification.
	Gpu *GpuSpec `json:"gpu,omitempty"`

	// Memory The amount of memory in GiB.
	Memory int `json:"memory"`
}

// Flavors A list of flavors.
type Flavors = []Flavor

// GpuModel A GPU model number.
type GpuModel = string

// GpuModelList A list of GPU model numbers.
type GpuModelList = []GpuModel

// GpuSpec GPU specification.
type GpuSpec struct {
	// LogicalCount The logical number of GPUs available as seen in the OS.
	LogicalCount int `json:"logicalCount"`

	// Memory GPU memory in GiB.
	Memory int `json:"memory"`

	// Model A GPU model.
	Model string `json:"model"`

	// PhysicalCount The physical number of GPUs (cards) available.
	PhysicalCount int `json:"physicalCount"`

	// Vendor The GPU vendor.
	Vendor GpuVendor `json:"vendor"`
}

// GpuVendor The GPU vendor.
type GpuVendor string

// IdentitiesRead A list of provider specific identities.
type IdentitiesRead = []IdentityRead

// IdentityRead A provider specific identity.
type IdentityRead struct {
	Metadata externalRef0.ProjectScopedResourceReadMetadata `json:"metadata"`

	// Spec A provider specific identity, while the client can list regions to infer the
	// type, we don't requires this and return it with the response.  That can then
	// be used in turn to determine which provider specification to examine.
	Spec IdentitySpec `json:"spec"`
}

// IdentitySpec A provider specific identity, while the client can list regions to infer the
// type, we don't requires this and return it with the response.  That can then
// be used in turn to determine which provider specification to examine.
type IdentitySpec struct {
	// Openstack Everything an OpenStack client needs to function.
	Openstack *IdentitySpecOpenStack `json:"openstack,omitempty"`

	// RegionId The region an identity is provisioned in.
	RegionId string `json:"regionId"`

	// Tags A list of tags.
	Tags *TagList `json:"tags,omitempty"`

	// Type The region's provider type.
	Type RegionType `json:"type"`
}

// IdentitySpecOpenStack Everything an OpenStack client needs to function.
type IdentitySpecOpenStack struct {
	// Cloud The name of the cloud in the cloud config.
	Cloud *string `json:"cloud,omitempty"`

	// CloudConfig A base64 encoded cloud config file.
	CloudConfig *string `json:"cloudConfig,omitempty"`

	// ProjectId Project identifier allocated for the infrastructure.
	ProjectId *string `json:"projectId,omitempty"`

	// ServerGroupId Server group identifier allocated for the intrastructure.
	ServerGroupId *string `json:"serverGroupId,omitempty"`

	// SshKeyName Ephemeral SSH key generated for the identity.
	SshKeyName *string `json:"sshKeyName,omitempty"`

	// UserId User identitifer allocated for the infrastructure.
	UserId *string `json:"userId,omitempty"`
}

// IdentityWrite An identity request.
type IdentityWrite struct {
	// Metadata Resource metadata valid for all API resource reads and writes.
	Metadata externalRef0.ResourceWriteMetadata `json:"metadata"`

	// Spec Request parameters for creating an identity.
	Spec IdentityWriteSpec `json:"spec"`
}

// IdentityWriteSpec Request parameters for creating an identity.
type IdentityWriteSpec struct {
	// RegionId The region an identity is provisioned in.
	RegionId string `json:"regionId"`

	// Tags A list of tags.
	Tags *TagList `json:"tags,omitempty"`
}

// Image An image.
type Image struct {
	// Metadata This metadata is for resources that just exist, and don't require
	// any provisioning and health status, but benefit from a standarized
	// metadata format.
	Metadata externalRef0.StaticResourceMetadata `json:"metadata"`

	// Spec An image.
	Spec ImageSpec `json:"spec"`
}

// ImageGpu The GPU driver if installed.
type ImageGpu struct {
	// Driver The GPU driver version, this is vendor specific.
	Driver string `json:"driver"`

	// Models A list of GPU model numbers.
	Models *GpuModelList `json:"models,omitempty"`

	// Vendor The GPU vendor.
	Vendor GpuVendor `json:"vendor"`
}

// ImageSpec An image.
type ImageSpec struct {
	// Gpu The GPU driver if installed.
	Gpu *ImageGpu `json:"gpu,omitempty"`

	// SoftwareVersions Image preinstalled version version metadata.
	SoftwareVersions *SoftwareVersions `json:"softwareVersions,omitempty"`

	// Virtualization What type of machine the image is for.
	Virtualization ImageVirtualization `json:"virtualization"`
}

// ImageVirtualization What type of machine the image is for.
type ImageVirtualization string

// Images A list of images that are compatible with this platform.
type Images = []Image

// Ipv4Address An IPv4 address.
type Ipv4Address = string

// Ipv4AddressList A list of IPv4 addresses.
type Ipv4AddressList = []Ipv4Address

// KubernetesNameParameter A Kubernetes name. Must be a valid DNS containing only lower case characters, numbers or hyphens, start and end with a character or number, and be at most 63 characters in length.
type KubernetesNameParameter = string

// PhysicalNetworkRead A physical network.
type PhysicalNetworkRead struct {
	Metadata externalRef0.ProjectScopedResourceReadMetadata `json:"metadata"`

	// Spec A phyical network's specification.
	Spec PhysicalNetworkReadSpec `json:"spec"`
}

// PhysicalNetworkReadSpec A phyical network's specification.
type PhysicalNetworkReadSpec struct {
	// DnsNameservers A list of IPv4 addresses.
	DnsNameservers Ipv4AddressList `json:"dnsNameservers"`

	// Openstack An openstack physical network.
	Openstack *PhysicalNetworkSpecOpenstack `json:"openstack,omitempty"`

	// Prefix An IPv4 prefix for the network.
	Prefix string `json:"prefix"`

	// RegionId The region an identity is provisioned in.
	RegionId string `json:"regionId"`

	// Tags A list of tags.
	Tags *TagList `json:"tags,omitempty"`

	// Type The region's provider type.
	Type RegionType `json:"type"`
}

// PhysicalNetworkSpecOpenstack An openstack physical network.
type PhysicalNetworkSpecOpenstack struct {
	// NetworkId The openstack network ID.
	NetworkId *string `json:"networkId,omitempty"`

	// RouterId The openstack router ID.
	RouterId *string `json:"routerId,omitempty"`

	// SubnetId The openstack subnet ID.
	SubnetId *string `json:"subnetId,omitempty"`

	// VlanId The allocated VLAN ID.
	VlanId *int `json:"vlanId,omitempty"`
}

// PhysicalNetworkWrite A physical network request.
type PhysicalNetworkWrite struct {
	// Metadata Resource metadata valid for all API resource reads and writes.
	Metadata externalRef0.ResourceWriteMetadata `json:"metadata"`

	// Spec A phyical network's specification.
	Spec *PhysicalNetworkWriteSpec `json:"spec,omitempty"`
}

// PhysicalNetworkWriteSpec A phyical network's specification.
type PhysicalNetworkWriteSpec struct {
	// DnsNameservers A list of IPv4 addresses.
	DnsNameservers Ipv4AddressList `json:"dnsNameservers"`

	// Prefix An IPv4 prefix for the network.
	Prefix string `json:"prefix"`

	// Tags A list of tags.
	Tags *TagList `json:"tags,omitempty"`
}

// PhysicalNetworksRead A list of physical networks.
type PhysicalNetworksRead = []PhysicalNetworkRead

// QuotasSpec defines model for quotasSpec.
type QuotasSpec struct {
	// Flavors A list of flavor quotas.
	Flavors *FlavorQuotaList `json:"flavors,omitempty"`
}

// RegionFeatures A set of features the region may provide to clients.
type RegionFeatures struct {
	// PhysicalNetworks If set, this indicates that the region supports physical networks and
	// one should be provisioned for clusters to use.  The impliciation here is
	// the region supports base-metal machines, and these must be provisioned
	// on a physical VLAN etc.
	PhysicalNetworks bool `json:"physicalNetworks"`
}

// RegionRead A region.
type RegionRead struct {
	// Metadata Resource metadata valid for all reads.
	Metadata externalRef0.ResourceReadMetadata `json:"metadata"`

	// Spec Information about the region.
	Spec RegionSpec `json:"spec"`
}

// RegionSpec Information about the region.
type RegionSpec struct {
	// Features A set of features the region may provide to clients.
	Features RegionFeatures `json:"features"`

	// Type The region's provider type.
	Type RegionType `json:"type"`
}

// RegionType The region's provider type.
type RegionType string

// Regions A list of regions.
type Regions = []RegionRead

// SecurityGroupRead A security group.
type SecurityGroupRead struct {
	Metadata externalRef0.ProjectScopedResourceReadMetadata `json:"metadata"`

	// Spec A security group's specification.
	Spec SecurityGroupReadSpec `json:"spec"`
}

// SecurityGroupReadSpec A security group's specification.
type SecurityGroupReadSpec struct {
	// RegionId The region an identity is provisioned in.
	RegionId string `json:"regionId"`

	// Tags A list of tags.
	Tags *TagList `json:"tags,omitempty"`
}

// SecurityGroupRulePort The port definition to allow traffic.
type SecurityGroupRulePort struct {
	// Number The port to allow.
	Number *int `json:"number,omitempty"`

	// Range The port range to allow traffic.
	Range *SecurityGroupRulePortRange `json:"range,omitempty"`
}

// SecurityGroupRulePortRange The port range to allow traffic.
type SecurityGroupRulePortRange struct {
	// End The end of the port range.
	End int `json:"end"`

	// Start The start of the port range.
	Start int `json:"start"`
}

// SecurityGroupRuleRead A security group rule.
type SecurityGroupRuleRead struct {
	Metadata externalRef0.ProjectScopedResourceReadMetadata `json:"metadata"`

	// Spec A security group rule's specification.
	Spec SecurityGroupRuleReadSpec `json:"spec"`
}

// SecurityGroupRuleReadSpec A security group rule's specification.
type SecurityGroupRuleReadSpec struct {
	// Cidr An IPv4 address.
	Cidr Ipv4Address `json:"cidr"`

	// Direction The direction of the rule.
	Direction SecurityGroupRuleReadSpecDirection `json:"direction"`

	// Port The port definition to allow traffic.
	Port SecurityGroupRulePort `json:"port"`

	// Protocol The protocol to allow.
	Protocol SecurityGroupRuleReadSpecProtocol `json:"protocol"`
}

// SecurityGroupRuleReadSpecDirection The direction of the rule.
type SecurityGroupRuleReadSpecDirection string

// SecurityGroupRuleReadSpecProtocol The protocol to allow.
type SecurityGroupRuleReadSpecProtocol string

// SecurityGroupRuleWrite A security group rule request.
type SecurityGroupRuleWrite struct {
	// Metadata Resource metadata valid for all API resource reads and writes.
	Metadata externalRef0.ResourceWriteMetadata `json:"metadata"`

	// Spec A security group rule's specification.
	Spec SecurityGroupRuleWriteSpec `json:"spec"`
}

// SecurityGroupRuleWriteSpec A security group rule's specification.
type SecurityGroupRuleWriteSpec struct {
	// Cidr An IPv4 address.
	Cidr Ipv4Address `json:"cidr"`

	// Direction The direction of the rule.
	Direction SecurityGroupRuleWriteSpecDirection `json:"direction"`

	// Port The port definition to allow traffic.
	Port SecurityGroupRulePort `json:"port"`

	// Protocol The protocol to allow.
	Protocol SecurityGroupRuleWriteSpecProtocol `json:"protocol"`
}

// SecurityGroupRuleWriteSpecDirection The direction of the rule.
type SecurityGroupRuleWriteSpecDirection string

// SecurityGroupRuleWriteSpecProtocol The protocol to allow.
type SecurityGroupRuleWriteSpecProtocol string

// SecurityGroupRulesRead A list of security group rules.
type SecurityGroupRulesRead = []SecurityGroupRuleRead

// SecurityGroupWrite A security group request.
type SecurityGroupWrite struct {
	// Metadata Resource metadata valid for all API resource reads and writes.
	Metadata externalRef0.ResourceWriteMetadata `json:"metadata"`

	// Spec A security group's specification.
	Spec *SecurityGroupWriteSpec `json:"spec,omitempty"`
}

// SecurityGroupWriteSpec A security group's specification.
type SecurityGroupWriteSpec struct {
	// Tags A list of tags.
	Tags *TagList `json:"tags,omitempty"`
}

// SecurityGroupsRead A list of security groups.
type SecurityGroupsRead = []SecurityGroupRead

// SoftwareVersions Image preinstalled version version metadata.
type SoftwareVersions struct {
	// Kubernetes A semantic version.
	Kubernetes *externalRef0.Semver `json:"kubernetes,omitempty"`
}

// Tag An arbitrary tag name and value.
type Tag struct {
	// Name A unique tag name.
	Name string `json:"name"`

	// Value The value of the tag.
	Value string `json:"value"`
}

// TagList A list of tags.
type TagList = []Tag

// IdentityIDParameter A Kubernetes name. Must be a valid DNS containing only lower case characters, numbers or hyphens, start and end with a character or number, and be at most 63 characters in length.
type IdentityIDParameter = KubernetesNameParameter

// OrganizationIDParameter defines model for organizationIDParameter.
type OrganizationIDParameter = string

// PhysicalNetworkIDParameter A Kubernetes name. Must be a valid DNS containing only lower case characters, numbers or hyphens, start and end with a character or number, and be at most 63 characters in length.
type PhysicalNetworkIDParameter = KubernetesNameParameter

// ProjectIDParameter A Kubernetes name. Must be a valid DNS containing only lower case characters, numbers or hyphens, start and end with a character or number, and be at most 63 characters in length.
type ProjectIDParameter = KubernetesNameParameter

// RegionIDParameter A Kubernetes name. Must be a valid DNS containing only lower case characters, numbers or hyphens, start and end with a character or number, and be at most 63 characters in length.
type RegionIDParameter = KubernetesNameParameter

// RuleIDParameter A Kubernetes name. Must be a valid DNS containing only lower case characters, numbers or hyphens, start and end with a character or number, and be at most 63 characters in length.
type RuleIDParameter = KubernetesNameParameter

// SecurityGroupIDParameter A Kubernetes name. Must be a valid DNS containing only lower case characters, numbers or hyphens, start and end with a character or number, and be at most 63 characters in length.
type SecurityGroupIDParameter = KubernetesNameParameter

// ExternalNetworksResponse A list of openstack external networks.
type ExternalNetworksResponse = ExternalNetworks

// FlavorsResponse A list of flavors.
type FlavorsResponse = Flavors

// IdentitiesResponse A list of provider specific identities.
type IdentitiesResponse = IdentitiesRead

// IdentityResponse A provider specific identity.
type IdentityResponse = IdentityRead

// ImagesResponse A list of images that are compatible with this platform.
type ImagesResponse = Images

// PhysicalNetworkResponse A physical network.
type PhysicalNetworkResponse = PhysicalNetworkRead

// PhysicalNetworksResponse A list of physical networks.
type PhysicalNetworksResponse = PhysicalNetworksRead

// QuotasResponse defines model for quotasResponse.
type QuotasResponse = QuotasSpec

// RegionsResponse A list of regions.
type RegionsResponse = Regions

// SecurityGroupResponse A security group.
type SecurityGroupResponse = SecurityGroupRead

// SecurityGroupRuleResponse A security group rule.
type SecurityGroupRuleResponse = SecurityGroupRuleRead

// SecurityGroupRulesResponse A list of security group rules.
type SecurityGroupRulesResponse = SecurityGroupRulesRead

// SecurityGroupsResponse A list of security groups.
type SecurityGroupsResponse = SecurityGroupsRead

// IdentityRequest An identity request.
type IdentityRequest = IdentityWrite

// PhysicalNetworkRequest A physical network request.
type PhysicalNetworkRequest = PhysicalNetworkWrite

// QuotasRequest defines model for quotasRequest.
type QuotasRequest = QuotasSpec

// SecurityGroupRequest A security group request.
type SecurityGroupRequest = SecurityGroupWrite

// SecurityGroupRuleRequest A security group rule request.
type SecurityGroupRuleRequest = SecurityGroupRuleWrite

// PostApiV1OrganizationsOrganizationIDProjectsProjectIDIdentitiesJSONRequestBody defines body for PostApiV1OrganizationsOrganizationIDProjectsProjectIDIdentities for application/json ContentType.
type PostApiV1OrganizationsOrganizationIDProjectsProjectIDIdentitiesJSONRequestBody = IdentityWrite

// PostApiV1OrganizationsOrganizationIDProjectsProjectIDIdentitiesIdentityIDPhysicalnetworksJSONRequestBody defines body for PostApiV1OrganizationsOrganizationIDProjectsProjectIDIdentitiesIdentityIDPhysicalnetworks for application/json ContentType.
type PostApiV1OrganizationsOrganizationIDProjectsProjectIDIdentitiesIdentityIDPhysicalnetworksJSONRequestBody = PhysicalNetworkWrite

// PutApiV1OrganizationsOrganizationIDProjectsProjectIDIdentitiesIdentityIDQuotasJSONRequestBody defines body for PutApiV1OrganizationsOrganizationIDProjectsProjectIDIdentitiesIdentityIDQuotas for application/json ContentType.
type PutApiV1OrganizationsOrganizationIDProjectsProjectIDIdentitiesIdentityIDQuotasJSONRequestBody = QuotasSpec

// PostApiV1OrganizationsOrganizationIDProjectsProjectIDIdentitiesIdentityIDSecuritygroupsJSONRequestBody defines body for PostApiV1OrganizationsOrganizationIDProjectsProjectIDIdentitiesIdentityIDSecuritygroups for application/json ContentType.
type PostApiV1OrganizationsOrganizationIDProjectsProjectIDIdentitiesIdentityIDSecuritygroupsJSONRequestBody = SecurityGroupWrite

// PutApiV1OrganizationsOrganizationIDProjectsProjectIDIdentitiesIdentityIDSecuritygroupsSecurityGroupIDJSONRequestBody defines body for PutApiV1OrganizationsOrganizationIDProjectsProjectIDIdentitiesIdentityIDSecuritygroupsSecurityGroupID for application/json ContentType.
type PutApiV1OrganizationsOrganizationIDProjectsProjectIDIdentitiesIdentityIDSecuritygroupsSecurityGroupIDJSONRequestBody = SecurityGroupWrite

// PostApiV1OrganizationsOrganizationIDProjectsProjectIDIdentitiesIdentityIDSecuritygroupsSecurityGroupIDRulesJSONRequestBody defines body for PostApiV1OrganizationsOrganizationIDProjectsProjectIDIdentitiesIdentityIDSecuritygroupsSecurityGroupIDRules for application/json ContentType.
type PostApiV1OrganizationsOrganizationIDProjectsProjectIDIdentitiesIdentityIDSecuritygroupsSecurityGroupIDRulesJSONRequestBody = SecurityGroupRuleWrite

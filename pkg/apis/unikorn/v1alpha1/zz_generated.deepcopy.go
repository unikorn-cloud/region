//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	unikornv1alpha1 "github.com/unikorn-cloud/core/pkg/apis/unikorn/v1alpha1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CPUSpec) DeepCopyInto(out *CPUSpec) {
	*out = *in
	if in.Count != nil {
		in, out := &in.Count, &out.Count
		*out = new(int)
		**out = **in
	}
	if in.Family != nil {
		in, out := &in.Family, &out.Family
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CPUSpec.
func (in *CPUSpec) DeepCopy() *CPUSpec {
	if in == nil {
		return nil
	}
	out := new(CPUSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalNetworks) DeepCopyInto(out *ExternalNetworks) {
	*out = *in
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = new(NetworkSelector)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalNetworks.
func (in *ExternalNetworks) DeepCopy() *ExternalNetworks {
	if in == nil {
		return nil
	}
	out := new(ExternalNetworks)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlavorMetadata) DeepCopyInto(out *FlavorMetadata) {
	*out = *in
	if in.CPU != nil {
		in, out := &in.CPU, &out.CPU
		*out = new(CPUSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Memory != nil {
		in, out := &in.Memory, &out.Memory
		x := (*in).DeepCopy()
		*out = &x
	}
	if in.GPU != nil {
		in, out := &in.GPU, &out.GPU
		*out = new(GPUSpec)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlavorMetadata.
func (in *FlavorMetadata) DeepCopy() *FlavorMetadata {
	if in == nil {
		return nil
	}
	out := new(FlavorMetadata)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlavorSelector) DeepCopyInto(out *FlavorSelector) {
	*out = *in
	if in.IDs != nil {
		in, out := &in.IDs, &out.IDs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlavorSelector.
func (in *FlavorSelector) DeepCopy() *FlavorSelector {
	if in == nil {
		return nil
	}
	out := new(FlavorSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GPUSpec) DeepCopyInto(out *GPUSpec) {
	*out = *in
	if in.Memory != nil {
		in, out := &in.Memory, &out.Memory
		x := (*in).DeepCopy()
		*out = &x
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GPUSpec.
func (in *GPUSpec) DeepCopy() *GPUSpec {
	if in == nil {
		return nil
	}
	out := new(GPUSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Identity) DeepCopyInto(out *Identity) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Identity.
func (in *Identity) DeepCopy() *Identity {
	if in == nil {
		return nil
	}
	out := new(Identity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Identity) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IdentityList) DeepCopyInto(out *IdentityList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Identity, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IdentityList.
func (in *IdentityList) DeepCopy() *IdentityList {
	if in == nil {
		return nil
	}
	out := new(IdentityList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IdentityList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IdentitySpec) DeepCopyInto(out *IdentitySpec) {
	*out = *in
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(TagList, len(*in))
		copy(*out, *in)
	}
	if in.OpenStack != nil {
		in, out := &in.OpenStack, &out.OpenStack
		*out = new(IdentitySpecOpenStack)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IdentitySpec.
func (in *IdentitySpec) DeepCopy() *IdentitySpec {
	if in == nil {
		return nil
	}
	out := new(IdentitySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IdentitySpecOpenStack) DeepCopyInto(out *IdentitySpecOpenStack) {
	*out = *in
	if in.CloudConfig != nil {
		in, out := &in.CloudConfig, &out.CloudConfig
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	if in.ServerGroupID != nil {
		in, out := &in.ServerGroupID, &out.ServerGroupID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IdentitySpecOpenStack.
func (in *IdentitySpecOpenStack) DeepCopy() *IdentitySpecOpenStack {
	if in == nil {
		return nil
	}
	out := new(IdentitySpecOpenStack)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IdentityStatus) DeepCopyInto(out *IdentityStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IdentityStatus.
func (in *IdentityStatus) DeepCopy() *IdentityStatus {
	if in == nil {
		return nil
	}
	out := new(IdentityStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageSelector) DeepCopyInto(out *ImageSelector) {
	*out = *in
	if in.Properties != nil {
		in, out := &in.Properties, &out.Properties
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SigningKey != nil {
		in, out := &in.SigningKey, &out.SigningKey
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageSelector.
func (in *ImageSelector) DeepCopy() *ImageSelector {
	if in == nil {
		return nil
	}
	out := new(ImageSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespacedObject) DeepCopyInto(out *NamespacedObject) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespacedObject.
func (in *NamespacedObject) DeepCopy() *NamespacedObject {
	if in == nil {
		return nil
	}
	out := new(NamespacedObject)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkSelector) DeepCopyInto(out *NetworkSelector) {
	*out = *in
	if in.IDs != nil {
		in, out := &in.IDs, &out.IDs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkSelector.
func (in *NetworkSelector) DeepCopy() *NetworkSelector {
	if in == nil {
		return nil
	}
	out := new(NetworkSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenstackFlavorsSpec) DeepCopyInto(out *OpenstackFlavorsSpec) {
	*out = *in
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = new(FlavorSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.Metadata != nil {
		in, out := &in.Metadata, &out.Metadata
		*out = make([]FlavorMetadata, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenstackFlavorsSpec.
func (in *OpenstackFlavorsSpec) DeepCopy() *OpenstackFlavorsSpec {
	if in == nil {
		return nil
	}
	out := new(OpenstackFlavorsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenstackProviderNetworkSpec) DeepCopyInto(out *OpenstackProviderNetworkSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenstackProviderNetworkSpec.
func (in *OpenstackProviderNetworkSpec) DeepCopy() *OpenstackProviderNetworkSpec {
	if in == nil {
		return nil
	}
	out := new(OpenstackProviderNetworkSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PhysicalNetwork) DeepCopyInto(out *PhysicalNetwork) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PhysicalNetwork.
func (in *PhysicalNetwork) DeepCopy() *PhysicalNetwork {
	if in == nil {
		return nil
	}
	out := new(PhysicalNetwork)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PhysicalNetwork) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PhysicalNetworkList) DeepCopyInto(out *PhysicalNetworkList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PhysicalNetwork, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PhysicalNetworkList.
func (in *PhysicalNetworkList) DeepCopy() *PhysicalNetworkList {
	if in == nil {
		return nil
	}
	out := new(PhysicalNetworkList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PhysicalNetworkList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PhysicalNetworkSpec) DeepCopyInto(out *PhysicalNetworkSpec) {
	*out = *in
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(TagList, len(*in))
		copy(*out, *in)
	}
	if in.ProviderNetwork != nil {
		in, out := &in.ProviderNetwork, &out.ProviderNetwork
		*out = new(OpenstackProviderNetworkSpec)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PhysicalNetworkSpec.
func (in *PhysicalNetworkSpec) DeepCopy() *PhysicalNetworkSpec {
	if in == nil {
		return nil
	}
	out := new(PhysicalNetworkSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PhysicalNetworkStatus) DeepCopyInto(out *PhysicalNetworkStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PhysicalNetworkStatus.
func (in *PhysicalNetworkStatus) DeepCopy() *PhysicalNetworkStatus {
	if in == nil {
		return nil
	}
	out := new(PhysicalNetworkStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProviderNetworks) DeepCopyInto(out *ProviderNetworks) {
	*out = *in
	if in.PhysicalNetwork != nil {
		in, out := &in.PhysicalNetwork, &out.PhysicalNetwork
		*out = new(string)
		**out = **in
	}
	if in.VLAN != nil {
		in, out := &in.VLAN, &out.VLAN
		*out = new(VLANSpec)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProviderNetworks.
func (in *ProviderNetworks) DeepCopy() *ProviderNetworks {
	if in == nil {
		return nil
	}
	out := new(ProviderNetworks)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Region) DeepCopyInto(out *Region) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Region.
func (in *Region) DeepCopy() *Region {
	if in == nil {
		return nil
	}
	out := new(Region)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Region) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegionList) DeepCopyInto(out *RegionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Region, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegionList.
func (in *RegionList) DeepCopy() *RegionList {
	if in == nil {
		return nil
	}
	out := new(RegionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RegionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegionOpenstackComputeSpec) DeepCopyInto(out *RegionOpenstackComputeSpec) {
	*out = *in
	if in.ServerGroupPolicy != nil {
		in, out := &in.ServerGroupPolicy, &out.ServerGroupPolicy
		*out = new(string)
		**out = **in
	}
	if in.Flavors != nil {
		in, out := &in.Flavors, &out.Flavors
		*out = new(OpenstackFlavorsSpec)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegionOpenstackComputeSpec.
func (in *RegionOpenstackComputeSpec) DeepCopy() *RegionOpenstackComputeSpec {
	if in == nil {
		return nil
	}
	out := new(RegionOpenstackComputeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegionOpenstackIdentitySpec) DeepCopyInto(out *RegionOpenstackIdentitySpec) {
	*out = *in
	if in.ClusterRoles != nil {
		in, out := &in.ClusterRoles, &out.ClusterRoles
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegionOpenstackIdentitySpec.
func (in *RegionOpenstackIdentitySpec) DeepCopy() *RegionOpenstackIdentitySpec {
	if in == nil {
		return nil
	}
	out := new(RegionOpenstackIdentitySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegionOpenstackImageSpec) DeepCopyInto(out *RegionOpenstackImageSpec) {
	*out = *in
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = new(ImageSelector)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegionOpenstackImageSpec.
func (in *RegionOpenstackImageSpec) DeepCopy() *RegionOpenstackImageSpec {
	if in == nil {
		return nil
	}
	out := new(RegionOpenstackImageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegionOpenstackNetworkSpec) DeepCopyInto(out *RegionOpenstackNetworkSpec) {
	*out = *in
	if in.ExternalNetworks != nil {
		in, out := &in.ExternalNetworks, &out.ExternalNetworks
		*out = new(ExternalNetworks)
		(*in).DeepCopyInto(*out)
	}
	if in.ProviderNetworks != nil {
		in, out := &in.ProviderNetworks, &out.ProviderNetworks
		*out = new(ProviderNetworks)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegionOpenstackNetworkSpec.
func (in *RegionOpenstackNetworkSpec) DeepCopy() *RegionOpenstackNetworkSpec {
	if in == nil {
		return nil
	}
	out := new(RegionOpenstackNetworkSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegionOpenstackSpec) DeepCopyInto(out *RegionOpenstackSpec) {
	*out = *in
	if in.ServiceAccountSecret != nil {
		in, out := &in.ServiceAccountSecret, &out.ServiceAccountSecret
		*out = new(NamespacedObject)
		**out = **in
	}
	if in.Identity != nil {
		in, out := &in.Identity, &out.Identity
		*out = new(RegionOpenstackIdentitySpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Compute != nil {
		in, out := &in.Compute, &out.Compute
		*out = new(RegionOpenstackComputeSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Image != nil {
		in, out := &in.Image, &out.Image
		*out = new(RegionOpenstackImageSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Network != nil {
		in, out := &in.Network, &out.Network
		*out = new(RegionOpenstackNetworkSpec)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegionOpenstackSpec.
func (in *RegionOpenstackSpec) DeepCopy() *RegionOpenstackSpec {
	if in == nil {
		return nil
	}
	out := new(RegionOpenstackSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegionSpec) DeepCopyInto(out *RegionSpec) {
	*out = *in
	if in.Openstack != nil {
		in, out := &in.Openstack, &out.Openstack
		*out = new(RegionOpenstackSpec)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegionSpec.
func (in *RegionSpec) DeepCopy() *RegionSpec {
	if in == nil {
		return nil
	}
	out := new(RegionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegionStatus) DeepCopyInto(out *RegionStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]unikornv1alpha1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegionStatus.
func (in *RegionStatus) DeepCopy() *RegionStatus {
	if in == nil {
		return nil
	}
	out := new(RegionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Tag) DeepCopyInto(out *Tag) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Tag.
func (in *Tag) DeepCopy() *Tag {
	if in == nil {
		return nil
	}
	out := new(Tag)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in TagList) DeepCopyInto(out *TagList) {
	{
		in := &in
		*out = make(TagList, len(*in))
		copy(*out, *in)
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TagList.
func (in TagList) DeepCopy() TagList {
	if in == nil {
		return nil
	}
	out := new(TagList)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VLANSegment) DeepCopyInto(out *VLANSegment) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VLANSegment.
func (in *VLANSegment) DeepCopy() *VLANSegment {
	if in == nil {
		return nil
	}
	out := new(VLANSegment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VLANSpec) DeepCopyInto(out *VLANSpec) {
	*out = *in
	if in.Segments != nil {
		in, out := &in.Segments, &out.Segments
		*out = make([]VLANSegment, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VLANSpec.
func (in *VLANSpec) DeepCopy() *VLANSpec {
	if in == nil {
		return nil
	}
	out := new(VLANSpec)
	in.DeepCopyInto(out)
	return out
}

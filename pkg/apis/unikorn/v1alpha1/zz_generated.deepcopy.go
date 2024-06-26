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
	if in.OpenStack != nil {
		in, out := &in.OpenStack, &out.OpenStack
		*out = new(IdentitySpecOpenStack)
		**out = **in
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
func (in *OpenstackFlavorExclude) DeepCopyInto(out *OpenstackFlavorExclude) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenstackFlavorExclude.
func (in *OpenstackFlavorExclude) DeepCopy() *OpenstackFlavorExclude {
	if in == nil {
		return nil
	}
	out := new(OpenstackFlavorExclude)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenstackFlavorInclude) DeepCopyInto(out *OpenstackFlavorInclude) {
	*out = *in
	if in.CPU != nil {
		in, out := &in.CPU, &out.CPU
		*out = new(CPUSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.GPU != nil {
		in, out := &in.GPU, &out.GPU
		*out = new(GPUSpec)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenstackFlavorInclude.
func (in *OpenstackFlavorInclude) DeepCopy() *OpenstackFlavorInclude {
	if in == nil {
		return nil
	}
	out := new(OpenstackFlavorInclude)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenstackFlavorsSpec) DeepCopyInto(out *OpenstackFlavorsSpec) {
	*out = *in
	if in.Include != nil {
		in, out := &in.Include, &out.Include
		*out = make([]OpenstackFlavorInclude, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Exclude != nil {
		in, out := &in.Exclude, &out.Exclude
		*out = make([]OpenstackFlavorExclude, len(*in))
		copy(*out, *in)
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
	if in.PropertiesInclude != nil {
		in, out := &in.PropertiesInclude, &out.PropertiesInclude
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

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

package v1alpha1

import (
	unikornv1core "github.com/unikorn-cloud/core/pkg/apis/unikorn/v1alpha1"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"
)

// Paused implements the ReconcilePauser interface.
func (c *Identity) Paused() bool {
	return c.Spec.Pause
}

// StatusConditionRead scans the status conditions for an existing condition whose type
// matches.
func (c *Identity) StatusConditionRead(t unikornv1core.ConditionType) (*unikornv1core.Condition, error) {
	return unikornv1core.GetCondition(c.Status.Conditions, t)
}

// StatusConditionWrite either adds or updates a condition in the cluster manager status.
// If the condition, status and message match an existing condition the update is
// ignored.
func (c *Identity) StatusConditionWrite(t unikornv1core.ConditionType, status corev1.ConditionStatus, reason unikornv1core.ConditionReason, message string) {
	unikornv1core.UpdateCondition(&c.Status.Conditions, t, status, reason, message)
}

// ResourceLabels generates a set of labels to uniquely identify the resource
// if it were to be placed in a single global namespace.
func (c *Identity) ResourceLabels() (labels.Set, error) {
	//nolint:nilnil
	return nil, nil
}

// Paused implements the ReconcilePauser interface.
func (c *Network) Paused() bool {
	return c.Spec.Pause
}

// StatusConditionRead scans the status conditions for an existing condition whose type
// matches.
func (c *Network) StatusConditionRead(t unikornv1core.ConditionType) (*unikornv1core.Condition, error) {
	return unikornv1core.GetCondition(c.Status.Conditions, t)
}

// StatusConditionWrite either adds or updates a condition in the cluster manager status.
// If the condition, status and message match an existing condition the update is
// ignored.
func (c *Network) StatusConditionWrite(t unikornv1core.ConditionType, status corev1.ConditionStatus, reason unikornv1core.ConditionReason, message string) {
	unikornv1core.UpdateCondition(&c.Status.Conditions, t, status, reason, message)
}

// ResourceLabels generates a set of labels to uniquely identify the resource
// if it were to be placed in a single global namespace.
func (c *Network) ResourceLabels() (labels.Set, error) {
	//nolint:nilnil
	return nil, nil
}

// Paused implements the ReconcilePauser interface.
func (c *SecurityGroup) Paused() bool {
	return c.Spec.Pause
}

// StatusConditionRead scans the status conditions for an existing condition whose type
// matches.
func (c *SecurityGroup) StatusConditionRead(t unikornv1core.ConditionType) (*unikornv1core.Condition, error) {
	return unikornv1core.GetCondition(c.Status.Conditions, t)
}

// StatusConditionWrite either adds or updates a condition in the cluster manager status.
// If the condition, status and message match an existing condition the update is
// ignored.
func (c *SecurityGroup) StatusConditionWrite(t unikornv1core.ConditionType, status corev1.ConditionStatus, reason unikornv1core.ConditionReason, message string) {
	unikornv1core.UpdateCondition(&c.Status.Conditions, t, status, reason, message)
}

// ResourceLabels generates a set of labels to uniquely identify the resource
// if it were to be placed in a single global namespace.
func (c *SecurityGroup) ResourceLabels() (labels.Set, error) {
	//nolint:nilnil
	return nil, nil
}

// Paused implements the ReconcilePauser interface.
func (c *SecurityGroupRule) Paused() bool {
	return c.Spec.Pause
}

// StatusConditionRead scans the status conditions for an existing condition whose type
// matches.
func (c *SecurityGroupRule) StatusConditionRead(t unikornv1core.ConditionType) (*unikornv1core.Condition, error) {
	return unikornv1core.GetCondition(c.Status.Conditions, t)
}

// StatusConditionWrite either adds or updates a condition in the cluster manager status.
// If the condition, status and message match an existing condition the update is
// ignored.
func (c *SecurityGroupRule) StatusConditionWrite(t unikornv1core.ConditionType, status corev1.ConditionStatus, reason unikornv1core.ConditionReason, message string) {
	unikornv1core.UpdateCondition(&c.Status.Conditions, t, status, reason, message)
}

// ResourceLabels generates a set of labels to uniquely identify the resource
// if it were to be placed in a single global namespace.
func (c *SecurityGroupRule) ResourceLabels() (labels.Set, error) {
	//nolint:nilnil
	return nil, nil
}

// Paused implements the ReconcilePauser interface.
func (c *Server) Paused() bool {
	return c.Spec.Pause
}

// StatusConditionRead scans the status conditions for an existing condition whose type
// matches.
func (c *Server) StatusConditionRead(t unikornv1core.ConditionType) (*unikornv1core.Condition, error) {
	return unikornv1core.GetCondition(c.Status.Conditions, t)
}

// StatusConditionWrite either adds or updates a condition in the cluster manager status.
// If the condition, status and message match an existing condition the update is
// ignored.
func (c *Server) StatusConditionWrite(t unikornv1core.ConditionType, status corev1.ConditionStatus, reason unikornv1core.ConditionReason, message string) {
	unikornv1core.UpdateCondition(&c.Status.Conditions, t, status, reason, message)
}

// ResourceLabels generates a set of labels to uniquely identify the resource
// if it were to be placed in a single global namespace.
func (c *Server) ResourceLabels() (labels.Set, error) {
	//nolint:nilnil
	return nil, nil
}

func (s *RegionOpenstackNetworkSpec) UseProviderNetworks() bool {
	return s != nil && s.ProviderNetworks != nil && s.ProviderNetworks.Network != nil
}

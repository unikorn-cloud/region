/*
Copyright 2022-2024 EscherCloud.
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

//nolint:revive
package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/unikorn-cloud/core/pkg/server/errors"
	"github.com/unikorn-cloud/core/pkg/server/util"
	identityclient "github.com/unikorn-cloud/identity/pkg/client"
	identityapi "github.com/unikorn-cloud/identity/pkg/openapi"
	"github.com/unikorn-cloud/identity/pkg/rbac"
	"github.com/unikorn-cloud/region/pkg/handler/identity"
	"github.com/unikorn-cloud/region/pkg/handler/network"
	"github.com/unikorn-cloud/region/pkg/handler/region"
	"github.com/unikorn-cloud/region/pkg/handler/securitygroup"
	"github.com/unikorn-cloud/region/pkg/handler/securitygrouprule"
	"github.com/unikorn-cloud/region/pkg/handler/server"
	"github.com/unikorn-cloud/region/pkg/openapi"

	"sigs.k8s.io/controller-runtime/pkg/client"
)

type Handler struct {
	// client gives cached access to Kubernetes.
	client client.Client

	// namespace is the namespace we are running in.
	namespace string

	// options allows behaviour to be defined on the CLI.
	options *Options

	// identity is an identity client for RBAC access.
	identity *identityclient.Client
}

func New(client client.Client, namespace string, options *Options, identity *identityclient.Client) (*Handler, error) {
	h := &Handler{
		client:    client,
		namespace: namespace,
		options:   options,
		identity:  identity,
	}

	return h, nil
}

func (h *Handler) setCacheable(w http.ResponseWriter) {
	w.Header().Add("Cache-Control", fmt.Sprintf("max-age=%d", h.options.CacheMaxAge/time.Second))
	w.Header().Add("Cache-Control", "private")
}

func (h *Handler) setUncacheable(w http.ResponseWriter) {
	w.Header().Add("Cache-Control", "no-cache")
}

func (h *Handler) GetApiV1OrganizationsOrganizationIDRegions(w http.ResponseWriter, r *http.Request, organizationID openapi.OrganizationIDParameter) {
	if err := rbac.AllowOrganizationScope(r.Context(), "region:regions", identityapi.Read, organizationID); err != nil {
		errors.HandleError(w, r, err)
		return
	}

	result, err := region.NewClient(h.client, h.namespace).List(r.Context())
	if err != nil {
		errors.HandleError(w, r, err)
		return
	}

	h.setUncacheable(w)
	util.WriteJSONResponse(w, r, http.StatusOK, result)
}

func (h *Handler) GetApiV1OrganizationsOrganizationIDRegionsRegionIDDetail(w http.ResponseWriter, r *http.Request, organizationID openapi.OrganizationIDParameter, regionID openapi.RegionIDParameter) {
	if err := rbac.AllowOrganizationScope(r.Context(), "region:regions/detail", identityapi.Read, organizationID); err != nil {
		errors.HandleError(w, r, err)
		return
	}

	result, err := region.NewClient(h.client, h.namespace).GetDetail(r.Context(), regionID)
	if err != nil {
		errors.HandleError(w, r, err)
		return
	}

	h.setUncacheable(w)
	util.WriteJSONResponse(w, r, http.StatusOK, result)
}

func (h *Handler) GetApiV1OrganizationsOrganizationIDRegionsRegionIDExternalnetworks(w http.ResponseWriter, r *http.Request, organizationID openapi.OrganizationIDParameter, regionID openapi.RegionIDParameter) {
	if err := rbac.AllowOrganizationScope(r.Context(), "region:externalnetworks", identityapi.Read, organizationID); err != nil {
		errors.HandleError(w, r, err)
		return
	}

	result, err := region.NewClient(h.client, h.namespace).ListExternalNetworks(r.Context(), regionID)
	if err != nil {
		errors.HandleError(w, r, err)
		return
	}

	h.setCacheable(w)
	util.WriteJSONResponse(w, r, http.StatusOK, result)
}

func (h *Handler) GetApiV1OrganizationsOrganizationIDRegionsRegionIDFlavors(w http.ResponseWriter, r *http.Request, organizationID openapi.OrganizationIDParameter, regionID openapi.RegionIDParameter) {
	if err := rbac.AllowOrganizationScope(r.Context(), "region:flavors", identityapi.Read, organizationID); err != nil {
		errors.HandleError(w, r, err)
		return
	}

	result, err := region.NewClient(h.client, h.namespace).ListFlavors(r.Context(), organizationID, regionID)
	if err != nil {
		errors.HandleError(w, r, err)
		return
	}

	h.setCacheable(w)
	util.WriteJSONResponse(w, r, http.StatusOK, result)
}

func (h *Handler) GetApiV1OrganizationsOrganizationIDRegionsRegionIDImages(w http.ResponseWriter, r *http.Request, organizationID openapi.OrganizationIDParameter, regionID openapi.RegionIDParameter) {
	if err := rbac.AllowOrganizationScope(r.Context(), "region:images", identityapi.Read, organizationID); err != nil {
		errors.HandleError(w, r, err)
		return
	}

	result, err := region.NewClient(h.client, h.namespace).ListImages(r.Context(), organizationID, regionID)
	if err != nil {
		errors.HandleError(w, r, err)
		return
	}

	h.setCacheable(w)
	util.WriteJSONResponse(w, r, http.StatusOK, result)
}

func (h *Handler) GetApiV1OrganizationsOrganizationIDIdentities(w http.ResponseWriter, r *http.Request, organizationID openapi.OrganizationIDParameter) {
	if err := rbac.AllowOrganizationScope(r.Context(), "region:identities", identityapi.Read, organizationID); err != nil {
		errors.HandleError(w, r, err)
		return
	}

	result, err := identity.New(h.client, h.namespace).List(r.Context(), organizationID)
	if err != nil {
		errors.HandleError(w, r, err)
		return
	}

	util.WriteJSONResponse(w, r, http.StatusOK, result)
}

func (h *Handler) PostApiV1OrganizationsOrganizationIDProjectsProjectIDIdentities(w http.ResponseWriter, r *http.Request, organizationID openapi.OrganizationIDParameter, projectID openapi.ProjectIDParameter) {
	if err := rbac.AllowProjectScope(r.Context(), "region:identities", identityapi.Create, organizationID, projectID); err != nil {
		errors.HandleError(w, r, err)
		return
	}

	request := &openapi.IdentityWrite{}

	if err := util.ReadJSONBody(r, request); err != nil {
		errors.HandleError(w, r, err)
		return
	}

	result, err := identity.New(h.client, h.namespace).Create(r.Context(), organizationID, projectID, request)
	if err != nil {
		errors.HandleError(w, r, err)
		return
	}

	util.WriteJSONResponse(w, r, http.StatusCreated, result)
}

func (h *Handler) GetApiV1OrganizationsOrganizationIDProjectsProjectIDIdentitiesIdentityID(w http.ResponseWriter, r *http.Request, organizationID openapi.OrganizationIDParameter, projectID openapi.ProjectIDParameter, identityID openapi.IdentityIDParameter) {
	if err := rbac.AllowProjectScope(r.Context(), "region:identities", identityapi.Read, organizationID, projectID); err != nil {
		errors.HandleError(w, r, err)
		return
	}

	result, err := identity.New(h.client, h.namespace).Get(r.Context(), organizationID, projectID, identityID)
	if err != nil {
		errors.HandleError(w, r, err)
		return
	}

	util.WriteJSONResponse(w, r, http.StatusOK, result)
}

func (h *Handler) DeleteApiV1OrganizationsOrganizationIDProjectsProjectIDIdentitiesIdentityID(w http.ResponseWriter, r *http.Request, organizationID openapi.OrganizationIDParameter, projectID openapi.ProjectIDParameter, identityID openapi.IdentityIDParameter) {
	if err := rbac.AllowProjectScope(r.Context(), "region:identities", identityapi.Delete, organizationID, projectID); err != nil {
		errors.HandleError(w, r, err)
		return
	}

	if err := identity.New(h.client, h.namespace).Delete(r.Context(), organizationID, projectID, identityID); err != nil {
		errors.HandleError(w, r, err)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}

func (h *Handler) GetApiV1OrganizationsOrganizationIDNetworks(w http.ResponseWriter, r *http.Request, organizationID openapi.OrganizationIDParameter) {
	if err := rbac.AllowOrganizationScope(r.Context(), "region:networks", identityapi.Read, organizationID); err != nil {
		errors.HandleError(w, r, err)
		return
	}

	result, err := network.New(h.client, h.namespace).List(r.Context(), organizationID)
	if err != nil {
		errors.HandleError(w, r, err)
		return
	}

	util.WriteJSONResponse(w, r, http.StatusOK, result)
}

func (h *Handler) PostApiV1OrganizationsOrganizationIDProjectsProjectIDIdentitiesIdentityIDNetworks(w http.ResponseWriter, r *http.Request, organizationID openapi.OrganizationIDParameter, projectID openapi.ProjectIDParameter, identityID openapi.IdentityIDParameter) {
	if err := rbac.AllowProjectScope(r.Context(), "region:networks", identityapi.Create, organizationID, projectID); err != nil {
		errors.HandleError(w, r, err)
		return
	}

	request := &openapi.NetworkWrite{}

	if err := util.ReadJSONBody(r, request); err != nil {
		errors.HandleError(w, r, err)
		return
	}

	result, err := network.New(h.client, h.namespace).Create(r.Context(), organizationID, projectID, identityID, request)
	if err != nil {
		errors.HandleError(w, r, err)
		return
	}

	util.WriteJSONResponse(w, r, http.StatusCreated, result)
}

func (h *Handler) GetApiV1OrganizationsOrganizationIDProjectsProjectIDIdentitiesIdentityIDNetworksNetworkID(w http.ResponseWriter, r *http.Request, organizationID openapi.OrganizationIDParameter, projectID openapi.ProjectIDParameter, identityID openapi.IdentityIDParameter, networkID openapi.NetworkIDParameter) {
	if err := rbac.AllowProjectScope(r.Context(), "region:networks", identityapi.Read, organizationID, projectID); err != nil {
		errors.HandleError(w, r, err)
		return
	}

	result, err := network.New(h.client, h.namespace).Get(r.Context(), organizationID, projectID, networkID)
	if err != nil {
		errors.HandleError(w, r, err)
		return
	}

	util.WriteJSONResponse(w, r, http.StatusOK, result)
}

func (h *Handler) DeleteApiV1OrganizationsOrganizationIDProjectsProjectIDIdentitiesIdentityIDNetworksNetworkID(w http.ResponseWriter, r *http.Request, organizationID openapi.OrganizationIDParameter, projectID openapi.ProjectIDParameter, identityID openapi.IdentityIDParameter, networkID openapi.NetworkIDParameter) {
	if err := rbac.AllowProjectScope(r.Context(), "region:networks", identityapi.Delete, organizationID, projectID); err != nil {
		errors.HandleError(w, r, err)
		return
	}

	if err := network.New(h.client, h.namespace).Delete(r.Context(), organizationID, projectID, networkID); err != nil {
		errors.HandleError(w, r, err)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}

func (h *Handler) GetApiV1OrganizationsOrganizationIDSecuritygroups(w http.ResponseWriter, r *http.Request, organizationID openapi.OrganizationIDParameter, params openapi.GetApiV1OrganizationsOrganizationIDSecuritygroupsParams) {
	if err := rbac.AllowOrganizationScope(r.Context(), "region:securitygroups", identityapi.Read, organizationID); err != nil {
		errors.HandleError(w, r, err)
		return
	}

	result, err := securitygroup.New(h.client, h.namespace).List(r.Context(), organizationID, params)
	if err != nil {
		errors.HandleError(w, r, err)
		return
	}

	util.WriteJSONResponse(w, r, http.StatusOK, result)
}

func (h *Handler) PostApiV1OrganizationsOrganizationIDProjectsProjectIDIdentitiesIdentityIDSecuritygroups(w http.ResponseWriter, r *http.Request, organizationID openapi.OrganizationIDParameter, projectID openapi.ProjectIDParameter, identityID openapi.IdentityIDParameter) {
	if err := rbac.AllowProjectScope(r.Context(), "region:securitygroups", identityapi.Create, organizationID, projectID); err != nil {
		errors.HandleError(w, r, err)
		return
	}

	request := &openapi.SecurityGroupWrite{}

	if err := util.ReadJSONBody(r, request); err != nil {
		errors.HandleError(w, r, err)
		return
	}

	result, err := securitygroup.New(h.client, h.namespace).Create(r.Context(), organizationID, projectID, identityID, request)
	if err != nil {
		errors.HandleError(w, r, err)
		return
	}

	util.WriteJSONResponse(w, r, http.StatusCreated, result)
}

func (h *Handler) DeleteApiV1OrganizationsOrganizationIDProjectsProjectIDIdentitiesIdentityIDSecuritygroupsSecurityGroupID(w http.ResponseWriter, r *http.Request, organizationID openapi.OrganizationIDParameter, projectID openapi.ProjectIDParameter, identityID openapi.IdentityIDParameter, securityGroupID openapi.SecurityGroupIDParameter) {
	if err := rbac.AllowProjectScope(r.Context(), "region:securitygroups", identityapi.Delete, organizationID, projectID); err != nil {
		errors.HandleError(w, r, err)
		return
	}

	if err := securitygroup.New(h.client, h.namespace).Delete(r.Context(), organizationID, projectID, securityGroupID); err != nil {
		errors.HandleError(w, r, err)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}

func (h *Handler) GetApiV1OrganizationsOrganizationIDProjectsProjectIDIdentitiesIdentityIDSecuritygroupsSecurityGroupID(w http.ResponseWriter, r *http.Request, organizationID openapi.OrganizationIDParameter, projectID openapi.ProjectIDParameter, identityID openapi.IdentityIDParameter, securityGroupID openapi.SecurityGroupIDParameter) {
	if err := rbac.AllowProjectScope(r.Context(), "region:securitygroups", identityapi.Read, organizationID, projectID); err != nil {
		errors.HandleError(w, r, err)
		return
	}

	result, err := securitygroup.New(h.client, h.namespace).Get(r.Context(), organizationID, projectID, securityGroupID)
	if err != nil {
		errors.HandleError(w, r, err)
		return
	}

	util.WriteJSONResponse(w, r, http.StatusOK, result)
}

//nolint:dupl
func (h *Handler) PutApiV1OrganizationsOrganizationIDProjectsProjectIDIdentitiesIdentityIDSecuritygroupsSecurityGroupID(w http.ResponseWriter, r *http.Request, organizationID openapi.OrganizationIDParameter, projectID openapi.ProjectIDParameter, identityID openapi.IdentityIDParameter, securityGroupID openapi.SecurityGroupIDParameter) {
	if err := rbac.AllowProjectScope(r.Context(), "region:securitygroups", identityapi.Update, organizationID, projectID); err != nil {
		errors.HandleError(w, r, err)
		return
	}

	request := &openapi.SecurityGroupWrite{}

	if err := util.ReadJSONBody(r, request); err != nil {
		errors.HandleError(w, r, err)
		return
	}

	result, err := securitygroup.New(h.client, h.namespace).Update(r.Context(), organizationID, projectID, identityID, securityGroupID, request)
	if err != nil {
		errors.HandleError(w, r, err)
		return
	}

	util.WriteJSONResponse(w, r, http.StatusAccepted, result)
}

func (h *Handler) GetApiV1OrganizationsOrganizationIDProjectsProjectIDIdentitiesIdentityIDSecuritygroupsSecurityGroupIDRules(w http.ResponseWriter, r *http.Request, organizationID openapi.OrganizationIDParameter, projectID openapi.ProjectIDParameter, identityID openapi.IdentityIDParameter, securityGroupID openapi.SecurityGroupIDParameter) {
	if err := rbac.AllowProjectScope(r.Context(), "region:securitygroups", identityapi.Read, organizationID, projectID); err != nil {
		errors.HandleError(w, r, err)
		return
	}

	// TODO: filtering???
	result, err := securitygrouprule.New(h.client, h.namespace).List(r.Context(), organizationID)
	if err != nil {
		errors.HandleError(w, r, err)
		return
	}

	util.WriteJSONResponse(w, r, http.StatusOK, result)
}

//nolint:dupl
func (h *Handler) PostApiV1OrganizationsOrganizationIDProjectsProjectIDIdentitiesIdentityIDSecuritygroupsSecurityGroupIDRules(w http.ResponseWriter, r *http.Request, organizationID openapi.OrganizationIDParameter, projectID openapi.ProjectIDParameter, identityID openapi.IdentityIDParameter, securityGroupID openapi.SecurityGroupIDParameter) {
	if err := rbac.AllowProjectScope(r.Context(), "region:securitygroups", identityapi.Create, organizationID, projectID); err != nil {
		errors.HandleError(w, r, err)
		return
	}

	request := &openapi.SecurityGroupRuleWrite{}

	if err := util.ReadJSONBody(r, request); err != nil {
		errors.HandleError(w, r, err)
		return
	}

	result, err := securitygrouprule.New(h.client, h.namespace).Create(r.Context(), organizationID, projectID, identityID, securityGroupID, request)
	if err != nil {
		errors.HandleError(w, r, err)
		return
	}

	util.WriteJSONResponse(w, r, http.StatusCreated, result)
}

func (h *Handler) DeleteApiV1OrganizationsOrganizationIDProjectsProjectIDIdentitiesIdentityIDSecuritygroupsSecurityGroupIDRulesRuleID(w http.ResponseWriter, r *http.Request, organizationID openapi.OrganizationIDParameter, projectID openapi.ProjectIDParameter, identityID openapi.IdentityIDParameter, securityGroupID openapi.SecurityGroupIDParameter, ruleID openapi.RuleIDParameter) {
	if err := rbac.AllowProjectScope(r.Context(), "region:securitygroups", identityapi.Delete, organizationID, projectID); err != nil {
		errors.HandleError(w, r, err)
		return
	}

	if err := securitygrouprule.New(h.client, h.namespace).Delete(r.Context(), organizationID, projectID, ruleID); err != nil {
		errors.HandleError(w, r, err)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}

func (h *Handler) GetApiV1OrganizationsOrganizationIDProjectsProjectIDIdentitiesIdentityIDSecuritygroupsSecurityGroupIDRulesRuleID(w http.ResponseWriter, r *http.Request, organizationID openapi.OrganizationIDParameter, projectID openapi.ProjectIDParameter, identityID openapi.IdentityIDParameter, securityGroupID openapi.SecurityGroupIDParameter, ruleID openapi.RuleIDParameter) {
	if err := rbac.AllowProjectScope(r.Context(), "region:securitygroups", identityapi.Read, organizationID, projectID); err != nil {
		errors.HandleError(w, r, err)
		return
	}

	result, err := securitygrouprule.New(h.client, h.namespace).Get(r.Context(), organizationID, projectID, ruleID)
	if err != nil {
		errors.HandleError(w, r, err)
		return
	}

	util.WriteJSONResponse(w, r, http.StatusOK, result)
}

func (h *Handler) GetApiV1OrganizationsOrganizationIDServers(w http.ResponseWriter, r *http.Request, organizationID openapi.OrganizationIDParameter, params openapi.GetApiV1OrganizationsOrganizationIDServersParams) {
	if err := rbac.AllowOrganizationScope(r.Context(), "region:servers", identityapi.Read, organizationID); err != nil {
		errors.HandleError(w, r, err)
		return
	}

	result, err := server.NewClient(h.client, h.namespace).List(r.Context(), organizationID, params)
	if err != nil {
		errors.HandleError(w, r, err)
		return
	}

	util.WriteJSONResponse(w, r, http.StatusOK, result)
}

func (h *Handler) PostApiV1OrganizationsOrganizationIDProjectsProjectIDIdentitiesIdentityIDServers(w http.ResponseWriter, r *http.Request, organizationID openapi.OrganizationIDParameter, projectID openapi.ProjectIDParameter, identityID openapi.IdentityIDParameter) {
	if err := rbac.AllowProjectScope(r.Context(), "region:servers", identityapi.Create, organizationID, projectID); err != nil {
		errors.HandleError(w, r, err)
		return
	}

	request := &openapi.ServerWrite{}

	if err := util.ReadJSONBody(r, request); err != nil {
		errors.HandleError(w, r, err)
		return
	}

	result, err := server.NewClient(h.client, h.namespace).Create(r.Context(), organizationID, projectID, identityID, request)
	if err != nil {
		errors.HandleError(w, r, err)
		return
	}

	util.WriteJSONResponse(w, r, http.StatusCreated, result)
}

//nolint:dupl
func (h *Handler) PutApiV1OrganizationsOrganizationIDProjectsProjectIDIdentitiesIdentityIDServersServerID(w http.ResponseWriter, r *http.Request, organizationID openapi.OrganizationIDParameter, projectID openapi.ProjectIDParameter, identityID openapi.IdentityIDParameter, serverID openapi.ServerIDParameter) {
	if err := rbac.AllowProjectScope(r.Context(), "region:servers", identityapi.Create, organizationID, projectID); err != nil {
		errors.HandleError(w, r, err)
		return
	}

	request := &openapi.ServerWrite{}

	if err := util.ReadJSONBody(r, request); err != nil {
		errors.HandleError(w, r, err)
		return
	}

	result, err := server.NewClient(h.client, h.namespace).Update(r.Context(), organizationID, projectID, identityID, serverID, request)
	if err != nil {
		errors.HandleError(w, r, err)
		return
	}

	util.WriteJSONResponse(w, r, http.StatusAccepted, result)
}

func (h *Handler) DeleteApiV1OrganizationsOrganizationIDProjectsProjectIDIdentitiesIdentityIDServersServerID(w http.ResponseWriter, r *http.Request, organizationID openapi.OrganizationIDParameter, projectID openapi.ProjectIDParameter, identityID openapi.IdentityIDParameter, serverID openapi.ServerIDParameter) {
	if err := rbac.AllowProjectScope(r.Context(), "region:servers", identityapi.Delete, organizationID, projectID); err != nil {
		errors.HandleError(w, r, err)
		return
	}

	err := server.NewClient(h.client, h.namespace).Delete(r.Context(), organizationID, projectID, serverID)
	if err != nil {
		errors.HandleError(w, r, err)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}

func (h *Handler) GetApiV1OrganizationsOrganizationIDProjectsProjectIDIdentitiesIdentityIDServersServerID(w http.ResponseWriter, r *http.Request, organizationID openapi.OrganizationIDParameter, projectID openapi.ProjectIDParameter, identityID openapi.IdentityIDParameter, serverID openapi.ServerIDParameter) {
	if err := rbac.AllowProjectScope(r.Context(), "region:servers", identityapi.Read, organizationID, projectID); err != nil {
		errors.HandleError(w, r, err)
		return
	}

	result, err := server.NewClient(h.client, h.namespace).Get(r.Context(), organizationID, projectID, serverID)
	if err != nil {
		errors.HandleError(w, r, err)
		return
	}

	util.WriteJSONResponse(w, r, http.StatusOK, result)
}

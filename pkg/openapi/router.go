// Package openapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.2 DO NOT EDIT.
package openapi

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/oapi-codegen/runtime"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (GET /api/v1/organizations/{organizationID}/projects/{projectID}/regions)
	GetApiV1OrganizationsOrganizationIDProjectsProjectIDRegions(w http.ResponseWriter, r *http.Request, organizationID OrganizationIDParameter, projectID ProjectIDParameter)

	// (GET /api/v1/organizations/{organizationID}/projects/{projectID}/regions/{regionID}/externalnetworks)
	GetApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDExternalnetworks(w http.ResponseWriter, r *http.Request, organizationID OrganizationIDParameter, projectID ProjectIDParameter, regionID RegionIDParameter)

	// (GET /api/v1/organizations/{organizationID}/projects/{projectID}/regions/{regionID}/flavors)
	GetApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDFlavors(w http.ResponseWriter, r *http.Request, organizationID OrganizationIDParameter, projectID ProjectIDParameter, regionID RegionIDParameter)

	// (POST /api/v1/organizations/{organizationID}/projects/{projectID}/regions/{regionID}/identities)
	PostApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDIdentities(w http.ResponseWriter, r *http.Request, organizationID OrganizationIDParameter, projectID ProjectIDParameter, regionID RegionIDParameter)

	// (POST /api/v1/organizations/{organizationID}/projects/{projectID}/regions/{regionID}/identities/{identityID}/physicalNetworks)
	PostApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDIdentitiesIdentityIDPhysicalNetworks(w http.ResponseWriter, r *http.Request, organizationID OrganizationIDParameter, projectID ProjectIDParameter, regionID RegionIDParameter, identityID IdentityIDParameter)

	// (GET /api/v1/organizations/{organizationID}/projects/{projectID}/regions/{regionID}/images)
	GetApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDImages(w http.ResponseWriter, r *http.Request, organizationID OrganizationIDParameter, projectID ProjectIDParameter, regionID RegionIDParameter)
}

// Unimplemented server implementation that returns http.StatusNotImplemented for each endpoint.

type Unimplemented struct{}

// (GET /api/v1/organizations/{organizationID}/projects/{projectID}/regions)
func (_ Unimplemented) GetApiV1OrganizationsOrganizationIDProjectsProjectIDRegions(w http.ResponseWriter, r *http.Request, organizationID OrganizationIDParameter, projectID ProjectIDParameter) {
	w.WriteHeader(http.StatusNotImplemented)
}

// (GET /api/v1/organizations/{organizationID}/projects/{projectID}/regions/{regionID}/externalnetworks)
func (_ Unimplemented) GetApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDExternalnetworks(w http.ResponseWriter, r *http.Request, organizationID OrganizationIDParameter, projectID ProjectIDParameter, regionID RegionIDParameter) {
	w.WriteHeader(http.StatusNotImplemented)
}

// (GET /api/v1/organizations/{organizationID}/projects/{projectID}/regions/{regionID}/flavors)
func (_ Unimplemented) GetApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDFlavors(w http.ResponseWriter, r *http.Request, organizationID OrganizationIDParameter, projectID ProjectIDParameter, regionID RegionIDParameter) {
	w.WriteHeader(http.StatusNotImplemented)
}

// (POST /api/v1/organizations/{organizationID}/projects/{projectID}/regions/{regionID}/identities)
func (_ Unimplemented) PostApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDIdentities(w http.ResponseWriter, r *http.Request, organizationID OrganizationIDParameter, projectID ProjectIDParameter, regionID RegionIDParameter) {
	w.WriteHeader(http.StatusNotImplemented)
}

// (POST /api/v1/organizations/{organizationID}/projects/{projectID}/regions/{regionID}/identities/{identityID}/physicalNetworks)
func (_ Unimplemented) PostApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDIdentitiesIdentityIDPhysicalNetworks(w http.ResponseWriter, r *http.Request, organizationID OrganizationIDParameter, projectID ProjectIDParameter, regionID RegionIDParameter, identityID IdentityIDParameter) {
	w.WriteHeader(http.StatusNotImplemented)
}

// (GET /api/v1/organizations/{organizationID}/projects/{projectID}/regions/{regionID}/images)
func (_ Unimplemented) GetApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDImages(w http.ResponseWriter, r *http.Request, organizationID OrganizationIDParameter, projectID ProjectIDParameter, regionID RegionIDParameter) {
	w.WriteHeader(http.StatusNotImplemented)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.Handler) http.Handler

// GetApiV1OrganizationsOrganizationIDProjectsProjectIDRegions operation middleware
func (siw *ServerInterfaceWrapper) GetApiV1OrganizationsOrganizationIDProjectsProjectIDRegions(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "organizationID" -------------
	var organizationID OrganizationIDParameter

	err = runtime.BindStyledParameterWithLocation("simple", false, "organizationID", runtime.ParamLocationPath, chi.URLParam(r, "organizationID"), &organizationID)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "organizationID", Err: err})
		return
	}

	// ------------- Path parameter "projectID" -------------
	var projectID ProjectIDParameter

	err = runtime.BindStyledParameterWithLocation("simple", false, "projectID", runtime.ParamLocationPath, chi.URLParam(r, "projectID"), &projectID)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "projectID", Err: err})
		return
	}

	ctx = context.WithValue(ctx, Oauth2AuthenticationScopes, []string{})

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetApiV1OrganizationsOrganizationIDProjectsProjectIDRegions(w, r, organizationID, projectID)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDExternalnetworks operation middleware
func (siw *ServerInterfaceWrapper) GetApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDExternalnetworks(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "organizationID" -------------
	var organizationID OrganizationIDParameter

	err = runtime.BindStyledParameterWithLocation("simple", false, "organizationID", runtime.ParamLocationPath, chi.URLParam(r, "organizationID"), &organizationID)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "organizationID", Err: err})
		return
	}

	// ------------- Path parameter "projectID" -------------
	var projectID ProjectIDParameter

	err = runtime.BindStyledParameterWithLocation("simple", false, "projectID", runtime.ParamLocationPath, chi.URLParam(r, "projectID"), &projectID)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "projectID", Err: err})
		return
	}

	// ------------- Path parameter "regionID" -------------
	var regionID RegionIDParameter

	err = runtime.BindStyledParameterWithLocation("simple", false, "regionID", runtime.ParamLocationPath, chi.URLParam(r, "regionID"), &regionID)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "regionID", Err: err})
		return
	}

	ctx = context.WithValue(ctx, Oauth2AuthenticationScopes, []string{})

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDExternalnetworks(w, r, organizationID, projectID, regionID)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDFlavors operation middleware
func (siw *ServerInterfaceWrapper) GetApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDFlavors(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "organizationID" -------------
	var organizationID OrganizationIDParameter

	err = runtime.BindStyledParameterWithLocation("simple", false, "organizationID", runtime.ParamLocationPath, chi.URLParam(r, "organizationID"), &organizationID)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "organizationID", Err: err})
		return
	}

	// ------------- Path parameter "projectID" -------------
	var projectID ProjectIDParameter

	err = runtime.BindStyledParameterWithLocation("simple", false, "projectID", runtime.ParamLocationPath, chi.URLParam(r, "projectID"), &projectID)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "projectID", Err: err})
		return
	}

	// ------------- Path parameter "regionID" -------------
	var regionID RegionIDParameter

	err = runtime.BindStyledParameterWithLocation("simple", false, "regionID", runtime.ParamLocationPath, chi.URLParam(r, "regionID"), &regionID)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "regionID", Err: err})
		return
	}

	ctx = context.WithValue(ctx, Oauth2AuthenticationScopes, []string{})

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDFlavors(w, r, organizationID, projectID, regionID)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// PostApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDIdentities operation middleware
func (siw *ServerInterfaceWrapper) PostApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDIdentities(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "organizationID" -------------
	var organizationID OrganizationIDParameter

	err = runtime.BindStyledParameterWithLocation("simple", false, "organizationID", runtime.ParamLocationPath, chi.URLParam(r, "organizationID"), &organizationID)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "organizationID", Err: err})
		return
	}

	// ------------- Path parameter "projectID" -------------
	var projectID ProjectIDParameter

	err = runtime.BindStyledParameterWithLocation("simple", false, "projectID", runtime.ParamLocationPath, chi.URLParam(r, "projectID"), &projectID)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "projectID", Err: err})
		return
	}

	// ------------- Path parameter "regionID" -------------
	var regionID RegionIDParameter

	err = runtime.BindStyledParameterWithLocation("simple", false, "regionID", runtime.ParamLocationPath, chi.URLParam(r, "regionID"), &regionID)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "regionID", Err: err})
		return
	}

	ctx = context.WithValue(ctx, Oauth2AuthenticationScopes, []string{})

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDIdentities(w, r, organizationID, projectID, regionID)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// PostApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDIdentitiesIdentityIDPhysicalNetworks operation middleware
func (siw *ServerInterfaceWrapper) PostApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDIdentitiesIdentityIDPhysicalNetworks(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "organizationID" -------------
	var organizationID OrganizationIDParameter

	err = runtime.BindStyledParameterWithLocation("simple", false, "organizationID", runtime.ParamLocationPath, chi.URLParam(r, "organizationID"), &organizationID)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "organizationID", Err: err})
		return
	}

	// ------------- Path parameter "projectID" -------------
	var projectID ProjectIDParameter

	err = runtime.BindStyledParameterWithLocation("simple", false, "projectID", runtime.ParamLocationPath, chi.URLParam(r, "projectID"), &projectID)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "projectID", Err: err})
		return
	}

	// ------------- Path parameter "regionID" -------------
	var regionID RegionIDParameter

	err = runtime.BindStyledParameterWithLocation("simple", false, "regionID", runtime.ParamLocationPath, chi.URLParam(r, "regionID"), &regionID)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "regionID", Err: err})
		return
	}

	// ------------- Path parameter "identityID" -------------
	var identityID IdentityIDParameter

	err = runtime.BindStyledParameterWithLocation("simple", false, "identityID", runtime.ParamLocationPath, chi.URLParam(r, "identityID"), &identityID)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "identityID", Err: err})
		return
	}

	ctx = context.WithValue(ctx, Oauth2AuthenticationScopes, []string{})

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDIdentitiesIdentityIDPhysicalNetworks(w, r, organizationID, projectID, regionID, identityID)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDImages operation middleware
func (siw *ServerInterfaceWrapper) GetApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDImages(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "organizationID" -------------
	var organizationID OrganizationIDParameter

	err = runtime.BindStyledParameterWithLocation("simple", false, "organizationID", runtime.ParamLocationPath, chi.URLParam(r, "organizationID"), &organizationID)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "organizationID", Err: err})
		return
	}

	// ------------- Path parameter "projectID" -------------
	var projectID ProjectIDParameter

	err = runtime.BindStyledParameterWithLocation("simple", false, "projectID", runtime.ParamLocationPath, chi.URLParam(r, "projectID"), &projectID)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "projectID", Err: err})
		return
	}

	// ------------- Path parameter "regionID" -------------
	var regionID RegionIDParameter

	err = runtime.BindStyledParameterWithLocation("simple", false, "regionID", runtime.ParamLocationPath, chi.URLParam(r, "regionID"), &regionID)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "regionID", Err: err})
		return
	}

	ctx = context.WithValue(ctx, Oauth2AuthenticationScopes, []string{})

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDImages(w, r, organizationID, projectID, regionID)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

type UnescapedCookieParamError struct {
	ParamName string
	Err       error
}

func (e *UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter '%s'", e.ParamName)
}

func (e *UnescapedCookieParamError) Unwrap() error {
	return e.Err
}

type UnmarshalingParamError struct {
	ParamName string
	Err       error
}

func (e *UnmarshalingParamError) Error() string {
	return fmt.Sprintf("Error unmarshaling parameter %s as JSON: %s", e.ParamName, e.Err.Error())
}

func (e *UnmarshalingParamError) Unwrap() error {
	return e.Err
}

type RequiredParamError struct {
	ParamName string
}

func (e *RequiredParamError) Error() string {
	return fmt.Sprintf("Query argument %s is required, but not found", e.ParamName)
}

type RequiredHeaderError struct {
	ParamName string
	Err       error
}

func (e *RequiredHeaderError) Error() string {
	return fmt.Sprintf("Header parameter %s is required, but not found", e.ParamName)
}

func (e *RequiredHeaderError) Unwrap() error {
	return e.Err
}

type InvalidParamFormatError struct {
	ParamName string
	Err       error
}

func (e *InvalidParamFormatError) Error() string {
	return fmt.Sprintf("Invalid format for parameter %s: %s", e.ParamName, e.Err.Error())
}

func (e *InvalidParamFormatError) Unwrap() error {
	return e.Err
}

type TooManyValuesForParamError struct {
	ParamName string
	Count     int
}

func (e *TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("Expected one value for %s, got %d", e.ParamName, e.Count)
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{})
}

type ChiServerOptions struct {
	BaseURL          string
	BaseRouter       chi.Router
	Middlewares      []MiddlewareFunc
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, r chi.Router) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseRouter: r,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, r chi.Router, baseURL string) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseURL:    baseURL,
		BaseRouter: r,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options ChiServerOptions) http.Handler {
	r := options.BaseRouter

	if r == nil {
		r = chi.NewRouter()
	}
	if options.ErrorHandlerFunc == nil {
		options.ErrorHandlerFunc = func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandlerFunc:   options.ErrorHandlerFunc,
	}

	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/api/v1/organizations/{organizationID}/projects/{projectID}/regions", wrapper.GetApiV1OrganizationsOrganizationIDProjectsProjectIDRegions)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/api/v1/organizations/{organizationID}/projects/{projectID}/regions/{regionID}/externalnetworks", wrapper.GetApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDExternalnetworks)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/api/v1/organizations/{organizationID}/projects/{projectID}/regions/{regionID}/flavors", wrapper.GetApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDFlavors)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/api/v1/organizations/{organizationID}/projects/{projectID}/regions/{regionID}/identities", wrapper.PostApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDIdentities)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/api/v1/organizations/{organizationID}/projects/{projectID}/regions/{regionID}/identities/{identityID}/physicalNetworks", wrapper.PostApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDIdentitiesIdentityIDPhysicalNetworks)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/api/v1/organizations/{organizationID}/projects/{projectID}/regions/{regionID}/images", wrapper.GetApiV1OrganizationsOrganizationIDProjectsProjectIDRegionsRegionIDImages)
	})

	return r
}

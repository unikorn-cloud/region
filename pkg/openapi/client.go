// Package openapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.2 DO NOT EDIT.
package openapi

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/oapi-codegen/runtime"
	externalRef0 "github.com/unikorn-cloud/core/pkg/openapi"
)

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A list of callbacks for modifying requests which are generated before sending over
	// the network.
	RequestEditors []RequestEditorFn
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(server string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Server: server,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = &http.Client{}
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, fn)
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// GetApiV1Regions request
	GetApiV1Regions(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV1RegionsRegionIDFlavors request
	GetApiV1RegionsRegionIDFlavors(ctx context.Context, regionID RegionIDParameter, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV1RegionsRegionIDImages request
	GetApiV1RegionsRegionIDImages(ctx context.Context, regionID RegionIDParameter, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) GetApiV1Regions(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetApiV1RegionsRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetApiV1RegionsRegionIDFlavors(ctx context.Context, regionID RegionIDParameter, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetApiV1RegionsRegionIDFlavorsRequest(c.Server, regionID)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetApiV1RegionsRegionIDImages(ctx context.Context, regionID RegionIDParameter, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetApiV1RegionsRegionIDImagesRequest(c.Server, regionID)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewGetApiV1RegionsRequest generates requests for GetApiV1Regions
func NewGetApiV1RegionsRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v1/regions")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetApiV1RegionsRegionIDFlavorsRequest generates requests for GetApiV1RegionsRegionIDFlavors
func NewGetApiV1RegionsRegionIDFlavorsRequest(server string, regionID RegionIDParameter) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "regionID", runtime.ParamLocationPath, regionID)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v1/regions/%s/flavors", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetApiV1RegionsRegionIDImagesRequest generates requests for GetApiV1RegionsRegionIDImages
func NewGetApiV1RegionsRegionIDImagesRequest(server string, regionID RegionIDParameter) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "regionID", runtime.ParamLocationPath, regionID)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v1/regions/%s/images", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (c *Client) applyEditors(ctx context.Context, req *http.Request, additionalEditors []RequestEditorFn) error {
	for _, r := range c.RequestEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	for _, r := range additionalEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	return nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(server string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(server, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Server = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// GetApiV1RegionsWithResponse request
	GetApiV1RegionsWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetApiV1RegionsResponse, error)

	// GetApiV1RegionsRegionIDFlavorsWithResponse request
	GetApiV1RegionsRegionIDFlavorsWithResponse(ctx context.Context, regionID RegionIDParameter, reqEditors ...RequestEditorFn) (*GetApiV1RegionsRegionIDFlavorsResponse, error)

	// GetApiV1RegionsRegionIDImagesWithResponse request
	GetApiV1RegionsRegionIDImagesWithResponse(ctx context.Context, regionID RegionIDParameter, reqEditors ...RequestEditorFn) (*GetApiV1RegionsRegionIDImagesResponse, error)
}

type GetApiV1RegionsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *RegionsResponse
	JSON401      *externalRef0.UnauthorizedResponse
	JSON500      *externalRef0.InternalServerErrorResponse
}

// Status returns HTTPResponse.Status
func (r GetApiV1RegionsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetApiV1RegionsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetApiV1RegionsRegionIDFlavorsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *FlavorsResponse
	JSON400      *externalRef0.BadRequestResponse
	JSON401      *externalRef0.UnauthorizedResponse
	JSON500      *externalRef0.InternalServerErrorResponse
}

// Status returns HTTPResponse.Status
func (r GetApiV1RegionsRegionIDFlavorsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetApiV1RegionsRegionIDFlavorsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetApiV1RegionsRegionIDImagesResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *ImagesResponse
	JSON400      *externalRef0.BadRequestResponse
	JSON401      *externalRef0.UnauthorizedResponse
	JSON500      *externalRef0.InternalServerErrorResponse
}

// Status returns HTTPResponse.Status
func (r GetApiV1RegionsRegionIDImagesResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetApiV1RegionsRegionIDImagesResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// GetApiV1RegionsWithResponse request returning *GetApiV1RegionsResponse
func (c *ClientWithResponses) GetApiV1RegionsWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetApiV1RegionsResponse, error) {
	rsp, err := c.GetApiV1Regions(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetApiV1RegionsResponse(rsp)
}

// GetApiV1RegionsRegionIDFlavorsWithResponse request returning *GetApiV1RegionsRegionIDFlavorsResponse
func (c *ClientWithResponses) GetApiV1RegionsRegionIDFlavorsWithResponse(ctx context.Context, regionID RegionIDParameter, reqEditors ...RequestEditorFn) (*GetApiV1RegionsRegionIDFlavorsResponse, error) {
	rsp, err := c.GetApiV1RegionsRegionIDFlavors(ctx, regionID, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetApiV1RegionsRegionIDFlavorsResponse(rsp)
}

// GetApiV1RegionsRegionIDImagesWithResponse request returning *GetApiV1RegionsRegionIDImagesResponse
func (c *ClientWithResponses) GetApiV1RegionsRegionIDImagesWithResponse(ctx context.Context, regionID RegionIDParameter, reqEditors ...RequestEditorFn) (*GetApiV1RegionsRegionIDImagesResponse, error) {
	rsp, err := c.GetApiV1RegionsRegionIDImages(ctx, regionID, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetApiV1RegionsRegionIDImagesResponse(rsp)
}

// ParseGetApiV1RegionsResponse parses an HTTP response from a GetApiV1RegionsWithResponse call
func ParseGetApiV1RegionsResponse(rsp *http.Response) (*GetApiV1RegionsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetApiV1RegionsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest RegionsResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest externalRef0.UnauthorizedResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest externalRef0.InternalServerErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParseGetApiV1RegionsRegionIDFlavorsResponse parses an HTTP response from a GetApiV1RegionsRegionIDFlavorsWithResponse call
func ParseGetApiV1RegionsRegionIDFlavorsResponse(rsp *http.Response) (*GetApiV1RegionsRegionIDFlavorsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetApiV1RegionsRegionIDFlavorsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest FlavorsResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest externalRef0.BadRequestResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest externalRef0.UnauthorizedResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest externalRef0.InternalServerErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParseGetApiV1RegionsRegionIDImagesResponse parses an HTTP response from a GetApiV1RegionsRegionIDImagesWithResponse call
func ParseGetApiV1RegionsRegionIDImagesResponse(rsp *http.Response) (*GetApiV1RegionsRegionIDImagesResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetApiV1RegionsRegionIDImagesResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest ImagesResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest externalRef0.BadRequestResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest externalRef0.UnauthorizedResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest externalRef0.InternalServerErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

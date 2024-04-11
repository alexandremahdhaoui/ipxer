// Package ipxerclient provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.1.0 DO NOT EDIT.
package ipxerclient

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/oapi-codegen/runtime"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

// Defines values for BuildarchSelector.
const (
	BuildarchSelectorArm32 BuildarchSelector = "arm32"
	BuildarchSelectorArm64 BuildarchSelector = "arm64"
	BuildarchSelectorI386  BuildarchSelector = "i386"
	BuildarchSelectorX8664 BuildarchSelector = "x86_64"
)

// Defines values for GetContentByIDParamsBuildarch.
const (
	GetContentByIDParamsBuildarchArm32 GetContentByIDParamsBuildarch = "arm32"
	GetContentByIDParamsBuildarchArm64 GetContentByIDParamsBuildarch = "arm64"
	GetContentByIDParamsBuildarchI386  GetContentByIDParamsBuildarch = "i386"
	GetContentByIDParamsBuildarchX8664 GetContentByIDParamsBuildarch = "x86_64"
)

// Defines values for GetIPXEBySelectorsParamsBuildarch.
const (
	Arm32 GetIPXEBySelectorsParamsBuildarch = "arm32"
	Arm64 GetIPXEBySelectorsParamsBuildarch = "arm64"
	I386  GetIPXEBySelectorsParamsBuildarch = "i386"
	X8664 GetIPXEBySelectorsParamsBuildarch = "x86_64"
)

// Error defines model for Error.
type Error struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

// UUID defines model for UUID.
type UUID = openapi_types.UUID

// Content Any content, e.g. a butane/ignition or cloud-init manifest.
type Content = string

// IPXE An iPXE manifest.
type IPXE = string

// BuildarchSelector defines model for buildarchSelector.
type BuildarchSelector string

// UuidSelector defines model for uuidSelector.
type UuidSelector = UUID

// N400 defines model for 400.
type N400 = Error

// N401 defines model for 401.
type N401 = Error

// N403 defines model for 403.
type N403 = Error

// N404 defines model for 404.
type N404 = Error

// N500 defines model for 500.
type N500 = Error

// N503 defines model for 503.
type N503 = Error

// GetContentByIDParams defines parameters for GetContentByID.
type GetContentByIDParams struct {
	Uuid      UuidSelector                  `form:"uuid" json:"uuid"`
	Buildarch GetContentByIDParamsBuildarch `form:"buildarch" json:"buildarch"`
}

// GetContentByIDParamsBuildarch defines parameters for GetContentByID.
type GetContentByIDParamsBuildarch string

// GetIPXEBySelectorsParams defines parameters for GetIPXEBySelectors.
type GetIPXEBySelectorsParams struct {
	Uuid      UuidSelector                      `form:"uuid" json:"uuid"`
	Buildarch GetIPXEBySelectorsParamsBuildarch `form:"buildarch" json:"buildarch"`
}

// GetIPXEBySelectorsParamsBuildarch defines parameters for GetIPXEBySelectors.
type GetIPXEBySelectorsParamsBuildarch string

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
	// GetIPXEBootstrap request
	GetIPXEBootstrap(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetContentByID request
	GetContentByID(ctx context.Context, contentID UUID, params *GetContentByIDParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetIPXEBySelectors request
	GetIPXEBySelectors(ctx context.Context, params *GetIPXEBySelectorsParams, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) GetIPXEBootstrap(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetIPXEBootstrapRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetContentByID(ctx context.Context, contentID UUID, params *GetContentByIDParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetContentByIDRequest(c.Server, contentID, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetIPXEBySelectors(ctx context.Context, params *GetIPXEBySelectorsParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetIPXEBySelectorsRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewGetIPXEBootstrapRequest generates requests for GetIPXEBootstrap
func NewGetIPXEBootstrapRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/boot.ipxe")
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

// NewGetContentByIDRequest generates requests for GetContentByID
func NewGetContentByIDRequest(server string, contentID UUID, params *GetContentByIDParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "contentID", runtime.ParamLocationPath, contentID)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/content/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		queryValues := queryURL.Query()

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "uuid", runtime.ParamLocationQuery, params.Uuid); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "buildarch", runtime.ParamLocationQuery, params.Buildarch); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

		queryURL.RawQuery = queryValues.Encode()
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetIPXEBySelectorsRequest generates requests for GetIPXEBySelectors
func NewGetIPXEBySelectorsRequest(server string, params *GetIPXEBySelectorsParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/ipxe")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		queryValues := queryURL.Query()

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "uuid", runtime.ParamLocationQuery, params.Uuid); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "buildarch", runtime.ParamLocationQuery, params.Buildarch); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

		queryURL.RawQuery = queryValues.Encode()
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
	// GetIPXEBootstrapWithResponse request
	GetIPXEBootstrapWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetIPXEBootstrapResponse, error)

	// GetContentByIDWithResponse request
	GetContentByIDWithResponse(ctx context.Context, contentID UUID, params *GetContentByIDParams, reqEditors ...RequestEditorFn) (*GetContentByIDResponse, error)

	// GetIPXEBySelectorsWithResponse request
	GetIPXEBySelectorsWithResponse(ctx context.Context, params *GetIPXEBySelectorsParams, reqEditors ...RequestEditorFn) (*GetIPXEBySelectorsResponse, error)
}

type GetIPXEBootstrapResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON400      *N400
	JSON401      *N401
	JSON403      *N403
	JSON404      *N404
	JSON500      *N500
	JSON503      *N503
}

// Status returns HTTPResponse.Status
func (r GetIPXEBootstrapResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetIPXEBootstrapResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetContentByIDResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON400      *N400
	JSON401      *N401
	JSON403      *N403
	JSON404      *N404
	JSON500      *N500
	JSON503      *N503
}

// Status returns HTTPResponse.Status
func (r GetContentByIDResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetContentByIDResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetIPXEBySelectorsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON400      *N400
	JSON401      *N401
	JSON403      *N403
	JSON404      *N404
	JSON500      *N500
	JSON503      *N503
}

// Status returns HTTPResponse.Status
func (r GetIPXEBySelectorsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetIPXEBySelectorsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// GetIPXEBootstrapWithResponse request returning *GetIPXEBootstrapResponse
func (c *ClientWithResponses) GetIPXEBootstrapWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetIPXEBootstrapResponse, error) {
	rsp, err := c.GetIPXEBootstrap(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetIPXEBootstrapResponse(rsp)
}

// GetContentByIDWithResponse request returning *GetContentByIDResponse
func (c *ClientWithResponses) GetContentByIDWithResponse(ctx context.Context, contentID UUID, params *GetContentByIDParams, reqEditors ...RequestEditorFn) (*GetContentByIDResponse, error) {
	rsp, err := c.GetContentByID(ctx, contentID, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetContentByIDResponse(rsp)
}

// GetIPXEBySelectorsWithResponse request returning *GetIPXEBySelectorsResponse
func (c *ClientWithResponses) GetIPXEBySelectorsWithResponse(ctx context.Context, params *GetIPXEBySelectorsParams, reqEditors ...RequestEditorFn) (*GetIPXEBySelectorsResponse, error) {
	rsp, err := c.GetIPXEBySelectors(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetIPXEBySelectorsResponse(rsp)
}

// ParseGetIPXEBootstrapResponse parses an HTTP response from a GetIPXEBootstrapWithResponse call
func ParseGetIPXEBootstrapResponse(rsp *http.Response) (*GetIPXEBootstrapResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetIPXEBootstrapResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest N400
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest N401
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest N403
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON403 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest N404
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest N500
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 503:
		var dest N503
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON503 = &dest

	}

	return response, nil
}

// ParseGetContentByIDResponse parses an HTTP response from a GetContentByIDWithResponse call
func ParseGetContentByIDResponse(rsp *http.Response) (*GetContentByIDResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetContentByIDResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest N400
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest N401
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest N403
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON403 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest N404
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest N500
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 503:
		var dest N503
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON503 = &dest

	}

	return response, nil
}

// ParseGetIPXEBySelectorsResponse parses an HTTP response from a GetIPXEBySelectorsWithResponse call
func ParseGetIPXEBySelectorsResponse(rsp *http.Response) (*GetIPXEBySelectorsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetIPXEBySelectorsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest N400
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest N401
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest N403
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON403 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest N404
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest N500
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 503:
		var dest N503
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON503 = &dest

	}

	return response, nil
}

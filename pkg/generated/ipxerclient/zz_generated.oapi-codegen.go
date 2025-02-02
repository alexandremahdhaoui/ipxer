// Package ipxerclient provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.3.0 DO NOT EDIT.
package ipxerclient

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
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

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/+xXXW/bOhL9K1w2jzIlf8RtBQQLp3ELA/0IkmZRoC4CWhpbbCVSJSknXsP//WJIyVYS",
	"IzFyc4H70CfL0uHMmeHhzHBNE1WUSoK0hsZrWnLNC7Cg3b9ZJfKU6yS7hBwSqzS+FJLG9HcFekUDKnkB",
	"NN4BaUA1/K6EhpTGVlcQUJNkUHBcCbIqaPydiv6bIQ3o7Zvh9XBAA8p10e/53+GA/gioXZVo1Vgt5IJu",
	"NgGtKpE+RQIxj/o/0jCnMX0V7mIO/VcTXl1NzugGXWkwpZIGXAYGUYQ/iZIWpMVHXpa5SLgVSoY/jZIu",
	"sFtelDl4ZAo0HkRRQAswhi+Q2SlPCdICYwNS5sANkCSD5BdZqUoTIcvK0s2hVMdaK+25pmASLUokU7u5",
	"8G7Q2iDqPo97t839SvLKZkqL/0O6JV9qtRQpkCXPRUoQANLWln045gXiGdWOG7NzpYv62ZBCGCPkgijM",
	"n+PhY+4/L+Z+O+b3Ss9EmoIMcINIqohUlmR8CaQE7TwrSawiPEnAGGIzYYgGoyqdwAsEvvXvQxo8L6RB",
	"O6SvGTQShHTLldxw42Kbq0qmL7FlxJSQiLloOxH3fBw/71Ad3z1UE2lBS54TA3oJmgBy2irU6hXhCy4k",
	"ybkF/QKhXUm4LSHB9Il9rn1k/edFdkd+l6CXIgFSSb7kIuezHP7BuPZ4Y2i2FUTr0cKtDcucCxfJYa6b",
	"5fucV+4Ezas8XxENVgtYQkrqFY6HOP82fgESzszhDBBOCi7FHAzy2ObZtQWfTOyXWpWgrfDdwu/mmvpC",
	"RWMqpHWdrW5oqJuF37Ttbq/3dLtdC/vube7wu+6oZj8hcYXe9a62rmi314fB8fB1B968nXW6vbTf4YPj",
	"YWfQGw67g+7rQRRFNNjxrFvnPSZ3NHDvnMtVs0kBAbZghJNZZbmEUCykcEVaaZLkqko7QgrbymXQIrrk",
	"WnBpYzJPlJnKJWgsrTHpsgGLprLkxtyk8VQSUhnQxj0R0iHY8mOSKA3+DSHGZNe7TnX9C1YN2q8wJuto",
	"w8loNBoxxqZyX7yN1u4He18O7RBe/UeUtzCVU2nAksuvF+PRJ2IsniP/6n/ji8vJl8+k/5b1ot4g6nZ7",
	"rI/R4cd3Xz6/n3y4uvhIMmtLE4dhbZklqsCTMxcLJhaysX86uhy30W7yMgwzoQybQ6o0L7VCaTClF2Gp",
	"VRoaq4EXJjxae3qbell4tK7JbUI/jaGbX6Al5ORoXfvahN5sxzvp7BZ1crGEjsd3vAGCm63TkwJLVM0K",
	"UUwrZefmutL5ycGW/RrmLTNRLEgjLjYX2tiZUnb3qsy5RUUzkZ4UYHm++1Tn0TvfpnwzlZ4t6XRQUMSR",
	"PpidW8uLOwQxf8hqn7zwZOMQ0xQznriDVU+voxxuuUw1kE88SzOuKkEDWumcxrTZ7IWwWTVzyuANvGjQ",
	"IcrQ1Za76v2Kw4nAIQXI6HxC5koTXiv60vUv1HMuEpAG2oRKnmRAeix6wOPm5oZx99lprF5rwo+Td+PP",
	"l+NOj0Uss0WOZKyw7pg4f6PzCQ1ofcqxTrGIRYhSJUheChrTPotYnwa05DZzRTXEfDIMDv8twCUNq65r",
	"qpOUxvQD2Mn5t/GpUtZYzUt6b4rv+YFjX3vY4uoGETQj/+NgBO1G7Kew3dZo+hS235r5nsIOWtPU49hj",
	"z/f4EA4Icv2uKgquVzSmF3Vb3OrGnyecgJOMC5krnuKfKXUi/G/OZ5CbkyXPKzBTikeBL4y7+uE2/kDr",
	"zVQQruuHydnmsQ1+51Gnq8mZU8funvp9/WBSE78rICLFS8lcgCZq7uS/HSsCf4FEje3uj1sef/cSGexH",
	"7TiHd66zB+AfXsQ3P56j8e0k9kfmj8k8XUleiITjSMgb1ZDZighryOSMtQTdZNRr+qAqtWo20TwU8r9U",
	"N39q4+G1sZkSUS+mtdP3S+Bm81cAAAD//z+wuVh9EwAA",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}

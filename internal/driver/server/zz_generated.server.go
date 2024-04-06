// Package server provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.1.0 DO NOT EDIT.
package server

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
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

// Defines values for GetIPXEBySelectorsParamsBuildarch.
const (
	GetIPXEBySelectorsParamsBuildarchArm32 GetIPXEBySelectorsParamsBuildarch = "arm32"
	GetIPXEBySelectorsParamsBuildarchArm64 GetIPXEBySelectorsParamsBuildarch = "arm64"
	GetIPXEBySelectorsParamsBuildarchI386  GetIPXEBySelectorsParamsBuildarch = "i386"
	GetIPXEBySelectorsParamsBuildarchX8664 GetIPXEBySelectorsParamsBuildarch = "x86_64"
)

// Defines values for GetConfigByIDParamsBuildarch.
const (
	Arm32 GetConfigByIDParamsBuildarch = "arm32"
	Arm64 GetConfigByIDParamsBuildarch = "arm64"
	I386  GetConfigByIDParamsBuildarch = "i386"
	X8664 GetConfigByIDParamsBuildarch = "x86_64"
)

// Error defines model for Error.
type Error struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

// UUID defines model for UUID.
type UUID = openapi_types.UUID

// Config Any butane/ignition or cloud-init manifest.
type Config = string

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

// GetIPXEBySelectorsParams defines parameters for GetIPXEBySelectors.
type GetIPXEBySelectorsParams struct {
	Uuid      UuidSelector                      `form:"uuid" json:"uuid"`
	Buildarch GetIPXEBySelectorsParamsBuildarch `form:"buildarch" json:"buildarch"`
}

// GetIPXEBySelectorsParamsBuildarch defines parameters for GetIPXEBySelectors.
type GetIPXEBySelectorsParamsBuildarch string

// GetConfigByIDParams defines parameters for GetConfigByID.
type GetConfigByIDParams struct {
	Uuid      UuidSelector                 `form:"uuid" json:"uuid"`
	Buildarch GetConfigByIDParamsBuildarch `form:"buildarch" json:"buildarch"`
}

// GetConfigByIDParamsBuildarch defines parameters for GetConfigByID.
type GetConfigByIDParamsBuildarch string

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Retrieve an iPXE config to chainload to "/ipxe?labels=values"
	// (GET /boot.ipxe)
	GetIPXEBootstrap(ctx echo.Context) error
	// Retrieve an iPXE manifest by selectors
	// (GET /ipxe)
	GetIPXEBySelectors(ctx echo.Context, params GetIPXEBySelectorsParams) error
	// Retrieve dynamically a configuration file by its profile name and config ID.
	// (GET /profile/{profileName}/config/{configID})
	GetConfigByID(ctx echo.Context, profileName string, configID UUID, params GetConfigByIDParams) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetIPXEBootstrap converts echo context to params.
func (w *ServerInterfaceWrapper) GetIPXEBootstrap(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetIPXEBootstrap(ctx)
	return err
}

// GetIPXEBySelectors converts echo context to params.
func (w *ServerInterfaceWrapper) GetIPXEBySelectors(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetIPXEBySelectorsParams
	// ------------- Required query parameter "uuid" -------------

	err = runtime.BindQueryParameter("form", true, true, "uuid", ctx.QueryParams(), &params.Uuid)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter uuid: %s", err))
	}

	// ------------- Required query parameter "buildarch" -------------

	err = runtime.BindQueryParameter("form", true, true, "buildarch", ctx.QueryParams(), &params.Buildarch)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter buildarch: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetIPXEBySelectors(ctx, params)
	return err
}

// GetConfigByID converts echo context to params.
func (w *ServerInterfaceWrapper) GetConfigByID(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "profileName" -------------
	var profileName string

	err = runtime.BindStyledParameterWithOptions("simple", "profileName", ctx.Param("profileName"), &profileName, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter profileName: %s", err))
	}

	// ------------- Path parameter "configID" -------------
	var configID UUID

	err = runtime.BindStyledParameterWithOptions("simple", "configID", ctx.Param("configID"), &configID, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter configID: %s", err))
	}

	// Parameter object where we will unmarshal all parameters from the context
	var params GetConfigByIDParams
	// ------------- Required query parameter "uuid" -------------

	err = runtime.BindQueryParameter("form", true, true, "uuid", ctx.QueryParams(), &params.Uuid)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter uuid: %s", err))
	}

	// ------------- Required query parameter "buildarch" -------------

	err = runtime.BindQueryParameter("form", true, true, "buildarch", ctx.QueryParams(), &params.Buildarch)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter buildarch: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetConfigByID(ctx, profileName, configID, params)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/boot.ipxe", wrapper.GetIPXEBootstrap)
	router.GET(baseURL+"/ipxe", wrapper.GetIPXEBySelectors)
	router.GET(baseURL+"/profile/:profileName/config/:configID", wrapper.GetConfigByID)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/+xY227bPBJ+FS6bS1mSD3FbAcHCadzCQJsGSbMoUAcBLY0tthKpkpQTr6F3Xwwp2Ypj",
	"JEY2Bf6LXuk0nPlm+M2BWtNY5oUUIIym0ZoWTLEcDCj7NCt5ljAVp1eQQWykwpdc0Ij+LkGtqEcFy4FG",
	"W0HqUQW/S64goZFRJXhUxynkDFeCKHMa/aC8/25IPXr/bng7HFCPMpX3e+46HNAbj5pVgVq1UVwsaFV5",
	"tCx58hwIlHnS/pGCOY3om2Drc+C+6uD6enJGKzSlQBdSaLARGIQhXmIpDAiDt6woMh4zw6UIfmoprGP3",
	"LC8ycJIJ0GgQhh7NQWu2QGSnLCEIC7TxSJEB00DiFOJfZCVLRbgoSkOrQ6GOlZLKYU1Ax4oXCKY2c+nM",
	"oLZB2H0Z9m4b+7VgpUml4v+FZAO+UHLJEyBLlvGEoAAIU2t27uhX8GdUG27UzqXK63tNcq41FwsiMX4W",
	"h/O5/zKf+22fP0o140kCwsMNIokkQhqSsiWQApS1LAUxkrA4Bq2JSbkmCrQsVQyv4PjGvnNp8DKXBm2X",
	"vqXQUBCSDVZyx7T1bS5LkbzGlhFdQMznvG2E79g4fllSHT9MqokwoATLiAa1BEUAMW0YatSKsAXjgmTM",
	"gHoF164F3BcQY/j4PtPOs/7LPHtAvytQSx4DKQVbMp6xWQZ/0K891nxUG0sx54sddwzcm6DIGLeOHGa5",
	"VrTPdGnzZ15m2YooMIrDEhLiFpTK5XrOBJ+DNhYUv/g+fgVIVs3hgFC8jWMTdNsjXGSxeSpZgDLctQ63",
	"tWvqqhaNKBfGtrm6uyGJFm4HN1u/3tP6tv3sh9O5ld+2Sjn7CbGt+raRtUlGu70+DI6Hbzvw7v2s0+0l",
	"/Q4bHA87g95w2B103w7CMKTeFmfdR3eQtAmxk/NiRWalYQICvhDcbppUJM5kmXS44KYVOq+Fa8kUZ8JE",
	"ZB5LPRVLUFhWI9L1B344FQXT+i6JpoKQUoPS9o6QDsF2H5FYKnBvCNE6vd12qdtfsGqk3Qqt047SjIxG",
	"o5Hv+1Oxz72GWrvO7e5+24U3/+LFPUzFVGgw5Orb5Xj0hWiDOeRe/Wd8eTX5ek767/1e2BuE3W7P76N3",
	"+PHD1/OPk0/Xl59JakyhoyCoNfuxzOu08flCNPpPR1fjtrSdurSPkZDan0MiFSuURCb4Ui2CQskk0EYB",
	"y3VwtHbwqnpZcLSuwVWBm8TQzC9QAjJytK5tVYFT23FGOttFnYwvoePkO04Bwc1WyUmO5alGhVK+ktLM",
	"9W2pspODNbs1vtPs83xBGnL5c660mUlptq+KjBkksM+TkxwMy7af6jg645uQV1Ph0JJOBwlFLOiD0dm1",
	"LH8AEOOHqPbRCxMZB5imdrHY1q56ch1lcM9EooB8YWmSMlly6tFSZTSizWYvuEnLmWUGa8TzRjpAGtpS",
	"8pC933Aw4TigABldTMhcKsJqRl/Z3oV8zngMQkMbUMHiFEjPDx/huLu785n9bDlWr9XB58mH8fnVuNPz",
	"Qz81eYZgDDc2Tay90cWEerTOcixLfuiHKCULEKzgNKJ9P/T71KMFM6mtoQHG00fn8GkBNmhYZG1rmCQ0",
	"op/ATC6+j0+lNNooVtCdCb7nho193WAjV/cDrxn3nxZGoe14/ZxstzWWPifbb817z8kOWpPU07LHDu/x",
	"IRhQyLa3Ms+ZWtGIXtZdcMMbl084/cYp4yKTLMGHKbUk/HfGZpDpkyXLStBTiqnAFtoe+3Abb1B7cNCO",
	"rpoDn7aU2B5Mf+z3YisSPDguVt6z8o8PutXNXx79WR41HZXMVkS3dnoPXQol5zyDYF3fnLMcqrpBBmt3",
	"nZxVTzHqgxU6XU3OHpNpd87nv0sgPMEj7ZyDInJuC2htHCum/f2AVWr796EF7cmfEI8GvIPt141sv/km",
	"CP/vDxDvH5tbzTHib3Y9lV3JSrCcxwwPMGznJIX8xGzjRjdktsM0YaI5dJHJmd/KwTrmN1VVVf8LAAD/",
	"/0agIYArFAAA",
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

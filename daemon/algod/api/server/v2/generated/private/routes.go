// Package private provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/algorand/oapi-codegen DO NOT EDIT.
package private

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"github.com/algorand/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Aborts a catchpoint catchup.
	// (DELETE /v2/catchup/{catchpoint})
	AbortCatchup(ctx echo.Context, catchpoint string) error
	// Starts a catchpoint catchup.
	// (POST /v2/catchup/{catchpoint})
	StartCatchup(ctx echo.Context, catchpoint string) error

	// (POST /v2/register-participation-keys/{address})
	RegisterParticipationKeys(ctx echo.Context, address string, params RegisterParticipationKeysParams) error

	// (POST /v2/shutdown)
	ShutdownNode(ctx echo.Context, params ShutdownNodeParams) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// AbortCatchup converts echo context to params.
func (w *ServerInterfaceWrapper) AbortCatchup(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error
	// ------------- Path parameter "catchpoint" -------------
	var catchpoint string

	err = runtime.BindStyledParameter("simple", false, "catchpoint", ctx.Param("catchpoint"), &catchpoint)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter catchpoint: %s", err))
	}

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.AbortCatchup(ctx, catchpoint)
	return err
}

// StartCatchup converts echo context to params.
func (w *ServerInterfaceWrapper) StartCatchup(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error
	// ------------- Path parameter "catchpoint" -------------
	var catchpoint string

	err = runtime.BindStyledParameter("simple", false, "catchpoint", ctx.Param("catchpoint"), &catchpoint)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter catchpoint: %s", err))
	}

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.StartCatchup(ctx, catchpoint)
	return err
}

// RegisterParticipationKeys converts echo context to params.
func (w *ServerInterfaceWrapper) RegisterParticipationKeys(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty":           true,
		"fee":              true,
		"key-dilution":     true,
		"round-last-valid": true,
		"no-wait":          true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error
	// ------------- Path parameter "address" -------------
	var address string

	err = runtime.BindStyledParameter("simple", false, "address", ctx.Param("address"), &address)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter address: %s", err))
	}

	ctx.Set("api_key.Scopes", []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params RegisterParticipationKeysParams
	// ------------- Optional query parameter "fee" -------------
	if paramValue := ctx.QueryParam("fee"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "fee", ctx.QueryParams(), &params.Fee)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter fee: %s", err))
	}

	// ------------- Optional query parameter "key-dilution" -------------
	if paramValue := ctx.QueryParam("key-dilution"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "key-dilution", ctx.QueryParams(), &params.KeyDilution)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter key-dilution: %s", err))
	}

	// ------------- Optional query parameter "round-last-valid" -------------
	if paramValue := ctx.QueryParam("round-last-valid"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "round-last-valid", ctx.QueryParams(), &params.RoundLastValid)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter round-last-valid: %s", err))
	}

	// ------------- Optional query parameter "no-wait" -------------
	if paramValue := ctx.QueryParam("no-wait"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "no-wait", ctx.QueryParams(), &params.NoWait)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter no-wait: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.RegisterParticipationKeys(ctx, address, params)
	return err
}

// ShutdownNode converts echo context to params.
func (w *ServerInterfaceWrapper) ShutdownNode(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty":  true,
		"timeout": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error

	ctx.Set("api_key.Scopes", []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params ShutdownNodeParams
	// ------------- Optional query parameter "timeout" -------------
	if paramValue := ctx.QueryParam("timeout"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "timeout", ctx.QueryParams(), &params.Timeout)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter timeout: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.ShutdownNode(ctx, params)
	return err
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}, si ServerInterface, m ...echo.MiddlewareFunc) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.DELETE("/v2/catchup/:catchpoint", wrapper.AbortCatchup, m...)
	router.POST("/v2/catchup/:catchpoint", wrapper.StartCatchup, m...)
	router.POST("/v2/register-participation-keys/:address", wrapper.RegisterParticipationKeys, m...)
	router.POST("/v2/shutdown", wrapper.ShutdownNode, m...)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file.
func GetSwagger() (*openapi3.Swagger, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	swagger, err := openapi3.NewSwaggerLoader().LoadSwaggerFromData(buf.Bytes())
	if err != nil {
		return nil, fmt.Errorf("error loading Swagger: %s", err)
	}
	return swagger, nil
}

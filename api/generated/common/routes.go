// Package common provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/algorand/oapi-codegen DO NOT EDIT.
package common

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Returns 200 if healthy.
	// (GET /health)
	MakeHealthCheck(ctx echo.Context) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// MakeHealthCheck converts echo context to params.
func (w *ServerInterfaceWrapper) MakeHealthCheck(ctx echo.Context) error {

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

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.MakeHealthCheck(ctx)
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

	router.GET("/health", wrapper.MakeHealthCheck, m...)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"uKYAAA==",
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

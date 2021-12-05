// Package handler provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.9.0 DO NOT EDIT.
package handler

import (
	"fmt"
	"net/http"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/labstack/echo/v4"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// joinRoom
	// (POST /rooms/join)
	JoinRoom(ctx echo.Context) error
	// createRoom
	// (POST /rooms/new)
	CreateRoom(ctx echo.Context) error
	// getRoom
	// (GET /rooms/{roomId})
	GetRoom(ctx echo.Context, roomId RoomId) error
	// Your GET endpoint
	// (GET /ws)
	Ws(ctx echo.Context) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// JoinRoom converts echo context to params.
func (w *ServerInterfaceWrapper) JoinRoom(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.JoinRoom(ctx)
	return err
}

// CreateRoom converts echo context to params.
func (w *ServerInterfaceWrapper) CreateRoom(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.CreateRoom(ctx)
	return err
}

// GetRoom converts echo context to params.
func (w *ServerInterfaceWrapper) GetRoom(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "roomId" -------------
	var roomId RoomId

	err = runtime.BindStyledParameterWithLocation("simple", false, "roomId", runtime.ParamLocationPath, ctx.Param("roomId"), &roomId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter roomId: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetRoom(ctx, roomId)
	return err
}

// Ws converts echo context to params.
func (w *ServerInterfaceWrapper) Ws(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.Ws(ctx)
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

	router.POST(baseURL+"/rooms/join", wrapper.JoinRoom)
	router.POST(baseURL+"/rooms/new", wrapper.CreateRoom)
	router.GET(baseURL+"/rooms/:roomId", wrapper.GetRoom)
	router.GET(baseURL+"/ws", wrapper.Ws)

}
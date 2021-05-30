// Package main provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package main

import (
	"fmt"
	"net/http"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/labstack/echo/v4"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Your GET endpoint
	// (GET /feeds)
	GetFeeds(ctx echo.Context) error
	// Your GET endpoint
	// (GET /feeds/{feedID})
	GetFeedsFeedID(ctx echo.Context, feedID string) error
	// Your GET endpoint
	// (GET /feeds/{feedID}/posts)
	GetFeedsFeedIDPosts(ctx echo.Context, feedID string) error
	// Your GET endpoint
	// (GET /group/{groupID})
	GetGroupGroupID(ctx echo.Context, groupID string) error
	// Your GET endpoint
	// (GET /groups)
	GetGroups(ctx echo.Context) error
	// Your GET endpoint
	// (GET /groups/{groupID}/feeds)
	GetGroupsGroupIDFeeds(ctx echo.Context, groupID string) error
	// Your GET endpoint
	// (GET /groups/{groupID}/posts)
	GetGroupGroupIDPosts(ctx echo.Context, groupID string) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetFeeds converts echo context to params.
func (w *ServerInterfaceWrapper) GetFeeds(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetFeeds(ctx)
	return err
}

// GetFeedsFeedID converts echo context to params.
func (w *ServerInterfaceWrapper) GetFeedsFeedID(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "feedID" -------------
	var feedID string

	err = runtime.BindStyledParameter("simple", false, "feedID", ctx.Param("feedID"), &feedID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter feedID: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetFeedsFeedID(ctx, feedID)
	return err
}

// GetFeedsFeedIDPosts converts echo context to params.
func (w *ServerInterfaceWrapper) GetFeedsFeedIDPosts(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "feedID" -------------
	var feedID string

	err = runtime.BindStyledParameter("simple", false, "feedID", ctx.Param("feedID"), &feedID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter feedID: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetFeedsFeedIDPosts(ctx, feedID)
	return err
}

// GetGroupGroupID converts echo context to params.
func (w *ServerInterfaceWrapper) GetGroupGroupID(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "groupID" -------------
	var groupID string

	err = runtime.BindStyledParameter("simple", false, "groupID", ctx.Param("groupID"), &groupID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter groupID: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetGroupGroupID(ctx, groupID)
	return err
}

// GetGroups converts echo context to params.
func (w *ServerInterfaceWrapper) GetGroups(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetGroups(ctx)
	return err
}

// GetGroupsGroupIDFeeds converts echo context to params.
func (w *ServerInterfaceWrapper) GetGroupsGroupIDFeeds(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "groupID" -------------
	var groupID string

	err = runtime.BindStyledParameter("simple", false, "groupID", ctx.Param("groupID"), &groupID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter groupID: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetGroupsGroupIDFeeds(ctx, groupID)
	return err
}

// GetGroupGroupIDPosts converts echo context to params.
func (w *ServerInterfaceWrapper) GetGroupGroupIDPosts(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "groupID" -------------
	var groupID string

	err = runtime.BindStyledParameter("simple", false, "groupID", ctx.Param("groupID"), &groupID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter groupID: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetGroupGroupIDPosts(ctx, groupID)
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

	router.GET(baseURL+"/feeds", wrapper.GetFeeds)
	router.GET(baseURL+"/feeds/:feedID", wrapper.GetFeedsFeedID)
	router.GET(baseURL+"/feeds/:feedID/posts", wrapper.GetFeedsFeedIDPosts)
	router.GET(baseURL+"/group/:groupID", wrapper.GetGroupGroupID)
	router.GET(baseURL+"/groups", wrapper.GetGroups)
	router.GET(baseURL+"/groups/:groupID/feeds", wrapper.GetGroupsGroupIDFeeds)
	router.GET(baseURL+"/groups/:groupID/posts", wrapper.GetGroupGroupIDPosts)

}

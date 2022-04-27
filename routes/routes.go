package routes

import "github.com/labstack/echo/v4"

var AppRouting = []Group{}

// Group .
type Group struct {
	Prefix      string
	Middlewares []echo.MiddlewareFunc
	Routes      []Route
}

// Route .
type Route struct {
	Method      string
	Path        string
	Handler     echo.HandlerFunc
	Middlewares []echo.MiddlewareFunc
}

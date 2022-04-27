package routes

import (
	c "app/config"
	"app/controllers/releases"
	"net/http"

	"github.com/labstack/echo/v4"
)

func init() {
	const path = "releases"
	handler := releases.Handler{}

	routes := Group{
		Prefix: "/" + c.Env.App.Version + "/" + path,
		Routes: []Route{
			{
				Method: http.MethodPost,
				Path:   "",
				Handler: func(c echo.Context) error {
					return handler.NewRelease(c)
				},
			},
		},
	}

	AppRouting = append(AppRouting, routes)
}

package api

import (
	"kafka-example/api/route"
	"kafka-example/helper"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// NewAPI struct ...
type NewAPI struct {
	E      *echo.Echo
	Helper helper.NewHelper
}

// Register ...
func (t *NewAPI) Register() *NewAPI {
	t.E.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowCredentials: true,
		AllowMethods:     []string{echo.OPTIONS, echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
	t.E.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))
	t.E.HTTPErrorHandler = func(err error, c echo.Context) {
		code := http.StatusNotFound
		if he, ok := err.(*echo.HTTPError); ok {
			code = he.Code
		}
		c.JSON(code, map[string]interface{}{
			"code":    code,
			"status":  "error",
			"message": http.StatusText(code),
		})
	}

	if true == t.Helper.Config.GetBool(`app.debug`) {
		t.E.Use(middleware.Logger())
		t.E.HideBanner = true
		t.E.Debug = true
	} else {
		t.E.HideBanner = true
		t.E.Debug = false
		t.E.Use(middleware.Recover())
	}
	route := route.NewRoute{
		E:      t.E,
		Helper: t.Helper,
	}
	route.Register()

	return t
}

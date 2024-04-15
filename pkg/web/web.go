package web

import (
	"context"

	"github.com/e9cryptteam/coinman/db"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Page struct {
	Title string
	Data  interface{}
}

func InitRoutes(e *echo.Echo, q db.Queries, ctx context.Context) {
	coinService := NewCoinCoinService(q, ctx)
	// marketWebController := NewMarketWebController(q, ctx)

	web := e.Group("")
	components := e.Group("/components")
	components.Use(middleware.CORSWithConfig(
    middleware.CORSConfig{
      AllowHeaders: []string{"HX-GET"},
    },
  ))

	web.GET("/", func(c echo.Context) error {
		return c.Redirect(302, "/home")
	})

	web.GET("/home", func(c echo.Context) error {
		data := Page{Title: "Home", Data: nil}
		return c.Render(200, "index.html", data)
	})

	// web.GET("/markets", marketWebController.Index)
}

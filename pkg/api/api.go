package api

import (
	"context"

	"github.com/e9cryptteam/coinman/db"
	"github.com/labstack/echo/v4"
)

func InitRoutes(e* echo.Echo, q db.Queries, ctx context.Context) {
  coinHandler := NewCoinHandler(q, ctx)

  api := e.Group("/api")

  api.GET("/coins", coinHandler.FindAll)
  api.GET("/coins/:id", coinHandler.GetById)
}

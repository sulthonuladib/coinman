package web

import (
	"context"
	"strconv"

	"github.com/e9cryptteam/coinman/db"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/labstack/echo/v4"
)

type MarketWebController struct {
	db  db.Queries
	ctx context.Context
}

func NewMarketWebController(db db.Queries, ctx context.Context) *MarketWebController {
	return &MarketWebController{db, ctx}
}

func (w *MarketWebController) Index(c echo.Context) error {
	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		limit = 10
	}
	offset, err := strconv.Atoi(c.QueryParam("offset"))
	if err != nil || offset < 0 {
		offset = 0
	}

	markets, err := w.db.GetMarkets(w.ctx, db.GetMarketsParams{
		Offset: pgtype.Int4{Int32: int32(offset), Valid: true},
		Limit:  pgtype.Int4{Int32: int32(limit), Valid: true},
	})
	if err != nil {
		return err
	}

	return c.Render(200, "market-table.html", Page{Title: "Markets", Data: markets})
}

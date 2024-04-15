package api

import (
	"context"
	"net/http"
	// "database/sql"
	"log"
	"strconv"

	"github.com/e9cryptteam/coinman/db"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/labstack/echo/v4"
)

type CoinHandler struct {
	db  db.Queries
	ctx context.Context
}

func NewCoinHandler(db db.Queries, ctx context.Context) *CoinHandler {
	return &CoinHandler{
		db:  db,
		ctx: ctx,
	}
}

func (h *CoinHandler) FindAll(c echo.Context) error {
	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		limit = 10
	}
	offset, err := strconv.Atoi(c.QueryParam("offset"))
	if err != nil {
		offset = 0
	}
  q := c.QueryParam("q")

	_, err = h.db.GetCoins(h.ctx, db.GetCoinsParams{
    Symbol: pgtype.Text{String: q, Valid: true},
		Limit:  pgtype.Int4{Int32: int32(limit), Valid: true},
		Offset: pgtype.Int4{Int32: int32(offset), Valid: true},
	})
	if err != nil {
    log.Println(err)
		return c.JSON(500, err)
	}

    return c.NoContent(http.StatusNoContent)
	// return c.JSON(200, coins)
}

func (h *CoinHandler) GetById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
    return echo.ErrBadRequest
    // return echo.NewHTTPError(400, "Invalid Coin ID")
	}

	coin, err := h.db.GetCoin(
		h.ctx,
		int32(id),
	)
	if err != nil {
		if err.Error() == pgx.ErrNoRows.Error() {
      return echo.NewHTTPError(404, "Coin not found")
		}
    return err
	}

	return c.JSON(200, coin)
}

package handlers

import (
	"context"
	"log"

	"github.com/e9cryptteam/coinman/db"
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

func (h *CoinHandler) Index(ctx echo.Context) error {
	coins, err := h.db.GetCoins(h.ctx, db.GetCoinsParams{})
	if err != nil {
		log.Fatal(err)
		return err
	}
	log.Println(coins)

	return ctx.Render(200, "index.html", struct {
		Title string
		Coins []db.Coin
	}{
		Title: "CoinMan",
		Coins: coins,
	})
}

package web

import (
	"context"

	"github.com/e9cryptteam/coinman/db"
	"github.com/jackc/pgx/v5/pgtype"
)

type CoinService struct {
	db  db.Queries
	ctx context.Context
}

func NewCoinCoinService(db db.Queries, ctx context.Context) *CoinService {
	return &CoinService{db, ctx}
}

func (w *CoinService) Find(q string, limit, offset int32) ([]db.Coin, error) {
  if limit == 0 {
    limit = 10
  }
	coins, err := w.db.GetCoins(w.ctx, db.GetCoinsParams{
		Symbol: pgtype.Text{String: q, Valid: true},
		Limit:  pgtype.Int4{Int32: int32(limit), Valid: true},
		Offset: pgtype.Int4{Int32: int32(offset), Valid: true},
	})
	if err != nil {
    return nil, err
	}

  return coins, nil
}

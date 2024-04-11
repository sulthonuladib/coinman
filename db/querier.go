// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"context"
)

type Querier interface {
	CreateCoin(ctx context.Context, arg CreateCoinParams) (Coin, error)
	DeleteCoin(ctx context.Context, id int32) (Coin, error)
	GetCoin(ctx context.Context, id int32) (Coin, error)
	GetCoins(ctx context.Context, arg GetCoinsParams) ([]Coin, error)
	UpdateCoin(ctx context.Context, arg UpdateCoinParams) (Coin, error)
}

var _ Querier = (*Queries)(nil)
// Code generated by sqlc. DO NOT EDIT.

package dao

import (
	"context"
)

type Querier interface {
	GetExchangeRate(ctx context.Context) (ExchangeRate, error)
	UpdateExchangeRate(ctx context.Context, arg UpdateExchangeRateParams) error
}

var _ Querier = (*Queries)(nil)

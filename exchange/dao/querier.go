// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package dao

import (
	"context"
)

type Querier interface {
	CreateExchangeRate(ctx context.Context, arg CreateExchangeRateParams) error
	GetExchangeCurrencies(ctx context.Context) ([]ExchangeCurrency, error)
	GetExchangeRateForConversion(ctx context.Context, arg GetExchangeRateForConversionParams) (ExchangeRate, error)
}

var _ Querier = (*Queries)(nil)

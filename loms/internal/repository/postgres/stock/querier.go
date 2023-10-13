// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0

package stock

import (
	"context"
)

type Querier interface {
	GetBySKU(ctx context.Context, sku int64) (int64, error)
	ReserveCancel(ctx context.Context, arg ReserveCancelParams) error
	ReserveRemove(ctx context.Context, arg ReserveRemoveParams) error
	ReserveStock(ctx context.Context, arg ReserveStockParams) (int64, error)
}

var _ Querier = (*Queries)(nil)
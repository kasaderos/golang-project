// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: cart_item.sql

package cart

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const addCartItem = `-- name: AddCartItem :exec
insert into cart_item (
    user_id,
    sku,
    name,
    price,
    count
) VALUES ($1, $2, $3, $4, $5)
`

type AddCartItemParams struct {
	UserID pgtype.Int8 `json:"user_id"`
	Sku    pgtype.Int4 `json:"sku"`
	Name   pgtype.Text `json:"name"`
	Price  pgtype.Int4 `json:"price"`
	Count  pgtype.Int4 `json:"count"`
}

func (q *Queries) AddCartItem(ctx context.Context, arg AddCartItemParams) error {
	_, err := q.db.Exec(ctx, addCartItem,
		arg.UserID,
		arg.Sku,
		arg.Name,
		arg.Price,
		arg.Count,
	)
	return err
}

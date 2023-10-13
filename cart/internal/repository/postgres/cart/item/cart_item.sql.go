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
    price,
    amount
) VALUES ($1, $2, $3, $4)
on conflict on constraint id
do update
set amount = EXCLUDED.amount + $4,
    price = $3
where EXCLUDED.user_id = $1 and EXCLUDED.sku = $2
`

type AddCartItemParams struct {
	UserID int64       `json:"user_id"`
	Sku    int64       `json:"sku"`
	Price  pgtype.Int4 `json:"price"`
	Amount pgtype.Int4 `json:"amount"`
}

func (q *Queries) AddCartItem(ctx context.Context, arg AddCartItemParams) error {
	_, err := q.db.Exec(ctx, addCartItem,
		arg.UserID,
		arg.Sku,
		arg.Price,
		arg.Amount,
	)
	return err
}

const deleteItem = `-- name: DeleteItem :exec
delete from cart_item
where user_id = $1 and sku = $2
`

type DeleteItemParams struct {
	UserID int64 `json:"user_id"`
	Sku    int64 `json:"sku"`
}

func (q *Queries) DeleteItem(ctx context.Context, arg DeleteItemParams) error {
	_, err := q.db.Exec(ctx, deleteItem, arg.UserID, arg.Sku)
	return err
}

const deleteItemByUser = `-- name: DeleteItemByUser :exec
delete from cart_item
where user_id = $1
`

func (q *Queries) DeleteItemByUser(ctx context.Context, userID int64) error {
	_, err := q.db.Exec(ctx, deleteItemByUser, userID)
	return err
}

const getItemsByUserID = `-- name: GetItemsByUserID :many
select user_id, sku, price, amount, created_at from cart_item
where user_id = $1
`

func (q *Queries) GetItemsByUserID(ctx context.Context, userID int64) ([]CartItem, error) {
	rows, err := q.db.Query(ctx, getItemsByUserID, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []CartItem
	for rows.Next() {
		var i CartItem
		if err := rows.Scan(
			&i.UserID,
			&i.Sku,
			&i.Price,
			&i.Amount,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
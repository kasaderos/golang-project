// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0

package cart

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type CartItem struct {
	UserID    int64              `json:"user_id"`
	Sku       int64              `json:"sku"`
	Price     pgtype.Int4        `json:"price"`
	Amount    pgtype.Int4        `json:"amount"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
}
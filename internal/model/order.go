package model

import "time"

type Order struct {
	ID        int64     `db:"id" json:"id"`
	Status    string    `db:"status" json:"status"`
	Amount    int64     `db:"amount" json:"amount"`
	ProductID int64     `db:"product_id" json:"product_id"`
	UserID    int64     `db:"user_id" json:"user_id"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"update_at" json:"updated_at"`
}

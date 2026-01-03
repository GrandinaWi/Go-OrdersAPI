package postgres

import (
	"OrderAPI/internal/model"
	"context"
	"database/sql"
)

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (repository *OrderRepository) GetList(ctx context.Context) ([]model.Order, error) {
	var orders []model.Order
	rows, err := repository.db.Query("SELECT id,status,amount,product_id,user_id,created_at,updated_at FROM orders ORDER BY id")
	if err == sql.ErrNoRows {
		return orders, nil
	}
	if err != nil {
		return orders, err
	}
	defer rows.Close()
	for rows.Next() {
		var order model.Order
		if err := rows.Scan(
			&order.ID,
			&order.Status,
			&order.Amount,
			&order.ProductID,
			&order.UserID,
			&order.CreatedAt,
			&order.UpdatedAt,
		); err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	return orders, nil

}
func (repository *OrderRepository) Create(ctx context.Context, order *model.Order) error {
	query := `
		INSERT INTO orders (status, amount,product_id, user_id)
		VALUES ($1, $2)
		RETURNING id, created_at, updated_at
	`
	return repository.db.QueryRowContext(ctx, query, order.Status, order.Amount).Scan(&order.ID, &order.CreatedAt, &order.UpdatedAt)
}
func (repository *OrderRepository) UpdateStatus(ctx context.Context, status string, id int64) error {
	query := `
		UPDATE orders
		SET status = $1, updated_at = now()
		WHERE id = $2
	`
	_, err := repository.db.ExecContext(ctx, query, status, id)
	return err
}
func (repository *OrderRepository) GetById(ctx context.Context, id int64) (*model.Order, error) {
	var order model.Order
	query := `
		SELECT id, status, amount, created_at, updated_at
		FROM orders
		WHERE id = $1
	`
	err := repository.db.QueryRowContext(ctx, query, id).Scan(&order.ID, &order.Status, &order.Amount, &order.CreatedAt, &order.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &order, nil
}

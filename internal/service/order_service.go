package service

import (
	"OrderAPI/internal/model"
	"context"
)

type OrderService interface {
	CreateOrder(ctx context.Context, amount int64, product_id int64, user_id int64) (*model.Order, error)
	GetOrder(ctx context.Context, orderId int64, userID int64) (*model.Order, error)
	UpdateOrder(ctx context.Context, status string, id int64) error
	GetOrders(ctx context.Context, userID int64) ([]model.Order, error)
}

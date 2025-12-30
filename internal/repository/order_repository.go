package repository

import (
	"OrderAPI/internal/model"
	"context"
)

type OrderRepository interface {
	Create(ctx context.Context, order *model.Order) error
	GetById(ctx context.Context, id int64) (*model.Order, error)
	GetList(ctx context.Context) ([]model.Order, error)
	UpdateStatus(ctx context.Context, status string, id int64) error
}

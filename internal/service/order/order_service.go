package order

import (
	"OrderAPI/internal/model"
	"OrderAPI/internal/repository"
	"context"
	"errors"
)

const (
	StatusNew        = "new"
	StatusProcessing = "processing"
	StatusCompleted  = "completed"
	StatusFailed     = "failed"
)

type Service struct {
	repo repository.OrderRepository
}

func NewOrderService(repo repository.OrderRepository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateOrder(ctx context.Context, amount int64, product_id int64, user_id int64) (*model.Order, error) {
	if amount <= 0 {
		return nil, errors.New("amount must be greater than zero")
	}
	if product_id <= 0 {
		return nil, errors.New("product_id must be greater than zero")
	}
	if user_id <= 0 {
		return nil, errors.New("user_id must be greater than zero")
	}
	order := &model.Order{
		Amount:    amount,
		ProductID: product_id,
		UserID:    user_id,
		Status:    StatusNew,
	}
	if err := s.repo.Create(ctx, order); err != nil {
		return nil, err
	}

	return order, nil
}
func (s *Service) GetOrder(ctx context.Context, id int64) (*model.Order, error) {
	if id <= 0 {
		return nil, errors.New("id must be greater than zero")
	}
	order, err := s.repo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	return order, nil
}
func (s *Service) UpdateOrder(ctx context.Context, status string, id int64) error {
	if id <= 0 {
		return errors.New("id must be greater than zero")
	}
	if !isValidStatus(status) {
		return errors.New("invalid status")
	}
	return s.repo.UpdateStatus(ctx, status, id)
}
func (s *Service) GetOrders(ctx context.Context) ([]model.Order, error) {
	return s.repo.GetList(ctx)
}

func isValidStatus(status string) bool {
	switch status {
	case StatusNew, StatusProcessing, StatusCompleted, StatusFailed:
		return true
	default:
		return false
	}
}

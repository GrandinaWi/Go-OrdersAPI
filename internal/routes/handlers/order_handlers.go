package handlers

import (
	"OrderAPI/internal/auth"
	"OrderAPI/internal/client"
	"OrderAPI/internal/service"
	"encoding/json"
	"net/http"
	"strconv"
)

type OrderHandler struct {
	service       service.OrderService
	userClient    *client.UserClient
	catalogClient *client.CatalogClient
}

func NewOrderHandler(service service.OrderService, userClient *client.UserClient, catalogClient *client.CatalogClient) *OrderHandler {
	return &OrderHandler{
		service:       service,
		userClient:    userClient,
		catalogClient: catalogClient,
	}
}

func (s *OrderHandler) CreateOrderHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req struct {
		Amount    int64 `json:"amount"`
		ProductID int64 `json:"productId"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid json", http.StatusBadRequest)
		return
	}
	if req.ProductID <= 0 || req.Amount <= 0 {
		http.Error(w, "invalid request data", http.StatusBadRequest)
		return
	}
	userID, err := auth.UserIDFromContext(ctx)
	if err != nil {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}
	order, err := s.service.CreateOrder(ctx, req.Amount, req.ProductID, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(order)
}
func (s *OrderHandler) GetOrderHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	userID, err := auth.UserIDFromContext(ctx)
	if err != nil {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}
	order, err := s.service.GetOrder(ctx, id, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if order == nil {
		http.Error(w, "Order not found", http.StatusNotFound)
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(order)
}
func (s *OrderHandler) GetOrders(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userID, err := auth.UserIDFromContext(ctx)
	if err != nil {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	orders, err := s.service.GetOrders(ctx, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(orders)
}
func (s *OrderHandler) UpdateOrderHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var req struct {
		Status string `json:"status"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid json", http.StatusBadRequest)
		return
	}
	if err := s.service.UpdateOrder(ctx, req.Status, id); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

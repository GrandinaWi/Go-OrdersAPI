package routes

import (
	"OrderAPI/internal/routes/handlers"
	"OrderAPI/internal/service"
	"net/http"
)

func NewRouter(orderService service.OrderService) http.Handler {
	mux := http.NewServeMux()
	handler := handlers.NewOrderHandler(orderService)
	mux.HandleFunc("POST /orders", handler.CreateOrderHandler)
	mux.HandleFunc("GET /orders/{id}", handler.GetOrderHandler)
	mux.HandleFunc("PUT /orders/{id}", handler.UpdateOrderHandler)
	mux.HandleFunc("GET /orders", handler.GetOrders)

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})
	return mux
}

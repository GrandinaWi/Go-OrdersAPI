package routes

import (
	"OrderAPI/internal/auth"
	"OrderAPI/internal/client"
	"OrderAPI/internal/routes/handlers"
	"OrderAPI/internal/service"
	"net/http"
)

func NewRouter(orderService service.OrderService, userClient *client.UserClient,
	catalogClient *client.CatalogClient,
) http.Handler {
	mux := http.NewServeMux()
	handler := handlers.NewOrderHandler(orderService, userClient, catalogClient)
	mux.Handle("POST /orders", auth.Middleware(http.HandlerFunc(handler.CreateOrderHandler)))
	mux.HandleFunc("GET /orders/{id}", handler.GetOrderHandler)
	mux.HandleFunc("PUT /orders/{id}", handler.UpdateOrderHandler)
	mux.HandleFunc("GET /orders", handler.GetOrders)

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})
	return mux
}

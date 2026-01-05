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
	mux.Handle("GET /orders/{id}", auth.Middleware(http.HandlerFunc(handler.GetOrderHandler)))
	mux.HandleFunc("PUT /orders/{id}", handler.UpdateOrderHandler)
	mux.Handle("GET /orders", auth.Middleware(http.HandlerFunc(handler.GetOrders)))

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})
	return mux
}

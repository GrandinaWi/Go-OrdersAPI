package main

import (
	"OrderAPI/internal/client"
	"OrderAPI/internal/config"
	"OrderAPI/pkg/postgres"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	repository "OrderAPI/internal/repository/postgres"
	"OrderAPI/internal/routes"
	service "OrderAPI/internal/service/order"

	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	// === root context ===
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	db, err := postgres.New(cfg.PostgresDSN)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to database")
	defer db.Close()

	userClient := client.NewUserClient(cfg.UserAPIURL)
	catalogClient := client.NewCatalogClient(cfg.ProductAPIURL)
	// === DI ===
	orderRepo := repository.NewOrderRepository(db)
	orderService := service.NewOrderService(orderRepo)

	router := routes.NewRouter(
		orderService,
		userClient,
		catalogClient,
	)

	// === HTTP server ===
	server := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  30 * time.Second,
	}

	// === graceful shutdown ===
	go func() {
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

		<-sigCh
		log.Println("shutdown signal received")

		shutdownCtx, shutdownCancel := context.WithTimeout(ctx, 10*time.Second)
		defer shutdownCancel()

		if err := server.Shutdown(shutdownCtx); err != nil {
			log.Printf("server shutdown error: %v", err)
		}
	}()

	log.Printf("server listening on %s", cfg.HTTPPort)

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server error: %v", err)
	}

	log.Println("server stopped gracefully")
}

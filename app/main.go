package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/suraboy/go-miniproject/app/internal"
	"github.com/suraboy/go-miniproject/app/internal/loan"
)

func main() {
	// Load configuration
	conf, err := internal.LoadConfig()
	if err != nil {
		log.Fatal("Error loading config:", err)
	}

	// Create Echo instance with optimizations
	e := echo.New()

	// Disable Echo banner for performance
	e.HideBanner = true
	e.HidePort = true

	// Add essential middleware only
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Add rate limiting for load testing
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(1000)))

	// Gzip compression for better throughput
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5, // Balanced compression
	}))

	// Set up validator
	e.Validator = internal.NewValidator()

	// Initialize services
	restSvc := loan.NewService()
	restHdl := loan.NewHandler(restSvc)

	// Register routes directly in main
	v1 := e.Group("/api/v1")
	v1.POST("/loans", restHdl.ApplyForLoan)

	// Create HTTP server with optimized settings
	server := &http.Server{
		Addr:           fmt.Sprintf(":%s", conf.Server.Port),
		Handler:        e,
		ReadTimeout:    conf.Server.GetReadTimeout(),
		WriteTimeout:   conf.Server.GetWriteTimeout(),
		IdleTimeout:    conf.Server.GetIdleTimeout(),
		MaxHeaderBytes: 1 << 20, // 1MB
	}

	// Start server in goroutine
	go func() {
		log.Printf("🚀 Server starting on port %s", conf.Server.Port)
		log.Printf("📊 Performance mode: GOMAXPROCS=%d", runtime.GOMAXPROCS(0))

		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("Failed to start server:", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("🛑 Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("✅ Server exited")
}

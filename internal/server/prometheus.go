package server

import (
    "context"
    "log"
    "net/http"
    "os"
    "os/signal"
    "syscall"
    "time"
)

func StartServer(metricsHandler http.Handler, address string) {
    server := &http.Server{
        Addr:    address,
        Handler: metricsHandler,
    }

    stop := make(chan os.Signal, 1)
    signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

    go func() {
        log.Printf("Starting server on %s", address)
        if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            log.Fatalf("Could not start server: %v", err)
        }
    }()


    <-stop
    log.Println("Shutting down server...")


    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    if err := server.Shutdown(ctx); err != nil {
        log.Fatalf("Server forced to shutdown: %v", err)
    }

    log.Println("Server stopped gracefully")
}

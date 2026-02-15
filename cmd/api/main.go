package main

import (
	"fmt"
	"net/http"

	"arifjehoh.com/ledgerlite/internal/config"
	"arifjehoh.com/ledgerlite/internal/router"
)

func main() {
	cfg := config.LoadConfig()
	fmt.Println("Starting API server on port", cfg.PORT)

	http.HandleFunc("/health", router.HealthHandler)

	http.ListenAndServe(cfg.PORT, nil)
}

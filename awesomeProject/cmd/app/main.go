package main

import (
	"fmt"
	"log"
	"net/http"

	"awesomeProject/internal/handler"
	"awesomeProject/internal/middleware"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/predict/hba1c", middleware.AuthMiddleware(handler.HandleHba1cPrediction))
	mux.HandleFunc("/predict/ldl", middleware.AuthMiddleware(handler.HandleLdlPrediction))
	mux.HandleFunc("/predict/tg", middleware.AuthMiddleware(handler.HandleTgPrediction))
	mux.HandleFunc("/predict/ferr", middleware.AuthMiddleware(handler.HandleFerrPrediction))
	mux.HandleFunc("/predict/hdl", middleware.AuthMiddleware(handler.HandleHdlPrediction))
	mux.HandleFunc("/predict/all", middleware.AuthMiddleware(handler.HandleAllPrediction))

	port := ":8080"
	fmt.Println("Server started on port", port)
	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatal("Server error:", err)
	}
}

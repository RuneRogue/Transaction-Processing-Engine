package main

import (
	"log"
	"net/http"
	"os"

	"github.com/RuneRogue/Transaction-Processing-Engine/handler"
	"github.com/RuneRogue/Transaction-Processing-Engine/repository"
	"github.com/RuneRogue/Transaction-Processing-Engine/service"
	"github.com/RuneRogue/Transaction-Processing-Engine/storage"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file or No .env file found")
	}

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not found in environment")
	}

	store := storage.InitializeMemoryStore()
	cardRepo := repository.NewCardRepository(store)
	transactionRepo := repository.NewTransactionRepository(store)

	cardService := service.NewCardService(cardRepo, transactionRepo)
	transactionService := service.NewTransactionService(cardRepo, transactionRepo)

	cardHandler := handler.NewCardHandler(cardService)
	txHandler := handler.NewTransactionHandler(transactionService)

	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	router.Get("/api/card/balance/{cardNumber}", cardHandler.GetBalance)
	router.Get("/api/card/transactions/{cardNumber}", cardHandler.GetTransactions)

	router.Post("/api/transaction", txHandler.HandleTransaction)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	log.Printf("Starting server on port %s\n", portString)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

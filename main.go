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
	//Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system env")
	}
	//Fetch PORT environment variable
	portString := os.Getenv("PORT")
	if portString == "" {
		portString = "8080"
		log.Println("PORT not set, defaulting to 8080")
	}
	//Initialize MemoryStore
	store := storage.InitializeMemoryStore()
	//Initialize Repositories
	cardRepo := repository.NewCardRepository(store)
	transactionRepo := repository.NewTransactionRepository(store)
	//Initialize Services
	cardService := service.NewCardService(cardRepo, transactionRepo)
	transactionService := service.NewTransactionService(cardRepo, transactionRepo)
	//Initialize Handlers
	cardHandler := handler.NewCardHandler(cardService)
	txHandler := handler.NewTransactionHandler(transactionService)

	router := chi.NewRouter()
	//Add CORS middleware
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	//Define API routes
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

package main

import (
	"log"
	"net/http"
	"todo-app/internal/handlers"
	"todo-app/internal/middleware"
	"todo-app/internal/store"

	"github.com/gorilla/mux"
)

func main() {
	// Initialize store
	todoStore := store.NewTodoStore()

	// Initialize handler
	todoHandler := handlers.NewTodoHandler(todoStore)

	// Setup router
	r := mux.NewRouter()

	// API routes
	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/todos", todoHandler.GetTodos).Methods("GET")
	api.HandleFunc("/todos", todoHandler.CreateTodo).Methods("POST")
	api.HandleFunc("/todos/{id}", todoHandler.GetTodo).Methods("GET")
	api.HandleFunc("/todos/{id}", todoHandler.UpdateTodo).Methods("PUT")
	api.HandleFunc("/todos/{id}", todoHandler.DeleteTodo).Methods("DELETE")

	// Apply CORS middleware
	handler := middleware.EnableCORS(r)

	// Start server
	log.Println("🚀 Server starting on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}

package main

import (
	"log"
	"net/http"

	"to-do/handlers"
	"to-do/storage"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	store := storage.NewMemoryStore()
	handler := handlers.NewTodoHandler(store)

	router.HandleFunc("/todos", handler.CreateTodo).Methods("POST")
	router.HandleFunc("/todos/{id}", handler.GetTodo).Methods("GET")
	router.HandleFunc("/todos/{id}", handler.UpdateTask).Methods("PUT")
	router.HandleFunc("/todos/{id}", handler.DeleteTask).Methods("DELETE")
	router.HandleFunc("/todos", handler.ListTask).Methods("GET")

	log.Println("Server running on port 8080")

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}

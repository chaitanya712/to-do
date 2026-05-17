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

	log.Println("Server running on port 8080")

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}

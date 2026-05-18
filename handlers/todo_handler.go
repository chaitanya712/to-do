package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"to-do/models"
	"to-do/storage"

	"github.com/gorilla/mux"
)

type TodoHandler struct {
	Store *storage.MemoryStore
}

func NewTodoHandler(store *storage.MemoryStore) *TodoHandler {
	return &TodoHandler{
		Store: store,
	}
}

func (h *TodoHandler) CreateTodo(w http.ResponseWriter, r *http.Request) {

	var req models.TodoRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	todo := models.Todo{
		Task:      req.Task,
		DueDate:   req.DueDate,
		Completed: false,
	}

	createdTodo := h.Store.Create(todo)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdTodo)

}

func (h *TodoHandler) GetTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	todo, err := h.Store.Get(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(todo)
}

func (h *TodoHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	var req models.TodoRequest

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	newTodo := models.Todo{
		Task:      req.Task,
		DueDate:   req.DueDate,
		Completed: req.Completed,
	}

	updatedTodo, err := h.Store.Update(id, newTodo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}
	json.NewEncoder(w).Encode(updatedTodo)
}

func (h *TodoHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	err = h.Store.Delete(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *TodoHandler) ListTask(w http.ResponseWriter, r *http.Request) {
	includeCompleted := r.URL.Query().Get("include_completed") == "true"

	todos := h.Store.List(includeCompleted)
	json.NewEncoder(w).Encode(todos)
}

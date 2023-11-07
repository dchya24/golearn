package todo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"dchya24/golearn/utils"

	"github.com/go-chi/chi/v5"
)

type Todo struct {
	Title       string
	Description string
}

var Todos = []Todo{
	{"Learn Go", "Learn Go from Internet"},
	{"Learn Blockchain", "Learn Blockchain for crypto"},
}

func GetTodo(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	fmt.Printf("Id is: %s \n", id)

	index, _ := strconv.Atoi(id)

	if id != "" {
		var todo Todo = Todos[index-1]

		w.Header().Add("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(todo)
	} else {
		w.Header().Add("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(Todos)
	}

}

func AddTodo(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Query().Get("title")
	description := r.URL.Query().Get("description")

	var newTodo Todo = Todo{title, description}

	Todos = append(Todos, newTodo)

	response := utils.Response{
		Status:  "Ok",
		Message: "Success Add Todo",
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	title := r.URL.Query().Get("title")
	description := r.URL.Query().Get("description")

	var updateTodo Todo = Todo{title, description}

	Todos[id-1] = updateTodo

	response := utils.Response{
		Status: "Ok", Message: "Todo updated success!",
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")

	id, _ := strconv.Atoi(idParam)

	Todos = append(Todos[:id-1], Todos[id:]...)

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Todos)
}

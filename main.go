package main

import (
	"net/http"

	"dchya24/golearn/email"
	"dchya24/golearn/todo"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	r.Get("/", todo.GetTodo)

	r.Get("/email", email.SendEmail)

	r.Get("/add", todo.AddTodo)

	r.Get("/update/{id}", todo.UpdateTodo)

	r.Get("/delete", todo.DeleteTodo)

	http.ListenAndServe(":8000", r)
}

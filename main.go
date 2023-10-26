package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Todo struct {
	Title       string
	Description string
}

type Response struct {
	Status  string
	Message string
}

var todos = []Todo{
	{"Learn Go", "Learn Go from Internet"},
	{"Learn Blockchain", "Learn Blockchain for crypto"},
}

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		id := c.QueryParams().Get("id")

		fmt.Printf("Id is: %s \n", id)

		index, _ := strconv.Atoi(id)

		if id != "" {
			var todo Todo = todos[index-1]

			return c.JSON(http.StatusOK, todo)
		}
		return c.JSON(http.StatusOK, todos)
	})

	e.GET("/add", func(c echo.Context) error {
		title := c.QueryParams().Get("title")
		description := c.QueryParams().Get("description")

		var newTodo Todo = Todo{title, description}

		todos = append(todos, newTodo)

		response := Response{"Ok", "Success Add Todo"}

		return c.JSON(http.StatusAccepted, response)
	})

	e.GET("update/:id", func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		title := c.QueryParams().Get("title")
		description := c.QueryParams().Get("description")

		var updateTodo Todo = Todo{title, description}

		todos[id-1] = updateTodo

		response := Response{"Ok", "Todo updated success!"}

		return c.JSON(http.StatusAccepted, response)
	})

	e.GET("delete", func(c echo.Context) error {
		idParam := c.QueryParams().Get("id")

		id, _ := strconv.Atoi(idParam)

		todos = append(todos[:id], todos[id+1:]...)

		return c.JSON(http.StatusAccepted, todos)
	})

	e.Logger.Fatal(e.Start(":8001"))
}

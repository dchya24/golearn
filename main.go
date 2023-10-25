package main

import (
	"net/http"

	"fmt"

	"github.com/labstack/echo/v4"
)

type Response struct {
	Message string `json:"message" xml:"message"`
}

func main() {
	fmt.Println("Hello")

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		response := Response{
			Message: "Hello World",
		}

		return c.JSON(http.StatusOK, response)
	})

	e.Logger.Fatal(e.Start(":8000"))
}

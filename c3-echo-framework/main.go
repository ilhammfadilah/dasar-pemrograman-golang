package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

type M map[string]interface{}

func main() {
	r := echo.New()

	r.GET("/", func(ctx echo.Context) error {
		data := "hello from index"
		return ctx.String(http.StatusOK, data)
	})

	r.GET("/json", func(ctx echo.Context) error {
		data := M{"message": "assalamualaikum", "Counter": 2}
		return ctx.JSON(http.StatusOK, data)
	})

	// parsing request
	r.GET("page1", func(ctx echo.Context) error {
		name := ctx.QueryParam("name")
		data := fmt.Sprintf("hllo %s", name)

		return ctx.String(http.StatusOK, data)
	})

	r.Start(":9000")
}

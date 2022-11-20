package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo"
)

type M map[string]interface{}

func main() {
	r := echo.New()

	r.Static("/static", "assets")

	r.GET("/", func(ctx echo.Context) error {
		data := "hello from index"
		return ctx.String(http.StatusOK, data)
	})

	// r.GET("/json", func(ctx echo.Context) error {
	// 	data := M{"message": "assalamualaikum", "Counter": 2}
	// 	return ctx.JSON(http.StatusOK, data)
	// })

	// parsing request
	// r.GET("page1", func(ctx echo.Context) error {
	// 	name := ctx.QueryParam("name")
	// 	data := fmt.Sprintf("hllo %s", name)

	// 	return ctx.String(http.StatusOK, data)
	// })

	// parsing url path param
	// r.GET("/page2/:name", func(ctx echo.Context) error {
	// 	name := ctx.Param("name")
	// 	data := fmt.Sprintf("assalamulaikum %s\n", name)

	// 	return ctx.String(http.StatusOK, data)
	// })

	// parsing url path dan seterusnya
	r.GET("/page3/:name/*", func(ctx echo.Context) error {
		name := ctx.Param("name")
		message := ctx.Param("*")
		data := fmt.Sprintf("assalamulaikum %s pesan: %s\n", name, message)

		return ctx.String(http.StatusOK, data)
	})

	// parsing form data
	r.POST("/page4", func(ctx echo.Context) error {
		name := ctx.FormValue("name")
		message := ctx.FormValue("message")
		data := fmt.Sprintf("Hello %s, I have message for you: %s", name, strings.Replace(message, "/", "", 1))
		return ctx.String(http.StatusOK, data)
	})

	r.Start(":9000")
}

package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	// middleware here
	e.Use(MiddlewareOne)
	e.Use(MiddlewareTwo)

	e.GET("/index", func(c echo.Context) (err error) {
		fmt.Println("threeee!")

		return c.JSON(http.StatusOK, true)
	})

	e.Logger.Fatal(e.Start(":9000"))
}

func MiddlewareOne(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("middleware one")
		return next(c)
	}
}

func MiddlewareTwo(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("middleware two")
		return next(c)
	}
}

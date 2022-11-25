package main

import (
	"fmt"
	"net/http"

	"github.com/alecthomas/kingpin"
	"github.com/labstack/echo"
)

var (
	argAppName = kingpin.Arg("name", "Application Name").Required().String()
	argPort    = kingpin.Arg("port", "Web Server Port").Default("9000").Int()
)

func main() {
	kingpin.Parse()

	appName := *argAppName
	port := fmt.Sprintf(":%d", *argPort)

	fmt.Printf("Starting %s at %s ", appName, port)

	e := echo.New()
	e.GET("/index", func(c echo.Context) error {
		return c.JSON(http.StatusOK, true)
	})
	e.Logger.Fatal(e.Start(port))
}

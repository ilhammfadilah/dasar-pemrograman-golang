package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/alecthomas/kingpin"
	"github.com/labstack/echo"
)

var (
	app        = kingpin.New("App", "Simple App")
	argAppName = app.Arg("name", "Application Name").Required().String()
	argPort    = app.Arg("port", "Web Server Port").Default("9000").Int()
)

func main() {
	comand, err := app.Parse(os.Args[1:])

	appName := *argAppName
	port := fmt.Sprintf(":%d", *argPort)

	fmt.Printf("Starting %s at %s ", appName, port)

	e := echo.New()
	e.GET("/index", func(c echo.Context) error {
		return c.JSON(http.StatusOK, true)
	})
	e.Logger.Fatal(e.Start(port))
}

package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
)

func MakeLogEntry(c echo.Context) *log.Entry {
	if c == nil {
		return log.WithFields(log.Fields{
			"at": time.Now().Format("2006-01-02 15:04:05"),
		})
	}

	return log.WithFields(log.Fields{
		"at":     time.Now().Format("2006-01-02 15:04:05"),
		"method": c.Request().Method,
		"url":    c.Request().URL.String(),
		"ip":     c.Request().RemoteAddr,
	})
}

func MiddlewareLogging(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		MakeLogEntry(c).Info("incoming request")
		return next(c)
	}
}

func ErrorHandler(err error, c echo.Context) {
	report, ok := err.(*echo.HTTPError)
	if ok {
		report.Message = fmt.Sprintf("http error %d - %v", report.Code, report.Message)
	} else {
		report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	MakeLogEntry(c).Error(report.Message)
	c.HTML(report.Code, report.Message.(string))
}

func main() {
	e := echo.New()

	// middleware here
	e.Use(MiddlewareOne)
	e.Use(MiddlewareTwo)
	e.Use(echo.WrapMiddleware(MiddlewareThree))

	e.Use(MiddlewareLogging)
	e.HTTPErrorHandler = ErrorHandler

	e.GET("/index", func(c echo.Context) error {
		return c.JSON(http.StatusOK, true)
	})

	lock := make(chan error)
	go func(locak chan error) {
		lock <- e.Start(":9000")
	}(lock)

	time.Sleep(1 * time.Millisecond)
	MakeLogEntry(nil).Warning("application started without ssl/lts enabled")

	err := <-lock
	if err != nil {
		MakeLogEntry(nil).Panic("failed to start application")
	}

	// e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
	// 	Format: "method=${method}, uri=${uri}, status=${status}\n",
	// }))

	// e.GET("/index", func(c echo.Context) (err error) {
	// 	fmt.Println("threeee!")

	// 	return c.JSON(http.StatusOK, true)
	// })

	// e.Logger.Fatal(e.Start(":9000"))
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

func MiddlewareThree(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("middleware three")
		next.ServeHTTP(w, r)
	})
}

package main

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
)

type User struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
	Age   int    `json:"age" validate:"gte=0,lte=80"`
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func main() {
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	// route
	e.POST("/users", func(c echo.Context) error {
		u := new(User)
		if err := c.Bind(u); err != nil {
			return err
		}
		if err := c.Validate(u); err != nil {
			return err
		}

		return c.JSON(http.StatusOK, true)
	})

	e.HTTPErrorHandler = func(err error, c echo.Context) {
		report, ok := err.(*echo.HTTPError)
		if !ok {
			report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		if castedObject, ok := err.(validator.ValidationErrors); ok {
			for _, err := range castedObject {
				switch err.Tag() {
				case "required":
					report.Message = fmt.Sprintf("%s tidak boleh kosong", err.Field())
				case "email":
					report.Message = fmt.Sprintf("%s tidak valid", err.Field())
				case "gte":
					report.Message = fmt.Sprintf("%s harus lebih dari / sama dengan %s", err.Field(), err.Param())
				case "lte":
					report.Message = fmt.Sprintf("%s harus lebih kecil / sama dengan %s", err.Field(), err.Param())
				}

				break
			}
		}

		errPage := fmt.Sprintf("%d.html", report.Code)
		if err := c.File(errPage); err != nil {
			c.HTML(report.Code, "Errrooorrr")
		}

		c.Logger().Error(report)
		c.JSON(report.Code, report)
	}

	e.Logger.Fatal(e.Start(":9000"))
}

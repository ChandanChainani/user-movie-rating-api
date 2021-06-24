package validators

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/go-playground/validator"
)

type CustomValidator struct {
	Validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
  if err := cv.Validator.Struct(i); err != nil {
    return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
  }
  return nil
}

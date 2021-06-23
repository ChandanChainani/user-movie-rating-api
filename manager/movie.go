package manager

import (
	"net/http"
	"github.com/labstack/echo/v4"

	e "usermovieratingapi/errors"
)

func (m *Manager) GetMovieInfoByName(c echo.Context) error {
	name := c.Param("name")

	if name == "" {
		return c.JSON(http.StatusInternalServerError, e.MovieNameFieldMissing)
	}

	result, err := m.DB.GetMovieInfoByName(name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, e.NoDataFound)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": result,
	})
}

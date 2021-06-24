package manager

import (
	"fmt"
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo-contrib/session"

	"usermovieratingapi/types/request"
	e "usermovieratingapi/errors"
)

func (m *Manager) GetUserRatedMoviesComments(c echo.Context) error {
	sess, _ := session.Get("session", c)

	// if _, ok := sess.Values["uID"]; !ok {
	// 	return echo.NewHTTPError(http.StatusUnauthorized, "Please provide validate credentials")
	// }

	u := &request.SearchUser{sess.Values["uID"].(string)}

	fmt.Println(u)
	result, err := m.DB.GetUserInfoByID(u)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, e.NoDataFound)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": result,
	})
}

func (m *Manager) SaveMovieRate(c echo.Context) error {
	sess, _ := session.Get("session", c)
	u := new(request.UserMovieRating)
	u.UserID = sess.Values["uID"].(string)

	if err := c.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	_, err := m.DB.InsertUserMovieRating(u)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, e.RateMovieFailed)
	}

	return c.JSON(http.StatusCreated, map[string]string{
		"data": "rating saved successfully",
	})
}

func (m *Manager) SaveMovieComment(c echo.Context) error {
	sess, _ := session.Get("session", c)
	u := new(request.UserMovieComment)
	u.UserID = sess.Values["uID"].(string)
	if err := c.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	_, err := m.DB.InsertUserMovieComment(u)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, e.RateMovieFailed)
	}

	// var b []interface{}
	// err := json.Unmarshal(result, &b)
	// fmt.Println(b)

	return c.JSON(http.StatusCreated, map[string]string{
		"data": "comment added successfully",
	})
}

func (m *Manager) AddMovie(c echo.Context) error {
	movie := new(request.Movie)
	fmt.Println(movie)
	if err := c.Bind(movie); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	_, err := m.DB.InsertMovie(movie)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, e.AddMovieFailed)
	}

	return c.JSON(http.StatusCreated, map[string]string{
		"data": "movie added successfully",
	})
}

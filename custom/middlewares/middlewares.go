package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo-contrib/session"
	e "usermovieratingapi/errors"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sess, _ := session.Get("session", c)

		if _, ok := sess.Values["uID"]; !ok {
			return echo.NewHTTPError(http.StatusUnauthorized, e.AuthFailed)
		}

		return next(c)
	}
}

func AdminMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sess, _ := session.Get("session", c)

		_, v1 := sess.Values["uID"]
		v2, _ := sess.Values["isAdmin"]
		t2, _ := v2.(int32)
		// fmt.Println(v2, reflect.TypeOf(v2), reflect.TypeOf(1), t2 == 1)

		if !v1 || t2 != 1 {
			// return nil
			return echo.NewHTTPError(http.StatusUnauthorized, e.AuthFailed)
		}
		// c.Response().Header().Set(echo.HeaderServer, "Echo/3.0")
		return next(c)
	}
}

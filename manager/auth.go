package manager

import (
	"fmt"
	// "reflect"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/labstack/echo/v4"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"

	"usermovieratingapi/types/request"
	e "usermovieratingapi/errors"
)

func (m *Manager) Login(c echo.Context) error {
	sess, _ := session.Get("session", c)
	// fmt.Println(sess.Values)
	if _, ok := sess.Values["uID"]; !ok {
		u := new(request.User)
		if err := c.Bind(u); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		user, err := m.DB.GetUserByEmail(u)
		if err != nil {
			// return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			return c.JSON(http.StatusBadRequest, e.AuthFailed)
		}
		// fmt.Println(user)
		// fmt.Println(reflect.TypeOf(user))
		// fmt.Println(user["_id"])
		// fmt.Println(reflect.TypeOf(user["_id"]))
		// fmt.Println(reflect.TypeOf(user["_id"].(primitive.ObjectID).Hex()))

		sess.Options = &sessions.Options{
			Path:     "/",
			MaxAge:   86400 * 7,
			HttpOnly: true,
		}
		fmt.Println(user)
		sess.Values["uID"]       = user["_id"].(primitive.ObjectID).Hex()
		sess.Values["isAdmin"]   = user["admin"]
		// sess.Values["uName"]  = user["name"].(string)
		// sess.Values["uEmail"] = user["email"].(string)
		// fmt.Println(sess.Values)

		sess.Save(c.Request(), c.Response())

		return c.JSON(http.StatusOK, map[string]string{
			"data": "LoggedIn Successfully",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"data": "Already LoggedIn",
	})
	// return c.Redirect(http.StatusMovedPermanently, "/")
	// return c.JSON(http.StatusInternalServerError, e.AuthFailed)
}

func (m *Manager) Logout(c echo.Context) error {
	sess, _ := session.Get("session", c)
	sess.Options.MaxAge = -1
	for k := range sess.Values { delete(sess.Values, k) }
	sess.Save(c.Request(), c.Response())

	return c.NoContent(http.StatusOK)
}

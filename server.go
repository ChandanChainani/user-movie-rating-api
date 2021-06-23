package main

import (
	// "fmt"
	// "reflect"

	"net/http"
	
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"

	"usermovieratingapi/config"
	"usermovieratingapi/db"
	"usermovieratingapi/manager"
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

func isAdminMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
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

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))

	d := db.DB{}

	conf := config.GetConfig()
	d.SetupConnection(&conf)
	d.RunMigrations()

	h := manager.Manager{&d}

	e.POST("/login", h.Login)

	e.GET("/logout", h.Logout)

	e.GET("/movie/:name", h.GetMovieInfoByName)

	g := e.Group("/", AuthMiddleware)

	g.POST("movie/rating", h.SaveMovieRate)

	g.POST("movie/comment", h.SaveMovieComment)

	g.GET("user/rated/movies", h.GetUserRatedMoviesComments)

	a := e.Group("/", isAdminMiddleware)

	a.POST("movie", h.AddMovie)

	e.Logger.Fatal(e.Start(":9001"))
}

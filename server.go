package main

import (
	// "fmt"
	// "reflect"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/go-playground/validator"

	"usermovieratingapi/config"
	"usermovieratingapi/db"
	"usermovieratingapi/manager"
	"usermovieratingapi/custom/validators"
	m "usermovieratingapi/custom/middlewares"
)

func main() {
	e := echo.New()
	e.Validator = &validators.CustomValidator{Validator: validator.New()}
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

	g := e.Group("/", m.AuthMiddleware)

	g.POST("movie/rating", h.SaveMovieRate)

	g.POST("movie/comment", h.SaveMovieComment)

	g.GET("user/rated/movies", h.GetUserRatedMoviesComments)

	a := e.Group("/", m.AdminMiddleware)

	a.POST("movie", h.AddMovie)

	e.Logger.Fatal(e.Start(":9001"))
}

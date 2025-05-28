package main

import (
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"go-auth-api/handlers"
	"go-auth-api/middleware"
	"go-auth-api/models"
	"log"
	"os"
)

func main() {
	models.InitDB()

	e := echo.New()
	e.Use(echoMiddleware.Logger())
	e.Use(echoMiddleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Auth API is running 🚀")
	})

	e.POST("/register", handlers.Register)
	e.POST("/login", handlers.Login)

	r := e.Group("/me")
	r.Use(middleware.JWTMiddleware())
	r.GET("", handlers.Me)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Fatal(e.Start(":" + port))
}

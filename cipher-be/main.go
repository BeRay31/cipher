package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mkamadeus/cipher/handlers/affine"
	"github.com/mkamadeus/cipher/handlers/hill"
	"github.com/mkamadeus/cipher/handlers/playfair"
	"github.com/mkamadeus/cipher/handlers/vigenere"
)

func main() {
	app := echo.New()
	app.Use(middleware.CORS())

	app.GET("/", func(c echo.Context) error {
		return c.JSON(200, "OK")
	})

	affine.Register(app)
	playfair.Register(app)
	vigenere.Register(app)
	hill.Register(app)

	if err := app.Start(":1337"); err != nil {
		log.Fatal(err)
	}
}

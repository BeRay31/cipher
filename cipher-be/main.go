package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/mkamadeus/cipher/handlers/affine"
	"github.com/mkamadeus/cipher/handlers/playfair"
)

func main() {
	app := echo.New()
	affine.Register(app)
	playfair.Register(app)

	if err := app.Start(":1337"); err != nil {
		log.Fatal(err)
	}
}

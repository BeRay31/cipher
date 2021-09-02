package vigenere

import (
	"github.com/labstack/echo/v4"
	"github.com/mkamadeus/cipher/handlers/vigenere/auto"
	"github.com/mkamadeus/cipher/handlers/vigenere/extended"
	"github.com/mkamadeus/cipher/handlers/vigenere/full"
	"github.com/mkamadeus/cipher/handlers/vigenere/standard"
)

func Register(e *echo.Echo) {
	group := e.Group("/vigenere")
	standard.Register(group)
	extended.Register(group)
	auto.Register(group)
	full.Register(group)
}

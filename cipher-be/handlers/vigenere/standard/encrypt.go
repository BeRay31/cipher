package standard

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mkamadeus/cipher/cipher/vigenere/standard"
)

func Encrypt(c echo.Context) error {
	plain := c.QueryParam("plain")
	key := c.QueryParam("key")

	encrypted := standard.Encrypt(plain, key)

	return c.JSON(http.StatusOK, encrypted)
}

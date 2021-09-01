package standard

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mkamadeus/cipher/cipher/vigenere/standard"
)

func Decrypt(c echo.Context) error {
	cipher := c.QueryParam("cipher")
	key := c.QueryParam("key")

	decrypted := standard.Decrypt(cipher, key)

	return c.JSON(http.StatusOK, decrypted)
}

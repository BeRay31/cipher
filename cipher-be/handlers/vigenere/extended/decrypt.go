package extended

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mkamadeus/cipher/cipher/vigenere/extended"
)

func Decrypt(c echo.Context) error {
	cipher := c.QueryParam("cipher")
	key := c.QueryParam("key")

	decrypted := extended.Decrypt([]byte(cipher), []byte(key))

	return c.JSON(http.StatusOK, string(decrypted))
}

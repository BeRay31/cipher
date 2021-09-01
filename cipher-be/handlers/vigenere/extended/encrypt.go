package extended

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mkamadeus/cipher/cipher/vigenere/extended"
)

func Encrypt(c echo.Context) error {
	plain := c.QueryParam("plain")
	key := c.QueryParam("key")

	encrypted := extended.Encrypt([]byte(plain), []byte(key))

	return c.JSON(http.StatusOK, string(encrypted))
}

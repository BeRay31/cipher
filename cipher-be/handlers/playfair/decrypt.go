package playfair

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mkamadeus/cipher/cipher/playfair"
)

func Decrypt(c echo.Context) error {
	cipher := c.QueryParam("cipher")
	key := c.QueryParam("key")

	decrypted, err := playfair.Decrypt(cipher, key)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, decrypted)
}

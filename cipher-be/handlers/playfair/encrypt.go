package playfair

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mkamadeus/cipher/cipher/playfair"
)

func Encrypt(c echo.Context) error {
	plain := c.QueryParam("plain")
	key := c.QueryParam("key")

	encrypted, err := playfair.Encrypt(plain, key)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, encrypted)
}

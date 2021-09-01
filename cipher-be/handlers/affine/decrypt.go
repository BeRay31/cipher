package affine

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/mkamadeus/cipher/cipher/affine"
)

func Decrypt(c echo.Context) error {
	plain := c.QueryParam("cipher")
	m, err := strconv.Atoi(c.QueryParam("m"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	b, err := strconv.Atoi(c.QueryParam("b"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	decrypted := affine.Decrypt(plain, m, b)

	return c.JSON(http.StatusOK, decrypted)
}

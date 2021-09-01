package affine

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/mkamadeus/cipher/cipher/affine"
)

func Encrypt(c echo.Context) error {
	plain := c.QueryParam("plain")
	m, err := strconv.Atoi(c.QueryParam("m"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	b, err := strconv.Atoi(c.QueryParam("b"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	encrypted, err := affine.Encrypt(plain, m, b)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, encrypted)
}

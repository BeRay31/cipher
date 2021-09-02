package affine

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mkamadeus/cipher/cipher/affine"
	"github.com/mkamadeus/cipher/models"
)

func Encrypt(c echo.Context) error {
	body := new(models.AffineRequest)
	err := c.Bind(body)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	content, err := body.Input.ParseContent()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	encrypted, err := affine.Encrypt(content, body.Key, body.Offset)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	payload := &models.AffineResponse{
		BaseResponse: models.BaseResponse{
			Content: encrypted,
		},
		Key:    body.Key,
		Offset: body.Offset,
	}
	return c.JSON(http.StatusOK, payload)
}

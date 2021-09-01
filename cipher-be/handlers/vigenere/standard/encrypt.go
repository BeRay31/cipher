package standard

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mkamadeus/cipher/cipher/vigenere/standard"
	"github.com/mkamadeus/cipher/models"
)

func Encrypt(c echo.Context) error {
	body := new(models.VigenereRequest)
	err := c.Bind(body)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	content, err := body.Input.ParseContent()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	encrypted := standard.Encrypt(content, body.Key)
	payload := &models.VigenereResponse{
		BaseResponse: models.BaseResponse{
			Content: encrypted,
		},
		Key: body.Key,
	}
	return c.JSON(http.StatusOK, payload)
}

package full

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mkamadeus/cipher/cipher/vigenere/full"
	"github.com/mkamadeus/cipher/models"
)

func Decrypt(c echo.Context) error {
	body := new(models.VigenereRequest)
	err := c.Bind(body)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	content, err := body.Input.ParseContent()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	decrypted := full.Decrypt(content, body.Key)
	payload := &models.VigenereResponse{
		BaseResponse: models.BaseResponse{
			Content: decrypted,
		},
		Key: body.Key,
	}
	return c.JSON(http.StatusOK, payload)
}

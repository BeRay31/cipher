package playfair

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mkamadeus/cipher/cipher/playfair"
	"github.com/mkamadeus/cipher/models"
)

func Decrypt(c echo.Context) error {
	body := new(models.PlayfairRequest)
	err := c.Bind(body)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	content, err := body.Input.ParseContent()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	decrypted, err := playfair.Decrypt(content, body.Key)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	payload := &models.PlayfairResponse{
		BaseResponse: models.BaseResponse{
			Content: decrypted,
		},
		Key: body.Key,
	}
	return c.JSON(http.StatusOK, payload)
}

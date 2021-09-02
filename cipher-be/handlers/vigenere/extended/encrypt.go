package extended

import (
	"encoding/base64"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mkamadeus/cipher/cipher/vigenere/extended"
	"github.com/mkamadeus/cipher/models"
)

func Encrypt(c echo.Context) error {
	body := new(models.VigenereExtendedRequest)
	err := c.Bind(body)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	content, err := base64.URLEncoding.DecodeString(body.Content)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	encrypted := extended.Encrypt(content, []byte(body.Key))
	encryptedBase64 := base64.URLEncoding.EncodeToString(encrypted)
	payload := &models.VigenereExtendedResponse{
		BaseResponse: models.BaseResponse{
			Content: encryptedBase64,
		},
		Key: body.Key,
	}
	return c.JSON(http.StatusOK, payload)
}

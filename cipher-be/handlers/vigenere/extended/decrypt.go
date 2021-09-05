package extended

import (
	"encoding/base64"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mkamadeus/cipher/cipher/vigenere/extended"
	"github.com/mkamadeus/cipher/models"
)

func Decrypt(c echo.Context) error {
	body := new(models.VigenereExtendedRequest)
	err := c.Bind(body)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	content, err := base64.StdEncoding.DecodeString(body.Content)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	decrypted := extended.Decrypt(content, []byte(body.Key))
	decryptedBase64 := base64.StdEncoding.EncodeToString(decrypted)

	payload := &models.VigenereExtendedResponse{
		BaseResponse: models.BaseResponse{
			Content: decryptedBase64,
		},
		Key: body.Key,
	}
	return c.JSON(http.StatusOK, payload)
}

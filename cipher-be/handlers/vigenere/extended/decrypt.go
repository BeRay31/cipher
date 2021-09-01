package extended

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mkamadeus/cipher/cipher/vigenere/extended"
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

	decrypted := extended.Decrypt([]byte(content), []byte(body.Key))

	payload := &models.VigenereExtendedResponse{
		Content: decrypted,
		Key:     body.Key,
	}
	return c.JSON(http.StatusOK, payload)
}

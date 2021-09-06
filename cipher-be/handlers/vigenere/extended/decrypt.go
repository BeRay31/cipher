package extended

import (
	"bytes"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mkamadeus/cipher/cipher/vigenere/extended"
)

func Decrypt(c echo.Context) error {
	content, err := c.FormFile("content")
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	key := c.FormValue("key")

	contentSrc, err := content.Open()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	defer contentSrc.Close()

	contentBuffer := bytes.NewBuffer(nil)
	if _, err := io.Copy(contentBuffer, contentSrc); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	decrypted := extended.Decrypt(contentBuffer.Bytes(), []byte(key))
	return c.Blob(http.StatusOK, "application/octet-stream", decrypted)
}

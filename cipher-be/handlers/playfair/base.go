package playfair

import "github.com/labstack/echo/v4"

func Register(e *echo.Echo) {
	group := e.Group("/playfair")
	group.POST("/encrypt", Encrypt)
	group.POST("/decrypt", Decrypt)
}

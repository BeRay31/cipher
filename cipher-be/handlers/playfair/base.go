package playfair

import "github.com/labstack/echo/v4"

func Register(e *echo.Echo) {
	group := e.Group("/playfair")
	group.GET("/encrypt", Encrypt)
	group.GET("/decrypt", Decrypt)
}

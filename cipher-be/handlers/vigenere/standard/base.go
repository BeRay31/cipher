package standard

import "github.com/labstack/echo/v4"

func Register(group *echo.Group) {
	group.GET("/standard/encrypt", Encrypt)
	group.GET("/standard/decrypt", Decrypt)
}

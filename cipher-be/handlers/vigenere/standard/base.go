package standard

import "github.com/labstack/echo/v4"

func Register(group *echo.Group) {
	group.POST("/standard/encrypt", Encrypt)
	group.POST("/standard/decrypt", Decrypt)
}

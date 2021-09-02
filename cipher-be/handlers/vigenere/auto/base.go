package auto

import "github.com/labstack/echo/v4"

func Register(group *echo.Group) {
	group.POST("/auto/encrypt", Encrypt)
	group.POST("/auto/decrypt", Decrypt)
}

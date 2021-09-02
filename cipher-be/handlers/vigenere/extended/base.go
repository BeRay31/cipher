package extended

import "github.com/labstack/echo/v4"

func Register(group *echo.Group) {
	group.POST("/extended/encrypt", Encrypt)
	group.POST("/extended/decrypt", Decrypt)
}

package extended

import "github.com/labstack/echo/v4"

func Register(group *echo.Group) {
	group.GET("/extended/encrypt", Encrypt)
	group.GET("/extended/decrypt", Decrypt)
}

package full

import "github.com/labstack/echo/v4"

func Register(group *echo.Group) {
	group.POST("/full/encrypt", Encrypt)
	group.POST("/full/decrypt", Decrypt)
}

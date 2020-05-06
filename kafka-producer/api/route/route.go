package route

import (
	"kafka-example/helper"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

// NewRoute Handler
type NewRoute struct {
	E           *echo.Echo
	DBOms       *gorm.DB
	DBOmsMaster *gorm.DB
	Helper      helper.NewHelper
}

// Register ...
func (t *NewRoute) Register() {
	group := t.E.Group("api/v1")
	t.KafkaRoute(group)
}

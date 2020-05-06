package route

import (
	"kafka-example/api/handler"

	"github.com/labstack/echo"
)

// KafkaRoute ...
func (t *NewRoute) KafkaRoute(g *echo.Group) {
	handler := handler.NewKafkaHandler{
		Helper: t.Helper,
	}
	g.GET("/test", handler.Test)
}

package http

// TODO arma el *gin.Engine y asigna rutas a los handlers
import (
	"github.com/gin-gonic/gin"
)

func NewRouter(ordersHandler *OrdersHandler) *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	r.GET("/orders/:order_id", ordersHandler.FindOrderById)
	r.POST("/orders", ordersHandler.SaveOrder)

	return r
}

package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mjeffers-998/interview-task/controllers"
)

func CreateRoutes(r *gin.Engine) {
	r.POST("/orders", controllers.CreateOrder)
	r.PUT("/orders/:id/cancel", controllers.CancelOrder)
	r.GET("/orders/:id", controllers.GetOrderByID)
	r.GET("/orders", controllers.ListAllOrders)
	r.GET("/orders/revenue", controllers.GetRevenue)
}

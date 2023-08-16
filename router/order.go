package router

import (
	"usecase_test/handler"
	"usecase_test/repository"
	"usecase_test/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewOrderRouter(v1 *gin.RouterGroup, db *gorm.DB) *gin.RouterGroup {
	orderRepository := repository.NewOrderRepository(db)
	orderService := service.NewOrderService(orderRepository)
	orderHandler := handler.NewOrderHandler(orderService)

	v1.GET("/orders", orderHandler.FindAll)
	v1.GET("/orders/:id", orderHandler.FindByID)
	v1.POST("/orders", orderHandler.Create)

	return v1
}

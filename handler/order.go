package handler

import (
	"net/http"
	"strconv"
	"usecase_test/service"
	"usecase_test/model"
	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	orderService service.OrderService
}

func NewOrderHandler(orderService service.OrderService) OrderHandler {
	return OrderHandler{orderService}
}

func (h *OrderHandler) FindAll(c *gin.Context) {
	orders, err := h.orderService.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, orders)
}

func (h *OrderHandler) FindByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order, err := h.orderService.FindByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, order)
}

func (h *OrderHandler) Create(c *gin.Context) {
	var order model.CreateOrderRequest
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.orderService.Create(order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Order created successfully",
	})
}
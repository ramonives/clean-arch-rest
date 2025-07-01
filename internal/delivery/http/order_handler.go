package http

import (
	"net/http"

	"clean-arch-rest/internal/domain"

	"github.com/gin-gonic/gin"
)

// OrderHandler gerencia as requisições HTTP relacionadas a orders
type OrderHandler struct {
	orderUseCase domain.OrderUseCase
}

func NewOrderHandler(orderUseCase domain.OrderUseCase) *OrderHandler {
	return &OrderHandler{
		orderUseCase: orderUseCase,
	}
}

// CreateOrder cria uma nova order via HTTP
func (h *OrderHandler) CreateOrder(c *gin.Context) {
	var order domain.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.orderUseCase.CreateOrder(&order); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, order)
}

// ListOrders retorna todas as orders via HTTP
func (h *OrderHandler) ListOrders(c *gin.Context) {
	orders, err := h.orderUseCase.ListOrders()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, orders)
}

package http

import (
	"clean-arch-rest/internal/domain"

	"github.com/gin-gonic/gin"
)

// SetupRouter configura as rotas HTTP da aplicação
func SetupRouter(orderUseCase domain.OrderUseCase) *gin.Engine {
	r := gin.Default()

	orderHandler := NewOrderHandler(orderUseCase)

	// Rotas para orders
	orders := r.Group("/order")
	{
		orders.POST("/", orderHandler.CreateOrder)
		orders.GET("/", orderHandler.ListOrders)
	}

	return r
}

package domain

import (
	"time"
)

// Order representa uma ordem/pedido no sistema
type Order struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	CustomerID uint      `json:"customer_id"`
	Total      float64   `json:"total"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type OrderRepository interface {
	Create(order *Order) error
	List() ([]Order, error)
}

type OrderUseCase interface {
	CreateOrder(order *Order) error
	ListOrders() ([]Order, error)
}

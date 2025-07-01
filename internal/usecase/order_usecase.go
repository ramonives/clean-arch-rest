package usecase

import (
	"clean-arch-rest/internal/domain"
)

// OrderUseCase implementa a lógica de negócio para orders
type OrderUseCase struct {
	orderRepo domain.OrderRepository
}

func NewOrderUseCase(orderRepo domain.OrderRepository) *OrderUseCase {
	return &OrderUseCase{
		orderRepo: orderRepo,
	}
}

// CreateOrder cria uma nova order no sistema
func (uc *OrderUseCase) CreateOrder(order *domain.Order) error {
	return uc.orderRepo.Create(order)
}

// ListOrders retorna todas as orders cadastradas
func (uc *OrderUseCase) ListOrders() ([]domain.Order, error) {
	return uc.orderRepo.List()
}

package repository

import (
	"clean-arch-rest/internal/domain"

	"gorm.io/gorm"
)

// OrderRepository implementa o repositório de orders usando GORM
type OrderRepository struct {
	db *gorm.DB
}

// NewOrderRepository cria uma nova instância do repositório
func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

// Create salva uma nova order no banco
func (r *OrderRepository) Create(order *domain.Order) error {
	return r.db.Create(order).Error
}

// List retorna todas as orders do banco
func (r *OrderRepository) List() ([]domain.Order, error) {
	var orders []domain.Order
	err := r.db.Find(&orders).Error
	return orders, err
}

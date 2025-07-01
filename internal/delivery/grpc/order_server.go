package grpc

import (
	"context"
	"time"

	"clean-arch-rest/internal/domain"
	pb "clean-arch-rest/proto/order"
)

// OrderServer implementa o servidor gRPC para orders
type OrderServer struct {
	pb.UnimplementedOrderServiceServer
	orderUseCase domain.OrderUseCase
}

// NewOrderServer cria uma nova instância do servidor gRPC
func NewOrderServer(orderUseCase domain.OrderUseCase) *OrderServer {
	return &OrderServer{
		orderUseCase: orderUseCase,
	}
}

// CreateOrder cria uma nova order via gRPC
func (s *OrderServer) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	order := &domain.Order{
		CustomerID: uint(req.CustomerId),
		Total:      req.Total,
		Status:     req.Status,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	if err := s.orderUseCase.CreateOrder(order); err != nil {
		return nil, err
	}

	return &pb.CreateOrderResponse{
		Id:         uint32(order.ID),
		CustomerId: uint32(order.CustomerID),
		Total:      order.Total,
		Status:     order.Status,
		CreatedAt:  order.CreatedAt.Format(time.RFC3339),
		UpdatedAt:  order.UpdatedAt.Format(time.RFC3339),
	}, nil
}

// ListOrders retorna todas as orders via gRPC
func (s *OrderServer) ListOrders(ctx context.Context, req *pb.ListOrdersRequest) (*pb.ListOrdersResponse, error) {
	orders, err := s.orderUseCase.ListOrders()
	if err != nil {
		return nil, err
	}

	var pbOrders []*pb.CreateOrderResponse
	for _, order := range orders {
		pbOrder := &pb.CreateOrderResponse{
			Id:         uint32(order.ID),
			CustomerId: uint32(order.CustomerID),
			Total:      order.Total,
			Status:     order.Status,
			CreatedAt:  order.CreatedAt.Format(time.RFC3339),
			UpdatedAt:  order.UpdatedAt.Format(time.RFC3339),
		}
		pbOrders = append(pbOrders, pbOrder)
	}

	return &pb.ListOrdersResponse{
		Orders: pbOrders,
	}, nil
}

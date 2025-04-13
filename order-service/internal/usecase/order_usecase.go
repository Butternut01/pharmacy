package usecase

import (
    "github.com/Butternut01/order-service/internal/entity"
    "github.com/Butternut01/order-service/internal/repository"
)

type OrderUseCase interface {
    CreateOrder(order *entity.Order) error
    GetOrder(id string) (*entity.Order, error)
    UpdateOrderStatus(id string, status entity.OrderStatus) error
    ListOrders(filter entity.OrderFilter) ([]entity.Order, error)
}

type orderUseCase struct {
    orderRepo repository.OrderRepository
}

func NewOrderUseCase(orderRepo repository.OrderRepository) OrderUseCase {
    return &orderUseCase{
        orderRepo: orderRepo,
    }
}

func (uc *orderUseCase) CreateOrder(order *entity.Order) error {
    return uc.orderRepo.Create(order)
}

func (uc *orderUseCase) GetOrder(id string) (*entity.Order, error) {
    return uc.orderRepo.FindByID(id)
}

func (uc *orderUseCase) UpdateOrderStatus(id string, status entity.OrderStatus) error {
    return uc.orderRepo.UpdateStatus(id, status)
}

func (uc *orderUseCase) ListOrders(filter entity.OrderFilter) ([]entity.Order, error) {
    return uc.orderRepo.FindAll(filter)
}
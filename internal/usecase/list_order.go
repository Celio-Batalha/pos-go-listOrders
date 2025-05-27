// filepath: /home/batalha/go/src/pos-go-listOrders/internal/usecase/list_order.go
package usecase

import "github.com/Celio-Batalha/pos-go-listOrders/internal/entity"

type ListOrderUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewListOrderUseCase(
	orderRepository entity.OrderRepositoryInterface,
) *ListOrderUseCase {
	return &ListOrderUseCase{
		OrderRepository: orderRepository,
	}
}

func (c *ListOrderUseCase) Execute() ([]*entity.Order, error) {
	orders, err := c.OrderRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return orders, nil
}

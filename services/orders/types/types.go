package types

import (
	"context"
	"github.com/fawwazmw/go-kitchen-api/services/common/genproto/orders"
)

type OrderService interface {
	CreateOrder(context.Context, *orders.Order) error
	GetOrders(ctx context.Context) []*orders.Order
}

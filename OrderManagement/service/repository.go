package service

import (
	"context"
	"github.com/akbarshoh/microOLX/OrderManagement/proto"
)

type Repository interface {
	GetOrder(ctx context.Context, request proto.Request) (proto.OrderResponse, error)
	Payment(ctx context.Context, request proto.Request) error
	Choose(ctx context.Context)
}

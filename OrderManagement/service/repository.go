package service

import (
	"context"
	"github.com/akbarshoh/microOLX/proto"
)

type Repository interface {
	GetOrder(ctx context.Context, request proto.Request) (proto.Order, error)
	Payment(ctx context.Context, request proto.Request) error
	Choose(ctx context.Context, request proto.Request) error
	MealList(ctx context.Context, request proto.Request) (proto.Meals, error)
	OrderList(ctx context.Context, request proto.Request) (proto.Orders, error)
}

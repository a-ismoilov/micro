package service

import (
	"context"
	"github.com/akbarshoh/microOLX/protos/orderproto"
)

type Repository interface {
	GetOrder(ctx context.Context, request orderproto.Request) (orderproto.Order, error)
	Payment(ctx context.Context, request orderproto.Request) (int, error)
	Choose(ctx context.Context, request orderproto.MealChoice) error
	MealList(ctx context.Context, request orderproto.Request) (orderproto.Meals, error)
	OrderList(ctx context.Context, request orderproto.Request) (orderproto.Orders, error)
	Cancel(context.Context, *orderproto.Request) error
	NewMeal(context.Context, *orderproto.Meal) error
	UpdateMeal(context.Context, *orderproto.Meal) error
}

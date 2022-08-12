package server

import (
	"context"
	"github.com/akbarshoh/microOLX/protos/orderproto"
)

type Service interface {
	GetOrder(context.Context, *orderproto.Request) (*orderproto.Order, error)
	Choose(context.Context, *orderproto.MealChoice) error
	Payment(context.Context, *orderproto.Request) error
	MealList(context.Context, *orderproto.Request) (*orderproto.Meals, error)
	OrderList(context.Context, *orderproto.Request) (*orderproto.Orders, error)
	Cancel(context.Context, *orderproto.Request) error
	NewMeal(context.Context, *orderproto.Meal) error
	UpdateMeal(context.Context, *orderproto.Meal) error
}

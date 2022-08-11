package service

import (
	"context"
	"github.com/akbarshoh/microOLX/protos/orderproto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	r Repository
}

func New(r Repository) Service {
	return Service{
		r: r,
	}
}

func (Service) GetOrder(context.Context, *orderproto.Request) (*orderproto.Order, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrder not implemented")
}
func (Service) Choose(context.Context, *orderproto.MealChoice) error {
	return status.Errorf(codes.Unimplemented, "method Choose not implemented")
}
func (Service) Payment(context.Context, *orderproto.Request) error {
	return status.Errorf(codes.Unimplemented, "method Payment not implemented")
}
func (Service) MealList(context.Context, *orderproto.Request) (*orderproto.Meals, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MealList not implemented")
}
func (Service) OrderList(context.Context, *orderproto.Request) (*orderproto.Orders, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OrderList not implemented")
}
func (Service) Cancel(context.Context, *orderproto.Request) error {
	return status.Errorf(codes.Unimplemented, "method Cancel not implemented")
}
func (Service) NewMeal(context.Context, *orderproto.Meal) error {
	return status.Errorf(codes.Unimplemented, "method NewMeal not implemented")
}
func (Service) NewOrder(context.Context, *orderproto.Order) error {
	return status.Errorf(codes.Unimplemented, "method NewOrder not implemented")
}

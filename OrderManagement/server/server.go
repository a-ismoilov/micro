package server

import (
	"context"
	"github.com/akbarshoh/microOLX/protos/orderproto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	orderproto.UnimplementedOrderServiceServer
	service Service
}

func New(s Service) *Server {
	return &Server{
		service: s,
	}
}

func (s *Server) GetOrder(context.Context, *orderproto.Request) (*orderproto.Order, error) {
	return nil, nil
}
func (s *Server) Choose(context.Context, *orderproto.MealChoice) (*orderproto.OK, error) {
	return nil, nil
}
func (s *Server) Payment(context.Context, *orderproto.Request) (*orderproto.OK, error) {
	return nil, nil
}
func (s *Server) MealList(context.Context, *orderproto.Request) (*orderproto.Meals, error) {
	return nil, nil
}
func (s *Server) OrderList(context.Context, *orderproto.Request) (*orderproto.Orders, error) {
	return nil, nil
}

func (s *Server) Cancel(context.Context, *orderproto.Request) (*orderproto.OK, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Cancel not implemented")
}
func (s *Server) NewMeal(context.Context, *orderproto.Meal) (*orderproto.OK, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewMeal not implemented")
}
func (s *Server) UpdateMeal(context.Context, *orderproto.Meal) (*orderproto.OK, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewOrder not implemented")
}

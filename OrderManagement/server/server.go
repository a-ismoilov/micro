package server

import (
	"context"
	"github.com/akbarshoh/microOLX/proto"
)

type Server struct {
	proto.UnimplementedOrderServiceServer
}

func (s *Server) GetOrder(context.Context, *proto.Request) (*proto.Order, error) {
	return nil, nil
}
func (s *Server) Choose(context.Context, *proto.MealChoice) (*proto.OK, error) {
	return nil, nil
}
func (s *Server) Payment(context.Context, *proto.Request) (*proto.OK, error) {
	return nil, nil
}
func (s *Server) MealList(context.Context, *proto.Request) (*proto.Meals, error) {
	return nil, nil
}
func (s *Server) OrderList(context.Context, *proto.Request) (*proto.Orders, error) {
	return nil, nil
}

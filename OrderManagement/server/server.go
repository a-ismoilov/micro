package server

import (
	"context"
	"github.com/akbarshoh/microOLX/protos/orderproto"
	"time"
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

func (s *Server) GetOrder(ctx context.Context, request *orderproto.Request) (*orderproto.Order, error) {
	c, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()
	o, err := s.service.GetOrder(c, request)
	if err != nil {
		return nil, err
	}
	return o, nil
}
func (s *Server) Choose(ctx context.Context, request *orderproto.MealChoice) (*orderproto.OK, error) {
	c, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()
	if err := s.service.Choose(c, request); err != nil {
		return nil, err
	}
	ok := orderproto.OK{OK: 1}
	return &ok, nil
}
func (s *Server) Payment(ctx context.Context, request *orderproto.Request) (*orderproto.OK, error) {
	c, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()
	price, err := s.service.Payment(c, request)
	if err != nil {
		return nil, err
	}
	ok := orderproto.OK{OK: int32(price)}
	return &ok, nil
}
func (s *Server) MealList(ctx context.Context, request *orderproto.Request) (*orderproto.Meals, error) {
	c, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()
	ms, err := s.service.MealList(c, request)
	if err != nil {
		return nil, err
	}
	return ms, nil
}
func (s *Server) OrderList(ctx context.Context, request *orderproto.Request) (*orderproto.Orders, error) {
	c, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()
	os, err := s.service.OrderList(c, request)
	if err != nil {
		return nil, err
	}
	return os, nil
}

func (s *Server) Cancel(ctx context.Context, request *orderproto.Request) (*orderproto.OK, error) {
	c, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()
	if err := s.service.Cancel(c, request); err != nil {
		return nil, err
	}
	ok := orderproto.OK{OK: 1}
	return &ok, nil
}
func (s *Server) NewMeal(ctx context.Context, request *orderproto.Meal) (*orderproto.OK, error) {
	c, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()
	if err := s.service.NewMeal(c, request); err != nil {
		return nil, err
	}
	ok := orderproto.OK{OK: 1}
	return &ok, nil
}
func (s *Server) UpdateMeal(ctx context.Context, request *orderproto.Meal) (*orderproto.OK, error) {
	c, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()
	if err := s.service.UpdateMeal(c, request); err != nil {
		return nil, err
	}
	ok := orderproto.OK{OK: 1}
	return &ok, nil
}

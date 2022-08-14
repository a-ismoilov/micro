package service

import (
	"context"
	"github.com/akbarshoh/microOLX/protos/orderproto"
	"time"
)

type Service struct {
	r Repository
}

func New(r Repository) Service {
	return Service{
		r: r,
	}
}

func (s Service) GetOrder(ctx context.Context, request *orderproto.Request) (*orderproto.Order, error) {
	c, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()
	o, err := s.r.GetOrder(c, *request)
	if err != nil {
		return &o, nil
	}
	return &o, nil
}
func (s Service) Choose(ctx context.Context, request *orderproto.MealChoice) error {
	c, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()
	if err := s.r.Choose(c, *request); err != nil {
		return err
	}
	return nil
}
func (s Service) Payment(ctx context.Context, request *orderproto.Request) (int, error) {
	c, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()
	price, err := s.r.Payment(c, *request)
	if err != nil {
		return 0, err
	}
	return price, nil
}
func (s Service) MealList(ctx context.Context, request *orderproto.Request) (*orderproto.Meals, error) {
	c, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()
	ms, err := s.r.MealList(c, *request)
	if err != nil {
		return nil, err
	}
	return &ms, nil
}
func (s Service) OrderList(ctx context.Context, request *orderproto.Request) (*orderproto.Orders, error) {
	c, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()
	os, err := s.r.OrderList(c, *request)
	if err != nil {
		return nil, err
	}
	return &os, nil
}
func (s Service) Cancel(ctx context.Context, request *orderproto.Request) error {
	c, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()
	if err := s.r.Cancel(c, request); err != nil {
		return err
	}
	return nil
}
func (s Service) NewMeal(ctx context.Context, request *orderproto.Meal) error {
	c, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()
	if err := s.r.NewMeal(c, request); err != nil {
		return err
	}
	return nil
}
func (s Service) UpdateMeal(ctx context.Context, request *orderproto.Meal) error {
	c, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()
	if err := s.r.UpdateMeal(c, request); err != nil {
		return err
	}
	return nil
}

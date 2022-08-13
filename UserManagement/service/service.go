package service

import (
	"context"
	"github.com/akbarshoh/microOLX/protos/userproto"
	"time"
)

type Service struct {
	R Repository
}

func New(r Repository) Service {
	return Service{
		R: r,
	}
}

func (s *Service) Log(ctx context.Context, request *userproto.User) error {
	c, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()
	if err := s.R.Log(c, request); err != nil {
		return err
	}
	return nil
}
func (s *Service) Payment(ctx context.Context, request *userproto.PayRequest) error {
	c, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()
	if err := s.R.Payment(c, request); err != nil {
		return err
	}
	return nil
}
func (s *Service) UserList(ctx context.Context, request *userproto.Admin) (*userproto.Users, error) {
	c, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()
	us, err := s.R.UserList(c, request)
	if err != nil {
		return nil, err
	}
	return us, nil
}
func (s *Service) LogAdmin(ctx context.Context, request *userproto.Admin) error {
	c, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()
	if err := s.R.LogAdmin(c, request); err != nil {
		return err
	}
	return nil
}

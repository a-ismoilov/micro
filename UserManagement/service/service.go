package service

import (
	"context"
	"github.com/akbarshoh/microOLX/protos/userproto"
)

type Service struct {
	R Repository
}

func New(r Repository) Service {
	return Service{
		R: r,
	}
}

func (s *Service) Log(context.Context, *userproto.User) error {
	return nil
}
func (s *Service) Payment(context.Context, *userproto.PayRequest) error {
	return nil
}
func (s *Service) UserList(context.Context, *userproto.Admin) (*userproto.Users, error) {
	return nil, nil
}
func (s *Service) LogAdmin(context.Context, *userproto.Admin) error {
	return nil
}

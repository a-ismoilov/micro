package service

import (
	"context"
	"github.com/akbarshoh/microOLX/protos/userproto"
)

type Repository interface {
	Log(ctx context.Context, request *userproto.User) error
	LogAdmin(context.Context, *userproto.Admin) error
	UserList(context.Context, *userproto.Admin) (*userproto.Users, error)
	Payment(context.Context, *userproto.PayRequest) error
}

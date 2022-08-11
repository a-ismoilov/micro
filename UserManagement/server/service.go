package server

import (
	"context"
	"github.com/akbarshoh/microOLX/protos/userproto"
)

type Service interface {
	Log(ctx context.Context, request userproto.User) (userproto.Response, error)
	LogAdmin(context.Context, *userproto.Admin) (*userproto.Response, error)
	UserList(context.Context, *userproto.Admin) (*userproto.Users, error)
	Payment(context.Context, *userproto.PayRequest) (*userproto.Response, error)
}

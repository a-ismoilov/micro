package server

import (
	"context"
	proto2 "github.com/akbarshoh/microOLX/protos/userproto"
	"github.com/akbarshoh/microOLX/service"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
)

type Server struct {
	proto2.UnimplementedUserServiceServer
	service service.Service
}

func New(s service.Service) *Server {
	return &Server{
		service: s,
	}
}

func (s *Server) Log(context.Context, *proto2.User) (*proto2.Response, error) {
	return nil, nil
}
func (s *Server) Payment(context.Context, *proto2.PayRequest) (*proto2.Response, error) {
	return nil, nil
}
func (s *Server) UserList(context.Context, *proto2.Admin) (*proto2.Users, error) {
	return nil, nil
}
func (s *Server) LogAdmin(context.Context, *proto2.Admin) (*proto2.Response, error) {
	return nil, nil
}

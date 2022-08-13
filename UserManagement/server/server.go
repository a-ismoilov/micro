package server

import (
	"context"
	proto2 "github.com/akbarshoh/microOLX/protos/userproto"
	"github.com/akbarshoh/microOLX/service"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	"log"
	"time"
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

var (
	ok proto2.Response
)

func (s *Server) Log(ctx context.Context, request *proto2.User) (*proto2.Response, error) {
	log.Println("keldik mana")
	c, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()
	if err := s.service.Log(c, request); err != nil {
		return nil, err
	}
	ok.OK = "status ok"
	return &ok, nil
}
func (s *Server) Payment(ctx context.Context, request *proto2.PayRequest) (*proto2.Response, error) {
	c, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()
	if err := s.service.Payment(c, request); err != nil {
		return nil, err
	}
	ok.OK = "status ok"
	return &ok, nil
}
func (s *Server) UserList(ctx context.Context, request *proto2.Admin) (*proto2.Users, error) {
	c, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()
	us, err := s.service.UserList(c, request)
	if err != nil {
		return nil, err
	}
	return us, nil
}
func (s *Server) LogAdmin(ctx context.Context, request *proto2.Admin) (*proto2.Response, error) {
	c, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()
	if err := s.service.LogAdmin(c, request); err != nil {
		return nil, err
	}
	ok.OK = "status ok"
	return &ok, nil
}

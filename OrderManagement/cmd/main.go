package main

import (
	"github.com/akbarshoh/microOLX/config"
	"github.com/akbarshoh/microOLX/protos/orderproto"
	"github.com/akbarshoh/microOLX/repository/postgres"
	"github.com/akbarshoh/microOLX/server"
	service2 "github.com/akbarshoh/microOLX/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	cfg := config.Load()
	db, err := postgres.ConnectDB(cfg)
	if err != nil {
		log.Fatalln(err)
	}
	lis, err := net.Listen("tcp", cfg.Port)
	if err != nil {
		log.Fatalln(err)
	}

	s := grpc.NewServer()
	orderproto.RegisterOrderServiceServer(s, server.New(service2.New(postgres.New(db))))

	if err := s.Serve(lis); err != nil {
		log.Fatalln(err, "bor")
	}
}

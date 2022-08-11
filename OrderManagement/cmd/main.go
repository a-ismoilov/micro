package main

import (
	"github.com/akbarshoh/microOLX/config"
	"github.com/akbarshoh/microOLX/proto"
	"github.com/akbarshoh/microOLX/repository/postgres"
	"github.com/akbarshoh/microOLX/server"
	"github.com/akbarshoh/microOLX/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	db, err := postgres.ConnectDB(config.Load())
	if err != nil {
		log.Fatalln(err, "ha")
	}
	p := postgres.PostgresRepository{
		DB: db,
	}
	service.Handle(p)

	lis, err := net.Listen("tcp", ":6666")
	if err != nil {
		log.Fatalln(err, "yoq")
	}

	s := grpc.NewServer()

	proto.RegisterOrderServiceServer(s, &server.Server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalln(err, "bor")
	}
}

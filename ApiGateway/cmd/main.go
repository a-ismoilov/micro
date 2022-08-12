package main

import (
	"github.com/akbarshoh/microOLX/api"
	"github.com/akbarshoh/microOLX/api/handlers/orderhandler"
	"github.com/akbarshoh/microOLX/api/handlers/userhandler"
	"github.com/akbarshoh/microOLX/protos/orderproto"
	"github.com/akbarshoh/microOLX/protos/userproto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	conn, err := grpc.Dial(":1111", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln(err)
	}
	o := orderhandler.New(orderproto.NewOrderServiceClient(conn))
	conn, err = grpc.Dial(":2222", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln(err)
	}
	u := userhandler.New(userproto.NewUserServiceClient(conn))
	api.Handle(u, o)
}

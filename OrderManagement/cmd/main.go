package main

import (
	"github.com/akbarshoh/microOLX/OrderManagement/config"
	"github.com/akbarshoh/microOLX/OrderManagement/repository/postgres"
	"github.com/akbarshoh/microOLX/OrderManagement/service"
	"log"
)

func main() {
	db, err := postgres.ConnectDB(config.Load())
	if err != nil {
		log.Fatalln("can't connect db")
	}
	p := postgres.Repository{
		DB: db,
	}
	service.Handle(p)
}

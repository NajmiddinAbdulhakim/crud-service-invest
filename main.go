package main

import (
	"log"
	"net"
	pb "github.com/NajmiddinAbdulhakim/iman/crud-service/genproto"

	"github.com/NajmiddinAbdulhakim/iman/crud-service/config"
	"github.com/NajmiddinAbdulhakim/iman/crud-service/pkg/db"
	"github.com/NajmiddinAbdulhakim/iman/crud-service/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	cfg := config.Load()

	connDB, err := db.ConnectDB(cfg)
	if err != nil {
		log.Fatal(`sqlx connection to postgers error: `, err)
	}

	service := service.NewService(connDB)

	lis, err := net.Listen("tcp", cfg.RPCPort)
	if err != nil {
		log.Fatal(`Error while listening: `, err)
	}

	s := grpc.NewServer()
	pb.RegisterCRUDServiceServer(s, service)

	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatal(`Error while listening :`, err)
	}

}

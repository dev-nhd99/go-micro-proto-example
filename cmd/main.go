package main

import (
	"context"
	_ "fmt"
	"log"

	pb "github.com/go-micro-proto-example/pkg/proto/grpc"
	"go-micro.dev/v4"
)


type Server struct {
}


func (s *Server) Sum(ctx context.Context, req *pb.CalculatorReq,res *pb.CalculatorRes) ( error) {
	res.Sum = int64(req.GetA())+ int64(req.GetB())
	return nil
}

func main() {
	// init server
	service := micro.NewService(
		micro.Name("Calculator"),
	)

	service.Init()
	// init grpc
	pb.RegisterCalculatorHandler(service.Server(),new(Server))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
package main

import (
	"context"
	_ "fmt"
	"log"
	"net/http"

	pb "github.com/go-micro-proto-example/pkg/proto/grpc"
	_ "github.com/go-micro/plugins/v4/registry/kubernetes"
	"github.com/labstack/echo/v4"
	"go-micro.dev/v4"
)


type Server struct {
}


func (s *Server) Sum(ctx context.Context, req *pb.CalculatorReq,res *pb.CalculatorRes) ( error) {
	res.Sum = int64(req.GetA())+ int64(req.GetB())
	return nil
}

func main() {
	//-- Start server 
	e := echo.New()
	e.GET("/check", func(c echo.Context) error {
		return c.String(http.StatusOK, "Go Micro Proto Example!")
	})

	go func() {
		if err := e.Start(":3030"); err != nil {
			e.Logger.Info("Shutting down the server:%v\n", err)
		}
	}()

	//-- Start grpc service
	// registry := kubernetes.NewRegistry()
	service := micro.NewService(
		micro.Name("calculator.service"),
		// micro.Registry(registry),
	)

	service.Init()
	pb.RegisterCalculatorHandler(service.Server(),new(Server))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
package main

import (
	"api-gateway/config"
	"api-gateway/handlers"
	"api-gateway/proto"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("config.LoadConfig failed: %v", err)
	}

	router := gin.Default()

	employeeConn, err := grpc.NewClient(cfg.Address+cfg.EmployeePort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect to employee service: %v", err)
	}
	defer employeeConn.Close()
	employeeClient := proto.NewEmployeeServiceClient(employeeConn)

	Handler := handlers.NewHandler(employeeClient)

	router.POST("/employees", Handler.AddEmployee)
	router.DELETE("/employees", Handler.RemoveEmployee)
	router.GET("/employees", Handler.GetEmployees)
	router.PUT("/employees", Handler.UpdateEmployee)

	log.Printf("Gateway service is listening on port %s", cfg.GatewayPort)
	if err := router.Run(cfg.GatewayPort); err != nil {
		log.Fatalf("could not start gateway service: %v", err)
	}
}

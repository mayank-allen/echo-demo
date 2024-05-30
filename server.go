package main

import (
	v1 "echo-demo/clients/userclient"
	"echo-demo/handlers"
	"github.com/labstack/echo"
	"google.golang.org/grpc"
	"log"
)

func main() {
	// Initialize Echo
	e := echo.New()
	// Create gRPC client connection
	conn, err := grpc.Dial("localhost:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)

	// Initialize gRPC client
	handlers.UserClient = v1.NewUserClient(conn)
	// Define route
	e.GET("/user/:id", handlers.GetUsers)
	e.GET("/users", handlers.GetAllUsers)
	e.POST("/user", handlers.CreatUser)
	e.PUT("/user", handlers.UpdateUser)
	e.DELETE("user/:id", handlers.DeleteUser)

	// Start the Echo server
	err = e.Start(":8080")
	if err != nil {
		return
	}
}

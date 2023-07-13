package main

import (
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	netListen, err := net.Listen(os.Getenv("APP_PROTOCOL"), ":"+os.Getenv("APP_PORT"))
	if err != nil {
		log.Fatalf("Could not listen to %s port: %v", os.Getenv("APP_PORT"), err)
	}

	gprcServer := grpc.NewServer()
	log.Printf("Starting %s server on %s port", os.Getenv("APP_PROTOCOL"), os.Getenv("APP_PORT"))
	if gprcServer.Serve(netListen); err != nil {
		log.Fatalf("Could not serve: %v", err)
	}
	log.Printf("Server Started")
}

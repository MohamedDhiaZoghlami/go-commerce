package main

import (
	"context"
	"fmt"
	"log"

	"github.com/MohamedDhiaZoghlami/go-commerce/storage/postgres"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
	logger := logrus.New()
	_, err = postgres.Open(context.Background(), logger)
	if err != nil {
		log.Fatalf("Error Opening DB connection : %s", err)
	}
	fmt.Println("Its running")
}

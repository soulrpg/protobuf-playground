package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
	"playground/agify"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	agifyClient := agify.NewAgifyClient(conn)

	args := os.Args[1:]
	if len(args) < 2 {
		fmt.Printf("you have to provide name and countryId as command line arguments")
		os.Exit(1)
	}
	name := args[0]
	countryId := args[1]
	person := &agify.Person{Name: name, CountryId: countryId}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	age, err := agifyClient.GetEstimatedAge(ctx, person)
	if err != nil {
		log.Fatalf("could not get estimated age: %v", err)
	}
	fmt.Printf("Estimated age: %v\n", age.Age)

	count, err := agifyClient.GetCount(ctx, person)
	if err != nil {
		log.Fatalf("could not get count: %v", err)
	}
	fmt.Printf("Count: %v\n", count.Count)
}

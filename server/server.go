package main

import (
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"playground/agify"
)

const apiUrl = "https://api.agify.io/"

type AgifyServer struct {
	agify.UnimplementedAgifyServer
}

type AgifyResponse struct {
	Count     int32  `json:"count"`
	Name      string `json:"name"`
	Age       int32  `json:"age"`
	CountryId string `json:"country_id`
}

func (s *AgifyServer) GetEstimatedAge(ctx context.Context, person *agify.Person) (*agify.Age, error) {
	agifyData, err := getAgifyData(person)
	if err != nil {
		fmt.Printf("could not get agify data")
		os.Exit(1)
	}
	age := &agify.Age{Age: agifyData.Age}
	return age, nil
}

func (s *AgifyServer) GetCount(ctx context.Context, person *agify.Person) (*agify.Count, error) {
	agifyData, err := getAgifyData(person)
	if err != nil {
		fmt.Printf("could not get agify data")
		os.Exit(1)
	}
	count := &agify.Count{Count: agifyData.Count}
	return count, nil
}

func getAgifyData(person *agify.Person) (*AgifyResponse, error) {
	requestURL := fmt.Sprintf("%s?name=%s&country_id=%s", apiUrl, person.Name, person.CountryId)
	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		return nil, err
	}

	fmt.Printf("client: got response!\n")
	fmt.Printf("client: status code: %d\n", res.StatusCode)

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		return nil, err
	}
	agifyResponse := &AgifyResponse{}
	err = json.Unmarshal(resBody, agifyResponse)
	if err != nil {
		fmt.Printf("could not bind json to agifyResponse %s\n", err)
		return nil, err
	}
	return agifyResponse, nil
}

func main() {
	fmt.Printf("start server\n")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 50051))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	agify.RegisterAgifyServer(s, &AgifyServer{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	fmt.Printf("close server\n")
}

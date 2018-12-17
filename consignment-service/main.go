package main

import (
	"fmt"
	pb "github.com/kalmeshbhavi/shipper/consignment-service/proto/consignment"
	vesselProto "github.com/kalmeshbhavi/shipper/vessel-service/proto/vessel"
	"github.com/micro/go-micro"
	"log"
	"os"
)

const defaultHost  = "localhost:27017"

func main() {

	host := os.Getenv("DB_HOST")

	if host == "" {
		host = defaultHost
	}

	session, err := CreateSession(host)

	defer session.Close()

	if err != nil {
		log.Panicf("Could not connect to datastore with host %s - %v", host, err)
	}

	srv := micro.NewService(
		micro.Name("go.micro.srv.consignment"),
		micro.Version("latest"),
	)

	vesselServiceClient := vesselProto.NewVesselServiceClient("go.micro.srv.vessel", srv.Client())

	srv.Init()

	pb.RegisterShippingServiceHandler(srv.Server(), &handler{session, vesselServiceClient})

	err= srv.Run()

	if err != nil {
		fmt.Println(err)
	}
}
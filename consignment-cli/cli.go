package main

import (
	"encoding/json"
	pb "github.com/kalmeshbhavi/shipper/consignment-service/proto/consignment"
	microclient "github.com/micro/go-micro/client"
	"github.com/micro/go-micro/cmd"
	"golang.org/x/net/context"
	"io/ioutil"
	"log"
	"os"
)

const address = "localhost:50051"
const defaultFilename  = "consignment.json"

func parseFile(file string) (*pb.Consignment, error)  {
	var consignment *pb.Consignment

	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &consignment)
	return consignment, err
}

func main() {

	cmd.Init()

	// Create new greeter client
	client := pb.NewShippingServiceClient("go.micro.srv.consignment", microclient.DefaultClient)

	// Contact the server and print out its response.
	file := defaultFilename
	if len(os.Args) > 1 {
		file = os.Args[1]
	}

	consignment, err := parseFile(file)
	if err != nil {
		log.Fatalf("Could not parse file: %v", err)
	}

	response, err := client.CreateConsignment(context.TODO(), consignment)
	if err != nil {
		log.Fatalf("Could not create the consignment: %v", err)
	}

	log.Printf("Created: %t", response.Created)

	consignments, err := client.GetConsignments(context.Background(), &pb.GetRequest{})
	if err != nil {
		log.Fatalf("Could not list the consignments: %v", err)
	}

	for _, v := range consignments.Consignments {
		log.Println(v)
	}

}
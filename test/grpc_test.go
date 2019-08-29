package benchmarks

import (
	"context"
	"encoding/json"
	pb "github.com/enrinal/hermes/order/proto/order"
	"google.golang.org/grpc"
	"io/ioutil"
	"log"
	"testing"
)

const (
	address         = "localhost:50051"
	defaultFilename = "order.json"
)

func parseFile(file string) (*pb.Order, error) {
	var order *pb.Order
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &order)
	return order, err
}


func BenchmarkCreateGRPC(b *testing.B) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	client := pb.NewOrderServiceClient(conn)

	for n := 0; n < b.N; n++ {
		createGRPC(client, b)
	}
}

func BenchmarkGetGRPC(b *testing.B) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	client := pb.NewOrderServiceClient(conn)

	for n := 0; n < b.N; n++ {
		getGRPC(client, b)
	}
}

func getGRPC(client pb.OrderServiceClient, b *testing.B) {
	//GO111MODULE=on go test -bench=. -benchmem
	_, err := client.GetOrders(context.Background(), &pb.GetRequest{})
	if err != nil {
		b.Fatalf("Could not list consignments: %v", err)
	}
}

func createGRPC(client pb.OrderServiceClient, b *testing.B) {

	order := &pb.Order{
		CourirId:"12as",
		Description:"hello",
	}
	_, err := client.CreateOrder(context.Background(), order)
	if err != nil {
		b.Fatalf("Could not greet: %v", err)
	}
}




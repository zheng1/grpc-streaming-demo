package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	pb "github.com/ridha/grpc-streaming-demo/protobuf"
	grpc "google.golang.org/grpc"
)

type flightServer struct{}

func (*flightServer) Flight(stream pb.Pilot_FlightServer) error {
	go func() {
		fmt.Println("Entering FlightServer")
		for {
			req, _ := stream.Recv()
			fmt.Printf("Received: %v, %v\n", req.Register, req.Report)
			// time.Sleep(time.Second)
		}
	}()
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		fmt.Println("You had input: " + input)
		res := pb.Response{}
		if input == "c" {
			res = pb.Response{Cmd: &pb.RegisterCmd{"1.2.3.4", "/bin/bash"}}
		}
		if input == "r" {
			res = pb.Response{Register: &pb.RegisterMetadata{"metadata_content"}}
		}
		if input == "s" {
			res = pb.Response{Start: &pb.StartConfd{"1.3.4.5"}}
		}
		if err := stream.Send(&res); err != nil {
			log.Fatalf("Failed to send a note: %v", err)
		}
	}

	fmt.Println("Leaving FlightServer")
	return nil
}

func main() {
	port := flag.Int("port", 50051, "Port for the server to run on")
	flag.Parse()

	fmt.Printf("Starting FlightServer Server on port %d...\n", *port)
	conn, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	grpcServer := grpc.NewServer()
	pb.RegisterPilotServer(grpcServer, &flightServer{})
	grpcServer.Serve(conn)
}

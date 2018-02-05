package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	pb "github.com/ridha/grpc-streaming-demo/protobuf"
	"golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

func main() {
	port := flag.String("port", "50051", "Port for the server to run on")
	flag.Parse()
	address := "localhost:" + *port
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewPilotClient(conn)
	stream, err := client.Flight(context.Background())
	if err != nil {
		log.Fatalf("%v.Flight(_) = _, %v", client, err)
	}
	waitc := make(chan struct{})
	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				// read done.
				close(waitc)
				return
			}
			if err != nil {
				log.Fatalf("Failed to receive a note : %v", err)
			}
			log.Printf("Got message %v, %v, %v", in.Register, in.Start, in.Cmd)
		}
	}()
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		fmt.Println("You had input: " + input)
		req := pb.Request{}
		if input == "i" {
			req = pb.Request{&pb.RegisterProxy{"999"}, nil}
		}
		if input == "r" {
			req = pb.Request{nil, &pb.ReportCmdStatus{"xyz", "success"}}
		}
		if err := stream.Send(&req); err != nil {
			log.Fatalf("Failed to send a note: %v", err)
		}
	}
	defer stream.CloseSend()
	<-waitc
}

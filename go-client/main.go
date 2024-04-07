package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/guycole/elder-mixer/proto"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

const banner = "elder-mixer 0.0"

func writeCommand(id, cmd string) {
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	cc := pb.NewMixerClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	rr, err := cc.EnqueueCommand(ctx, &pb.EnqueueRequest{Message: cmd, ClientId: id})
	if err != nil {
		log.Fatalf("could not write: %v", err)
	}

	log.Printf("Result: %s", rr.GetReceiptId())
}

func receiveTraffic(id string) {
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	cc := pb.NewMixerClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	er, err := cc.EnqueueCommand(ctx, &pb.EnqueueRequest{ClientId: id, Message: "traffic"})
	if err != nil {
		log.Fatalf("could not read: %v", err)
	}

	log.Printf("Result: %s", er.GetReceiptId())
}

func main() {
	flag.Parse()

	log.Println(banner)

	clientId := uuid.NewString()
	log.Printf("client id:%s", clientId)

	var input string

	counter := 0
	runFlag := true

	for runFlag {
		// fmt.Print("prompt>")
		fmt.Print("top top top\n")

		input = fmt.Sprintf("command %d\n", counter)
		print(input)

		counter++

		//		fmt.Scanln(&input)
		//		if input == "quit" {
		//			runFlag = false
		//			continue
		//		}

		writeCommand(clientId, input)
		receiveTraffic(clientId)

		time.Sleep(8 * time.Second)
	}
}

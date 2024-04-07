package main

import (
	"context"
	"flag"
	"log"
	"strconv"
	"time"

	"github.com/google/uuid"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/guycole/elder-mixer/proto"
)

var (
	addr = flag.String("addr", "localhost:50053", "the address to connect to")
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

	er, err := cc.EnqueueCommand(ctx, &pb.EnqueueRequest{Command: cmd, ClientId: id})
	if err != nil {
		log.Fatalf("could not write: %v", err)
	}

	log.Printf("Result: %s", er.GetReceiptId())
}

func main() {
	flag.Parse()

	log.Println(banner)

	clientId := uuid.NewString()
	log.Printf("client id:%s", clientId)

	counter := 0
	runFlag := true

	for runFlag {
		counter++
		command := "simple" + strconv.Itoa(counter)
		   
		writeCommand(clientId, command) 

		time.Sleep(8 * time.Second)
	}
}

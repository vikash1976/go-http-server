package main

import (
	"io"
	"log"
	"os"
	"time"

	pb "github.com/vikash1976/go-rpc/simple-rpc/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Couldn't connect: %v\n", err)
	}
	defer conn.Close()
	c := pb.NewHelloWorldClient(conn)

	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	callSayHello(c, name)
	log.Println("..................Uniary Communication Ends..................")

	//Stream Server API
	callSayMoreHellos(c, name)
	log.Println("..................Server Streaming Communication Ends..................")
	//Stream Client API
	names := []string{"Allan", "Brain", "Charlie", "Derek", "Elli"}
	callSayHelloToMany(c, names)
	log.Println("..................Client Streaming Communication Ends..................")
	//Bidirectional Stream API
	callLetsTalk(c, names)
	log.Println("..................2 way Streaming Communication Ends..................")

}

func callSayHello(c pb.HelloWorldClient, name string) {

	ctx, cancel := context.WithTimeout(context.Background(), 150*time.Millisecond)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.ToWhom{Name: name})
	if err != nil {
		log.Printf("Couldn't greet: %v\n", err) //changed from Fatalf to Printf,
		//as i wish to continue other functions, Fatalf calls os.Exit(1) internally.
		return
	}
	log.Printf("Greeting Message: %v\n", r.GreetMessage)
}

func callSayMoreHellos(c pb.HelloWorldClient, name string) {
	stream, err := c.SayMoreHellos(context.Background(), &pb.ToWhom{Name: name})
	if err != nil {
		log.Fatalf("Couldn't greet: %v\n", err)
	}
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			log.Println("No more data on stream, breaking")
			break
		}
		if err != nil {
			log.Fatalf("%v.SayMoreHellos(_) = _, %v", c, err)
		}
		log.Println(msg.GreetMessage)
	}
}

func callSayHelloToMany(c pb.HelloWorldClient, names []string) {

	stream, err := c.SayHelloToMany(context.Background())
	if err != nil {
		log.Fatalf("%v.SayHelloToMany(_) = _, %v", c, err)
	}
	for _, name := range names {
		if err := stream.Send(&pb.ToWhom{Name: name}); err != nil {
			log.Fatalf("%v.Send(%v) = %v", stream, name, err)
		}
	}
	reply, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("%v.sayHelloToMany() got error %v, want %v", stream, err, nil)
	}
	log.Printf("Server responded: %v", reply.GreetMessage)
}

func callLetsTalk(c pb.HelloWorldClient, names []string) {
	stream, err := c.LetsTalk(context.Background())
	if err != nil {
		log.Fatalf("%v.LetsTalk(_) = _, %v", c, err)
	}

	waitc := make(chan struct{})
	//go routine to keep receiving from server's stream
	go func() {
		for {

			msg, err := stream.Recv()
			if err == io.EOF {
				// read done. close it so that function can exit
				close(waitc)
				return
			}
			if err != nil {
				log.Fatalf("Failed to receive a talk : %v", err)
			}
			log.Printf("Getting message from server: %v \n", msg.GreetMessage)

		}
	}()
	//streaming to server
	for _, name := range names {
		log.Println("Sending at client's stream")
		if err := stream.Send(&pb.ToWhom{Name: name}); err != nil {
			log.Fatalf("Failed to send a note: %v", err)
		}
	}
	stream.CloseSend()
	<-waitc
}

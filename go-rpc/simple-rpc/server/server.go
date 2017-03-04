package main

import (
	"io"
	"log"
	"math/rand"
	"net"
	"net/http"
	"strings"
	"time"

	pb "github.com/vikash1976/go-rpc/simple-rpc/pb"
	"golang.org/x/net/context"
	_ "golang.org/x/net/trace"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, in *pb.ToWhom) (*pb.Greeting, error) {
	d := time.Millisecond * time.Duration((50 * rand.Intn(5)))
	log.Printf("Waiting time is: %v\n", d)
	select {
	case <-time.After(d):
		return &pb.Greeting{GreetMessage: "Hello " + in.Name}, nil
	case <-ctx.Done():
		log.Println("Context Timedout")
		return nil, ctx.Err()
	}
}

func (s *server) SayMoreHellos(in *pb.ToWhom, stream pb.HelloWorld_SayMoreHellosServer) error {

	arrOfMsg := []string{"Howdy " + in.Name, "How are you?", "How is it being a gopher?"}
	for _, msg := range arrOfMsg {

		if err := stream.Send(&pb.Greeting{GreetMessage: msg}); err != nil {
			return err
		}
		time.Sleep(time.Millisecond * 50)
	}
	log.Println("Server Streaming: No more data to send over stream")
	return nil
}

func (s *server) SayHelloToMany(stream pb.HelloWorld_SayHelloToManyServer) error {
	allNames := []string{}
	for {
		toWhom, err := stream.Recv()

		if toWhom != nil {
			//concatinating all client streamed names
			allNames = append(allNames, toWhom.Name)
		}
		if err == io.EOF {
			return stream.SendAndClose(&pb.Greeting{GreetMessage: "Hello guys " + strings.Join(allNames, ",")})
		}
		if err != nil {
			return err
		}

	}

}

func (s *server) LetsTalk(stream pb.HelloWorld_LetsTalkServer) error {
	for {
		toWhom, err := stream.Recv()

		if err == io.EOF {
			//endTime := time.Now()
			return nil

		}
		if err != nil {
			return err
		}
		arrOfMsg := []string{"Howdy " + toWhom.Name, "How are you?", "How is it being a gopher?"}
		for _, msg := range arrOfMsg {

			if err := stream.Send(&pb.Greeting{GreetMessage: msg}); err != nil {
				return err
			}
			time.Sleep(time.Millisecond * 100)
		}

	}
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalln("Failure in listeneing: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterHelloWorldServer(s, &server{})
	go s.Serve(lis)
	log.Fatal(http.ListenAndServe(":50055", nil)) // For net/trace
}

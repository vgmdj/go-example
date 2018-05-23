package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"

	pb "grpc/cmd/helloworld"
	grpclb "grpc/etcdv3"
)

var (
	serv = flag.String("service", "hello_service", "service name")
	host = flag.String("host", "localhost", "listening host")
	port = flag.String("port", "50001", "listening port")
	reg  = flag.String("reg", "http://localhost:2379", "register etcd address")
)

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", net.JoinHostPort(*host, *port))
	if err != nil {
		panic(err)
	}

	err = grpclb.Register(*serv, *host, *port, *reg, 15)
	if err != nil {
		panic(err)
	}

	log.Printf("starting hello service at %s", *port)
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	s.Serve(lis)
}

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	fmt.Printf("%v: Receive is %s\n", time.Now(), in.Name)
	return &pb.HelloReply{Message: "Hello " + in.Name + " from " + net.JoinHostPort(*host, *port)}, nil
}

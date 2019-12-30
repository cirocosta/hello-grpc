package main

import (
	"context"
	"flag"
	"fmt"
	"net"

	hello "github.com/cirocosta/hello-grpc/messaging"
	"google.golang.org/grpc"
)

var (
	isServer = flag.Bool("server", false, "to act as a server or not")
	addr     = flag.String("address", "127.0.0.1:1337", "server addr (to listen on or connect to)")
)

func client() (err error) {
	conn, err := grpc.Dial(*addr, grpc.WithInsecure())
	if err != nil {
		return
	}

	defer conn.Close()

	c := hello.NewHelloServiceClient(conn)
	resp, err := c.SayHi(context.TODO(), &hello.Request{Message: "heey"})
	if err != nil {
		return
	}

	fmt.Println("RESP.MSG: " + resp.GetMessage())

	return nil
}

type HelloServer struct{}

func (s *HelloServer) SayHi(context context.Context, req *hello.Request) (res *hello.Response, err error) {
	fmt.Println("REQ.MSG: " + req.GetMessage())

	res = &hello.Response{
		Message: "yoo",
	}
	return
}

func server() (err error) {
	var opts []grpc.ServerOption

	lis, err := net.Listen("tcp", *addr)
	if err != nil {
		return
	}

	grpcServer := grpc.NewServer(opts...)
	helloServer := &HelloServer{}

	hello.RegisterHelloServiceServer(grpcServer, helloServer)

	err = grpcServer.Serve(lis)
	if err != nil {
		return
	}

	return nil
}

func main() {
	flag.Parse()

	var runner func() error = client

	if *isServer {
		runner = server
	}

	err := runner()
	if err != nil {
		panic(err)
	}
}

package main

import (
	"context"
	"errors"
	"io"
	"log"
	"net"
	"strconv"

	"grpctest/demo"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type demoServer struct {
}

// 单次调用
func (s *demoServer) SayHello(ctx context.Context, req *demo.Request) (*demo.Reply, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	log.Println("SayHello got call:", req, " meta:", md)

	reply := &demo.Reply{
		Message: req.Name,
	}

	return reply, nil
}

// 客户端流
func (s *demoServer) SayHelloCliStream(cliStream demo.Demo_SayHelloCliStreamServer) error {
	md, _ := metadata.FromIncomingContext(cliStream.Context())
	log.Println("SayHelloCliStream got call meta:", md)

	for {
		req, err := cliStream.Recv()

		if err == io.EOF {
			log.Println("SayHelloCliStream end")

			reply := &demo.Reply{
				Message: req.Name,
			}

			return cliStream.SendAndClose(reply)

		}

		log.Println("SayHelloCliStream got:", req)
	}

	return nil
}

// 服务端流
func (s *demoServer) SayHelloServStream(req *demo.Request, stream demo.Demo_SayHelloServStreamServer) error {
	md, _ := metadata.FromIncomingContext(stream.Context())
	log.Println("SayHelloServStream got call meta:", md)

	reqCount, err := strconv.Atoi(req.Name)
	if err != nil {
		return errors.New("req.Name should be a integer")
	}

	for i := 0; i < reqCount; i++ {
		reply := &demo.Reply{
			Message: strconv.Itoa(i),
		}

		stream.Send(reply)
	}

	return nil
}

// 双向流
func (s *demoServer) SayHelloBidiStream(bistream demo.Demo_SayHelloBidiStreamServer) error {
	md, _ := metadata.FromIncomingContext(bistream.Context())
	log.Println("SayHelloBidiStream got call meta:", md)

	for {
		req, err := bistream.Recv()

		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		reply := &demo.Reply{
			Message: req.Name,
		}
		bistream.Send(reply)
	}
}

func main() {

	grpcServer := grpc.NewServer()

	demo.RegisterDemoServer(grpcServer, &demoServer{})

	l, err := net.Listen("tcp", ":6000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Println("Listening on tcp://localhost:6000")
	grpcServer.Serve(l)

}

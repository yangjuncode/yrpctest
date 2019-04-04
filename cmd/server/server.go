package main

import (
	"context"
	"errors"
	"io"
	"log"
	"net"
	"strconv"
	"time"

	"grpctest/demo"

	"github.com/golang/glog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type demoServer struct {
}

// 单次调用
func (s *demoServer) SayHello(ctx context.Context, req *demo.Request) (*demo.Reply, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	glog.Info("SayHello got call:", req, " meta:", md)

	reply := &demo.Reply{
		Message: req.Name,
	}

	return reply, nil
}

// 客户端流
func (s *demoServer) SayHelloCliStream(cliStream demo.Demo_SayHelloCliStreamServer) error {
	md, _ := metadata.FromIncomingContext(cliStream.Context())
	glog.Info("SayHelloCliStream got call meta:", md)

	for {
		req, err := cliStream.Recv()

		if err == io.EOF {
			glog.Info("SayHelloCliStream end")

			reply := &demo.Reply{
				Message: "server SayHelloCliStream end",
			}

			return cliStream.SendAndClose(reply)

		}

		glog.Info("SayHelloCliStream got:", req)
	}

	return nil
}

// 服务端流
func (s *demoServer) SayHelloServStream(req *demo.Request, stream demo.Demo_SayHelloServStreamServer) error {
	md, _ := metadata.FromIncomingContext(stream.Context())
	glog.Info("SayHelloServStream got call meta:", md)

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
	glog.Info("SayHelloBidiStream got call meta:", md)
	replyCount := 0
	for {
		req, err := bistream.Recv()

		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		reply := &demo.Reply{
			Message: req.Name + "reply:" + strconv.Itoa(replyCount),
		}
		replyCount++
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

	glog.Info("Listening on tcp://localhost:6000")

	go grpcClientTest()

	grpcServer.Serve(l)

}

func grpcClientTest() {
	time.Sleep(time.Second)
	conn, err := grpc.Dial("localhost:6000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %s", err)
	}
	defer conn.Close()

	client := demo.NewDemoClient(conn)

	go func() {
		//单次调用
		req := &demo.Request{
			Name: "say hello test 1",
		}
		reply, err := client.SayHello(context.Background(), req)
		if err != nil {
			glog.Info("cli SayHello call err:", err)
		} else {
			glog.Info("cli SayHello result:", reply)
		}

		req.Name = "sya hello test 2"
		reply, err = client.SayHello(context.Background(), req)
		if err != nil {
			glog.Info("cli SayHello call err:", err)
		} else {
			glog.Info("cli SayHello result:", reply)
		}

	}()

	go func() {
		// 客户端流
		req := &demo.Request{
			Name: "cli stream",
		}
		cliStream, err := client.SayHelloCliStream(context.Background())
		if err != nil {
			glog.Info("cli SayHelloCliStream err:", err)
			return
		}

		for i := 0; i < 11; i++ {
			req.Name = req.Name + strconv.Itoa(i) + ","
			err = cliStream.Send(req)
			if err != nil {
				glog.Info("cli stream send err:", err)
			}
		}
		reply, err := cliStream.CloseAndRecv()
		if err != nil {
			glog.Info("cli CloseAndRecv err:", err)
			return
		}

		glog.Info("cli CloseAndRecv result:", reply)
	}()

	go func() {
		// 服务端流
		req := &demo.Request{
			Name: "22",
		}
		serStream, err := client.SayHelloServStream(context.Background(), req)
		if err != nil {
			glog.Info("cli init SayHelloServStream err:", err)
			return
		}
		for {
			reply, err := serStream.Recv()
			if err == io.EOF {
				glog.Info("cli SayHelloServStream end")
				return
			}
			if err != nil {
				glog.Info("cli SayHelloServStream err:", err)
			}

			glog.Info("cli SayHelloServStream got:", reply)
		}
	}()

	go func() {
		// 双向流
		req := &demo.Request{
			Name: "bi",
		}
		biStream, err := client.SayHelloBidiStream(context.Background())
		if err != nil {
			glog.Info("cli init SayHelloServStream err:", err)
			return
		}
		for i := 0; i < 8; i++ {
			req.Name = "bi " + strconv.Itoa(i)
			err = biStream.Send(req)
			if err != nil {
				glog.Info("cli SayHelloBidiStream send err:", err)
			}

			reply, err := biStream.Recv()

			if err != nil {
				glog.Info("cli SayHelloBidiStream recv err:", err)
			}
			glog.Info("cli SayHelloBidiStream got:", reply)
		}

		biStream.CloseSend()

		for {
			reply, err := biStream.Recv()
			if err == io.EOF {
				glog.Info("cli SayHelloBidiStream got end")
				return
			}
			if err != nil {
				glog.Info("cli SayHelloBidiStream recv err:", err)
				return
			}

			glog.Info("cli SayHelloBidiStream got:", reply)

		}
	}()

	for {
		time.Sleep(time.Minute)
	}
}

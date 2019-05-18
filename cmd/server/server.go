package main

import (
	"context"
	"errors"
	"io"
	"net"
	"strconv"
	"time"

	"github.com/yangjuncode/yrpc-test/demo"

	"github.com/golang/glog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
)

type demoServer struct {
}

// 单次调用
func (s *demoServer) SayHello(ctx context.Context, req *demo.Request) (*demo.Reply, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	p, _ := peer.FromContext(ctx)
	glog.Info("sayhello peer", p)
	glog.Info("ser SayHello got call:", req, " meta:", md)

	reply := &demo.Reply{
		Message: req.Name + ":reply",
	}

	return reply, nil
}

// 客户端流
func (s *demoServer) SayHelloCliStream(cliStream demo.Demo_SayHelloCliStreamServer) error {
	md, _ := metadata.FromIncomingContext(cliStream.Context())
	glog.Info("SayHelloCliStream got call meta:", md)

	glog.Info("clistream context:", cliStream.Context())

	p, _ := peer.FromContext(cliStream.Context())
	glog.Info("clistream peer", p)

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
		glog.Fatalf("failed to listen: %v", err)
	}

	glog.Info("Listening on tcp://localhost:6000")

	go grpcClientTest()

	go func() {
		time.Sleep(time.Second * 10)
		glog.Info("forward test begin")
		forwardGrpcClientTest()
	}()

	grpcServer.Serve(l)

}

func grpcClientTest() {
	time.Sleep(time.Second)
	conn, err := grpc.Dial("localhost:6000", grpc.WithInsecure())
	if err != nil {
		glog.Fatalf("failed to connect: %s", err)
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
				glog.Info("cli SayHelloCliStream send err:", err)
			}
		}
		reply, err := cliStream.CloseAndRecv()
		if err != nil {
			glog.Info("cli SayHelloCliStream CloseAndRecv err:", err, reply)
			return
		}

		glog.Info("cli SayHelloCliStream result:", reply)
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
				glog.Info("cli SayHelloServStream end", reply)
				return
			}
			if err != nil {
				glog.Info("cli SayHelloServStream err:", err, reply)
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
				glog.Info("cli SayHelloBidiStream got end", reply)
				return
			}
			if err != nil {
				glog.Info("cli SayHelloBidiStream recv err:", err, reply)
				return
			}

			glog.Info("cli SayHelloBidiStream got:", reply)

		}
	}()

	for {
		time.Sleep(time.Minute)
	}
}

func forwardGrpcClientTest() {
	time.Sleep(time.Second)
	conn, err := grpc.Dial("localhost:6000", grpc.WithInsecure())
	if err != nil {
		glog.Fatalf("failed to connect: %s", err)
	}
	defer conn.Close()

	go func() {
		//单次调用
		req := &demo.Request{
			Name: "say hello test 1 f",
		}
		reply := &demo.Reply{}

		reqB, _ := req.Marshal()
		replyB, err := conn.InvokeForward(context.Background(), "/demo.demo/SayHello", reqB)
		if err != nil {
			glog.Info("clif SayHello call err:", err)
		} else {
			err = reply.Unmarshal(replyB)
			glog.Info("clif SayHello result:", reply, err)
		}

		req.Name = "sya hello test 2 f"
		reqB, _ = req.Marshal()
		replyB, err = conn.InvokeForward(context.Background(), "/demo.demo/SayHello", reqB)
		if err != nil {
			glog.Info("clif SayHello call err:", err)
		} else {
			err = reply.Unmarshal(replyB)
			glog.Info("clif SayHello result:", reply, err)
		}

	}()

	go func() {
		// 客户端流
		req := &demo.Request{
			Name: "cli stream",
		}
		reply := &demo.Reply{}

		sreamDesc := &grpc.StreamDesc{
			//StreamName:    "SayHelloCliStream",
			ClientStreams: true,
		}

		streamfwd, err := grpc.NewClientStream(context.Background(), sreamDesc, conn, "/demo.demo/SayHelloCliStream")
		if err != nil {
			glog.Info("clif NewClientStream err:", err)
			return
		}

		for i := 0; i < 11; i++ {
			req.Name = req.Name + strconv.Itoa(i) + ","
			reqB, _ := req.Marshal()

			err = streamfwd.SendMsgForward(reqB)
			if err != nil {
				glog.Info("clif SayHelloCliStream sendfwd err:", err)
			}
		}
		err = streamfwd.CloseSend()
		if err != nil {
			glog.Info("clif SayHelloCliStream CloseSend err:", err)
		}
		replyB, err := streamfwd.RecvMsgForward()
		if err != nil {
			glog.Info("clif SayHelloCliStream RecvMsgForward err:", err)
			return
		}

		err = reply.Unmarshal(replyB)

		glog.Info("clif SayHelloCliStream result:", reply, err)
	}()

	go func() {
		// 服务端流
		req := &demo.Request{
			Name: "22",
		}
		reply := &demo.Reply{}

		reqB, _ := req.Marshal()

		sreamDesc := &grpc.StreamDesc{
			//StreamName:    "SayHelloCliStream",
			//ClientStreams: true,
			ServerStreams: true,
		}

		serStream, err := grpc.NewClientStream(context.Background(), sreamDesc, conn, "/demo.demo/SayHelloServStream")
		if err != nil {
			glog.Info("clif SayHelloServStream err:", err)
			return
		}

		err = serStream.SendMsgForward(reqB)
		if err != nil {
			glog.Info("clif SayHelloServStream SendMsgForward err:", err)
			return
		}

		err = serStream.CloseSend()
		if err != nil {
			glog.Info("clif SayHelloServStream CloseSend err:", err)
			return
		}

		for {
			replyB, err := serStream.RecvMsgForward()
			if err == io.EOF {
				glog.Info("clif SayHelloServStream end")
				return
			}
			if err != nil {
				glog.Info("clif SayHelloServStream err:", err)
			}

			err = reply.Unmarshal(replyB)

			glog.Info("clif SayHelloServStream got:", reply)
		}
	}()

	go func() {
		// 双向流
		req := &demo.Request{
			Name: "bi",
		}
		reply := &demo.Reply{}

		sreamDesc := &grpc.StreamDesc{
			//StreamName:    "SayHelloCliStream",
			ClientStreams: true,
			ServerStreams: true,
		}

		biStream, err := grpc.NewClientStream(context.Background(), sreamDesc, conn, "/demo.demo/SayHelloBidiStream")
		if err != nil {
			glog.Info("clif SayHelloBidiStream err:", err)
			return
		}

		for i := 0; i < 8; i++ {
			req.Name = "bi " + strconv.Itoa(i)
			reqB, _ := req.Marshal()
			err = biStream.SendMsgForward(reqB)
			if err != nil {
				glog.Info("clif SayHelloBidiStream send err:", err)
			}

			replyB, err := biStream.RecvMsgForward()

			if err != nil {
				glog.Info("clif SayHelloBidiStream recv err:", err)
			}
			if err == io.EOF {
				glog.Info("clif SayHelloBidiStream got end")
				break
			}
			err = reply.Unmarshal(replyB)
			glog.Info("clif SayHelloBidiStream got:", reply)
		}

		biStream.CloseSend()

		for {
			replyB, err := biStream.RecvMsgForward()
			if err == io.EOF {
				glog.Info("clif SayHelloBidiStream got end")
				return
			}
			if err != nil {
				glog.Info("clif SayHelloBidiStream recv err:", err)
				return
			}

			err = reply.Unmarshal(replyB)

			glog.Info("clif SayHelloBidiStream got:", reply)

		}
	}()

	for {
		time.Sleep(time.Minute)
		return
	}
}

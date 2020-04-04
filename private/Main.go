
package main

import (
	"context"
	"flag"
	"fmt"
	"gRPC-Gateway/private/pb"
	"gRPC-Gateway/private/service"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"net"
	"net/http"
	"sync"
	"github.com/jeanphorn/log4go"
)

const PORT = ":9192"

func GRPCStart() {
	listen, err := net.Listen("tcp", PORT)
	if err != nil {
		fmt.Printf("net.Listen failed, err=%+v\n", err)
	}

	grpc_server := grpc.NewServer()

	pb.RegisterMyServiceNameServer(grpc_server, new(service.Server))

	grpc_server.Serve(listen)
}

func GatewayStart() error {
	var (
		echoEndpoint = flag.String("echo_endpoint", PORT, "endpoint of Gateway")
	)

	var ctx = context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := pb.RegisterMyServiceNameHandlerFromEndpoint(ctx, mux, *echoEndpoint, opts)
	if err != nil {
		return err
	}

	return http.ListenAndServe(":8080", mux)	// 对外提供HTTP服务的端口
}

func main() {
	log4go.LoadConfiguration("../config/logConfig.json")

	var wg sync.WaitGroup	// 计数器，计数器不为0时当前线程一直阻塞

	go GRPCStart()

	go GatewayStart()


	wg.Add(1)

	wg.Wait()
}

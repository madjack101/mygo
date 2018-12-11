package main

import (
    "mygo/grpctest/hello"
    "log"
    "net"
    "runtime"
    _ "strconv"

    "golang.org/x/net/context"
    "google.golang.org/grpc"
)

const (
    port = "41005"
)

type Server struct{}

func main() {
    runtime.GOMAXPROCS(runtime.NumCPU())    

    //起服务
    lis, err := net.Listen("tcp", ":"+port) 
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    s := grpc.NewServer()   
    hello.RegisterHelloServer(s, &Server{})
    s.Serve(lis)

    log.Println("grpc server in: %s", port)
}


// 定义方法
func (t *Server) Hello(ctx context.Context, request *hello.User) (response *hello.Msg, err error) {
    response = &hello.Msg{
        Text: "hi " + request.Name,
    }
    return response, err
}
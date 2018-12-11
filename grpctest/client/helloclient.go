package main

import (
	"mygo/grpctest/hello"	
    _ "fmt"
    "os"
    "log"
    "runtime"
    "strconv"
    _ "strings"
    "sync"
    "time"

    "math/rand"

    "golang.org/x/net/context"
    "google.golang.org/grpc"
)

var (
    wg sync.WaitGroup   
)

var (
    networkType = "tcp"
    server      = "127.0.0.1"
    port        = "41005"
    parallel    = 50        //连接并行度
    times       = 10000    //每连接请求次数
)

func main() {
	if len(os.Args) >= 5 {
		server = os.Args[1]
		port = os.Args[2]
		parallel, _ = strconv.Atoi(os.Args[3])
		times, _ = strconv.Atoi(os.Args[4])
		log.Printf("set params: %s %s %d %d", server, port, parallel, times)
	}

    runtime.GOMAXPROCS(runtime.NumCPU())
    currTime := time.Now()

    //并行请求
    for i := 0; i < int(parallel); i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            exe()
        }()
    }
    wg.Wait()

    log.Printf("time taken: %.2f ", time.Now().Sub(currTime).Seconds())
}

func exe() {
    //建立连接
    conn, err := grpc.Dial(server + ":" + port, grpc.WithInsecure())
    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    
    defer conn.Close()
    client := hello.NewHelloClient(conn)

    for i := 0; i < int(times); i++ {
        sayHello(client)
    }

    // time.Sleep(1000000000)
}

func sayHello(client hello.HelloClient) {
    var request hello.User
    r := rand.Intn(parallel)

    request.Name = "robot " + strconv.Itoa(int(r))

    response, _ := client.Hello(context.Background(), &request) //调用远程方法

    if r == 10 {
    	log.Printf(response.Text)
	}
    //判断返回结果是否正确
    // if id, _ := strconv.Itoa(strings.Split(response.Name, ":")[0]); id != r {
    //     log.Printf("response error  %#v", response)
    // }

}

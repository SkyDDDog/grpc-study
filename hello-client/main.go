package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "grpc-study/hello-server/proto"
)

func main() {
	// 连接服务
	// 第二个参数为安全配置(这里新建了空的加密配置，即不进行安全加密)
	conn, err := grpc.Dial("127.0.0.1:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic("fail to connect server")
	}
	// 关闭grpc连接，都记得关闭！
	defer conn.Close()
	// 与服务端建立连接
	client := pb.NewSayHelloClient(conn)
	resp, _ := client.SayHello(context.Background(),
		&pb.HelloRequest{
			RequestName: "二火",
		})
	fmt.Println(resp.GetResponseMsg())
}

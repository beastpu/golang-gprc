package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pb "src/api/pb/hello"
)

func main() {
	// 连接服务器
	conn, err := grpc.Dial(":8972", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("faild to connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewGreeterClient(conn)
	// 调用服务端的SayHello
	//r, err := c.SayBye(context.Background(), &pb.ByeRequest{Info: "test"})
	//r2, _ := c.SayHello(context.Background(), &pb.HelloRequest{Name: "test"})
	r, err := c.GetInfo(context.Background(), &pb.Person{Name: "test",Id: 111,Email: "qq",Type: 2})
	if err != nil {
		fmt.Printf("could not greet: %v", err)
	}
	//fmt.Println()
	fmt.Printf("Bye: %s !\n", r.Data)
}
package main

import (
	"fmt"
	"google.golang.org/grpc"
	pb "src/api/pb/pb_pc"
	"src/service"
	"net"
)

func main(){
	laptopServer:=service.NewLaptopServer()
	laptopServer.Store=service.NewInMemoryLaptopStore()
	grpcServer:=grpc.NewServer()
	pb.RegisterCreateLaptopServiceServer(grpcServer,laptopServer)

	listen,err:=net.Listen("tcp",":8888")
	fmt.Printf("listen tcp 8888")
	if err!=nil{
		fmt.Println("grpc error")
	}
	err=grpcServer.Serve(listen)
	if err!=nil{
		fmt.Println("eepr:",err)
		return
	}

}


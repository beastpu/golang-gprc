package main

import (
	pb "src/api/pb/pb_pc"
	"src/sample"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"time"
)
func CreateLaptop(c pb.CreateLaptopServiceClient){
	laptop := sample.NewLaptop()
	laptop.Id = "2"
	req := &pb.CreateLaptopRequest{Laptop: laptop}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	r, err := c.CreateLaptop(ctx, req)
	//r, err := c.CreateLaptop(context.Background(),req)
	if err != nil {
		fmt.Printf("could not create laptop: %v", err)
	}
	fmt.Printf("data: %s !\n", r.GetId())
}
func SearchLaptop(c pb.CreateLaptopServiceClient){
	fmt.Println("Search Laptop ..")
	filter := &pb.Filter{
		MaxPriceUsd: 200,
		MinCpuCores: 2,
	}
	searchReq:= &pb.SearchLaptopRequest{Filter: filter}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	stream, err := c.SearchLaptop(ctx, searchReq)
	if err != nil {
		fmt.Println("can not search laptop")
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Println("can not receive response")
		}
		laptop:=res.GetLaptop()
		fmt.Println("found",laptop.GetId(),laptop.GetPriceUsd(),laptop.GetCpu().GetNumberCores())
	}
}

func main() {
	conn, err := grpc.Dial(":8888", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("faild to connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewCreateLaptopServiceClient(conn)
	for i:=0;i<5;i++{
		CreateLaptop(c)
	}
   SearchLaptop(c)

}
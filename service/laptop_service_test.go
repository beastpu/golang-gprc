package service

import (
	"context"
	"fmt"
	pb "src/api/pb/pb_pc"
	"src/sample"
	"testing"
)

func TestCreateLaptop(t *testing.T){
	server:=NewLaptopServer()
	server.Store=NewInMemoryLaptopStore()
	laptop:=sample.NewLaptop()

	req:=&pb.CreateLaptopRequest{
		Laptop: laptop,
	}
	res,err:=server.CreateLaptop(context.Background(),req)
	if err!=nil {
		t.Error("create labtop failed",err)
	}
	fmt.Println(res)

}

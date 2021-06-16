package service

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	pb "src/api/pb/pb_pc"
)
type laptopServer struct{
	Store LaptopStore
}


func NewLaptopServer()*laptopServer{
	return &laptopServer{}
}

func(l *laptopServer)CreateLaptop(ctx context.Context, req *pb.CreateLaptopRequest)(res *pb.CreateLaptopResponse,err error){
	fmt.Println("receive laptop object..",req.Laptop)
	id:=req.Laptop.Id
	 if len(id)>2 {
		 if _, err := uuid.Parse(id); err != nil {
			 return nil, status.Error(codes.InvalidArgument, "not a valid uuid")
		 }
	 }else{
		id,err:=uuid.NewRandom()
		if err!=nil{
			return nil,status.Error(codes.Internal,"can not generate new id")
		}
		fmt.Println(id.String())
		req.Laptop.Id=id.String()
	}
	if ctx.Err()==context.DeadlineExceeded{
		fmt.Println("context deadlineExceed")
		return nil,status.Error(codes.DeadlineExceeded,"deadline is exceed")

	}
	if ctx.Err()==context.Canceled{
		fmt.Println("context have been canceled")
		return nil,status.Error(codes.Canceled,"deadline is canceled")
	}
	err=l.Store.Save(req.Laptop)
	if err!=nil{
		return nil,status.Error(codes.Internal,"can not save")

	}
	fmt.Println("InMemory Data:",l.Store.GetData())
	res=&pb.CreateLaptopResponse{
		Id: req.Laptop.Id,
	}
	err=nil
	return
}

func (l *laptopServer)SearchLaptop(req *pb.SearchLaptopRequest, stream pb.CreateLaptopService_SearchLaptopServer) error {
	filter := req.GetFilter()
	err := l.Store.Search(filter, func(laptop *pb.Laptop) error {
		 res:=&pb.SearchLaptopResponse{Laptop: laptop}
		 err:=stream.Send(res)
		 if err!=nil{
		 	return err
		 }
		 return nil
	} )
	if err != nil {
		return status.Errorf(codes.Unimplemented, "method SearchLaptop not implemented")
	}
 return nil
}


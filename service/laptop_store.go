package service

import (
	"errors"
	"fmt"
	"github.com/jinzhu/copier"
	pb "src/api/pb/pb_pc"
	"sync"
)
var ErrorAlreadyExist=errors.New("alreadyExist")
type LaptopStore interface{
	Save(laptop *pb.Laptop) error
	GetData()map[string]*pb.Laptop
	Search(filter *pb.Filter,found func(laptop *pb.Laptop)error)error

}
type InMemoryLaptopStore struct {
	mutex sync.RWMutex
	data map[string]*pb.Laptop

}
func NewInMemoryLaptopStore()*InMemoryLaptopStore{
	return &InMemoryLaptopStore{
		data: make(map[string]*pb.Laptop),
	}
}
func(m *InMemoryLaptopStore)Save(laptop *pb.Laptop)error {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	if m.data[laptop.Id] != nil {
		return ErrorAlreadyExist
	}
	other:=&pb.Laptop{}
	err:=copier.Copy(other,laptop)
	if  err!= nil{
		return fmt.Errorf("copy error")

	}
	m.data[laptop.Id]=laptop
	return nil
}
func(m *InMemoryLaptopStore)GetData()map[string]*pb.Laptop{
	return m.data
}

func (m *InMemoryLaptopStore)Search(filter *pb.Filter,found func(laptop *pb.Laptop)error)error{
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	for _,laptop:=range m.data{
		if isQualified(filter,laptop){
			other,err:=deepCopy(laptop)
			if err!=nil{
				return err
			}
			err=found(other)
			if err!=nil{
				return err
			}
		}
	}
	return nil
}
func isQualified(filter *pb.Filter,laptop *pb.Laptop)bool{
	if laptop.GetPriceUsd()>filter.GetMaxPriceUsd(){
		return false
	}
	if laptop.GetCpu().GetNumberCores()<filter.MinCpuCores{
		return false
	}
	return true
}
func deepCopy(laptop *pb.Laptop)(*pb.Laptop,error){
	other:=&pb.Laptop{}
	err:=copier.Copy(other,laptop)
	if err!=nil{
		return nil,fmt.Errorf("can not copy laptop")

	}
	return other,nil
}
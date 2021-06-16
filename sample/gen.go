package sample

import (
	pb "src/api/pb/pb_pc"
	"math/rand"
	"time"
)
func init(){
	rand.Seed(time.Now().UnixNano())
}
func randomLayout() pb.Keyboard_Layout{
	n := rand.Int31n(3)
	switch n {
	case 1:
		return pb.Keyboard_l1
	case 2:
		return pb.Keyboard_l2
	default:
		return pb.Keyboard_UNKNOWN
	}
}
func NewKeyboard()*pb.Keyboard {

	return &pb.Keyboard{
		Layout: randomLayout(),
	}
}
func NewScreen()*pb.Screen{
	r:=&pb.Screen_Resolution{
		Width: rand.Uint32(),
		Height: rand.Uint32(),
	}
	panel:=pb.Screen_IPS
	size_inch:=float32(14.5)
	multitouch:=false
	return &pb.Screen{
		Resolution: r,
		SizeInch: size_inch,
		Panel: panel,
		Multitouch: multitouch,
	}

}
func randomMem() *pb.Memory {
	unit:=pb.Memory_Unit(rand.Int31n(7))
	value:=rand.Uint64()
	return &pb.Memory{
		Value:value,
		Unit: unit,
	}

}
func NewStorage() *pb.Storage {
	driver:=pb.Storage_Driver(rand.Int31n(3))
	return &pb.Storage{
		Driver: driver,
		Memory: randomMem(),
	}
}
func NewProcessor()(*pb.CPU,*pb.GPU){
	cpu:=&pb.CPU{
		Brand: "mac",
		Name: "mac pro",
		NumberCores: uint32(rand.Int31n(8)),
		NumberThreads: uint32(rand.Int31n(8)),
		MaxGhz:5,
		MinGhz: 1,
	}
	gpu:=&pb.GPU{
		Brand: "microsoft",
		Name: "win10",
		MinGhz: 2,
		MaxGhz: 5,
		Memory: randomMem(),
	}
	return cpu,gpu
}

func NewLaptop()*pb.Laptop{
	cpu,gpu:=NewProcessor()
	gpus:=[]*pb.GPU{gpu}
	storage:=[]*pb.Storage{NewStorage()}
	return &pb.Laptop{
		Id: "1",
		Brand: "dell",
		Name: "dell2",
		Cpu: cpu,
		Ram: randomMem(),
		Gpus:gpus,
		Storage: storage,
		Screen: NewScreen(),
		Keyboard: NewKeyboard(),
		Weight: &pb.Laptop_WeightKg{100},
		PriceUsd: 100,
		ReleaseYear: 2018,
		UpdatedAt: nil,




	}
}
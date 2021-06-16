package serializer

import (
	"fmt"
	"github.com/gogo/protobuf/proto"
	"github.com/stretchr/testify/require"
	pb "src/api/pb/pb_pc"
	"src/sample"
	"io/ioutil"
	"testing"
)

func TestFileSerializer(t*testing.T){
	t.Parallel()
	binaryfile:="./laptop.bin"
	laptop:=sample.NewLaptop()
	err:= WritePbToBinaryFile(laptop,binaryfile)
	require.NoError(t, err)
	laptop2:=&pb.Laptop{}
	err2:=ReadPbFromBinaryFile(binaryfile,laptop2)
	require.NoError(t, err2)
	require.True(t, proto.Equal(laptop,laptop2))
}

func TestPbToJSON(t *testing.T){

	data,error:=ProtobufToJSON(sample.NewLaptop())
	if error!=nil{
		fmt.Println(error)
	}
	ioutil.WriteFile("lapton.json",[]byte(data),0644)
}
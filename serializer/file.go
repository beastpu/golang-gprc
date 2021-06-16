package serializer

import (
"fmt"
	"github.com/gogo/protobuf/jsonpb"
	"github.com/gogo/protobuf/proto"
"io/ioutil"
)

func WritePbToBinaryFile(message proto.Message,filename string)error{
	data,_:=proto.Marshal(message)
	err:=ioutil.WriteFile(filename,data,0644)
	if err!=nil{
		return fmt.Errorf("can not write %w",err)
	}
	return nil
}
func ReadPbFromBinaryFile(filename string,message proto.Message)error{
	data,error:=ioutil.ReadFile(filename)
	if error!=nil{
		fmt.Println(error)
	}
	err:=proto.Unmarshal(data,message)
	if err !=nil{
		fmt.Println(err)
	}
	return nil
}
func ProtobufToJSON(message proto.Message)(string,error){
	marsharler:=jsonpb.Marshaler{
		EnumsAsInts: false,
		EmitDefaults: true,
		Indent: " ",
		OrigName:true,
	}

	return marsharler.MarshalToString(message)
}
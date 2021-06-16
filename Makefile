

mod:
	go mod tidy
proto-hello:
	tools/proto/bin/protoc -I=api/msg/pc -I=api/msg/hello  api/msg/hello/*.proto --gogo_out=plugins=grpc:api/pb/hello

proto-pc:
	tools/proto/bin/protoc -I=api/msg/pc -I=api/msg/hello -I=api/third api/msg/pc/*.proto --gogo_out=plugins=grpc:api/pb/pb_pc

build-hello:
	go build -o _bin/hello/client.exe ./client/cmd/helloworld/ && go build -o _bin/hello/server.exe ./server/cmd/helloworld/

build-pc:
	go build -o _bin/pc/client.exe ./client/cmd/pc/ && go build -o _bin/pc/server.exe ./server/cmd/pc/

test-serializer:
	go test -v ./serializer/

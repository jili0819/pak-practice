app_name=demo
GREETING := "Hello, World!"
.PHONY: build clean run proto
#.PHONY后面跟的目标都被称为伪目标，也就是说我们 make 命令后面跟的参数如果出现在.PHONY 定义的伪目标中，
#那就直接在Makefile中就执行伪目标的依赖和命令。不管Makefile同级目录下是否有该伪目标同名的文件，即使有也不会产生冲突。另一个就是提高执行makefile时的效率
clean:
	go clean
	rm -rf ./bin/*

build:
	@echo "Building $(GREETING)..."
	go build -o ./bin/$(app_name) main.go

run:
	go run -race main.go

proto:
	export PATH="$PATH:$(go env GOPATH)/bin"
	protoc --go_out=. --go-grpc_out=. grpc/proto/rpcpb1/helloworld.proto grpc/proto/rpcpb2/helloworlds.proto
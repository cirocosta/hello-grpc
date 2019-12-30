GRPC_DSTS := $(patsubst %.proto,%.pb.go,$(shell find ./messaging -name "*.proto" -type f))


build: $(GRPC_DSTS)
	go build -v ./cmd/hello

%.pb.go: %.proto
	protoc $^ --go_out=plugins=grpc:.


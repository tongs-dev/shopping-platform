# Define variables
GOPATH := $(shell go env GOPATH)

.PHONY: proto
proto:
	protoc --plugin=protoc-gen-go=$(GOPATH)/bin/protoc-gen-go --plugin=protoc-gen-micro=$(GOPATH)/bin/protoc-gen-micro --proto_path=. --micro_out=. --go-grpc_out=./ --go_out=.  ./proto/user/user.proto

.PHONY: build
build:
	go build -o $(BINARY_NAME) *.go

.PHONY: release
release:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $(BINARY_NAME) *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: dockerBuild
dockerBuild:
	docker build -t $(BINARY_NAME):latest .

.PHONY: clean
clean:
	rm -rf $(BINARY_NAME) $(OUTPUT_DIR)/*.pb.go
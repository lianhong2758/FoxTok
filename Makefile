NAME=FoxTok
EXE_NAME=${NAME}.exe
PROTOPATH=server/proto
VERSION=0.0.1

build:
	@echo "build!"
	@go version
	@go env -w GOPROXY=https://goproxy.cn,direct
	@go mod tidy
	@go build
	@echo "Done!"

run:
	@echo "run!"
	@go version
	@go env -w GOPROXY=https://goproxy.cn,direct
	@go mod tidy
	@go run main.go

debug:
	@echo "debug"
	@go version
	@go env -w GOPROXY=https://goproxy.cn,direct
	@go mod tidy
	@go run main.go -d
	
build_proto:
	@echo "build_proto!"
	@cd ${PROTOPATH} && protoc --go_out=. --go_opt=paths=source_relative *.proto
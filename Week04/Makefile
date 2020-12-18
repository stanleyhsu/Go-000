
app: proto
	go build -o hello cmd/hello/main.go

.FORCE:
proto:.FORCE
	export GO111MODULE=on
	protoc --go_out=. --go-grpc_out=. ./api/hello/v1/*.proto

clean:
	rm -rf proto/hello/v1/*.pb.go


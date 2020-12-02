
app: proto
	go build -o cmd/user_client/client cmd/user_client/main.go
	go build -o cmd/user_server/server cmd/user_server/main.go

.FORCE:
proto:.FORCE
	export GO111MODULE=on
	protoc --go_out=. --go-grpc_out=. ./proto/user/*.proto

clean:
	rm -rf proto/user/*.pb.go


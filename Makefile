BUILD_FLAGS = -ldflags="-s -w" 
CLIENT_BINARY = bin/client
SERVER_BINARY = bin/server
build:
	#build client ...
	env GOOS=linux CGO_ENABLED=0 go build -a -installsuffix nocgo $(BUILD_FLAGS) -o $(CLIENT_BINARY) ./client/.
	#build server
	env GOOS=linux CGO_ENABLED=0 go build -a -installsuffix nocgo $(BUILD_FLAGS) -o $(SERVER_BINARY) ./server/.

dependecies:
	#installing dependecies ...
	go mod tidy

protoc:
	#generate protocol buffer(.proto) files ...
	protoc -I proto/notif proto/notif/notif.proto --go_out=plugins=grpc:models/notif

start-client:
	#starting client ...
	./bin/client

start-server:
	#starting server ...
	./bin/server
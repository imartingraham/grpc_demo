SERVER_OUT := "bin/server"
CLIENT_OUT := "bin/client"
API_OUT := "api/api.pb.go"
API_REST_OUT := "api/api.pb.gw.go"
PKG := "github.com/imartingraham/grpc_demo"
SERVER_PKG_BUILD := "${PKG}/server"
CLIENT_PKG_BUILD := "${PKG}/client"
PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/)
.P
build:
	@go install ./server

run:
	@-pkill server
	@server

watch:
	xnotify \
		-i . \
		-e "(vendor|\.git)$$" \
		--batch 100 \
		--trigger \
		-- make -s build \
		-- make -s run
api:
	protoc -I api/ -I${GOPATH}/src --go_out=plugins=grpc:api api/api.proto
server: dep api ## Build the binary file for server
	@go build -i -v -o bin/client github.com/imartingraham/grpc_demo/server

client: dep api ## Build the binary file for client
	@go build -i -v -o bin/client github.com/imartingraham/grpc_demo/client

clean: ## Remove previous builds
	@rm $(SERVER_OUT) $(CLIENT_OUT) $(API_OUT) $(API_REST_OUT) $(API_SWAG_OUT)


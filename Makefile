go_mock_error:
	export GOPATH="$HOME/go"
	export PATH="$GOPATH/bin:$PATH"

gen_grpc:
	protoc --go_out=pkg/ --go_opt=paths=import --go-grpc_out=pkg/ --go-grpc_out=paths=import api/proto/balance_ms.proto
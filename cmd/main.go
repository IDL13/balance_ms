package cmd

import (
	"github.com/IDL13/balance_ms/internal/handlers"
	"github.com/IDL13/balance_ms/pkg/api"
	"google.golang.org/grpc"
)

func main() {
	s := grpc.NewServer()
	srv := &handlers.GRPCServer{}
	api.RegisterBalanceMsServer(s, srv)
}

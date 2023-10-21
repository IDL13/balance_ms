package cmd

import (
	"fmt"
	"github.com/IDL13/balance_ms/internal/CSV"
	"github.com/IDL13/balance_ms/internal/handlers"
	"github.com/IDL13/balance_ms/pkg/api"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

func main() {
	s := grpc.NewServer()
	srv := &handlers.GRPCServer{}
	api.RegisterBalanceMsServer(s, srv)

	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error from listen server: %e", err)
		os.Exit(1)
	}

	if err := s.Serve(listen); err != nil {
		log.Fatal(err)
	} else {
		if err = CSV.CreateCSV(); err != nil {
			log.Fatal(err)
		}
	}
}

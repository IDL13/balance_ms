package main

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
	//srv := &handlers.GRPCServer{request: requests.New()}
	srv := handlers.New()
	api.RegisterBalanceMsServer(s, srv)

	fmt.Println(`
╔══╗╔══╗─╔╗───╔╗╔══╗
╚╗╔╝║╔╗╚╗║║──╔╝║╚═╗║
─║║─║║╚╗║║║──╚╗║╔═╝║
─║║─║║─║║║║───║║╚═╗║
╔╝╚╗║╚═╝║║╚═╗─║║╔═╝║
╚══╝╚═══╝╚══╝─╚╝╚══╝
	`)
	fmt.Println("[SERVER STARTED]")
	fmt.Println("http://127.0.0.1:8085")

	listen, err := net.Listen("tcp", ":8085")
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

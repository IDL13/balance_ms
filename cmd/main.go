package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/IDL13/balance_ms/gql/graph"
	"github.com/IDL13/balance_ms/internal/CSV"
	"github.com/IDL13/balance_ms/internal/handlers"
	"github.com/IDL13/balance_ms/pkg/api"
	"google.golang.org/grpc"
)

func RunGrpc(stopGrpc chan bool) {
	go func() {
		<-stopGrpc
		os.Exit(1)
	}()

	s := grpc.NewServer()
	//srv := &handlers.GRPCServer{request: requests.New()}
	srv := handlers.New()
	api.RegisterBalanceMsServer(s, srv)

	err := CSV.CreateCSV()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when create log file csv format")
	}

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

func RunGrapql(stopGrapql chan bool) {
	go func() {
		<-stopGrapql
		os.Exit(1)
	}()

	port := "8086"

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: graph.New()}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

	select {
	case <-stopGrapql:
		os.Exit(1)
	}
}

func main() {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	stopGrpc := make(chan bool, 1)
	stopGrapql := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		fmt.Println(sig)
		os.Remove("./csv_log.csv")
		done <- true
	}()

	go func() {
		<-done
		stopGrpc <- true
		stopGrapql <- true
	}()

	go RunGrpc(stopGrpc)
	go RunGrapql(stopGrapql)
}

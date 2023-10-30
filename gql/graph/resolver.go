package graph

import "github.com/IDL13/balance_ms/internal/handlers"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

func New() *Resolver {
	return &Resolver{
		handler: handlers.New(),
	}
}

type Resolver struct {
	handler *handlers.GRPCServer
}

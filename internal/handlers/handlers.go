package handlers

import (
	"context"
	"github.com/IDL13/balance_ms/pkg/api"
)

// GRPCServer ...
type GRPCServer struct {
	api.UnimplementedBalanceMsServer
}

// AddBalance ...
func (s *GRPCServer) AddBalance(ctx context.Context, req *api.AddBalanceRequest) (*api.AddBalanceResponse, error) {
	return &api.AddBalanceResponse{Status: 0}, nil
}

// GetBalance ...
func (s *GRPCServer) GetBalance(context.Context, *api.GetBalanceRequest) (*api.GetBalanceResponse, error) {
	return &api.GetBalanceResponse{Balance: "100"}, nil
}

// Reserve ...
func (s *GRPCServer) Reserve(context.Context, *api.ReserveRequest) (*api.ReserveResponse, error) {
	return &api.ReserveResponse{Status: 0}, nil
}

// GetRevenue ...
func (s *GRPCServer) GetRevenue(context.Context, *api.GetRevenueRequest) (*api.GetRevenueResponse, error) {
	return &api.GetRevenueResponse{Status: 0}, nil
}

// mustEmbedUnimplementedBalanceMsServer ...
func (s *GRPCServer) mustEmbedUnimplementedBalanceMsServer() {}

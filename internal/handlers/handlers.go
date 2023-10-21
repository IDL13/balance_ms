package handlers

import (
	"context"
	"fmt"
	"github.com/IDL13/balance_ms/internal/requests"
	"github.com/IDL13/balance_ms/pkg/api"
	"os"
)

// GRPCServer ...
type GRPCServer struct {
	api.UnimplementedBalanceMsServer
	request requests.Request
}

// AddBalance ...
func (s *GRPCServer) AddBalance(ctx context.Context, req *api.AddBalanceRequest) (*api.AddBalanceResponse, error) {
	err := s.request.AddBalanceRequest(req.Id, req.Money)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error the during execut AddBalance request")
		panic(err)
	}
	return &api.AddBalanceResponse{Status: 0}, nil
}

// GetBalance ...
func (s *GRPCServer) GetBalance(ctx context.Context, req *api.GetBalanceRequest) (*api.GetBalanceResponse, error) {
	balance, err := s.request.GetBalanceRequest(req.Id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error the during execut AddGetBalance request")
		panic(err)
	}
	return &api.GetBalanceResponse{Balance: balance}, nil
}

// Reserve ...
func (s *GRPCServer) Reserve(ctx context.Context, req *api.ReserveRequest) (*api.ReserveResponse, error) {
	err := s.request.AddReserveRequest(req.Id, req.IdService, req.IdOrder, req.Money)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error the during execut AddReserve request")
		panic(err)
	}
	return &api.ReserveResponse{Status: 0}, nil
}

// GetRevenue ...
func (s *GRPCServer) GetRevenue(context.Context, *api.GetRevenueRequest) (*api.GetRevenueResponse, error) {
	re, err := s.request.GetReserveRequest()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error the during execut AddReserve request")
		panic(err)
	}
	var info []string
	for i := 0; i < len(re); i++ {
		str := fmt.Sprintf("Id: %d; IdService: %s; IdOrder: %s; money: %s", re[i].Id, re[i].IdService, re[i].IdOrder, re[i].Money)
		info = append(info, str)
	}

	return &api.GetRevenueResponse{Ans: info, Status: 0}, nil
}

// mustEmbedUnimplementedBalanceMsServer ...
func (s *GRPCServer) mustEmbedUnimplementedBalanceMsServer() {}

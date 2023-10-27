package handlers

import (
	"context"
	"fmt"
	"testing"

	mock_serv "github.com/IDL13/balance_ms/mock"
	"github.com/IDL13/balance_ms/pkg/api"
	"github.com/golang/mock/gomock"
)

func TestAddBalanceRequest(t *testing.T) {
	tests := []struct {
		name         string
		inputId      int64
		inputBalance string
		inputService string
		f            func(s *mock_serv.MockRequest, id int64, balance string)
		expectedBody error
	}{{
		name:         "OK",
		inputId:      1,
		inputBalance: "1000",
		inputService: "test",
		f: func(s *mock_serv.MockRequest, id int64, balance string) {
			s.EXPECT().AddBalanceRequest(id, balance).Return(nil).AnyTimes()
		},
		expectedBody: nil,
	},
		{
			name:         "Error",
			inputId:      10,
			inputBalance: "1000",
			inputService: "test",
			f: func(s *mock_serv.MockRequest, id int64, balance string) {
				s.EXPECT().AddBalanceRequest(id, balance).Return(nil).AnyTimes()
			},
			expectedBody: nil,
		}}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			srvc := mock_serv.NewMockRequest(ctrl)
			test.f(srvc, test.inputId, test.inputBalance)

			service := New()

			api := api.AddBalanceRequest{Id: test.inputId, Money: test.inputBalance}

			_, err := service.AddBalance(context.Background(), &api)

			fmt.Println(err)

			if err != test.expectedBody {
				t.Errorf("Not equal")
			}
		})
	}
}

func TestGetBalanceRequest(t *testing.T) {
	tests := []struct {
		name         string
		inputId      int64
		inputService string
		f            func(s *mock_serv.MockRequest, id int64)
		expectedBody error
	}{{
		name:         "OK",
		inputId:      1,
		inputService: "test",
		f: func(s *mock_serv.MockRequest, id int64) {
			s.EXPECT().GetBalanceRequest(id).Return(nil).AnyTimes()
		},
		expectedBody: nil,
	},
		{
			name:         "Error",
			inputId:      10,
			inputService: "test",
			f: func(s *mock_serv.MockRequest, id int64) {
				s.EXPECT().GetBalanceRequest(id).Return(nil).AnyTimes()
			},
			expectedBody: nil,
		}}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			srvc := mock_serv.NewMockRequest(ctrl)
			test.f(srvc, test.inputId)

			service := New()

			api := api.GetBalanceRequest{Id: test.inputId}

			_, err := service.GetBalance(context.Background(), &api)

			fmt.Println(err)

			if err != test.expectedBody {
				t.Errorf("Not equal")
			}
		})
	}
}

func TestAddReserveRequest(t *testing.T) {
	tests := []struct {
		name           string
		inputId        int64
		inputIdService string
		inputIdOrder   string
		inputMoney     string
		inputService   string
		f              func(s *mock_serv.MockRequest, id int64, idService, idOrder, money string)
		expectedBody   error
	}{{
		name:           "OK",
		inputId:        1,
		inputIdService: "2",
		inputIdOrder:   "3",
		inputMoney:     "1000",
		inputService:   "test",
		f: func(s *mock_serv.MockRequest, id int64, idService, idOrder, money string) {
			s.EXPECT().AddReserveRequest(id, idService, idOrder, money).Return(nil).AnyTimes()
		},
		expectedBody: nil,
	},
		{
			name:           "Error",
			inputId:        1,
			inputIdService: "2",
			inputIdOrder:   "3",
			inputMoney:     "1000",
			inputService:   "test",
			f: func(s *mock_serv.MockRequest, id int64, idService, idOrder, money string) {
				s.EXPECT().AddReserveRequest(id, idService, idOrder, money).Return(nil).AnyTimes()
			},
			expectedBody: nil,
		}}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			srvc := mock_serv.NewMockRequest(ctrl)
			test.f(srvc, test.inputId, test.inputIdService, test.inputIdOrder, test.inputMoney)

			service := New()

			api := api.ReserveRequest{Id: test.inputId, IdService: test.inputIdService, IdOrder: test.inputIdOrder, Money: test.inputMoney}

			_, err := service.Reserve(context.Background(), &api)

			fmt.Println(err)

			if err != test.expectedBody {
				t.Errorf("Not equal")
			}
		})
	}
}

func TestGetReserveRequest(t *testing.T) {
	tests := []struct {
		name           string
		inputId        int64
		inputIdService string
		inputIdOrder   string
		inputMoney     string
		inputService   string
		f              func(s *mock_serv.MockRequest, id int64, idService, idOrder, money string)
		expectedBody   error
	}{{
		name:           "OK",
		inputId:        1,
		inputIdService: "2",
		inputIdOrder:   "3",
		inputMoney:     "1000",
		inputService:   "test",
		f: func(s *mock_serv.MockRequest, id int64, idService, idOrder, money string) {
			s.EXPECT().AddReserveRequest(id, idService, idOrder, money).Return(nil).AnyTimes()
		},
		expectedBody: nil,
	},
		{
			name:           "Error",
			inputId:        1,
			inputIdService: "2",
			inputIdOrder:   "3",
			inputMoney:     "1000",
			inputService:   "test",
			f: func(s *mock_serv.MockRequest, id int64, idService, idOrder, money string) {
				s.EXPECT().AddReserveRequest(id, idService, idOrder, money).Return(nil).AnyTimes()
			},
			expectedBody: nil,
		}}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			srvc := mock_serv.NewMockRequest(ctrl)
			test.f(srvc, test.inputId, test.inputIdService, test.inputIdOrder, test.inputMoney)

			service := New()

			api := api.ReserveRequest{Id: test.inputId, IdService: test.inputIdService, IdOrder: test.inputIdOrder, Money: test.inputMoney}

			_, err := service.Reserve(context.Background(), &api)

			fmt.Println(err)

			if err != test.expectedBody {
				t.Errorf("Not equal")
			}
		})
	}
}

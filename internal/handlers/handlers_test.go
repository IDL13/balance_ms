package handlers

import (
	"context"
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
			inputId:      1,
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

			if err != nil {
				t.Errorf("Not equal")
			}
		})
	}
}

package requests

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddBalanceRequest(t *testing.T) {
	var inputId int64 = 10
	var inputBalance string = "1000"

	req := New()

	err := req.AddBalanceRequest(inputId, inputBalance)

	req.DelUser(inputId)

	assert.Nil(t, err)
}

func TestGetBalanceRequest(t *testing.T) {
	var inputId int64 = 1
	var ExpectedBalance string = "100"

	req := New()

	actualBalance, err := req.GetBalanceRequest(inputId)

	assert.Nil(t, err)
	assert.Equal(t, actualBalance, ExpectedBalance)
}

func TestReserveRequest(t *testing.T) {
	var id int64 = 2
	var idService string = "2"
	var idOrder string = "2"
	var money string = "200"

	req := New()

	err := req.AddReserveRequest(id, idService, idOrder, money)

	req.DelReserve(id)

	assert.Nil(t, err)
}

func TestGetReserveRequest(t *testing.T) {
	expectedData := DataStruct{Id: 1, IdService: "1", IdOrder: "2", Money: "100"}

	req := New()

	actualData, err := req.GetReserveRequest()

	assert.Nil(t, err)
	assert.Equal(t, expectedData, actualData[1])
}

package entity_test

import (
	"testing"

	"github.com/fabiocaettano74/calculate-price/internal/order/entity"
	"github.com/stretchr/testify/assert"
)

func Test_GivenEmptyId_WhenCreateNewOrder_ThenShoulReceivedError(t *testing.T) {
	order := entity.Order{}
	assert.Error(t, order.IsValid(), "Invalid id")
}

func Test_GivenEmptyPrice_WhenCreateNewOrder_ThenShouldReceivedError(t *testing.T) {
	order := entity.Order{ID: "123"}
	assert.Error(t, order.IsValid(), "Invalid Price")
}

func Test_GivenZeroPrice_WhenCreateNewOrder_ThenShouldReceivedError(t *testing.T) {
	order := entity.Order{ID: "123", Price: 0}
	assert.Error(t, order.IsValid(), "Invalid Price")
}

func Test_GivenEmptyTax_WhenCreateNewOrder_ThenShouldReceivedError(t *testing.T) {
	order := entity.Order{ID: "123", Price: 1.19}
	assert.Error(t, order.IsValid(), "Invalid Tax")
}

func Test_GivenZeroTax_WhenCreateNewOrder_ThenShouldReceivedError(t *testing.T) {
	order := entity.Order{ID: "123", Price: 1.19, Tax: 0}
	assert.Error(t, order.IsValid(), "Invalid Tax")
}

func Test_GivenPriceZero_WhenCallCalculateFinalPrice_ThenShoulReceiveError(t *testing.T) {
	order, err := entity.NewOrder("123", 10, 2)
	assert.NoError(t, err)
	err = order.CalculateFinalPrice()
	assert.NoError(t, err)
	assert.Equal(t, 20.0, order.FinalPrice)
}

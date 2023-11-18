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

func Test_GivenEmptyTax_WhenCreateNewOrder_ThenShouldReceivedError(t *testing.T) {
	order := entity.Order{ID: "123", Price: 1.19}
	assert.Error(t, order.IsValid(), "Invalid Tax")
}

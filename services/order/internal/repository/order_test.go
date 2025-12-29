package repository

import (
	"context"

	"github.com/GabrielMoody/MikroNet/services/order/internal/model"
	"github.com/stretchr/testify/mock"
)

type MockDB struct {
	mock.Mock
}

func (m *MockDB) MakeOrder(ctx context.Context, order model.Order)

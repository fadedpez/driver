package common

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type UUIDGenerator struct{}

func (u *UUIDGenerator) NewUUID() string {
	return uuid.New().String()
}

type MockUUIDGenerator struct {
	mock.Mock
}

func (m *MockUUIDGenerator) NewUUID() string {
	args := m.Called()

	return args.Get(0).(string)
}

type UUIDInterface interface {
	NewUUID() string
}

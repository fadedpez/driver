package handlers

import (
	"context"
	"errors"
	"testing"

	"github.com/fadedpez/driver/internal/entities"
	driverapialpha "github.com/fadedpez/driver/protos"
	"github.com/stretchr/testify/mock"

	"github.com/stretchr/testify/assert"

	"github.com/fadedpez/driver/internal/repositories/driver"
)

func setupFixture() *Alpha {
	return &Alpha{
		driverRepo: &driver.MockRepo{},
	}
}

func TestNewAlpha(t *testing.T) {
	t.Run("it requires a config", func(t *testing.T) {
		actual, err := NewAlpha(nil)

		expError := errors.New("a config is required.")

		assert.Nil(t, actual)
		assert.NotNil(t, err)
		assert.Equal(t, expError, err)

	})

	t.Run("it requires a driver repo", func(t *testing.T) {
		actual, err := NewAlpha(&AlphaConfig{
			DriverRepo: nil,
		})

		expError := errors.New("a driver repo is required.")

		assert.Nil(t, actual)
		assert.NotNil(t, err)
		assert.Equal(t, expError, err)

	})
}

func TestAlpha_CreateDriver(t *testing.T) {
	t.Run("it calls the driver repo correctly", func(t *testing.T) {
		handler := setupFixture()
		m := handler.driverRepo.(*driver.MockRepo)

		expDriver := &entities.Driver{
			Name: "stan",
		}

		retDriver := &entities.Driver{
			Name: "stan",
			ID:   "0",
		}

		m.On("CreateDriver", expDriver).Return(retDriver, nil)

		actual, err := handler.CreateDriver(context.Background(), &driverapialpha.CreateDriverRequest{
			Name: "stan",
		})

		assert.Nil(t, err)
		assert.NotNil(t, actual)
		assert.Equal(t, &driverapialpha.CreateDriverResponse{
			Driver: &driverapialpha.Driver{
				Name: retDriver.Name,
				Id:   retDriver.ID,
			},
		}, actual)
		m.AssertExpectations(t)
	})

	t.Run("it returns an error when the driverRepo errors", func(t *testing.T) {
		handler := setupFixture()
		m := handler.driverRepo.(*driver.MockRepo)

		expErr := errors.New("mock driver create failed")

		m.On("CreateDriver", mock.Anything).Return(nil, expErr)

		actual, err := handler.CreateDriver(context.Background(), &driverapialpha.CreateDriverRequest{
			Name: "stan",
		})

		assert.Nil(t, actual)
		assert.NotNil(t, err)
		assert.Equal(t, expErr, err)
		m.AssertExpectations(t)
	})

	t.Run("it requires a name", func(t *testing.T) {
		handler := setupFixture()

		expErr := errors.New("name cannot be empty.")

		actual, err := handler.CreateDriver(context.Background(), &driverapialpha.CreateDriverRequest{
			Name: "",
		})

		assert.Nil(t, actual)
		assert.NotNil(t, err)
		assert.Equal(t, expErr, err)
	})
}

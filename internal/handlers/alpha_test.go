package handlers

import (
	"errors"
	"testing"

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

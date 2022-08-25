package drivers

import (
	"errors"
	"testing"

	"github.com/fadedpez/driver/internal/common"

	"github.com/KirkDiggler/go-projects/dynamo"

	"github.com/stretchr/testify/assert"
)

const (
	testTable = "test_table"
)

func setupFixture() *Dynamo {
	return &Dynamo{
		client:        &dynamo.Mock{},
		tableName:     testTable,
		uuidGenerator: &common.MockUUIDGenerator{},
	}
}

func TestNewDynamo(t *testing.T) {
	t.Run("it requires a config", func(t *testing.T) {
		actual, err := NewDynamo(nil)

		expErr := errors.New("a config is required.")

		assert.Nil(t, actual)
		assert.NotNil(t, err)
		assert.Equal(t, expErr, err)
	})

	t.Run("it requires a client", func(t *testing.T) {
		actual, err := NewDynamo(&DynamoConfig{
			Client: &dynamo.Mock{},
		})

		expErr := errors.New("a client is required.")

		assert.Nil(t, actual)
		assert.NotNil(t, err)
		assert.Equal(t, expErr, err)
	})

	t.Run("it requires a table name", func(t *testing.T) {
		actual, err := NewDynamo(&DynamoConfig{
			Client: &dynamo.Mock{},
		})

		expErr := errors.New("a table name is required.")

		assert.Nil(t, actual)
		assert.NotNil(t, err)
		assert.Equal(t, expErr, err)
	})

	t.Run("it returns a valid handler", func(t *testing.T) {
		actual, err := NewDynamo(&DynamoConfig{
			Client:    &dynamo.Mock{},
			TableName: "abc",
		})

		assert.Nil(t, err)
		assert.NotNil(t, actual)
		assert.NotNil(t, actual.client)
		assert.NotNil(t, actual.uuidGenerator)
		assert.Equal(t, "abc", actual.tableName)
	})
}

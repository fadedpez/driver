package drivers

import (
	"context"
	"errors"
	"testing"

	"github.com/KirkDiggler/go-projects/dynamo/inputs/putitem"

	"github.com/fadedpez/driver/internal/entities"

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

func TestCreateDriver(t *testing.T) {
	t.Run("it requires a driver", func(t *testing.T) {
		fixture := setupFixture()

		actual, err := fixture.CreateDriver(context.Background(), nil)

		assert.Nil(t, actual)
		assert.NotNil(t, err)
		assert.Equal(t, errors.New("a driver is required."), err)
	})

	t.Run("it requires a name", func(t *testing.T) {
		fixture := setupFixture()
		ctx := context.Background()

		actual, err := fixture.CreateDriver(ctx, &entities.Driver{})

		assert.Nil(t, actual)
		assert.NotNil(t, err)
		assert.Equal(t, errors.New("a driver name is required."), err)

	})

	t.Run("it generates a uuid and saves it in the db", func(t *testing.T) {
		fixture := setupFixture()
		ctx := context.Background()

		mclient := fixture.client.(*dynamo.Mock)
		muuid := fixture.uuidGenerator.(*common.MockUUIDGenerator)

		muuid.On("NewUUID").Return("uuid-1")

		mclient.On("PutItem", ctx, fixture.tableName, putitem.NewOptions(
			putitem.WithEntity(&entities.Driver{
				Name: "stan",
				ID:   "uuid-1",
			}))).Return(&putitem.Result{}, nil)

		actual, err := fixture.CreateDriver(context.Background(), &entities.Driver{Name: "stan"})

		assert.Nil(t, err)
		assert.NotNil(t, actual)

		mclient.AssertExpectations(t)
		muuid.AssertExpectations(t)

	})

	t.Run("it returns an error when the client returns an error", func(t *testing.T) {
		fixture := setupFixture()
		ctx := context.Background()

		mclient := fixture.client.(*dynamo.Mock)
		muuid := fixture.uuidGenerator.(*common.MockUUIDGenerator)

		muuid.On("NewUUID").Return("uuid-1")

		mclient.On("PutItem", ctx, fixture.tableName, putitem.NewOptions(
			putitem.WithEntity(&entities.Driver{
				Name: "stan",
				ID:   "uuid-1",
			}))).Return(nil, errors.New("client has returned error."))

		actual, err := fixture.CreateDriver(ctx, &entities.Driver{Name: "stan"})

		assert.Nil(t, actual)
		assert.NotNil(t, err)

		muuid.AssertExpectations(t)
		mclient.AssertExpectations(t)
	})
}

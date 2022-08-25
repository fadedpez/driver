package drivers

import (
	"context"
	"errors"

	"github.com/KirkDiggler/go-projects/dynamo/inputs/putitem"

	"github.com/fadedpez/driver/internal/entities"

	"github.com/fadedpez/driver/internal/common"

	"github.com/KirkDiggler/go-projects/dynamo"
)

type Dynamo struct {
	client        dynamo.Interface
	tableName     string
	uuidGenerator common.UUIDInterface
}

type DynamoConfig struct {
	Client    dynamo.Interface
	TableName string
}

func NewDynamo(cfg *DynamoConfig) (*Dynamo, error) {
	if cfg == nil {
		return nil, errors.New("a config is required.")
	}

	if cfg.Client == nil {
		return nil, errors.New("a client is required.")
	}

	if cfg.TableName == "" {
		return nil, errors.New("a table name is required.")
	}

	return &Dynamo{
		client:        cfg.Client,
		tableName:     cfg.TableName,
		uuidGenerator: &common.UUIDGenerator{},
	}, nil
}

func (r *Dynamo) CreateDriver(ctx context.Context, driver *entities.Driver) (*entities.Driver, error) {
	if driver == nil {
		return nil, errors.New("a driver is required.")
	}

	if driver.Name == "" {
		return nil, errors.New("a driver name is required.")
	}

	driver.ID = r.uuidGenerator.NewUUID()

	_, err := r.client.PutItem(ctx, r.tableName,
		putitem.WithEntity(driver))
	if err != nil {
		return nil, err
	}

	return driver, nil
}

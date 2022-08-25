package drivers

import (
	"errors"

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

package drivers

import (
	"context"

	"github.com/fadedpez/driver/internal/entities"
)

type Repository interface {
	CreateDriver(ctx context.Context, driver *entities.Driver) (*entities.Driver, error)
}

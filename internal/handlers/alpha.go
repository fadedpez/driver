package handlers

import (
	"context"
	"errors"

	"github.com/fadedpez/driver/internal/entities"

	"github.com/fadedpez/driver/internal/repositories/driver"
	driverapialpha "github.com/fadedpez/driver/protos"
)

type Alpha struct {
	driverRepo driver.Repository
}

type AlphaConfig struct {
	DriverRepo driver.Repository
}

func NewAlpha(cfg *AlphaConfig) (*Alpha, error) {
	if cfg == nil {
		return nil, errors.New("a config is required.")
	}

	if cfg.DriverRepo == nil {
		return nil, errors.New("a driver repo is required.")
	}

	return &Alpha{
		driverRepo: cfg.DriverRepo,
	}, nil
}

func (h *Alpha) CreateDriver(ctx context.Context, req *driverapialpha.CreateDriverRequest) (*driverapialpha.CreateDriverResponse, error) {
	if req == nil {
		return nil, errors.New("req is required but not passed in.")
	}

	if req.Name == "" {
		return nil, errors.New("name cannot be empty.")
	}

	driver, err := h.driverRepo.CreateDriver(ctx, &entities.Driver{
		Name: req.Name,
	})

	if err != nil {
		return nil, err
	}

	return &driverapialpha.CreateDriverResponse{
		Driver: &driverapialpha.Driver{
			Id:   driver.ID,
			Name: driver.Name,
		},
	}, nil

}

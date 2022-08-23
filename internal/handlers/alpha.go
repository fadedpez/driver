package handlers

import (
	"context"
	"errors"
	driverapialpha "github.com/fadedpez/driver/protos"
)

type Alpha struct {

}

func (h *Alpha) CreateDriver(ctx context.Context, req *driverapialpha.CreateDriverRequest) (*driverapialpha.CreateDriverResponse, error) {
	return nil, errors.New("not yet implemented")
}
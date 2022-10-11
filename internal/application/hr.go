package model

import (
	hrPb "awesomeProject/internal/proto/hr"
	"context"
)

func (app *Application) CreateHr(
	ctx context.Context,
	in *hrPb.CreateHrRequest,
) (*hrPb.CreateHrResponse, error) {
	panic("implement me")
}

func (app *Application) DeleteHr(
	ctx context.Context,
	in *hrPb.DeleteHrRequest,
) (*hrPb.DeleteHrResponse, error) {
	panic("implement me")
}

func (app *Application) GetHr(
	ctx context.Context,
	in *hrPb.GetHrRequest,
) (*hrPb.GetHrResponse, error) {
	panic("implement me")
}

func (app *Application) CreateHiring(
	ctx context.Context,
	in *hrPb.CreateHiringRequest,
) (*hrPb.CreateHiringResponse, error) {
	panic("implement me")
}

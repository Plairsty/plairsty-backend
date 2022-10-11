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

func (app *Application) DeleteHiring(
	ctx context.Context,
	in *hrPb.DeleteHiringRequest,
) (*hrPb.DeleteHiringResponse, error) {
	panic("implement me")
}

func (app *Application) GetJob(
	ctx context.Context,
	in *hrPb.GetJobRequest,
) (*hrPb.GetJobResponse, error) {
	panic("implement me")
}

func (app *Application) GetAllJobs(
	ctx context.Context,
	in *hrPb.GetAllJobsRequest,
) (*hrPb.GetAllJobsResponse, error) {
	panic("implement me")
}

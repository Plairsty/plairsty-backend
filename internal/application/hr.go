package model

import (
	hrPb "awesomeProject/internal/proto/hr"
	"awesomeProject/utils"
	"context"
	"errors"
)

func (app *Application) CreateHr(
	_ context.Context,
	in *hrPb.CreateHrRequest,
) (*hrPb.CreateHrResponse, error) {
	if !utils.EmailValidator(in.Hr.Email) {
		return nil, errors.New("invalid email")
	}
	id, err := app.persistence.Hr.Insert(in.Hr)
	if err != nil {
		return nil, err
	}
	return &hrPb.CreateHrResponse{
		Id:      int64(id),
		Success: true,
	}, nil

}

func (app *Application) DeleteHr(
	_ context.Context,
	in *hrPb.DeleteHrRequest,
) (*hrPb.DeleteHrResponse, error) {
	err := app.persistence.Hr.Delete(in.GetId())
	if err != nil {
		return nil, err
	}
	return &hrPb.DeleteHrResponse{
		Success: true,
	}, nil
}

func (app *Application) GetHr(
	_ context.Context,
	in *hrPb.GetHrRequest,
) (*hrPb.GetHrResponse, error) {
	dbRes, err := app.persistence.Hr.Get(in.GetId())
	if err != nil {
		return nil, err
	}
	return &hrPb.GetHrResponse{
		Hr: dbRes,
	}, nil
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

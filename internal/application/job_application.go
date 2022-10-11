package model

import (
	jobApplicationPb "awesomeProject/internal/proto/application"
	"context"
)

func (app *Application) AddJobApplication(_ context.Context, in *jobApplicationPb.JobApplicationRequest) (*jobApplicationPb.JobApplicationResponse, error) {
	err := app.persistence.JobApplication.Insert(int(in.Application.GetStudentId()), int(in.Application.GetJobId()))
	if err != nil {
		return nil, err
	}
	return &jobApplicationPb.JobApplicationResponse{
		Success: true,
	}, nil
}

func (app *Application) GetJobApplicationStatus(_ context.Context, in *jobApplicationPb.JobApplicationStatusRequest) (*jobApplicationPb.JobApplicationStatusResponse, error) {
	status, err := app.persistence.JobApplication.Get(int(in.GetId()))
	if err != nil {
		return nil, err
	}
	return &jobApplicationPb.JobApplicationStatusResponse{
		Status: status,
	}, nil
}

func (app *Application) GetAllJobApplicationStatus(_ context.Context, in *jobApplicationPb.AllJobApplicationRequest) (*jobApplicationPb.AllJobApplicationStatusResponse, error) {
	applications, err := app.persistence.JobApplication.GetAll(int(in.GetId()))
	if err != nil {
		return nil, err
	}
	return &jobApplicationPb.AllJobApplicationStatusResponse{
		Application: applications,
	}, nil
}

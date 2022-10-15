package model

import (
	internshipPb "awesomeProject/internal/proto/internship"
	"context"
)

func (app *Application) AddInternship(
	_ context.Context,
	req *internshipPb.AddInternshipRequest,
) (*internshipPb.AddInternshipResponse, error) {
	err := app.persistence.Internship.Insert(req.GetUserId(), req.GetInternship())
	if err != nil {
		return nil, err
	}
	return &internshipPb.AddInternshipResponse{
		Message: "Internship added",
	}, nil
}

func (app *Application) GetInternship(
	_ context.Context,
	req *internshipPb.GetInternshipRequest,
) (*internshipPb.GetInternshipResponse, error) {
	internship, err := app.persistence.Internship.Get(req.GetUserId(), req.GetInternshipId())
	if err != nil {
		return nil, err
	}
	return &internshipPb.GetInternshipResponse{
		Internship: internship,
	}, nil
}

func (app *Application) GetAllInternships(
	_ context.Context,
	req *internshipPb.GetAllInternshipsRequest,
) (*internshipPb.GetAllInternshipsResponse, error) {
	internships, err := app.persistence.Internship.GetAll(req.GetUserId())
	if err != nil {
		return nil, err
	}
	return &internshipPb.GetAllInternshipsResponse{
		Internships: internships,
	}, nil
}

func (app *Application) UpdateInternship(
	_ context.Context,
	req *internshipPb.UpdateInternshipRequest,
) (*internshipPb.UpdateInternshipResponse, error) {
	err := app.persistence.Internship.Update(req.GetUserId(), req.GetInternship())
	if err != nil {
		return nil, err
	}
	return &internshipPb.UpdateInternshipResponse{
		Message: "Internship updated",
	}, nil
}

func (app *Application) DeleteInternship(
	_ context.Context,
	req *internshipPb.DeleteInternshipRequest,
) (*internshipPb.DeleteInternshipResponse, error) {
	err := app.persistence.Internship.Delete(req.GetUserId(), req.GetInternshipId())
	if err != nil {
		return nil, err
	}
	return &internshipPb.DeleteInternshipResponse{
		Message: "Internship deleted",
	}, nil
}

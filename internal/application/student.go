package model

import (
	studentPb "awesomeProject/internal/proto/student"
	"context"
)

func (app *Application) CreateStudent(
	ctx context.Context,
	in *studentPb.CreateStudentRequest,
) (*studentPb.CreateStudentResponse, error) {
	return &studentPb.CreateStudentResponse{
		Success: true,
	}, nil
}

func (app *Application) GetStudent(
	ctx context.Context,
	in *studentPb.GetStudentRequest,
) (*studentPb.GetStudentResponse, error) {
	return &studentPb.GetStudentResponse{
		Student: &studentPb.Student{
			Id:         0,
			FirstName:  "",
			LastName:   "",
			MiddleName: "",
			Email:      "",
			Phone:      0,
		},
	}, nil
}

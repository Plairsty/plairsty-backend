package model

import (
	studentPb "awesomeProject/internal/proto/student"
	"awesomeProject/utils"
	"context"
	"errors"
)

func (app *Application) CreateStudent(
	ctx context.Context,
	in *studentPb.CreateStudentRequest,
) (*studentPb.CreateStudentResponse, error) {
	if !utils.EmailValidator(in.Student.Email) {
		return nil, errors.New("invalid email")
	}
	err := app.persistence.Student.Insert(in.Student)
	if err != nil {
		return nil, err
	}
	return &studentPb.CreateStudentResponse{
		Success: true,
		Message: "student created",
	}, nil
}

func (app *Application) GetStudent(
	ctx context.Context,
	in *studentPb.GetStudentRequest,
) (*studentPb.GetStudentResponse, error) {
	var students []*studentPb.Student
	_, err := app.persistence.Student.Get(10)
	if err != nil {
		return nil, err
	}
	return &studentPb.GetStudentResponse{
		Student: students,
	}, nil
}

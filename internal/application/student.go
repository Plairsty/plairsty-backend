package model

import (
	studentPb "awesomeProject/internal/proto/student"
	"awesomeProject/utils"
	"context"
	"database/sql"
	"errors"
)

func (app *Application) CreateStudent(
	_ context.Context,
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
	_ context.Context,
	in *studentPb.GetStudentRequest,
) (*studentPb.GetStudentResponse, error) {
	var students []*studentPb.Student
	dbRes, err := app.persistence.Student.Get(int64(in.Id))
	if err != nil {
		return nil, err
	}
	students = append(students, dbRes)
	return &studentPb.GetStudentResponse{
		Student: students,
	}, nil
}

func (app *Application) UpdateStudent(
	_ context.Context,
	in *studentPb.UpdateStudentRequest,
) (*studentPb.UpdateStudentResponse, error) {
	_, err := app.persistence.Student.Get(int64(in.GetStudent().GetId()))
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, errors.New("student not found")
		default:
			return nil, err
		}
	}
	err = app.persistence.Student.Update(in.GetStudent())
	if err != nil {
		return nil, err
	}
	return &studentPb.UpdateStudentResponse{
		Success: true,
		Message: "student updated",
	}, nil
}

func (app *Application) IsProfileCompleted(
	_ context.Context,
	in *studentPb.IsProfileCompletedRequest,
) (*studentPb.IsProfileCompletedResponse, error) {
	dbRes, err := app.persistence.Student.CheckProfileStatus(int64(in.GetId()))
	if err != nil {
		return nil, err
	}
	return &studentPb.IsProfileCompletedResponse{
		IsProfileCompleted: dbRes,
	}, nil
}

func (app *Application) GetGPA(
	_ context.Context,
	in *studentPb.GPARequest,
) (*studentPb.GPAResponse, error) {
	dbRes, err := app.persistence.Student.GetGPA(int64(in.GetId()))
	if err != nil {
		return nil, err
	}
	return &studentPb.GPAResponse{
		Gpa: dbRes,
	}, nil
}

func (app *Application) UpdateGPA(
	_ context.Context,
	in *studentPb.UpdateGPARequest,
) (*studentPb.UpdateGPAResponse, error) {
	err := app.persistence.Student.UpdateGPA(int64(in.GetId()), in.GetGpa())
	if err != nil {
		return nil, err
	}
	return &studentPb.UpdateGPAResponse{
		Success: true,
		Message: "gpa updated",
	}, nil
}

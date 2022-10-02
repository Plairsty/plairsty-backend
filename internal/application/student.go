package model

import (
	studentPb "awesomeProject/internal/proto/student"
	"awesomeProject/utils"
	"context"
	"errors"
	"fmt"
	"log"
)

func (app *Application) CreateStudent(
	ctx context.Context,
	in *studentPb.CreateStudentRequest,
) (*studentPb.CreateStudentResponse, error) {
	if !utils.EmailValidator(in.Student.Email) {
		return nil, errors.New("invalid email")
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
	for i := 0; i < 10; i++ {
		var phone []*studentPb.Student_PhoneNumber
		go func() {
			for j := 1; j < 56; j++ {
				phone = append(phone, &studentPb.Student_PhoneNumber{
					Number: fmt.Sprintf("%d", j),
					Type:   studentPb.Student_HOME,
				})
			}
		}()
		students = append(students, &studentPb.Student{
			Id:         int32(i),
			FirstName:  fmt.Sprintf("first_name_%d", i),
			LastName:   fmt.Sprintf("last_name_%d", i),
			MiddleName: fmt.Sprintf("middle_name_%d", i),
			Email:      fmt.Sprintf("mail_%d@gmail.com", i),
			Phones: []*studentPb.Student_PhoneNumber{
				{Number: "555-4321", Type: studentPb.Student_HOME},
			},
		})
	}
	log.Println("Array length: ", len(students))
	return &studentPb.GetStudentResponse{
		Student: students,
	}, nil
}

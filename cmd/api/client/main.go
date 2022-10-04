package main

import (
	sys "awesomeProject/internal/proto/health"
	studentPb "awesomeProject/internal/proto/student"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"math/rand"
	"time"
)

func main() {
	conn1, err := grpc.Dial(
		"localhost:4000",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		panic(err)
	}
	defer func(conn1 *grpc.ClientConn) {
		err := conn1.Close()
		if err != nil {
			panic(err)
		}
	}(conn1)

	healthClient := sys.NewHealthCheckClient(conn1)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	response, err := healthClient.Check(ctx, &sys.HealthCheckRequest{})
	if err != nil {
		return
	}
	println(response.Status)

	conn2, err := grpc.Dial(
		"localhost:4000",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		panic(err)
	}
	defer func(conn2 *grpc.ClientConn) {
		err := conn2.Close()
		if err != nil {
			panic(err)
		}
	}(conn2)

	studentClient := studentPb.NewStudentServiceClient(conn2)
	GetStudent(context.Background(), studentClient)
	CreateStudent(context.Background(), studentClient)
}
func GetStudent(
	ctx context.Context,
	in studentPb.StudentServiceClient,
) {
	r, err := in.GetStudent(ctx, &studentPb.GetStudentRequest{
		Id: 10,
	})
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(r)
}

func CreateStudent(ctx context.Context, in studentPb.StudentServiceClient) {
	r, err := in.CreateStudent(ctx, &studentPb.CreateStudentRequest{
		Student: &studentPb.Student{
			Id:         rand.Int31(),
			FirstName:  "first_name",
			LastName:   "last_name",
			MiddleName: "middle_name",
			Email:      fmt.Sprintf("some_email%d@gmail.com", rand.Int31()),
			Phones: []*studentPb.Student_PhoneNumber{
				{
					Number: "7977421559",
					Type:   1,
				},
			},
			Status: &studentPb.Student_AcademicStatus{
				Status:   studentPb.Student_FRESHMAN,
				Semester: 2,
			},
			FieldOfStudy: "Computer Science",
			IsBanned:     false,
		},
	})
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(r)
}

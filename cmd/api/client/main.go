package main

import (
	sys "awesomeProject/internal/proto/health"
	resumePb "awesomeProject/internal/proto/resume"
	studentPb "awesomeProject/internal/proto/student"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"math/rand"
	"os"
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
	log.Println("Server Status: ", response.Status)

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

	//studentClient := studentPb.NewStudentServiceClient(conn2)
	//GetStudent(context.Background(), studentClient)
	//CreateStudent(context.Background(), studentClient)
	//UpdateStudent(context.Background(), studentClient)
	resumeClient := resumePb.NewResumeServiceClient(conn2)
	UploadPdf(context.Background(), resumeClient)
}
func GetStudent(
	ctx context.Context,
	in studentPb.StudentServiceClient,
) {
	r, err := in.GetStudent(ctx, &studentPb.GetStudentRequest{
		Id: 1298498081,
	})
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Response: ", r)
}

func CreateStudent(ctx context.Context, in studentPb.StudentServiceClient) {
	rand.Seed(time.Now().UnixNano())
	r, err := in.CreateStudent(ctx, &studentPb.CreateStudentRequest{
		Student: &studentPb.Student{
			Id:         rand.Int31(),
			FirstName:  "first_name",
			LastName:   "last_name",
			MiddleName: "middle_name",
			Email:      fmt.Sprintf("some_email%d@gmail.com", rand.Int31()),
			Phones: []*studentPb.Student_PhoneNumber{
				{
					Number: "797742155",
					Type:   1,
				}, {
					Number: "797742155",
					Type:   0,
				}, {
					Number: "097742155",
					Type:   2,
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

func UpdateStudent(ctx context.Context, in studentPb.StudentServiceClient) {
	r, err := in.UpdateStudent(ctx, &studentPb.UpdateStudentRequest{
		Student: &studentPb.Student{
			Id:         1298498081,
			FirstName:  "first_name",
			LastName:   "last_name",
			MiddleName: "middle_name",
			Email:      "",
			Phones: []*studentPb.Student_PhoneNumber{
				{
					Number: "797742155",
					Type:   1,
				}, {
					Number: "797742155",
					Type:   0,
				}, {
					Number: "097742155",
					Type:   2,
				},
			},
			Status:       nil,
			FieldOfStudy: "",
			IsBanned:     false,
			ReasonForBan: "",
			ImageUrl:     "",
		},
	})
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(r)
}

func UploadPdf(ctx context.Context, in resumePb.ResumeServiceClient) {
	stream, err := in.UploadResume(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	// Read a file and send it to the server
	file, err := os.Open("Makefile")
	if err != nil {
		log.Fatalln(err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	buf := make([]byte, 2<<20) // 2 MB
	for {
		n, err := file.Read(buf)
		if err != nil {
			break
		}
		if err := stream.Send(&resumePb.ResumeUploadRequest{
			Resume: &resumePb.Resume{
				Data: buf[:n],
			},
		}); err != nil {
			log.Fatalln(err)
		}
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(res)
}

package main

import (
	sys "awesomeProject/internal/proto/health"
	studentPb "awesomeProject/internal/proto/student"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
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

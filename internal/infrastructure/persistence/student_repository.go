package persistence

import (
	studentPb "awesomeProject/internal/proto/student"
	"database/sql"
	"log"
)

type studentRepository struct {
	DB *sql.DB
}

// Insert injects the student into db
/// Currently I'm assuming that we won't be getting student phone
func (r studentRepository) Insert(student *studentPb.Student) error {
	query := `
		INSERT INTO students (id, first_name, middle_name, last_name, email)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, created_at`
	args := []interface{}{
		student.Id,
		student.FirstName,
		student.MiddleName,
		student.LastName,
		student.Email,
	}
	return r.DB.QueryRow(query, args...).Scan(&student.Id, &student.ImageUrl)
}

func (r studentRepository) Get(id int64) (*studentPb.Student, error) {
	log.Println("Get student")
	return nil, nil
}

func (r studentRepository) Update(student *studentPb.Student) error {
	panic("implement me")
}

func (r studentRepository) Delete(id int64) error {
	panic("implement me")
}

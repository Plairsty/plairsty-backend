package persistence

import (
	studentPb "awesomeProject/internal/proto/student"
	"database/sql"
)

type Repositories struct {
	Student interface {
		Insert(student *studentPb.Student) error
		Get(id int64) (*studentPb.Student, error)
		Update(student *studentPb.Student) error
		Delete(id int64) error
	}
}

func NewRepositories(db *sql.DB) *Repositories {
	return &Repositories{
		Student: studentRepository{DB: db},
	}
}

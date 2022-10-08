package persistence

import (
	resumePb "awesomeProject/internal/proto/resume"
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
	Resume interface {
		Insert(resume *resumePb.Resume, id int) error
		Get(id int64) (string, error)
		Delete(id int64) error
	}
}

func NewRepositories(db *sql.DB, s3 *S3) *Repositories {
	return &Repositories{
		Student: studentRepository{DB: db},
		Resume:  ResumeRepository{DB: db, S3: s3},
	}
}

package persistence

import (
	jobApplicationPb "awesomeProject/internal/proto/application"
	authPb "awesomeProject/internal/proto/auth"
	hrPb "awesomeProject/internal/proto/hr"
	resumePb "awesomeProject/internal/proto/resume"
	studentPb "awesomeProject/internal/proto/student"
	"database/sql"
)

type Repositories struct {
	User interface {
		Insert(user *authPb.UserFields, hashedPassword string) error
		GetByUsername(username string) (*authPb.UserFields, error)
		DeleteByUsername(username string) error
	}
	Student interface {
		Insert(student *studentPb.Student) error
		Get(id int64) (*studentPb.Student, error)
		Update(student *studentPb.Student) error
		Delete(id int64) error

		CheckProfileStatus(id int64) (bool, error)
		GetGPA(id int64) (*studentPb.Gpa, error)
		UpdateGPA(id int64, gpa *studentPb.Gpa) error
	}
	Resume interface {
		Insert(resume *resumePb.Resume, id int) error
		Get(id int64) (string, error)
		Delete(id int64) error
	}

	Hr interface {
		Insert(hr *hrPb.Hr) (int, error)
		Get(id int64) ([]*hrPb.Hr, error)
		Update(hr *hrPb.Hr) error
		Delete(id int64) error
	}

	Job interface {
		Insert(hrId int, job *hrPb.JobFields) error
		Get(job *hrPb.JobSearchQuery) ([]*hrPb.JobFields, error)
		GetAll() ([]*hrPb.JobFields, error)
		Update(job *hrPb.JobFields) error
		Delete(id int64) error
		GetJobById(userId int) (*hrPb.JobFields, error)
	}
	JobApplication interface {
		Insert(userId, jobId int) error
		Get(applicationId int) (bool, error)
		GetAll(userId int) ([]*jobApplicationPb.AllJobApplicationStatus, error)
	}
}

func NewRepositories(db *sql.DB, s3 *S3) *Repositories {
	return &Repositories{
		User:           userRepository{DB: db},
		Student:        studentRepository{DB: db},
		Resume:         ResumeRepository{DB: db, S3: s3},
		Hr:             hrRepository{DB: db},
		Job:            jobRepository{DB: db},
		JobApplication: JobApplicationRepository{DB: db},
	}
}

package persistence

import (
	jobApplicationPb "awesomeProject/internal/proto/application"
	authPb "awesomeProject/internal/proto/auth"
	certificatePb "awesomeProject/internal/proto/certificates"
	hrPb "awesomeProject/internal/proto/hr"
	internshipPb "awesomeProject/internal/proto/internship"
	projectPb "awesomeProject/internal/proto/project"
	resumePb "awesomeProject/internal/proto/resume"
	studentPb "awesomeProject/internal/proto/student"
	"database/sql"
)

type Repositories struct {
	User interface {
		Insert(user *authPb.UserFields, hashedPassword string) error
		GetByUsername(username string) (*authPb.UserFields, error)
		DeleteByUsername(username string) error

		// Get Functions solely for auth
		Get(username string) (hashedPassword, role string, err error)
	}
	Student interface {
		// Insert Please note that, id means here I'm refering to moodle id
		Insert(student *studentPb.Student) error
		Get(id string) (*studentPb.Student, error)
		Update(student *studentPb.Student) error
		Delete(id int64) error

		CheckProfileStatus(id string) (bool, error)
		GetGPA(id string) (*studentPb.Gpa, error)
		UpdateGPA(id string, gpa *studentPb.Gpa) error
	}
	Resume interface {
		Insert(resume *resumePb.Resume, id int) error
		Get(id int64) (string, error)
		Delete(id int64) error
		InsertUrl(url string, id int) error
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

	Certificate interface {
		Insert(userId int64,
			certificateData *certificatePb.CertificateData,
			certificate *certificatePb.CertificateFields) (int64, error)
		// Get certificate by certificate id
		Get(id int64) (*certificatePb.CertificateFields, error)
		// GetAll certificate by student id
		GetAll(id int64) ([]*certificatePb.CertificateFields, error)
		// Update Certificate id, and certificate fields
		Update(userId, certId int64, certificate *certificatePb.CertificateFields) error
		Delete(userId, certId int64) error
		ChangeStatus(id int64, status certificatePb.STATUS) error
	}

	Internship interface {
		Insert(userId int64, internship *internshipPb.InternshipFields) error
		Get(userId, internshipId int64) (*internshipPb.InternshipFields, error)
		GetAll(userId int64) ([]*internshipPb.InternshipFields, error)
		Update(userId int64, internship *internshipPb.InternshipFields) error // Internship id in field
		Delete(userId, internshipId int64) error
	}

	Project interface {
		Insert(username string, project *projectPb.ProjectFields) error
		Get(username string, projectId int64) (*projectPb.ProjectFields, error)
		GetAll(username string) ([]*projectPb.ProjectFields, error)
		Update(username string, project *projectPb.ProjectFields) error // Project id in field
		Delete(username string, projectId int64) error
		GetProjectsBySemester(username string, semester int64) ([]*projectPb.ProjectFields, error)
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
		Certificate:    CertificateRepository{DB: db, S3: s3},
		Internship:     InternshipRepository{DB: db},
		Project:        ProjectRepository{DB: db},
	}
}

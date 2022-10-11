package persistence

import (
	jobApplicationPb "awesomeProject/internal/proto/application"
	"database/sql"
)

type JobApplicationRepository struct {
	DB *sql.DB
}

func (r JobApplicationRepository) Insert(userId, jobId int) error {
	query := `INSERT INTO application (user_id, job_id) VALUES ($1, $2)`
	err := r.DB.QueryRow(query, userId, jobId).Scan()
	if err != nil {
		return err
	}
	return nil
}

func (r JobApplicationRepository) Get(applicationId int) (bool, error) {
	query := `SELECT status FROM application WHERE id = $1`
	var status bool
	err := r.DB.QueryRow(query, applicationId).Scan(&status)
	if err != nil {
		return false, err
	}
	return status, nil
}

func (r JobApplicationRepository) GetAll(userId int) ([]*jobApplicationPb.AllJobApplicationStatus, error) {
	query := `
		SELECT application.id, jobs.company, jobs.role, application.status
		FROM application
		INNER JOIN jobs ON application.job_id = jobs.id
		WHERE application.user_id = $1
		`
	var applications []*jobApplicationPb.AllJobApplicationStatus
	rows, err := r.DB.Query(query, userId)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var application jobApplicationPb.AllJobApplicationStatus
		err := rows.Scan(&application.Id, &application.CompanyName, &application.Role, &application.Status)
		if err != nil {
			return nil, err
		}
		applications = append(applications, &application)
	}
	return applications, nil
}

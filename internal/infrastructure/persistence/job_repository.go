package persistence

import (
	hrPb "awesomeProject/internal/proto/hr"
	"database/sql"
)

type jobRepository struct {
	DB *sql.DB
}

func (r jobRepository) Insert(job *hrPb.Job) error {
	panic("implement me")
}

func (r jobRepository) Get(query *hrPb.JobSearchQuery) (*hrPb.Job, error) {
	panic("implement me")
}

func (r jobRepository) Update(job *hrPb.Job) error {
	panic("implement me")
}

func (r jobRepository) Delete(id int64) error {
	panic("implement me")
}

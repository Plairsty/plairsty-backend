package persistence

import (
	hrPb "awesomeProject/internal/proto/hr"
	"database/sql"
)

type hrRepository struct {
	DB *sql.DB
}

func (r hrRepository) Insert(hr *hrPb.Hr) error {
	panic("implement me")
}

func (r hrRepository) Get(id int64) (*hrPb.Hr, error) {
	panic("implement me")
}

func (r hrRepository) Update(hr *hrPb.Hr) error {
	panic("implement me")
}

func (r hrRepository) Delete(id int64) error {
	panic("implement me")
}

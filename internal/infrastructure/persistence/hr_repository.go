package persistence

import (
	hrPb "awesomeProject/internal/proto/hr"
	"database/sql"
)

type hrRepository struct {
	DB *sql.DB
}

func (r hrRepository) Insert(hr *hrPb.Hr) error {
	Query := `
		INSERT INTO hr (name, phone, email, company) 
		VALUES ($1, $2, $3, $4)
		RETURNING id
		`
	err := r.DB.QueryRow(Query, hr.Name, hr.Phone, hr.Email, hr.Company).Scan(&hr.Email)
	if err != nil {
		return err
	}
	return nil
}

func (r hrRepository) Get(id int64) (*hrPb.Hr, error) {
	Query := `
		SELECT id, name, phone, email, company
		FROM hr
		WHERE id = $1
		`
	var hr hrPb.Hr
	err := r.DB.QueryRow(Query, id).Scan(&hr.Id, &hr.Name, &hr.Phone, &hr.Email, &hr.Company)
	if err != nil {
		return nil, err
	}
	return &hr, nil
}

func (r hrRepository) Update(hr *hrPb.Hr) error {
	Query := `
		UPDATE hr
		SET name = $1, phone = $2, email = $3, company = $4
		WHERE id = $5
		`
	_, err := r.DB.Exec(Query, hr.Name, hr.Phone, hr.Email, hr.Company, hr.Id)
	if err != nil {
		return err
	}
	return nil
}

func (r hrRepository) Delete(id int64) error {
	Query := `
		DELETE FROM hr
		WHERE id = $1
		`
	_, err := r.DB.Exec(Query, id)
	if err != nil {
		return err
	}
	return nil
}

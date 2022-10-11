package persistence

import (
	hrPb "awesomeProject/internal/proto/hr"
	"database/sql"
)

type hrRepository struct {
	DB *sql.DB
}

func (r hrRepository) Insert(hr *hrPb.Hr) (int, error) {
	Query := `
		INSERT INTO hr (name, phone, email, company) 
		VALUES ($1, $2, $3, $4)
		RETURNING id
		`
	var id int
	err := r.DB.QueryRow(Query, hr.Name, hr.Phone, hr.Email, hr.Company).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r hrRepository) Get(id int64) ([]*hrPb.Hr, error) {
	if id != 0 {
		Query := `
		SELECT id, name, phone, email, company
		FROM hr
		WHERE id = $1
		`
		var hrs []*hrPb.Hr
		rows, err := r.DB.Query(Query, id)
		if err != nil {
			return nil, err
		}
		for rows.Next() {
			var hr hrPb.Hr
			err := rows.Scan(&hr.Id, &hr.Name, &hr.Phone, &hr.Email, &hr.Company)
			if err != nil {
				return nil, err
			}
			hrs = append(hrs, &hr)
		}
		return hrs, nil
	}
	// If no one is present return all
	Query := `
		SELECT id, name, phone, email, company
		FROM hr
		`
	var hrs []*hrPb.Hr
	rows, err := r.DB.Query(Query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var hr hrPb.Hr
		err := rows.Scan(&hr.Id, &hr.Name, &hr.Phone, &hr.Email, &hr.Company)
		if err != nil {
			return nil, err
		}
		hrs = append(hrs, &hr)
	}
	return hrs, nil
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

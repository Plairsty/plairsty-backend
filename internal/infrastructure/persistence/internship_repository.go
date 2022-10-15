package persistence

import (
	internshipPb "awesomeProject/internal/proto/internship"
	"database/sql"
)

type InternshipRepository struct {
	DB *sql.DB
}

func (i InternshipRepository) Insert(
	userId int64,
	internship *internshipPb.InternshipFields) error {
	query := `
			INSERT INTO internship (user_id, title, description, start_date, end_date, company_name, location, mentor_name) 
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
			`
	_, err := i.DB.Exec(
		query,
		userId,
		internship.GetName(),
		internship.GetDescription(),
		internship.GetStartDate(),
		internship.GetEndDate(),
		internship.GetCompany(),
		internship.GetLocation(),
		internship.GetMentorName(),
	)
	return err
}

func (i InternshipRepository) Get(userId, internshipId int64) (*internshipPb.InternshipFields, error) {
	query := `
			SELECT title, description, start_date, end_date, company_name, location, mentor_name
			FROM internship
			WHERE id = $1 AND user_id = $2
			`
	internship := &internshipPb.InternshipFields{}
	err := i.DB.QueryRow(query, internshipId, userId).Scan(
		&internship.Name,
		&internship.Description,
		&internship.StartDate,
		&internship.EndDate,
		&internship.Company,
		&internship.Location,
		&internship.MentorName,
	)
	return internship, err
}

func (i InternshipRepository) GetAll(userId int64) ([]*internshipPb.InternshipFields, error) {
	query := `
			SELECT id, title, description, start_date, end_date, company_name, location, mentor_name
			FROM internship
			WHERE user_id = $1
			`
	rows, err := i.DB.Query(query, userId)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			return
		}
	}(rows)

	var internships []*internshipPb.InternshipFields
	for rows.Next() {
		internship := &internshipPb.InternshipFields{}
		err := rows.Scan(
			&internship.Id,
			&internship.Name,
			&internship.Description,
			&internship.StartDate,
			&internship.EndDate,
			&internship.Company,
			&internship.Location,
			&internship.MentorName,
		)
		if err != nil {
			return nil, err
		}
		internships = append(internships, internship)
	}
	return internships, nil
}

func (i InternshipRepository) Update(userId int64, internship *internshipPb.InternshipFields) error {
	query := `
			UPDATE internship
			SET title = $1, description = $2, start_date = $3, end_date = $4, company_name = $5, location = $6, mentor_name = $7
			WHERE id = $8 AND user_id = $9
			`
	_, err := i.DB.Exec(
		query,
		internship.GetName(),
		internship.GetDescription(),
		internship.GetStartDate(),
		internship.GetEndDate(),
		internship.GetCompany(),
		internship.GetLocation(),
		internship.GetMentorName(),
		internship.GetId(),
		userId,
	)
	return err
}

func (i InternshipRepository) Delete(userId, internshipId int64) error {
	query := `
			DELETE FROM internship
			WHERE id = $1 AND user_id = $2
			`
	_, err := i.DB.Exec(query, internshipId, userId)
	return err
}

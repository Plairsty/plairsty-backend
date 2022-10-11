package persistence

import (
	hrPb "awesomeProject/internal/proto/hr"
	"database/sql"
)

type jobRepository struct {
	DB *sql.DB
}

func (r jobRepository) Insert(hrId int, job *hrPb.JobFields) error {
	query := `
		INSERT INTO jobs (role, department, skills, experience, required_cgpa, description, location, certifications, title, company, hr_id) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
		RETURNING id`
	args := []interface{}{
		job.GetRole(),
		job.GetDepartment(),
		job.GetSkills(),
		job.GetExperience(),
		job.GetRequiredCgpa(),
		job.GetDescription(),
		job.GetLocation(),
		job.GetCertifications(),
		job.GetTitle(),
		job.GetCompany(),
		hrId,
	}
	_, err := r.DB.Exec(query, args...)
	if err != nil {
		return err
	}
	return nil
}

func (r jobRepository) Get(query *hrPb.JobSearchQuery) ([]*hrPb.JobFields, error) {
	Query := ``
	var jobs []*hrPb.JobFields
	var rows *sql.Rows
	var err error
	if query.GetId() != 0 {
		Query = `
		SELECT id, certifications, company, description, department, experience, location, required_cgpa, role, skills, title, hr_id
		FROM jobs
		WHERE LOWER(id) = LOWER($1)
		`
		rows, err = r.DB.Query(Query, query.GetId())
		if err != nil {
			return nil, err
		}
	}
	if query.GetTitle() != "" {
		Query = `
		SELECT id, certifications, company, description, department, experience, location, required_cgpa, role, skills, title, hr_id
		FROM jobs
		WHERE LOWER(title) like LOWER($1)
		`
		rows, err = r.DB.Query(Query, query.GetTitle())
		if err != nil {
			return nil, err
		}
	}
	if query.GetLocation() != "" {
		Query = `
		SELECT id, certifications, company, description, department, experience, location, required_cgpa, role, skills, title, hr_id
		FROM jobs
		WHERE LOWER(location) = LOWER($1)
		`
		rows, err = r.DB.Query(Query, query.GetLocation())
		if err != nil {
			return nil, err
		}
	}
	if query.GetCompany() != "" {
		// Ignore upper and lower case
		Query = `
		SELECT id, certifications, company, description, department, experience, location, required_cgpa, role, skills, title, hr_id
		FROM jobs
		WHERE LOWER(company) = LOWER($1)
		`
		rows, err = r.DB.Query(Query, query.GetCompany())
		if err != nil {
			return nil, err
		}
	}
	for rows.Next() {
		var job hrPb.JobFields
		err := rows.Scan(
			&job.Id,
			&job.Certifications,
			&job.Company,
			&job.Description,
			&job.Department,
			&job.Experience,
			&job.Location,
			&job.RequiredCgpa,
			&job.Role,
			&job.Skills,
			&job.Title,
			&job.HrId,
		)
		if err != nil {
			return nil, err
		}
		jobs = append(jobs, &job)
	}
	return jobs, nil
}

func (r jobRepository) GetAll() ([]*hrPb.JobFields, error) {
	Query := `
		SELECT id, certifications, company, description, department, experience, location, required_cgpa, role, skills, title, hr_id
		FROM jobs
		`
	var jobs []*hrPb.JobFields
	rows, err := r.DB.Query(Query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var job hrPb.JobFields
		err := rows.Scan(
			&job.Id,
			&job.Certifications,
			&job.Company,
			&job.Description,
			&job.Department,
			&job.Experience,
			&job.Location,
			&job.RequiredCgpa,
			&job.Role,
			&job.Skills,
			&job.Title,
			&job.HrId,
		)
		if err != nil {
			return nil, err
		}
		jobs = append(jobs, &job)
	}
	return jobs, nil
}

func (r jobRepository) Update(job *hrPb.JobFields) error {
	panic("implement me")
}

func (r jobRepository) Delete(id int64) error {
	panic("implement me")
}

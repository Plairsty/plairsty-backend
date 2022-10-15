package persistence

import (
	studentPb "awesomeProject/internal/proto/student"
	"database/sql"
	"errors"
	"strings"
)

type studentRepository struct {
	DB *sql.DB
}

// Insert injects the student into db
/// Currently I'm assuming that we won't be getting student phone
func (r studentRepository) Insert(student *studentPb.Student) error {
	query := `
		INSERT INTO students (first_name, middle_name, last_name, email, moodle, status, image_url)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, created_at`
	args := []interface{}{
		strings.Trim(student.FirstName, " "),
		strings.Trim(student.MiddleName, " "),
		strings.Trim(student.LastName, " "),
		strings.Trim(student.Email, " "),
		strings.Trim(student.Username, " "),
		true,
		strings.Trim(student.ImageUrl, " "),
	}
	err := r.DB.QueryRow(query, args...).Scan(&student.Id, &student.ImageUrl)
	if err != nil {
		return err
	}
	// Create a new student phone
	phoneQuery := `
		INSERT INTO phone (student_id, phone, phone_type)
		VALUES ($1, $2, $3)
		RETURNING id`
	for _, phone := range student.Phones {
		// Check if phone number is valid
		if phone.Number == "" {
			continue
		}
		_, err := r.DB.Exec(phoneQuery, student.Id, phone.Number, phone.Type)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r studentRepository) Get(id string) (*studentPb.Student, error) {
	// if length of id is not 8 then return error

	// TODO: Uncomment this when using in production
	//if id < 10000000 || id > 99999999 {
	//	return nil, errors.New("invalid id")
	//}
	query := `
		SELECT id, first_name, middle_name, last_name, email, image_url, moodle, status
		FROM students
		WHERE moodle = $1`
	var student studentPb.Student
	err := r.DB.QueryRow(query, id).Scan(
		&student.Id,
		&student.FirstName,
		&student.MiddleName,
		&student.LastName,
		&student.Email,
		&student.ImageUrl,
		&student.Username,
		&student.IsProfileCompleted,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, errors.New("student not found")
		default:
			return nil, err
		}
	}
	return &student, nil
}

func (r studentRepository) Update(student *studentPb.Student) error {
	query := `
		UPDATE students
		SET first_name = $1, middle_name = $2, last_name = $3, email = $4, image_url = $5
		WHERE moodle = $6
		RETURNING moodle`
	args := []interface{}{
		student.FirstName,
		student.MiddleName,
		student.LastName,
		student.Email,
		student.ImageUrl,
		student.Id,
	}
	return r.DB.QueryRow(query, args...).Scan(&student.Id)
}

func (r studentRepository) Delete(id int64) error {
	query := `
		DELETE FROM students
		WHERE moodle = $1`
	_, err := r.DB.Exec(query, id)
	return err
}

func (r studentRepository) CheckProfileStatus(id string) (bool, error) {
	query := `
		SELECT status
		FROM students
		WHERE moodle = $1`
	var exists bool
	err := r.DB.QueryRow(query, id).Scan(&exists)
	return exists, err
}

func (r studentRepository) GetGPA(id string) (*studentPb.Gpa, error) {
	query := `
		SELECT semester1, semester2, semester3, semester4, semester5, semester6, semester7, semester8
		FROM students
		WHERE moodle = $1`

	var gpa studentPb.Gpa
	err := r.DB.QueryRow(query, id).Scan(
		&gpa.Gpa_1,
		&gpa.Gpa_2,
		&gpa.Gpa_3,
		&gpa.Gpa_4,
		&gpa.Gpa_5,
		&gpa.Gpa_6,
		&gpa.Gpa_7,
		&gpa.Gpa_8,
	)
	if err != nil {
		return nil, err
	}
	return &gpa, nil
}

func (r studentRepository) UpdateGPA(id string, gpa *studentPb.Gpa) error {
	query := `
		UPDATE students
		SET semester1 = $1, semester2 = $2, semester3 = $3, semester4 = $4, semester5 = $5, semester6 = $6, semester7 = $7, semester8 = $8
		WHERE moodle = $9`
	args := []interface{}{
		gpa.Gpa_1,
		gpa.Gpa_2,
		gpa.Gpa_3,
		gpa.Gpa_4,
		gpa.Gpa_5,
		gpa.Gpa_6,
		gpa.Gpa_7,
		gpa.Gpa_8,
		id,
	}
	_, err := r.DB.Exec(query, args...)
	return err
}

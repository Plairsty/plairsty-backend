package persistence

import (
	authPb "awesomeProject/internal/proto/auth"
	"database/sql"
)

type userRepository struct {
	DB *sql.DB
}

/*
Insert(user *authPb.UserFields) error
GetByUsername(username string) (*authPb.UserFields, error)
DeleteByUsername(username string) error
*/

func (u userRepository) Insert(user *authPb.UserFields, hashedPassword string) error {
	query := `
		INSERT INTO "user" (username, hashedpassword, type) 
		VALUES ($1, $2, $3)
		RETURNING id
		`
	var id int
	err := u.DB.QueryRow(query, user.Username, hashedPassword, user.Type).Scan(&id)
	if err != nil {
		return err
	}
	return nil
}

func (u userRepository) GetByUsername(username string) (*authPb.UserFields, error) {
	query := `
		SELECT username, type
		FROM "user"
		WHERE username = $1
		`
	var user authPb.UserFields
	err := u.DB.QueryRow(query, username).Scan(&user.Username, &user.Type)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u userRepository) Get(username string) (hashedPassword, role string, err error) {
	query := `
		SELECT hashedpassword, type
		FROM "user"
		WHERE username = $1
		`
	err = u.DB.QueryRow(query, username).Scan(&hashedPassword, &role)
	if err != nil {
		return "", "", err
	}
	return hashedPassword, role, nil
}

func (u userRepository) DeleteByUsername(username string) error {
	query := `
		DELETE FROM "user"
		WHERE username = $1
		`
	_, err := u.DB.Exec(query, username)
	if err != nil {
		return err
	}
	return nil
}

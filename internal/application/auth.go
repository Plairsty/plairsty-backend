package model

import (
	"awesomeProject/internal/infrastructure/service"
	authPb "awesomeProject/internal/proto/auth"
	"context"
	"database/sql"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

func (app *Application) Register(
	_ context.Context,
	req *authPb.RegisterRequest,
) (*authPb.RegisterResponse, error) {
	generatedPassword := service.GeneratePassword()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(generatedPassword), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	err = app.persistence.User.Insert(req.GetUser(), string(hashedPassword))
	if err != nil {
		return nil, err
	}
	return &authPb.RegisterResponse{
		Password: generatedPassword,
		Success:  false,
	}, nil
}

func (app *Application) Login(
	_ context.Context,
	req *authPb.LoginRequest,
) (*authPb.LoginResponse, error) {
	hashedPassword, role, err := app.persistence.User.Get(req.GetUsername())
	if err == sql.ErrNoRows {
		return nil, errors.New("user not found")
	}
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(req.GetPassword()))
	if err != nil {
		return nil, errors.New("invalid password")
	}
	token, err := app.jwtManager.GenerateToken(req.GetUsername(), role)
	if err != nil {
		return nil, err
	}
	return &authPb.LoginResponse{
		AccessToken: token,
	}, nil
}

func (app *Application) Logout(
	_ context.Context,
	_ *authPb.LogoutRequest,
) (*authPb.LogoutResponse, error) {
	panic("implement me")
}

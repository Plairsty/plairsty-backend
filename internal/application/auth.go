package model

import (
	authPb "awesomeProject/internal/proto/auth"
	"context"
	"database/sql"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log"
	"strings"
)

func (app *Application) Register(
	_ context.Context,
	req *authPb.RegisterRequest,
) (*authPb.RegisterResponse, error) {
	//generatedPassword := service.GeneratePassword()
	generatedPassword := req.GetUser().GetUsername() + "@123"
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
	hashedPassword, role, err := app.persistence.User.Get(strings.Trim(req.GetUsername(), " "))
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
	// Remove whitespace from username
	username := strings.Trim(req.GetUsername(), " ")
	token, err := app.jwtManager.GenerateToken(username, role)
	if err != nil {
		return nil, err
	}
	log.Println("User logged in")
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

func (app *Application) Authorize(
	ctx context.Context,
	method string,
) error {
	accessibleRole, ok := app.accessibleRoles[method]
	if !ok {
		return nil // no need to authorize this route which is not present in accessibleRoles
	}
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return status.Error(codes.Unauthenticated, "metadata not found")
	}
	values := md["authorization"]
	if len(values) == 0 {
		return status.Error(codes.Unauthenticated, "authorization token not found")
	}
	accessToken := values[0]
	claims, err := app.jwtManager.ValidateToken(accessToken)
	if err != nil {
		return status.Error(codes.Unauthenticated, "invalid token")
	}
	for _, role := range accessibleRole {
		if role == claims.Role {
			return nil
		}
	}
	return status.Error(codes.PermissionDenied, "permission denied")
}

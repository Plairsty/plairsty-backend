package service

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"log"

	"time"
)

type JwtManager struct {
	secretKey     string
	issuer        string
	tokenDuration time.Duration
}

type UserClaims struct {
	jwt.StandardClaims
	Username string `json:"username"` // moodle or hr username
	Role     string `json:"role"`
}

type JwtToken struct {
	AccessToken string `json:"token"`
}

func NewJwtManager(
	secretKey string,
	issuer string,
	tokenDuration time.Duration,
) *JwtManager {
	log.Println("Jwt manager created 👾")
	return &JwtManager{
		secretKey:     secretKey,
		issuer:        issuer,
		tokenDuration: tokenDuration,
	}
}

func userClaimProvider(
	username, role string,
	expiresAt time.Duration,
) UserClaims {
	return UserClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(expiresAt).Unix(),
			Issuer:    "plairsty",
			Subject:   "access_token",
		},
		Username: username,
		Role:     role,
	}
}

func (m *JwtManager) GenerateToken(username, role string) (string, error) {
	claims := userClaimProvider(username, role, m.tokenDuration)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(m.secretKey))
}

func (m *JwtManager) ValidateToken(tokenString string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&UserClaims{},
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(m.secretKey), nil
		},
	)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*UserClaims)
	if !ok {
		return nil, err
	}
	return claims, nil
}

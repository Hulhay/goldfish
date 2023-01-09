package usecase

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/Hulhay/goldfish/model"
	"github.com/Hulhay/goldfish/usecase/token"
	"github.com/dgrijalva/jwt-go"
)

type Token interface {
	GenerateToken(ctx context.Context, user *model.User) (*token.ResultResponse, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtService struct {
	secretKey string
}

func NewTokenUc() Token {
	return &jwtService{
		secretKey: getSecretKey(),
	}
}

func getSecretKey() string {
	secretKey := os.Getenv("SECRET_KEY")
	return secretKey
}

func (u *jwtService) GenerateToken(ctx context.Context, user *model.User) (*token.ResultResponse, error) {

	now := time.Now()
	end := now.Add(time.Minute * 60)
	claims := &token.AccessCustomClaim{
		ID:    user.ID,
		Name:  user.UserName,
		Email: user.UserEmail,
		Role:  user.UserRole,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  now.Unix(),
			ExpiresAt: end.Unix(),
		},
	}

	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := newToken.SignedString([]byte(u.secretKey))
	if err != nil {
		return nil, err
	}

	res := &token.ResultResponse{
		Name:      user.UserName,
		Token:     tokenStr,
		Role:      user.UserRole,
		ExpiredAt: end.Format(time.RFC3339),
	}

	return res, nil
}

func (u *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", t.Header["alg"])
		}
		return []byte(u.secretKey), nil
	})
}

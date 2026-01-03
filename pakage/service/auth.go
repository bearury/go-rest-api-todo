package service

import (
	"crypto/sha1"
	"fmt"
	"time"

	"github.com/bearury/go-rest-api-postgres-todo/entitys"
	"github.com/bearury/go-rest-api-postgres-todo/pakage/repository"
	"github.com/dgrijalva/jwt-go"
)

const (
	solt       = "dsffsfdsgv45ttrdsfg"
	signingKey = "secret"
	tokenTTL   = time.Hour * 12
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

type AuthService struct {
	repo repository.AuthorizationRepository
}

func NewAuthService(repo repository.AuthorizationRepository) *AuthService {
	return &AuthService{repo: repo}
}

func (service *AuthService) CreateUser(user entitys.User) (int, error) {
	user.Password = service.generatePasswordHash(user.Password)
	return service.repo.CreateUser(user)
}

func (service *AuthService) GenerateToken(username, password string) (string, error) {

	user, err := service.repo.GetUser(username, service.generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 12).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})

	return token.SignedString([]byte(signingKey))
}

func (service *AuthService) ParseToken(token string) (int, error) {
	accessToken, err := jwt.ParseWithClaims(token, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := accessToken.Claims.(*tokenClaims)
	if !ok {
		return 0, fmt.Errorf("Invalid token claims")
	}

	return claims.UserId, nil
}

func (service *AuthService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(solt)))
}

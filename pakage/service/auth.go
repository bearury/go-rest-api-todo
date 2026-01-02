package service

import (
	"crypto/sha1"
	"fmt"

	"github.com/bearury/go-rest-api-postgres-todo/entitys"
	"github.com/bearury/go-rest-api-postgres-todo/pakage/repository"
)

const solt = "dsffsfdsgv45ttrdsfg"

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

func (service *AuthService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(solt)))
}

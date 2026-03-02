package service

import (
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/te-fa-bene/api-go/internal/app/repository"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var ErrInvalidCredentials = errors.New("Invalid credentials")

type AuthService struct {
	employees *repository.EmployeeRepository
}

func NewAuthService(employees *repository.EmployeeRepository) *AuthService {
	return &AuthService{
		employees: employees,
	}
}

type LoginResult struct {
	AccessToken string
	ExpiresIn   int
}

func (s *AuthService) Login(storeID, email, password string) (*LoginResult, error) {
	employee, err := s.employees.FindActiveByStoreAndEmail(storeID, email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrInvalidCredentials
		}
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(employee.PasswordHash), []byte(password)); err != nil {
		return nil, ErrInvalidCredentials
	}

	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return nil, errors.New("JWT_SECRET is not set")
	}

	expiresIn := 28800
	if v := os.Getenv("JWT_EXPIRES_IN_SECONDS"); v != "" {
		if n, convErr := strconv.Atoi(v); convErr == nil && n > 0 {
			expiresIn = n
		}
	}

	now := time.Now().UTC()
	claims := jwt.MapClaims{
		"sub":      employee.ID,
		"store_id": employee.StoreID,
		"role":     employee.Role,
		"iat":      now.Unix(),
		"exp":      now.Add(time.Duration(expiresIn) * time.Second).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString([]byte(secret))
	if err != nil {
		return nil, err
	}

	return &LoginResult{
		AccessToken: signed,
		ExpiresIn:   expiresIn,
	}, nil
}

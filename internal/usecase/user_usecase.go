package usecase

import (
	"blog-api/internal/domain"
	"blog-api/internal/repository"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecase struct {
	repo      repository.UserRepository
	secretKey string
}

func NewUserUsecase(repo repository.UserRepository, secretKey string) *UserUsecase {
	return &UserUsecase{repo: repo, secretKey: secretKey}
}

func (u *UserUsecase) Register(user domain.User) (domain.User, error) {

	if user.Username == "" {
		return domain.User{}, errors.New("username is required")
	}
	if user.Password == "" {
		return domain.User{}, errors.New("password is required")
	}
	if len(user.Username) < 3 {
		return domain.User{}, errors.New("username must be at least 3 character")
	}
	if len(user.Password) < 6 {
		return domain.User{}, errors.New("password must be at least 6 characters")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return domain.User{}, err
	}
	user.Password = string(hashedPassword)

	return u.repo.Create(user)
}

func (u *UserUsecase) Login(username string, password string) (string, error) {

	if username == "" {
		return "", errors.New("username is required")
	}
	if password == "" {
		return "", errors.New("password is required")
	}

	user, err := u.repo.GetByUsername(username)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(password),
	)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	claims := jwt.MapClaims{
		"user_id":  user.ID,
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(u.secretKey))
	if err != nil {
		return "", errors.New("failed to generate token")
	}
	return tokenString, nil
}

func (u *UserUsecase) GetById(id int) (domain.User, error) {
	if id <= 0 {
		return domain.User{}, errors.New("invalid user id")
	}
	user, err := u.repo.GetByID(id)
	if err != nil {
		return domain.User{}, errors.New("user not found")
	}
	return user, nil
}

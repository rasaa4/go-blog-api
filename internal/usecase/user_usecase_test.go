package usecase

import (
	"blog-api/internal/domain"
	"testing"
)

type MockUserRepo struct {
}

func (m *MockUserRepo) Create(user domain.User) (domain.User, error) {
	return user, nil
}
func (m *MockUserRepo) GetByUsername(username string) (domain.User, error) {
	return domain.User{}, nil
}
func (m *MockUserRepo) GetByID(id int) (domain.User, error) {
	return domain.User{
		ID:       id,
		Username: "rasa",
	}, nil
}

func TestRegisterEmptyUsername(t *testing.T) {
	repo := &MockUserRepo{}
	uc := NewUserUsecase(
		repo,
		"secret",
	)
	_, err := uc.Register(domain.User{
		Username: "",
		Password: "123456",
	})
	if err == nil {
		t.Fatal("expected error")
	}

}
func TestRegisterEmptyPassword(t *testing.T) {
	repo := &MockUserRepo{}
	uc := NewUserUsecase(repo, "secret")
	_, err := uc.Register(domain.User{
		Username: "Rasa",
		Password: "",
	})
	if err == nil {
		t.Fatal("expected error")
	}

}
func TestRegisterShortPassword(t *testing.T) {
	repo := &MockUserRepo{}
	uc := NewUserUsecase(repo, "secret")
	_, err := uc.Register(domain.User{
		Username: "Rasa",
		Password: "123",
	})
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestRegisterSuccess(t *testing.T) {

	repo := &MockUserRepo{}
	uc := NewUserUsecase(repo, "secret")
	user, err := uc.Register(domain.User{
		Username: "rasa",
		Password: "123456",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if user.Username != "rasa" {
		t.Fatal("username mismatch")
	}
	if user.Password == "123456" {
		t.Fatal("password should be hashed")
	}

}

func TestLoginEmptyUsername(t *testing.T) {
	repo := &MockUserRepo{}
	uc := NewUserUsecase(repo, "secret")
	_, err := uc.Login("", "1234")
	if err.Error() != "username is required" {
		t.Fatalf("unexpected error: %v", err)
	}
}

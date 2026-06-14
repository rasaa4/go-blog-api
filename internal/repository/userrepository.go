package repository

import "blog-api/internal/domain"

type UserRepository interface {
	Create(user domain.User) (domain.User, error)
	GetByUsername(username string) (domain.User, error)
	GetByID(id int) (domain.User, error)
}

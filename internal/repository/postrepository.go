package repository

import "blog-api/internal/domain"

type PostRepository interface {
	Create(post domain.Post) (domain.Post, error)
	GetPosts() ([]domain.Post, error)
	GetById(id int) (domain.Post, error)
	Update(post domain.Post) error
	Delete(id int) error
}

package usecase

import (
	"blog-api/internal/domain"
	"blog-api/internal/repository"
)

type PostUsecase struct {
	repo repository.PostRepository
}

func NewPostUsecase(repo repository.PostRepository) *PostUsecase {
	return &PostUsecase{repo: repo}
}

func (u *PostUsecase) Create(post domain.Post) (domain.Post, error) {
	return u.repo.Create(post)
}

func (u *PostUsecase) GetPosts() ([]domain.Post, error) {
	return u.repo.GetPosts()
}

func (u *PostUsecase) GetByID(id int) (domain.Post, error) {
	return u.repo.GetById(id)
}

func (u *PostUsecase) Update(post domain.Post) error {
	return u.repo.Update(post)
}

func (u *PostUsecase) Delete(id int) error {
	return u.repo.Delete(id)
}

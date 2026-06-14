package repository

import (
	"blog-api/internal/domain"
	"database/sql"
)

type MySQLPostRepository struct {
	db *sql.DB
}

func NewMySQLPostRepository(db *sql.DB) *MySQLPostRepository {
	return &MySQLPostRepository{db: db}
}

func (r *MySQLPostRepository) Create(post domain.Post) (domain.Post, error) {

	result, err := r.db.Exec(
		"INSERT INTO posts (title, content, user_id) VALUES (?, ?, ?)",
		post.Title, post.Content, post.UserID,
	)
	if err != nil {
		return domain.Post{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return domain.Post{}, err
	}

	post.ID = int(id)
	return post, nil
}
func (r *MySQLPostRepository) GetPosts() ([]domain.Post, error) {

	rows, err := r.db.Query("SELECT id, title, content, user_id FROM posts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []domain.Post

	for rows.Next() {
		var p domain.Post
		err := rows.Scan(&p.ID, &p.Title, &p.Content, &p.UserID)
		if err != nil {
			return nil, err
		}
		posts = append(posts, p)
	}

	return posts, nil
}
func (r *MySQLPostRepository) GetById(id int) (domain.Post, error) {

	var post domain.Post

	err := r.db.QueryRow(
		"SELECT id, title, content, user_id FROM posts WHERE id = ?",
		id,
	).Scan(&post.ID, &post.Title, &post.Content, &post.UserID)

	if err != nil {
		return domain.Post{}, err
	}

	return post, nil
}
func (r *MySQLPostRepository) Update(post domain.Post) error {

	_, err := r.db.Exec(
		"UPDATE posts SET title=?, content=? WHERE id=?",
		post.Title, post.Content, post.ID,
	)

	return err
}
func (r *MySQLPostRepository) Delete(id int) error {

	_, err := r.db.Exec(
		"DELETE FROM posts WHERE id=?",
		id,
	)

	return err
}

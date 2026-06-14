package repository

import (
	"blog-api/internal/domain"
	"database/sql"
	"errors"
)

type MySQLUserRepository struct {
	db *sql.DB
}

func NewMySQLUserRepository(db *sql.DB) *MySQLUserRepository {
	return &MySQLUserRepository{db: db}
}

func (r *MySQLUserRepository) Create(user domain.User) (domain.User, error) {
	result, err := r.db.Exec("INSERT INTO users (username,password) VALUES (?,?)",
		user.Username, user.Password)
	if err != nil {
		return domain.User{}, err
	}
	id, err := result.LastInsertId()
	user.ID = int(id)

	return user, nil

}

func (r *MySQLUserRepository) GetByUsername(username string) (domain.User, error) {
	var user domain.User
	err := r.db.QueryRow("SELECT id, username, password FROM users WHERE username = ?",
		username).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.User{}, errors.New("user not found")
		}
		return domain.User{}, err
	}
	return user, nil
}

func (r *MySQLUserRepository) GetByID(id int) (domain.User, error) {
	var user domain.User
	err := r.db.QueryRow("SELECT id, username, password FROM users WHERE id = ?",
		id).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {

		return domain.User{}, err
	}
	return user, nil
}

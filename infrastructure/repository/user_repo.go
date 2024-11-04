package repository

import (
	"database/sql"
	"errors"

	"github.com/christianluer/golang-backend-hex/domain"
)

type MysqlUserRepo struct {
	DB *sql.DB
}

func NewMySQLUserRepo(db *sql.DB) *MysqlUserRepo {
	return &MysqlUserRepo{DB: db}
}

func (repo MysqlUserRepo) GetByUsername(username string) (*domain.User, error) {
	var user domain.User
	query := "SELECT id, username, password FROM users WHERE username = ?"
	err := repo.DB.QueryRow(query, username).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

func (repo *MysqlUserRepo) Save(user *domain.User) error {
	query := "INSERT INTO users (username, password) VALUES (?, ?)"
	_, err := repo.DB.Exec(query, user.Username, user.Password)
	if err != nil {
		return err
	}
	return nil
}

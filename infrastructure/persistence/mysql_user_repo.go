package persistence

import (
	"database/sql"
	"errors"

	"github.com/christianluer/golang-backend-hex/domain/model"
	"github.com/christianluer/golang-backend-hex/domain/repository"
)

type MysqlUserRepo struct {
	DB *sql.DB
}

func NewMySQLUserRepo(db *sql.DB) repository.UserRepository {
	return &MysqlUserRepo{DB: db}
}

func (repo MysqlUserRepo) GetByUsername(username string) (*model.User, error) {
	var user model.User
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

func (repo *MysqlUserRepo) GetById(id int) (*model.User, error) {
	var user model.User
	query := "SELECT id, username, password FROM users WHERE id = ?"
	err := repo.DB.QueryRow(query, id).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

func (repo *MysqlUserRepo) Create(user *model.User) error {
	query := "INSERT INTO users (username, password) VALUES (?, ?)"
	_, err := repo.DB.Exec(query, user.Username, user.Password)
	if err != nil {
		return err
	}
	return nil
}

func (repo *MysqlUserRepo) Update(user *model.User) error {
	query := "UPDATE users SET username = ?, password = ? WHERE id = ?"
	_, err := repo.DB.Exec(query, user.Username, user.Password, user.ID)
	if err != nil {
		return err
	}
	return nil
}

func (repo *MysqlUserRepo) Delete(user *model.User) error {
	query := "DELETE from users WHERE id=?"
	_, err := repo.DB.Exec(query, user.ID)
	if err != nil {
		return err
	}
	return nil
}

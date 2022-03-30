package repository

import (
	"errors"
	"fmt"

	"github.com/Argento-prg/NoteBook/server/models"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user models.User) (int, error) {
	//проверяем уникальность логина пользователя
	if !r.CheckUniquenessUser(user.Login) {
		err := "user with the same login already exist"
		return 0, errors.New(err)
	}
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, login, password) VALUES ($1, $2, $3) RETURNING id", usersTable)

	row := r.db.QueryRow(query, user.Name, user.Login, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthPostgres) GetUser(login, password string) (models.User, error) {
	var user models.User
	query := fmt.Sprintf("SELECT id, name FROM %s WHERE login=$1 AND password=$2", usersTable)
	err := r.db.Get(&user, query, login, password)

	return user, err
}

func (r *AuthPostgres) CheckUniquenessUser(login string) bool {
	var user models.User
	var decision bool = true
	query := fmt.Sprintf(
		"SELECT id FROM %s WHERE login=$1",
		usersTable,
	)
	err := r.db.Get(&user, query, login)
	if err == nil {
		decision = false
	}
	return decision
}

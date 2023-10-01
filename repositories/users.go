package repositories

import (
	"database/sql"
	"devbook-api/models"
	"fmt"
)

type users struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *users {
	return &users{db}
}

func (usersRepository users) Create(user models.User) (uint64, error) {
	statements, err := usersRepository.db.Prepare(
		"insert into users (name, nick, email, password) values(?, ?, ?, ?)",
	)

	if err != nil {
		return 0, err
	}
	defer statements.Close()

	result, err := statements.Exec(user.Name, user.Nick, user.Email, user.Password)
	if err != nil {
		return 0, err
	}

	ID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(ID), nil
}

func (usersRepository users) Find(user string) ([]models.User, error) {
	user = fmt.Sprintf("%%%s%%", user) // %user%

	rows, err := usersRepository.db.Query(
		"select id, name, nick, email, createdAt from users where name LIKE ? or nick LIKE ?",
		user, user,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var foundUsers []models.User
	for rows.Next() {
		var user models.User

		if err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}

		foundUsers = append(foundUsers, user)
	}

	return foundUsers, nil
}

func (usersRepository users) FindByID(ID uint64) (models.User, error) {
	rows, err := usersRepository.db.Query(
		"select id, name, nick, email, createdAt from users where id LIKE ?",
		ID,
	)
	if err != nil {
		return models.User{}, nil
	}
	defer rows.Close()

	var user models.User

	if rows.Next() {
		if err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}

func (usersRepository users) Update(ID uint64, user models.User) error {
	statement, err := usersRepository.db.Prepare("update users set name = ?, email = ?, nick = ? where id = ?")
	if err != nil {
		return err
	}

	defer statement.Close()

	if _, err := statement.Exec(user.Name, user.Email, user.Nick, ID); err != nil {
		return err
	}

	return nil
}

func (usersRepository users) Delete(ID uint64) error {
	statement, err := usersRepository.db.Prepare("delete from users where id = ?")

	if err != nil {
		return err
	}

	defer statement.Close()
	if _, err := statement.Exec(ID); err != nil {
		return err
	}

	return nil
}

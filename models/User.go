package models

import (
	"devbook-api/helpers"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

const (
	SIGN_UP_STEP = iota
	UPDATE_USER_STEP
)

func (user *User) validate(step int) error {
	if user.Name == "" {
		return errors.New("User must have a name")
	}
	if user.Nick == "" {
		return errors.New("User must have a nick")
	}
	if user.Email == "" {
		return errors.New("User must have a email")
	}
	if err := checkmail.ValidateFormat(user.Email); err != nil {
		return errors.New("User must have a valid email")
	}
	if step == SIGN_UP_STEP && user.Password == "" {
		return errors.New("User must have a password")
	}

	return nil
}

func (user *User) format(step int) error {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)

	if step == SIGN_UP_STEP {
		passwordWithHash, err := helpers.Hash(user.Password)
		if err != nil {
			return err
		}

		user.Password = string(passwordWithHash)
	}

	return nil
}

func (user *User) Prepare(step int) error {
	if err := user.validate(step); err != nil {
		return err
	}

	if err := user.format(step); err != nil {
		return err
	}

	return nil
}

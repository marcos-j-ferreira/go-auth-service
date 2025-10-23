package domain

/*
Domain of user, responsibility to create own data of user;
has simple verification in data of users
*/

import (
	"strings"
	"errors"
)

type User struct {
	Nome 	string
	Email 	string
	Senha 	string
}

func NewUser(nome, email, senha string) (*User, error) {
	if !strings.Contains(email, "@") {
		return nil, errors.New("invalid email")
	}

	if strings.TrimSpace(nome) == "" {
		return nil, errors.New("name cannot be empty")
	}

	if len(senha) < 6 {
		return nil, errors.New("password must have at least 6 characters")
	}

	return &User{
		Nome:  strings.TrimSpace(nome),
		Email: email,
		Senha: senha,
	}, nil
}

// Validates email to be valid and updates it.
func (u *User) ChangeEmail(newEmail string) error {
	if !strings.Contains(newEmail, "@") {
		return errors.New("invalid email format")
	}

	u.Email = newEmail
	return nil
}

// Validates password length and updates it.
func (u *User) ChangePassword(newPassword string) error {
	if len(newPassword) < 6 {
		return errors.New("password must have at least 6 characters")
	}
	
	u.Senha = newPassword
	return nil
}


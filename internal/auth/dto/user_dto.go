package dto

/*
DTO layer for authentication service.
Provides request and response structures for user creation, login, and user info.
Includes validation functions for requests.
*/

import (
	"errors"
	"strings"
	//"encoding/json" // Para testes de marshal/unmarshal, se necessário
)

type CreateUserRequest struct {
	Nome  string `json:"nome"`
	Email string `json:"email"`
	Senha string `json:"senha"`
}

func ValidateCreateUserRequest(req *CreateUserRequest) error {
	if req == nil {
		return errors.New("request cannot be nil")
	}

	if strings.TrimSpace(req.Nome) == "" {
		return errors.New("name cannot be empty")
	}

	if !strings.Contains(req.Email, "@") {
		return errors.New("invalid email")
	}

	if len(req.Senha) < 6 {
		return errors.New("password must have at least 6 characters")
	}

	return nil
}

type LoginRequest struct {
	Email string `json:"email"`
	Senha string `json:"senha"`
}

func ValidateLoginRequest(req *LoginRequest) error {
	if req == nil {
		return errors.New("request cannot be nil")
	}

	if !strings.Contains(req.Email, "@") {
		return errors.New("invalid email")
	}

	if strings.TrimSpace(req.Senha) == "" {
		return errors.New("password cannot be empty")
	}

	return nil
}

type UserResponse struct {
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

func NewUserResponse(nome, email string) *UserResponse {
	return &UserResponse{
		Nome:  strings.TrimSpace(nome),
		Email: email,
	}
}

type AuthResponse struct {
	Token string       `json:"token"`
	User  *UserResponse `json:"user,omitempty"`
}

func NewAuthResponse(token string, user *UserResponse) *AuthResponse {
	return &AuthResponse{
		Token: token,
		User:  user,
	}
}
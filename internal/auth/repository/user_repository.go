package repository

import (
	"context"
	"auth/internal/auth/domain"  // Ajuste path
	"github.com/google/uuid"
)

type UserRepository interface {
	Create(ctx context.Context, user *domain.User) error
	GetByEmail(ctx context.Context, email string) (*domain.User, error)
	UpdatePassword(ctx context.Context, email, newPassword string) error
	UpdateEmail(ctx context.Context, oldEmail, newEmail string) error
}
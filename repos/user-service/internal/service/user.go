package service

import (
	"context"
	"fmt"
	"time"

	commonv1 "github.com/imkhoirularifin/proto-contracts/gen/go/common/v1"
	userv1 "github.com/imkhoirularifin/proto-contracts/gen/go/user/v1"
	"github.com/google/uuid"
)

type User struct {
	ID        string
	Email     string
	Name      string
	CreatedAt time.Time
}

type UserService interface {
	GetUser(ctx context.Context, id string) (*User, error)
	ListUsers(ctx context.Context, page, pageSize int32) ([]*User, int64, error)
}

type userService struct {
	users map[string]*User
}

func NewUserService() UserService {
	now := time.Now().UTC()
	users := map[string]*User{
		"1": {
			ID:        "1",
			Email:     "alice@example.com",
			Name:      "Alice",
			CreatedAt: now,
		},
		"2": {
			ID:        "2",
			Email:     "bob@example.com",
			Name:      "Bob",
			CreatedAt: now,
		},
	}
	return &userService{users: users}
}

func (s *userService) GetUser(_ context.Context, id string) (*User, error) {
	user, ok := s.users[id]
	if !ok {
		return nil, fmt.Errorf("user not found")
	}
	return user, nil
}

func (s *userService) ListUsers(_ context.Context, page, pageSize int32) ([]*User, int64, error) {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	all := make([]*User, 0, len(s.users))
	for _, user := range s.users {
		all = append(all, user)
	}

	start := int((page - 1) * pageSize)
	if start >= len(all) {
		return []*User{}, int64(len(all)), nil
	}

	end := start + int(pageSize)
	if end > len(all) {
		end = len(all)
	}

	return all[start:end], int64(len(all)), nil
}

func ToProto(user *User) *userv1.User {
	return &userv1.User{
		Id:        user.ID,
		Email:     user.Email,
		Name:      user.Name,
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
	}
}

func NewSeedUser(email, name string) *User {
	return &User{
		ID:        uuid.NewString(),
		Email:     email,
		Name:      name,
		CreatedAt: time.Now().UTC(),
	}
}

func ToPagination(page, pageSize int32, total int64) *commonv1.PaginationResponse {
	return &commonv1.PaginationResponse{
		Page:     page,
		PageSize: pageSize,
		Total:    total,
	}
}

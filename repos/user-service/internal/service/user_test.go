package service

import (
	"context"
	"testing"
)

func TestUserService_GetUser(t *testing.T) {
	svc := NewUserService()

	user, err := svc.GetUser(context.Background(), "1")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if user.Email != "alice@example.com" {
		t.Fatalf("unexpected email: %s", user.Email)
	}
}

func TestUserService_GetUser_NotFound(t *testing.T) {
	svc := NewUserService()

	_, err := svc.GetUser(context.Background(), "missing")
	if err == nil {
		t.Fatal("expected error for missing user")
	}
}

func TestUserService_ListUsers(t *testing.T) {
	svc := NewUserService()

	users, total, err := svc.ListUsers(context.Background(), 1, 10)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if total != 2 {
		t.Fatalf("expected total 2, got %d", total)
	}
	if len(users) != 2 {
		t.Fatalf("expected 2 users, got %d", len(users))
	}
}

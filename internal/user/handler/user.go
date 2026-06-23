package handler

import (
	"context"

	commonv1 "github.com/imkhoirularifin/go-grpc-microservice-template/gen/go/common/v1"
	userv1 "github.com/imkhoirularifin/go-grpc-microservice-template/gen/go/user/v1"
	"github.com/imkhoirularifin/go-grpc-microservice-template/internal/user/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserHandler struct {
	userv1.UnimplementedUserServiceServer
	service service.UserService
}

func NewUserHandler(svc service.UserService) *UserHandler {
	return &UserHandler{service: svc}
}

func (h *UserHandler) GetUser(ctx context.Context, req *userv1.GetUserRequest) (*userv1.GetUserResponse, error) {
	user, err := h.service.GetUser(ctx, req.GetId())
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return &userv1.GetUserResponse{User: service.ToProto(user)}, nil
}

func (h *UserHandler) ListUsers(ctx context.Context, req *userv1.ListUsersRequest) (*userv1.ListUsersResponse, error) {
	page := int32(1)
	pageSize := int32(10)
	if req.GetPagination() != nil {
		page = req.GetPagination().GetPage()
		pageSize = req.GetPagination().GetPageSize()
	}

	users, total, err := h.service.ListUsers(ctx, page, pageSize)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	protoUsers := make([]*userv1.User, 0, len(users))
	for _, user := range users {
		protoUsers = append(protoUsers, service.ToProto(user))
	}

	return &userv1.ListUsersResponse{
		Users: protoUsers,
		Pagination: &commonv1.PaginationResponse{
			Page:     page,
			PageSize: pageSize,
			Total:    total,
		},
	}, nil
}

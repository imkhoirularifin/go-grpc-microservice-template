package handler

import (
	commonv1 "github.com/imkhoirularifin/go-grpc-microservice-template/gen/go/common/v1"
	userv1 "github.com/imkhoirularifin/go-grpc-microservice-template/gen/go/user/v1"
	"github.com/imkhoirularifin/go-grpc-microservice-template/lib/dto"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	client userv1.UserServiceClient
}

func NewUserHandler(client userv1.UserServiceClient) *UserHandler {
	return &UserHandler{client: client}
}

func (h *UserHandler) Register(r fiber.Router) {
	r.Get("/:id", h.GetUser)
	r.Get("/", h.ListUsers)
}

func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	resp, err := h.client.GetUser(c.UserContext(), &userv1.GetUserRequest{
		Id: c.Params("id"),
	})
	if err != nil {
		return fiber.NewError(fiber.StatusBadGateway, "failed to fetch user")
	}

	return c.JSON(dto.Response{
		Message: "user found",
		Data:    resp.User,
	})
}

func (h *UserHandler) ListUsers(c *fiber.Ctx) error {
	resp, err := h.client.ListUsers(c.UserContext(), &userv1.ListUsersRequest{
		Pagination: &commonv1.PaginationRequest{
			Page:     1,
			PageSize: 10,
		},
	})
	if err != nil {
		return fiber.NewError(fiber.StatusBadGateway, "failed to list users")
	}

	return c.JSON(dto.Response{
		Message: "users listed",
		Data:    resp,
	})
}

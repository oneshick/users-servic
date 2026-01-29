package transportgrpc

import (
	"context"

	userpb "github.com/oneshick/project-protos/proto/user"
	"github.com/oneshick/users-service/internal/user"
)

type Handler struct {
	svc user.Service
	userpb.UnimplementedUserServiceServer
}

func NewHandler(svc user.Service) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) CreateUser(ctx context.Context, req *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {
	defaultPassword := "default123"

	createdUser, err := h.svc.CreateUser(req.GetEmail(), defaultPassword)
	if err != nil {
		return nil, err
	}

	return &userpb.CreateUserResponse{
		User: toProtoUser(createdUser),
	}, nil
}

func (h *Handler) GetUser(ctx context.Context, req *userpb.GetUserRequest) (*userpb.User, error) {
	user, err := h.svc.GetUserByID(req.GetId())
	if err != nil {
		return nil, err
	}

	return toProtoUser(user), nil
}

func (h *Handler) ListUsers(ctx context.Context, req *userpb.ListUsersRequest) (*userpb.ListUsersResponse, error) {
	users, err := h.svc.GetAllUsers()
	if err != nil {
		return nil, err
	}

	protoUsers := make([]*userpb.User, len(users))
	for i, u := range users {
		protoUsers[i] = toProtoUser(u)
	}

	return &userpb.ListUsersResponse{
		Users: protoUsers,
		Total: uint32(len(protoUsers)),
	}, nil
}

func (h *Handler) UpdateUser(ctx context.Context, req *userpb.UpdateUserRequest) (*userpb.UpdateUserResponse, error) {
	updatedUser, err := h.svc.UpdateUser(req.GetId(), req.GetEmail(), "")
	if err != nil {
		return nil, err
	}

	return &userpb.UpdateUserResponse{
		User: toProtoUser(updatedUser),
	}, nil
}

func (h *Handler) DeleteUser(ctx context.Context, req *userpb.DeleteUserRequest) (*userpb.DeleteUserResponse, error) {
	err := h.svc.DeleteUser(req.GetId())
	if err != nil {
		return &userpb.DeleteUserResponse{Success: false}, err
	}

	return &userpb.DeleteUserResponse{Success: true}, nil
}

func toProtoUser(u *user.User) *userpb.User {

	return &userpb.User{
		Id:    u.ID,
		Email: u.Email,
	}
}

// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/user/user.proto

package userpb

import (
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for User service

func NewUserEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for User service

type UserService interface {
	Register(ctx context.Context, in *UserRegisterRequest, opts ...client.CallOption) (*UserRegisterResponse, error)
	Login(ctx context.Context, in *UserLoginRequest, opts ...client.CallOption) (*UserLoginResponse, error)
	GetUserInfo(ctx context.Context, in *UserInfoRequest, opts ...client.CallOption) (*UserInfoResponse, error)
}

type userService struct {
	c    client.Client
	name string
}

func NewUserService(name string, c client.Client) UserService {
	return &userService{
		c:    c,
		name: name,
	}
}

func (c *userService) Register(ctx context.Context, in *UserRegisterRequest, opts ...client.CallOption) (*UserRegisterResponse, error) {
	req := c.c.NewRequest(c.name, "User.Register", in)
	out := new(UserRegisterResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) Login(ctx context.Context, in *UserLoginRequest, opts ...client.CallOption) (*UserLoginResponse, error) {
	req := c.c.NewRequest(c.name, "User.Login", in)
	out := new(UserLoginResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) GetUserInfo(ctx context.Context, in *UserInfoRequest, opts ...client.CallOption) (*UserInfoResponse, error) {
	req := c.c.NewRequest(c.name, "User.GetUserInfo", in)
	out := new(UserInfoResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for User service

type UserHandler interface {
	Register(context.Context, *UserRegisterRequest, *UserRegisterResponse) error
	Login(context.Context, *UserLoginRequest, *UserLoginResponse) error
	GetUserInfo(context.Context, *UserInfoRequest, *UserInfoResponse) error
}

func RegisterUserHandler(s server.Server, hdlr UserHandler, opts ...server.HandlerOption) error {
	type user interface {
		Register(ctx context.Context, in *UserRegisterRequest, out *UserRegisterResponse) error
		Login(ctx context.Context, in *UserLoginRequest, out *UserLoginResponse) error
		GetUserInfo(ctx context.Context, in *UserInfoRequest, out *UserInfoResponse) error
	}
	type User struct {
		user
	}
	h := &userHandler{hdlr}
	return s.Handle(s.NewHandler(&User{h}, opts...))
}

type userHandler struct {
	UserHandler
}

func (h *userHandler) Register(ctx context.Context, in *UserRegisterRequest, out *UserRegisterResponse) error {
	return h.UserHandler.Register(ctx, in, out)
}

func (h *userHandler) Login(ctx context.Context, in *UserLoginRequest, out *UserLoginResponse) error {
	return h.UserHandler.Login(ctx, in, out)
}

func (h *userHandler) GetUserInfo(ctx context.Context, in *UserInfoRequest, out *UserInfoResponse) error {
	return h.UserHandler.GetUserInfo(ctx, in, out)
}

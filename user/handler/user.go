package handler

import (
	"context"
	"github.com/tongs-dev/shopping-platform/user/domain/model"
	"github.com/tongs-dev/shopping-platform/user/domain/service"
	userpb "github.com/tongs-dev/shopping-platform/user/proto/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// UserHandler is the gRPC service handler for user-related operations.
type UserHandler struct {
	UserDataService service.IUserDataService
}

// Register creates a new user account based on the provided request data.
func (u *UserHandler) Register(ctx context.Context, userRegisterRequest *userpb.UserRegisterRequest, userRegisterResponse *userpb.UserRegisterResponse) error {
	if userRegisterRequest.UserName == "" || userRegisterRequest.Pwd == "" {
		return status.Errorf(codes.InvalidArgument, "username and password are required")
	}

	userRegister := &model.User{
		UserName:     userRegisterRequest.UserName,
		FirstName:    userRegisterRequest.FirstName,
		HashPassword: userRegisterRequest.Pwd,
	}

	_, err := u.UserDataService.AddUser(userRegister)
	if err != nil {
		return status.Errorf(codes.Internal, "failed to register user: %v", err)
	}

	userRegisterResponse.Message = "User registered successfully"
	return nil
}

// Login verifies the user's credentials and returns a success response if valid.
func (u *UserHandler) Login(ctx context.Context, userLogin *userpb.UserLoginRequest, loginResponse *userpb.UserLoginResponse) error {
	if userLogin.UserName == "" || userLogin.Pwd == "" {
		return status.Errorf(codes.InvalidArgument, "username and password are required")
	}

	isOk, err := u.UserDataService.CheckPwd(userLogin.UserName, userLogin.Pwd)
	if err != nil {
		return status.Errorf(codes.Internal, "login failed: %v", err)
	}

	loginResponse.IsSuccess = isOk
	return nil
}

// GetUserInfo retrieves user details based on the provided username.
func (u *UserHandler) GetUserInfo(ctx context.Context, userInfoRequest *userpb.UserInfoRequest, userInfoResponse *userpb.UserInfoResponse) error {
	if userInfoRequest.UserName == "" {
		return status.Errorf(codes.InvalidArgument, "username is required")
	}

	userInfo, err := u.UserDataService.FindUserByName(userInfoRequest.UserName)
	if err != nil {
		return status.Errorf(codes.NotFound, "user not found: %v", err)
	}

	*userInfoResponse = *UserForResponse(userInfo)
	return nil
}

// UserForResponse converts a model.User struct into a userpb.UserInfoResponse.
func UserForResponse(userModel *model.User) *userpb.UserInfoResponse {
	return &userpb.UserInfoResponse{
		UserName:  userModel.UserName,
		FirstName: userModel.FirstName,
		UserId:    userModel.ID,
	}
}

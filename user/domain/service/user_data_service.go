package service

import (
	"errors"
	"github.com/tongs-dev/shopping-platform/user/domain/model"
	"github.com/tongs-dev/shopping-platform/user/domain/repository"
	"golang.org/x/crypto/bcrypt"
)

// IUserDataService defines the contract for user-related operations.
type IUserDataService interface {
	AddUser(*model.User) (int64, error)
	DeleteUser(int64) error
	UpdateUser(user *model.User, isChangePwd bool) (err error)
	FindUserByName(string) (*model.User, error)
	CheckPwd(userName string, pwd string) (isOk bool, err error)
}

// NewUserDataService returns an implementation of IUserDataService.
func NewUserDataService(userRepository repository.IUserRepository) IUserDataService {
	return &UserDataService{UserRepository: userRepository}
}

// UserDataService is the concrete implementation of IUserDataService.
type UserDataService struct {
	UserRepository repository.IUserRepository
}

// GeneratePassword hashes the provided plaintext password using bcrypt.
func GeneratePassword(userPassword string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(userPassword), bcrypt.DefaultCost)
}

// ValidatePassword compares a plaintext password with a stored bcrypt hash.
func ValidatePassword(userPassword string, hashed string) (isOk bool, err error) {
	if err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(userPassword)); err != nil {
		return false, errors.New("password mismatch")
	}
	return true, nil
}

// AddUser hashes the password and saves a new user in the database.
func (u *UserDataService) AddUser(user *model.User) (int64, error) {
	// Hashes the user's password before storing
	pwdByte, err := GeneratePassword(user.HashPassword)
	if err != nil {
		return 0, err
	}
	user.HashPassword = string(pwdByte)

	// Saves the user to the repository
	userID, err := u.UserRepository.CreateUser(user)
	if err != nil {
		return 0, err
	}

	return userID, nil
}

// DeleteUser removes a user from the database by ID.
func (u *UserDataService) DeleteUser(userID int64) error {
	return u.UserRepository.DeleteUserByID(userID)
}

// UpdateUser updates user details, with an option to change the password.
func (u *UserDataService) UpdateUser(user *model.User, isChangePwd bool) (err error) {
	if isChangePwd {
		// Hashes the new password before updating
		pwdByte, err := GeneratePassword(user.HashPassword)
		if err != nil {
			return err
		}
		user.HashPassword = string(pwdByte)
	}

	// Updates user information in the repository
	return u.UserRepository.UpdateUser(user)
}

// FindUserByName retrieves a user by username.
func (u *UserDataService) FindUserByName(userName string) (user *model.User, err error) {
	return u.UserRepository.FindUserByName(userName)
}

// CheckPwd validates a user's password by retrieving their hashed password from storage.
func (u *UserDataService) CheckPwd(userName string, pwd string) (isOk bool, err error) {
	// Fetches user details
	user, err := u.UserRepository.FindUserByName(userName)
	if err != nil {
		return false, err
	}

	// Validates password
	return ValidatePassword(pwd, user.HashPassword)
}

package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"shopping-platform/user/domain/model"
)

// MockUserRepository is a mock implementation of IUserRepository.
type MockUserRepository struct {
	mock.Mock
}

// Mock method implementations
func (m *MockUserRepository) InitTable() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockUserRepository) FindUserByID(userID int64) (*model.User, error) {
	args := m.Called(userID)
	return args.Get(0).(*model.User), args.Error(1)
}

func (m *MockUserRepository) FindAll() ([]model.User, error) {
	args := m.Called()
	return args.Get(0).([]model.User), args.Error(1)
}

func (m *MockUserRepository) CreateUser(user *model.User) (int64, error) {
	args := m.Called(user)
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockUserRepository) DeleteUserByID(userID int64) error {
	args := m.Called(userID)
	return args.Error(0)
}

func (m *MockUserRepository) UpdateUser(user *model.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepository) FindUserByName(userName string) (*model.User, error) {
	args := m.Called(userName)
	return args.Get(0).(*model.User), args.Error(1)
}

// Test AddUser: Ensures the user is created successfully.
func TestAddUser(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := NewUserDataService(mockRepo)

	user := &model.User{
		UserName:     "testuser",
		FirstName:    "John",
		HashPassword: "securepassword",
	}

	mockRepo.On("CreateUser", mock.Anything).Return(int64(1), nil)

	userID, err := service.AddUser(user)
	assert.NoError(t, err)
	assert.Equal(t, int64(1), userID)
}

// Test DeleteUser: Ensures a user is deleted correctly.
func TestDeleteUser(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := NewUserDataService(mockRepo)

	mockRepo.On("DeleteUserByID", int64(1)).Return(nil)

	err := service.DeleteUser(1)
	assert.NoError(t, err)
	mockRepo.AssertCalled(t, "DeleteUserByID", int64(1))
}

// Test UpdateUser: Ensures a user update works correctly.
func TestUpdateUser(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := NewUserDataService(mockRepo)

	user := &model.User{
		ID:        1,
		UserName:  "testuser",
		FirstName: "UpdatedName",
	}

	mockRepo.On("UpdateUser", user).Return(nil)

	err := service.UpdateUser(user, false)
	assert.NoError(t, err)
	mockRepo.AssertCalled(t, "UpdateUser", user)
}

// Test FindUserByName: Ensures a user can be found by username.
func TestFindUserByName(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := NewUserDataService(mockRepo)

	expectedUser := &model.User{
		ID:        1,
		UserName:  "testuser",
		FirstName: "John",
	}

	mockRepo.On("FindUserByName", "testuser").Return(expectedUser, nil)

	user, err := service.FindUserByName("testuser")
	assert.NoError(t, err)
	assert.Equal(t, expectedUser, user)
}

// Test CheckPwd: Ensures password validation works.
func TestCheckPwd(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := NewUserDataService(mockRepo)

	// Hash password
	hashedPwd, _ := GeneratePassword("securepassword")

	expectedUser := &model.User{
		ID:           1,
		UserName:     "testuser",
		HashPassword: string(hashedPwd),
	}

	mockRepo.On("FindUserByName", "testuser").Return(expectedUser, nil)

	isValid, err := service.CheckPwd("testuser", "securepassword")
	assert.NoError(t, err)
	assert.True(t, isValid)
}

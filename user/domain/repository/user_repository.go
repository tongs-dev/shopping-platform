package repository

import (
	"github.com/jinzhu/gorm"
	"shopping-platform/user/domain/model"
)

// IUserRepository defines the contract for user-related database operations.
type IUserRepository interface {
	InitTable() error
	FindUserByName(string) (*model.User, error)
	FindUserByID(int64) (*model.User, error)
	CreateUser(*model.User) (int64, error)
	DeleteUserByID(int64) error
	UpdateUser(*model.User) error
	FindAll() ([]model.User, error)
}

// NewUserRepository returns an implementation of IUserRepository using GORM.
func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{mysqlDb: db}
}

// UserRepository is the concrete implementation of IUserRepository.
type UserRepository struct {
	mysqlDb *gorm.DB
}

// InitTable creates the users table if it doesn't exist.
func (u *UserRepository) InitTable() error {
	return u.mysqlDb.CreateTable(&model.User{}).Error
}

// FindUserByName retrieves a user from the database based on username.
func (u *UserRepository) FindUserByName(name string) (*model.User, error) {
	user := &model.User{}
	err := u.mysqlDb.Where("user_name = ?", name).First(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

// FindUserByID retrieves a user from the database using their unique ID.
func (u *UserRepository) FindUserByID(userID int64) (*model.User, error) {
	user := &model.User{}
	err := u.mysqlDb.First(user, userID).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

// CreateUser inserts a new user into the database and returns the generated ID.
func (u *UserRepository) CreateUser(user *model.User) (int64, error) {
	err := u.mysqlDb.Create(user).Error
	if err != nil {
		return 0, err // Returns 0 if user creation fails
	}
	return user.ID, nil
}

// DeleteUserByID removes a user from the database using their unique ID.
func (u *UserRepository) DeleteUserByID(userID int64) error {
	result := u.mysqlDb.Where("id = ?", userID).Delete(&model.User{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

// UpdateUser modifies an existing user record in the database.
func (u *UserRepository) UpdateUser(user *model.User) error {
	return u.mysqlDb.Model(user).Updates(user).Error
}

// FindAll retrieves all users from the database.
func (u *UserRepository) FindAll() ([]model.User, error) {
	var users []model.User
	err := u.mysqlDb.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

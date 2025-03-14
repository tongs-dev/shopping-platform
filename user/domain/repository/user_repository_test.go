package repository

import (
	"fmt"
	"log"
	"testing"

	"github.com/tongs-dev/shopping-platform/user/domain/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // Import MySQL dialect
	"github.com/stretchr/testify/assert"
)

// setupTestDB initializes a real MySQL test database for unit tests.
func setupTestDB(t *testing.T) *gorm.DB {
	dsn := "root:123456@tcp(localhost:3306)/userdb?charset=utf8mb4&parseTime=True&loc=Local"

	// Opens MySQL connection
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to MySQL: %v", err)
	}

	// Drops and recreates the user table for a clean test state
	err = db.DropTableIfExists(&model.User{}).Error
	assert.NoError(t, err, "Failed to drop test table")

	err = db.AutoMigrate(&model.User{}).Error
	assert.NoError(t, err, "Failed to migrate test table")

	fmt.Println("MySQL test database setup complete")
	return db
}

// TestCreateUser verifies that a user can be inserted into MySQL.
func TestCreateUser(t *testing.T) {
	db := setupTestDB(t)
	repo := NewUserRepository(db)

	// Create test user
	user := &model.User{
		UserName:  "testuser",
		FirstName: "John",
	}

	userID, err := repo.CreateUser(user)
	assert.NoError(t, err, "Failed to create user")
	assert.NotZero(t, userID, "User ID should not be zero")
}

// TestFindUserByName verifies retrieving a user by username.
func TestFindUserByName(t *testing.T) {
	db := setupTestDB(t)
	repo := NewUserRepository(db)

	// Insert test user
	user := &model.User{UserName: "testuser", FirstName: "John"}
	repo.CreateUser(user)

	// Retrieve user
	foundUser, err := repo.FindUserByName("testuser")
	assert.NoError(t, err, "Failed to find user")
	assert.Equal(t, "testuser", foundUser.UserName)
}

// TestFindUserByID verifies retrieving a user by ID.
func TestFindUserByID(t *testing.T) {
	db := setupTestDB(t)
	repo := NewUserRepository(db)

	// Insert test user
	user := &model.User{UserName: "testuser", FirstName: "John"}
	userID, _ := repo.CreateUser(user)

	// Retrieve user by ID
	foundUser, err := repo.FindUserByID(userID)
	assert.NoError(t, err, "Failed to find user")
	assert.Equal(t, userID, foundUser.ID)
}

// TestUpdateUser verifies updating an existing user.
func TestUpdateUser(t *testing.T) {
	db := setupTestDB(t)
	repo := NewUserRepository(db)

	// Insert user
	user := &model.User{UserName: "testuser", FirstName: "OldName"}
	userID, _ := repo.CreateUser(user)

	// Update user details
	user.ID = userID
	user.FirstName = "NewName"
	err := repo.UpdateUser(user)
	assert.NoError(t, err, "Failed to update user")

	// Fetch updated user
	updatedUser, _ := repo.FindUserByID(userID)
	assert.Equal(t, "NewName", updatedUser.FirstName)
}

// TestDeleteUserByID verifies user deletion.
func TestDeleteUserByID(t *testing.T) {
	db := setupTestDB(t)
	repo := NewUserRepository(db)

	// Insert user
	user := &model.User{UserName: "testuser", FirstName: "John"}
	userID, _ := repo.CreateUser(user)

	// Delete user
	err := repo.DeleteUserByID(userID)
	assert.NoError(t, err, "Failed to delete user")

	// Try to retrieve deleted user
	_, err = repo.FindUserByID(userID)
	assert.Error(t, err, "Expected error for deleted user")
}

// TestFindAll verifies retrieving all users.
func TestFindAll(t *testing.T) {
	db := setupTestDB(t)
	repo := NewUserRepository(db)

	// Insert multiple users
	repo.CreateUser(&model.User{UserName: "user1", FirstName: "Alice"})
	repo.CreateUser(&model.User{UserName: "user2", FirstName: "Bob"})

	// Retrieve all users
	users, err := repo.FindAll()
	assert.NoError(t, err, "Failed to retrieve all users")
	assert.Len(t, users, 2, "Expected 2 users")
}

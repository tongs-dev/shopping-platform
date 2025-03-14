package repository

import (
	"fmt"
	"log"
	"math/rand"
	"testing"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // Import MySQL dialect
	"github.com/stretchr/testify/assert"
	"github.com/tongs-dev/shopping-platform/user/domain/model"
)

// TestUserRepository contains all the unit tests for the UserRepository.
func TestUserRepository(t *testing.T) {
	// Initializes database and repository
	db := setupTestDB(t)
	repo := &UserRepository{mysqlDb: db}

	t.Run("CreateUser", func(t *testing.T) {
		user := &model.User{
			UserName:  generateRandomString(5),
			FirstName: generateRandomString(5),
		}

		userID, err := repo.CreateUser(user)
		assert.NoError(t, err)
		assert.NotZero(t, userID, "User ID should not be zero")
	})

	t.Run("FindUserByName", func(t *testing.T) {
		userName := generateRandomString(5)
		user := &model.User{UserName: userName, FirstName: "John"}
		_, err := repo.CreateUser(user)
		assert.NoError(t, err, "Failed to create user")

		foundUser, err := repo.FindUserByName(userName)
		assert.NoError(t, err)
		assert.Equal(t, userName, foundUser.UserName)
	})

	t.Run("FindUserByID", func(t *testing.T) {
		user := &model.User{UserName: generateRandomString(5), FirstName: "John"}
		userID, _ := repo.CreateUser(user)

		foundUser, err := repo.FindUserByID(userID)
		assert.NoError(t, err)
		assert.Equal(t, userID, foundUser.ID)
	})

	t.Run("UpdateUser", func(t *testing.T) {
		user := &model.User{UserName: generateRandomString(5), FirstName: "OldName"}
		userID, _ := repo.CreateUser(user)

		user.ID = userID
		user.FirstName = "NewName"
		err := repo.UpdateUser(user)
		assert.NoError(t, err, "Failed to update user")

		updatedUser, _ := repo.FindUserByID(userID)
		assert.Equal(t, "NewName", updatedUser.FirstName)
	})

	t.Run("DeleteUserByID", func(t *testing.T) {
		user := &model.User{UserName: generateRandomString(5), FirstName: "John"}
		userID, _ := repo.CreateUser(user)

		err := repo.DeleteUserByID(userID)
		assert.NoError(t, err, "Failed to delete user")

		_, err = repo.FindUserByID(userID)
		assert.Error(t, err, "Expected error for deleted user")
	})

	t.Run("FindAll", func(t *testing.T) {
		clearTable(t, db)

		repo.CreateUser(&model.User{UserName: generateRandomString(5), FirstName: "Alice"})
		repo.CreateUser(&model.User{UserName: generateRandomString(5), FirstName: "Bob"})

		users, err := repo.FindAll()
		assert.NoError(t, err, "Failed to retrieve all users")
		assert.Len(t, users, 2, "Expected 2 users")
	})

	t.Run("CreateUserWithSameUsername", func(t *testing.T) {
		user := &model.User{UserName: generateRandomString(5), FirstName: "John"}
		_, err := repo.CreateUser(user)
		assert.NoError(t, err, "Failed to create user")

		_, err = repo.CreateUser(user)
		assert.Error(t, err, "Expected error for duplicate user creation")
	})

	t.Run("FindUserByNameNotFound", func(t *testing.T) {
		_, err := repo.FindUserByName("nonexistentuser")
		assert.Error(t, err, "Expected error for non-existent user")
	})
}

// setupTestDB initializes a real MySQL test database for unit tests.
func setupTestDB(t *testing.T) *gorm.DB {
	dsn := "root:123456@tcp(localhost:3306)/userdb?charset=utf8mb4&parseTime=True&loc=Local"

	// Opens MySQL connection
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to MySQL: %v", err)
	}

	// Clear the 'users' table before each test
	err = db.Exec("DROP TABLE IF EXISTS users").Error
	if err != nil {
		log.Fatalf("Failed to drop 'users' table: %v", err)
	}

	// Automatically migrate the User model (creating the table)
	err = db.AutoMigrate(&model.User{}).Error
	assert.NoError(t, err, "Failed to migrate test table")

	fmt.Println("MySQL test database setup complete")
	return db
}

// clearTable clears the users table before each test
func clearTable(t *testing.T, db *gorm.DB) {
	err := db.Exec("TRUNCATE TABLE users").Error
	assert.NoError(t, err, "Failed to clear 'users' table")
}

func generateRandomString(length int) string {
	// Create a new random source with the current Unix timestamp
	randSource := rand.NewSource(time.Now().UnixNano())
	r := rand.New(randSource)

	// Define the characters to use in the random string
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[r.Intn(len(charset))]
	}
	return string(result)
}

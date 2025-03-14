package repository

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // Import MySQL dialect
	"github.com/stretchr/testify/assert"
	"github.com/tongs-dev/shopping-platform/category/domain/model"
	"log"
	"testing"
)

// TestCategoryRepository tests the CategoryRepository methods using MySQL database.
func TestCategoryRepository(t *testing.T) {
	// Initializes database and repository
	db := setupTestDB(t)
	repo := &CategoryRepository{mysqlDb: db}

	t.Run("FindCategoryByID", func(t *testing.T) {
		category := &model.Category{
			CategoryName: "Test Category",
		}

		_, err := repo.CreateCategory(category)
		assert.NoError(t, err)

		category, err = repo.FindCategoryByID(category.ID)
		assert.NoError(t, err)
		assert.Equal(t, "Test Category", category.CategoryName)
	})

	t.Run("CreateCategory", func(t *testing.T) {
		category := &model.Category{
			CategoryName: "New Category",
		}
		id, err := repo.CreateCategory(category)
		assert.NoError(t, err)
		assert.NotZero(t, id)
	})

	t.Run("DeleteCategoryByID", func(t *testing.T) {
		category := &model.Category{
			CategoryName: "Category To Be Deleted",
		}

		id, err := repo.CreateCategory(category)
		assert.NoError(t, err)

		err = repo.DeleteCategoryByID(id)
		assert.NoError(t, err)

		_, err = repo.FindCategoryByID(id)
		assert.Error(t, err)
	})

	t.Run("UpdateCategory", func(t *testing.T) {
		category := &model.Category{
			CategoryName: "Old Category Name",
		}
		id, err := repo.CreateCategory(category)
		assert.NoError(t, err)

		category.CategoryName = "Updated Category Name"
		err = repo.UpdateCategory(category)
		assert.NoError(t, err)

		updatedCategory, err := repo.FindCategoryByID(id)
		assert.NoError(t, err)
		assert.Equal(t, "Updated Category Name", updatedCategory.CategoryName)
	})

	t.Run("FindAll", func(t *testing.T) {
		clearTable(t, db)

		category1 := &model.Category{
			CategoryName: "Category 1",
		}
		category2 := &model.Category{
			CategoryName: "Category 2",
		}

		// Create two categories
		_, err := repo.CreateCategory(category1)
		assert.NoError(t, err)
		_, err = repo.CreateCategory(category2)
		assert.NoError(t, err)

		// Find all categories
		categories, err := repo.FindAll()
		assert.NoError(t, err)
		assert.Len(t, categories, 2)
	})

	t.Run("FindCategoryByName", func(t *testing.T) {
		category := &model.Category{
			CategoryName: "Unique Category",
		}

		// Create the category
		_, err := repo.CreateCategory(category)
		assert.NoError(t, err)

		// Find the category by name
		foundCategory, err := repo.FindCategoryByName("Unique Category")
		assert.NoError(t, err)
		assert.Equal(t, "Unique Category", foundCategory.CategoryName)
	})

	t.Run("FindCategoryByLevel", func(t *testing.T) {
		category := &model.Category{
			CategoryName:  "Level 1 Category",
			CategoryLevel: 1,
		}

		_, err := repo.CreateCategory(category)
		assert.NoError(t, err)

		categories, err := repo.FindCategoryByLevel(1)
		assert.NoError(t, err)
		assert.Len(t, categories, 1)
	})

	t.Run("FindCategoryByParent", func(t *testing.T) {
		parentCategory := &model.Category{
			CategoryName: "Parent Category",
		}
		parentID, err := repo.CreateCategory(parentCategory)
		assert.NoError(t, err)

		childCategory := &model.Category{
			CategoryName:   "Child Category",
			CategoryParent: parentID,
		}
		_, err = repo.CreateCategory(childCategory)
		assert.NoError(t, err)

		categories, err := repo.FindCategoryByParent(parentID)
		assert.NoError(t, err)
		assert.Len(t, categories, 1)
	})
}

// setupTestDB initializes a real MySQL test database for unit tests.
func setupTestDB(t *testing.T) *gorm.DB {
	dsn := "root:123456@tcp(localhost:3306)/categorydb?charset=utf8mb4&parseTime=True&loc=Local"

	// Opens MySQL connection
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to MySQL: %v", err)
	}

	// Clear the 'users' table before each test
	err = db.Exec("DROP TABLE IF EXISTS categories").Error
	if err != nil {
		log.Fatalf("Failed to drop 'categories' table: %v", err)
	}

	// Automatically migrate the User model (creating the table)
	err = db.AutoMigrate(&model.Category{}).Error
	assert.NoError(t, err, "Failed to migrate test table")

	fmt.Println("MySQL test database setup complete")
	return db
}

// clearTable clears the users table before each test
func clearTable(t *testing.T, db *gorm.DB) {
	err := db.Exec("TRUNCATE TABLE categories").Error
	assert.NoError(t, err, "Failed to clear 'categories' table")
}

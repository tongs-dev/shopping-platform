package service

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/tongs-dev/shopping-platform/category/domain/model"
)

// MockCategoryRepository is a mock type for the ICategoryRepository interface
type MockCategoryRepository struct {
	mock.Mock
}

// Mock the InitTable method
func (m *MockCategoryRepository) InitTable() error {
	args := m.Called()
	return args.Error(0) // Return the error as set up in the mock
}

func (m *MockCategoryRepository) CreateCategory(category *model.Category) (int64, error) {
	args := m.Called(category)
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockCategoryRepository) DeleteCategoryByID(categoryID int64) error {
	args := m.Called(categoryID)
	return args.Error(0)
}

func (m *MockCategoryRepository) UpdateCategory(category *model.Category) error {
	args := m.Called(category)
	return args.Error(0)
}

func (m *MockCategoryRepository) FindCategoryByID(categoryID int64) (*model.Category, error) {
	args := m.Called(categoryID)
	return args.Get(0).(*model.Category), args.Error(1)
}

func (m *MockCategoryRepository) FindAll() ([]model.Category, error) {
	args := m.Called()
	return args.Get(0).([]model.Category), args.Error(1)
}

func (m *MockCategoryRepository) FindCategoryByName(categoryName string) (*model.Category, error) {
	args := m.Called(categoryName)
	return args.Get(0).(*model.Category), args.Error(1)
}

func (m *MockCategoryRepository) FindCategoryByLevel(level uint32) ([]model.Category, error) {
	args := m.Called(level)
	return args.Get(0).([]model.Category), args.Error(1)
}

func (m *MockCategoryRepository) FindCategoryByParent(parent int64) ([]model.Category, error) {
	args := m.Called(parent)
	return args.Get(0).([]model.Category), args.Error(1)
}

// Helper function to create a mock repository and service
func newCategoryService() (*MockCategoryRepository, ICategoryService) {
	mockRepo := new(MockCategoryRepository)
	service := NewCategoryService(mockRepo)
	return mockRepo, service
}

// TestCategoryService tests the CategoryService
func TestCategoryService(t *testing.T) {

	t.Run("TestCategoryService", func(t *testing.T) {
		// Each subtest gets its own fresh mock and service
		mockRepo, service := newCategoryService()

		// Ensure cleanup after each subtest to reset expectations
		t.Cleanup(func() {
			mockRepo.AssertExpectations(t)
		})

		t.Run("AddCategory", func(t *testing.T) {
			category := &model.Category{CategoryName: "Test Category"}
			mockRepo.On("CreateCategory", category).Return(int64(1), nil)

			userID, err := service.AddCategory(category)
			assert.NoError(t, err)
			assert.Equal(t, int64(1), userID)
		})

		t.Run("DeleteCategory", func(t *testing.T) {
			categoryID := int64(1)
			mockRepo.On("DeleteCategoryByID", categoryID).Return(nil)

			err := service.DeleteCategory(categoryID)
			assert.NoError(t, err)
		})

		t.Run("UpdateCategory", func(t *testing.T) {
			category := &model.Category{ID: 1, CategoryName: "Updated Category"}
			mockRepo.On("UpdateCategory", category).Return(nil)

			err := service.UpdateCategory(category)
			assert.NoError(t, err)
		})

		t.Run("FindCategoryByID", func(t *testing.T) {
			categoryID := int64(1)
			expectedCategory := &model.Category{ID: categoryID, CategoryName: "Test Category"}
			mockRepo.On("FindCategoryByID", categoryID).Return(expectedCategory, nil)

			category, err := service.FindCategoryByID(categoryID)
			assert.NoError(t, err)
			assert.Equal(t, expectedCategory, category)
		})

		t.Run("FindAllCategory", func(t *testing.T) {
			expectedCategories := []model.Category{
				{ID: 1, CategoryName: "Category 1"},
				{ID: 2, CategoryName: "Category 2"},
			}
			mockRepo.On("FindAll").Return(expectedCategories, nil)

			categories, err := service.FindAllCategory()
			assert.NoError(t, err)
			assert.Equal(t, expectedCategories, categories)
		})

		t.Run("FindCategoryByName", func(t *testing.T) {
			categoryName := "Test Category"
			expectedCategory := &model.Category{ID: 1, CategoryName: categoryName}
			mockRepo.On("FindCategoryByName", categoryName).Return(expectedCategory, nil)

			category, err := service.FindCategoryByName(categoryName)
			assert.NoError(t, err)
			assert.Equal(t, expectedCategory, category)
		})
	})

	t.Run("Error Handling", func(t *testing.T) {
		// Each subtest gets its own fresh mock and service
		mockRepo, service := newCategoryService()

		// Ensure cleanup after each subtest to reset expectations
		t.Cleanup(func() {
			mockRepo.AssertExpectations(t)
		})

		// Prepare the mock to return an error when CreateCategory is called
		category := &model.Category{CategoryName: "Test Category"}
		mockRepo.On("CreateCategory", category).Return(int64(0), errors.New("database error"))

		// Call AddCategory, which should propagate the error
		_, err := service.AddCategory(category)

		// Verify the error is returned correctly
		assert.Error(t, err)
		assert.Equal(t, "database error", err.Error())
	})
}

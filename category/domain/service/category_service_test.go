package service

import (
	"errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"github.com/tongs-dev/shopping-platform/category/domain/model"
	"testing"
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

// CategoryServiceTestSuite is the test suite for CategoryService
type CategoryServiceTestSuite struct {
	suite.Suite
	mockRepo *MockCategoryRepository
	service  ICategoryService
}

// SetupTest runs before each test
func (suite *CategoryServiceTestSuite) SetupTest() {
	suite.T().Logf("Setup Test")
	suite.mockRepo, suite.service = newCategoryService()
}

// TestCreateCategory tests the CreateCategory method of CategoryService
func (suite *CategoryServiceTestSuite) TestCreateCategory() {
	category := &model.Category{CategoryName: "Test Category"}
	suite.mockRepo.On("CreateCategory", category).Return(int64(1), nil)

	userID, err := suite.service.AddCategory(category)

	suite.NoError(err)
	suite.Equal(int64(1), userID)
	suite.mockRepo.AssertExpectations(suite.T())
}

// TestDeleteCategory tests the DeleteCategory method of CategoryService
func (suite *CategoryServiceTestSuite) TestDeleteCategory() {
	categoryID := int64(1)
	suite.mockRepo.On("DeleteCategoryByID", categoryID).Return(nil)

	err := suite.service.DeleteCategory(categoryID)

	suite.NoError(err)
	suite.mockRepo.AssertExpectations(suite.T())
}

// TestUpdateCategory tests the UpdateCategory method of CategoryService
func (suite *CategoryServiceTestSuite) TestUpdateCategory() {
	category := &model.Category{ID: 1, CategoryName: "Updated Category"}
	suite.mockRepo.On("UpdateCategory", category).Return(nil)

	err := suite.service.UpdateCategory(category)

	suite.NoError(err)
	suite.mockRepo.AssertExpectations(suite.T())
}

// TestFindCategoryByID tests the FindCategoryByID method of CategoryService
func (suite *CategoryServiceTestSuite) TestFindCategoryByID() {
	categoryID := int64(1)
	expectedCategory := &model.Category{ID: categoryID, CategoryName: "Test Category"}
	suite.mockRepo.On("FindCategoryByID", categoryID).Return(expectedCategory, nil)

	category, err := suite.service.FindCategoryByID(categoryID)

	suite.NoError(err)
	suite.Equal(expectedCategory, category)
	suite.mockRepo.AssertExpectations(suite.T())
}

// TestFindAllCategory tests the FindAllCategory method of CategoryService
func (suite *CategoryServiceTestSuite) TestFindAllCategory() {
	expectedCategories := []model.Category{
		{ID: 1, CategoryName: "Category 1"},
		{ID: 2, CategoryName: "Category 2"},
	}
	suite.mockRepo.On("FindAll").Return(expectedCategories, nil)

	categories, err := suite.service.FindAllCategory()

	suite.NoError(err)
	suite.Equal(expectedCategories, categories)
	suite.mockRepo.AssertExpectations(suite.T())
}

// TestErrorHandling tests error scenarios in the service methods
func (suite *CategoryServiceTestSuite) TestErrorHandling() {
	category := &model.Category{CategoryName: "Test Category"}
	suite.mockRepo.On("CreateCategory", category).Return(int64(0), errors.New("database error"))

	_, err := suite.service.AddCategory(category)

	suite.Error(err)
	suite.Equal("database error", err.Error())
	suite.mockRepo.AssertExpectations(suite.T())
}

// Run the tests
func TestCategoryServiceTestSuite(t *testing.T) {
	suite.Run(t, new(CategoryServiceTestSuite))
}

package handler

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"github.com/tongs-dev/shopping-platform/category/domain/model"
	categorypb "github.com/tongs-dev/shopping-platform/category/proto/category"
)

// MockCategoryService is a mock type for the ICategoryService interface
type MockCategoryService struct {
	mock.Mock
}

func (m *MockCategoryService) AddCategory(category *model.Category) (int64, error) {
	args := m.Called(category)
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockCategoryService) UpdateCategory(category *model.Category) error {
	args := m.Called(category)
	return args.Error(0)
}

func (m *MockCategoryService) DeleteCategory(categoryID int64) error {
	args := m.Called(categoryID)
	return args.Error(0)
}

func (m *MockCategoryService) FindCategoryByID(categoryID int64) (*model.Category, error) {
	args := m.Called(categoryID)
	return args.Get(0).(*model.Category), args.Error(1)
}

func (m *MockCategoryService) FindCategoryByName(categoryName string) (*model.Category, error) {
	args := m.Called(categoryName)
	return args.Get(0).(*model.Category), args.Error(1)
}

func (m *MockCategoryService) FindCategoryByLevel(level uint32) ([]model.Category, error) {
	args := m.Called(level)
	return args.Get(0).([]model.Category), args.Error(1)
}

func (m *MockCategoryService) FindCategoryByParent(parent int64) ([]model.Category, error) {
	args := m.Called(parent)
	return args.Get(0).([]model.Category), args.Error(1)
}

func (m *MockCategoryService) FindAllCategory() ([]model.Category, error) {
	args := m.Called()
	return args.Get(0).([]model.Category), args.Error(1)
}

// CategoryHandlerTestSuite is the test suite for CategoryHandler
type CategoryHandlerTestSuite struct {
	suite.Suite
	mockService *MockCategoryService
	handler     *CategoryHandler
}

// SetupTest initializes the test environment for each test
func (suite *CategoryHandlerTestSuite) SetupTest() {
	suite.mockService = new(MockCategoryService)
	suite.handler = &CategoryHandler{CategoryService: suite.mockService}
}

// TestCreateCategory tests the CreateCategory method
func (suite *CategoryHandlerTestSuite) TestCreateCategory() {
	categoryRequest := &categorypb.CategoryRequest{CategoryName: "New Category"}
	response := &categorypb.CreateCategoryResponse{}
	suite.mockService.On("AddCategory", mock.AnythingOfType("*model.Category")).Return(int64(1), nil)

	err := suite.handler.CreateCategory(context.Background(), categoryRequest, response)

	suite.NoError(err)
	suite.Equal("Category created successfully", response.Message)
	suite.Equal(int64(1), response.CategoryId)
	suite.mockService.AssertExpectations(suite.T())
}

// TestCreateCategoryError tests error handling for CreateCategory
func (suite *CategoryHandlerTestSuite) TestCreateCategoryError() {
	categoryRequest := &categorypb.CategoryRequest{CategoryName: "New Category"}
	response := &categorypb.CreateCategoryResponse{}
	suite.mockService.On("AddCategory", mock.AnythingOfType("*model.Category")).Return(int64(0), errors.New("database error"))

	err := suite.handler.CreateCategory(context.Background(), categoryRequest, response)

	suite.Error(err)
	suite.Equal("database error", err.Error())
	suite.mockService.AssertExpectations(suite.T())
}

// TestUpdateCategory tests the UpdateCategory method
func (suite *CategoryHandlerTestSuite) TestUpdateCategory() {
	categoryRequest := &categorypb.CategoryRequest{CategoryName: "Updated Category"}
	response := &categorypb.UpdateCategoryResponse{}
	category := &model.Category{CategoryName: "Updated Category"}

	suite.mockService.On("UpdateCategory", category).Return(nil)

	err := suite.handler.UpdateCategory(context.Background(), categoryRequest, response)

	suite.NoError(err)
	suite.Equal("Category updated successfully", response.Message)
	suite.mockService.AssertExpectations(suite.T())
}

// TestDeleteCategory tests the DeleteCategory method
func (suite *CategoryHandlerTestSuite) TestDeleteCategory() {
	categoryRequest := &categorypb.DeleteCategoryRequest{CategoryId: 1}
	response := &categorypb.DeleteCategoryResponse{}

	suite.mockService.On("DeleteCategory", int64(1)).Return(nil)

	err := suite.handler.DeleteCategory(context.Background(), categoryRequest, response)

	suite.NoError(err)
	suite.Equal("Category deleted successfully", response.Message)
	suite.mockService.AssertExpectations(suite.T())
}

// TestFindCategoryByName tests the FindCategoryByName method
func (suite *CategoryHandlerTestSuite) TestFindCategoryByName() {
	categoryRequest := &categorypb.FindByNameRequest{CategoryName: "Test Category"}
	response := &categorypb.CategoryResponse{}
	category := &model.Category{ID: 1, CategoryName: "Test Category"}

	suite.mockService.On("FindCategoryByName", "Test Category").Return(category, nil)

	err := suite.handler.FindCategoryByName(context.Background(), categoryRequest, response)

	suite.NoError(err)
	suite.Equal(int64(1), response.Id)
	suite.Equal("Test Category", response.CategoryName)
	suite.mockService.AssertExpectations(suite.T())
}

// TestFindCategoryByID tests the FindCategoryByID method
func (suite *CategoryHandlerTestSuite) TestFindCategoryByID() {
	categoryRequest := &categorypb.FindByIdRequest{CategoryId: 1}
	response := &categorypb.CategoryResponse{}
	category := &model.Category{ID: 1, CategoryName: "Test Category"}

	suite.mockService.On("FindCategoryByID", int64(1)).Return(category, nil)

	err := suite.handler.FindCategoryByID(context.Background(), categoryRequest, response)

	suite.NoError(err)
	suite.Equal(int64(1), response.Id)
	suite.Equal("Test Category", response.CategoryName)
	suite.mockService.AssertExpectations(suite.T())
}

// TestFindAllCategory tests the FindAllCategory method
func (suite *CategoryHandlerTestSuite) TestFindAllCategory() {
	categoryRequest := &categorypb.FindAllRequest{}
	response := &categorypb.FindAllResponse{}
	categories := []model.Category{
		{ID: 1, CategoryName: "Category 1"},
		{ID: 2, CategoryName: "Category 2"},
	}

	suite.mockService.On("FindAllCategory").Return(categories, nil)

	err := suite.handler.FindAllCategory(context.Background(), categoryRequest, response)

	suite.NoError(err)
	suite.Len(response.Category, 2)
	suite.mockService.AssertExpectations(suite.T())
}

// Run the tests
func TestCategoryHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(CategoryHandlerTestSuite))
}

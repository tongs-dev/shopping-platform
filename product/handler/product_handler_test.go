package handler

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"github.com/tongs-dev/shopping-platform/product/domain/model"
	productpb "github.com/tongs-dev/shopping-platform/product/proto/product"
)

// MockProductService is a mock type for the IProductService interface
type MockProductService struct {
	mock.Mock
}

func (m *MockProductService) AddProduct(product *model.Product) (int64, error) {
	args := m.Called(product)
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockProductService) DeleteProduct(productID int64) error {
	args := m.Called(productID)
	return args.Error(0)
}

func (m *MockProductService) UpdateProduct(product *model.Product) error {
	args := m.Called(product)
	return args.Error(0)
}

func (m *MockProductService) FindProductByID(productID int64) (*model.Product, error) {
	args := m.Called(productID)
	// Return nil if no product found and error if needed
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Product), args.Error(1)
}

func (m *MockProductService) FindAllProduct() ([]model.Product, error) {
	args := m.Called()
	// If the first argument is nil, we should return an empty slice instead of nil to avoid the panic
	if args.Get(0) == nil {
		return nil, args.Error(1) // Return nil and the error that was set up
	}
	return args.Get(0).([]model.Product), args.Error(1) // Return the mocked products and error
}

// ProductHandlerTestSuite is the test suite for ProductHandler
type ProductHandlerTestSuite struct {
	suite.Suite
	mockService *MockProductService
	handler     *ProductHandler
}

// SetupTest runs before each test
func (suite *ProductHandlerTestSuite) SetupTest() {
	// Initialize mock service and handler
	suite.mockService = new(MockProductService)
	suite.handler = &ProductHandler{ProductService: suite.mockService}
}

// TearDownTest runs after each test to clear mock expectations
func (suite *ProductHandlerTestSuite) TearDownTest() {
	suite.mockService.AssertExpectations(suite.T())
}

// TestAddProduct tests the AddProduct handler
func (suite *ProductHandlerTestSuite) TestAddProduct() {
	product := &productpb.ProductInfo{
		ProductName: "Test Product",
		ProductSku:  "ABC123",
	}
	expectedProduct := &model.Product{
		ProductName: "Test Product",
		ProductSku:  "ABC123",
	}

	// Set up the expectation for AddProduct method
	suite.mockService.On("AddProduct", expectedProduct).Return(int64(1), nil)

	// Prepare response object
	response := &productpb.ResponseProduct{}

	// Call the handler method
	err := suite.handler.AddProduct(nil, product, response)

	// Assert expectations and verify result
	suite.NoError(err)
	suite.Equal(int64(1), response.ProductId)
}

// TestAddProductMissingName tests the AddProduct handler with missing product name
func (suite *ProductHandlerTestSuite) TestAddProductMissingName() {
	// Simulating a missing ProductName
	product := &productpb.ProductInfo{
		ProductSku: "ABC123", // Missing ProductName
	}

	// Prepare response object
	response := &productpb.ResponseProduct{}

	// Mock AddProduct to handle this scenario
	suite.mockService.On("AddProduct", mock.Anything).Return(int64(0), errors.New("product name and SKU are required"))

	// Call the handler method
	err := suite.handler.AddProduct(nil, product, response)

	// Assert error for missing name
	suite.Error(err)
	suite.Equal("failed to add product: product name and SKU are required", err.Error())
}

// TestFindProductByID tests the FindProductByID handler
func (suite *ProductHandlerTestSuite) TestFindProductByID() {
	expectedProduct := &model.Product{
		ID:          1,
		ProductName: "Test Product",
		ProductSku:  "ABC123",
	}

	// Set up the expectation for FindProductByID method
	suite.mockService.On("FindProductByID", int64(1)).Return(expectedProduct, nil)

	// Prepare response object
	response := &productpb.ProductInfo{}

	// Call the handler method
	err := suite.handler.FindProductByID(nil, &productpb.RequestID{ProductId: 1}, response)

	// Assert expectations and verify result
	suite.NoError(err)
	suite.Equal(expectedProduct.ProductName, response.ProductName)
}

// TestFindProductByIDNotFound tests the FindProductByID handler when the product is not found
func (suite *ProductHandlerTestSuite) TestFindProductByIDNotFound() {
	// Set up the expectation for FindProductByID method
	suite.mockService.On("FindProductByID", int64(1)).
		Return(nil, errors.New("error finding product with ID 1: product not found"))

	// Prepare response object
	response := &productpb.ProductInfo{}

	// Call the handler method
	err := suite.handler.FindProductByID(nil, &productpb.RequestID{ProductId: 1}, response)

	// Assert error for product not found
	suite.Error(err)
	suite.Equal("error finding product with ID 1: product not found", err.Error())
}

// TestFindAllProduct tests the FindAllProduct handler
func (suite *ProductHandlerTestSuite) TestFindAllProduct() {
	expectedProducts := []model.Product{
		{ID: 1, ProductName: "Product 1", ProductSku: "SKU1"},
		{ID: 2, ProductName: "Product 2", ProductSku: "SKU2"},
	}

	// Set up the expectation for FindAllProduct method
	suite.mockService.On("FindAllProduct").Return(expectedProducts, nil)

	// Prepare response object
	response := &productpb.AllProduct{}

	// Call the handler method
	err := suite.handler.FindAllProduct(nil, &productpb.RequestAll{}, response)

	// Assert expectations and verify result
	suite.NoError(err)
	suite.Len(response.ProductInfo, 2)
}

// TestFindAllProductError tests the FindAllProduct handler when an error occurs
func (suite *ProductHandlerTestSuite) TestFindAllProductError() {
	// Set up the expectation for FindAllProduct method to return an error
	suite.mockService.On("FindAllProduct").Return(nil, errors.New("failed to fetch products"))

	// Prepare response object
	response := &productpb.AllProduct{}

	// Call the handler method
	err := suite.handler.FindAllProduct(nil, &productpb.RequestAll{}, response)

	// Assert error for failing to fetch products
	suite.Error(err)
	suite.Equal("failed to fetch products", err.Error())
}

// Run the tests using suite
func TestProductHandlerSuite(t *testing.T) {
	suite.Run(t, new(ProductHandlerTestSuite))
}

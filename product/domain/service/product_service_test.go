package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/tongs-dev/shopping-platform/product/domain/model"
)

// MockProductRepository is a mock implementation of the IProductRepository interface
type MockProductRepository struct {
	mock.Mock
}

func (m *MockProductRepository) InitTable() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockProductRepository) FindProductByID(productID int64) (*model.Product, error) {
	args := m.Called(productID)
	return args.Get(0).(*model.Product), args.Error(1)
}

func (m *MockProductRepository) CreateProduct(product *model.Product) (int64, error) {
	args := m.Called(product)
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockProductRepository) DeleteProductByID(productID int64) error {
	args := m.Called(productID)
	return args.Error(0)
}

func (m *MockProductRepository) UpdateProduct(product *model.Product) error {
	args := m.Called(product)
	return args.Error(0)
}

func (m *MockProductRepository) FindAll() ([]model.Product, error) {
	args := m.Called()
	return args.Get(0).([]model.Product), args.Error(1)
}

func TestProductService(t *testing.T) {
	// Initialize mock repository
	mockRepo := new(MockProductRepository)
	service := NewProductService(mockRepo)

	t.Run("AddProduct - Valid", func(t *testing.T) {
		product := mockProduct(1)

		// Setup expectations
		mockRepo.On("CreateProduct", product).Return(int64(1), nil)

		// Call the service method
		productID, err := service.AddProduct(product)

		// Assert the results
		assert.NoError(t, err)
		assert.Equal(t, int64(1), productID)
		mockRepo.AssertExpectations(t)
	})

	t.Run("AddProduct - Missing Name", func(t *testing.T) {
		product := &model.Product{
			ProductSku: "ABC123",
		}

		// Call the service method
		productID, err := service.AddProduct(product)

		// Assert the results
		assert.Error(t, err)
		assert.Equal(t, "product name and SKU are required", err.Error())
		assert.Equal(t, int64(0), productID)
	})

	t.Run("DeleteProduct - Valid ID", func(t *testing.T) {
		productID := int64(1)

		// Setup expectations
		mockRepo.On("DeleteProductByID", productID).Return(nil)

		// Call the service method
		err := service.DeleteProduct(productID)

		// Assert the results
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("DeleteProduct - Invalid ID", func(t *testing.T) {
		productID := int64(-1)

		// Call the service method
		err := service.DeleteProduct(productID)

		// Assert the results
		assert.Error(t, err)
		assert.Equal(t, "invalid product ID", err.Error())
	})

	t.Run("UpdateProduct - Valid", func(t *testing.T) {
		product := mockProduct(1)

		// Setup expectations
		mockRepo.On("UpdateProduct", product).Return(nil)

		// Call the service method
		err := service.UpdateProduct(product)

		// Assert the results
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("UpdateProduct - Invalid ID", func(t *testing.T) {
		product := &model.Product{
			ID:          0,
			ProductName: "Invalid Product",
		}

		// Call the service method
		err := service.UpdateProduct(product)

		// Assert the results
		assert.Error(t, err)
		assert.Equal(t, "invalid product or product ID", err.Error())
	})

	t.Run("FindProductByID - Valid ID", func(t *testing.T) {
		productID := int64(1)
		expectedProduct := mockProduct(1)

		// Setup expectations
		mockRepo.On("FindProductByID", productID).Return(expectedProduct, nil)

		// Call the service method
		product, err := service.FindProductByID(productID)

		// Assert the results
		assert.NoError(t, err)
		assert.Equal(t, expectedProduct, product)
		mockRepo.AssertExpectations(t)
	})

	t.Run("FindProductByID - Invalid ID", func(t *testing.T) {
		productID := int64(-1)

		// Call the service method
		product, err := service.FindProductByID(productID)

		// Assert the results
		assert.Error(t, err)
		assert.Nil(t, product)
		assert.Equal(t, "invalid product ID", err.Error())
	})

	t.Run("FindAllProduct - Success", func(t *testing.T) {
		// Setup expectations
		mockRepo.On("FindAll").Return([]model.Product{
			*mockProduct(1),
			*mockProduct(2),
		}, nil)

		// Call the service method
		products, err := service.FindAllProduct()

		// Assert the results
		assert.NoError(t, err)
		assert.Len(t, products, 2)
		mockRepo.AssertExpectations(t)
	})
}

func mockProduct(id int, name ...string) *model.Product {
	// Set a default value if the name is not provided
	defaultName := "Default Product Name"
	if len(name) > 0 && name[0] != "" {
		defaultName = name[0]
	}

	return &model.Product{
		ID:          int64(id),
		ProductName: defaultName,
		ProductSku:  "ABC123",
	}
}

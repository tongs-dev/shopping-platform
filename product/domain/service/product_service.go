package service

import (
	"errors"
	"log"

	"github.com/tongs-dev/shopping-platform/product/domain/model"
	"github.com/tongs-dev/shopping-platform/product/domain/repository"
)

type IProductService interface {
	AddProduct(*model.Product) (int64, error)
	DeleteProduct(int64) error
	UpdateProduct(*model.Product) error
	FindProductByID(int64) (*model.Product, error)
	FindAllProduct() ([]model.Product, error)
}

func NewProductService(productRepository repository.IProductRepository) IProductService {
	return &ProductService{productRepository}
}

type ProductService struct {
	ProductRepository repository.IProductRepository
}

func (u *ProductService) AddProduct(product *model.Product) (int64, error) {
	if product == nil {
		return 0, errors.New("product cannot be nil")
	}

	// Check if the product has the necessary fields filled
	if product.ProductName == "" || product.ProductSku == "" {
		return 0, errors.New("product name and SKU are required")
	}

	// Call repository to add the product
	productID, err := u.ProductRepository.CreateProduct(product)
	if err != nil {
		log.Printf("error creating product: %v", err)
		return 0, err
	}

	return productID, nil
}

func (u *ProductService) DeleteProduct(productID int64) error {
	if productID <= 0 {
		return errors.New("invalid product ID")
	}

	// Call repository to delete the product
	err := u.ProductRepository.DeleteProductByID(productID)
	if err != nil {
		log.Printf("error deleting product with ID %d: %v", productID, err)
		return err
	}

	return nil
}

func (u *ProductService) UpdateProduct(product *model.Product) error {
	if product == nil || product.ID == 0 {
		return errors.New("invalid product or product ID")
	}

	// Call repository to update the product
	err := u.ProductRepository.UpdateProduct(product)
	if err != nil {
		log.Printf("error updating product with ID %d: %v", product.ID, err)
		return err
	}

	return nil
}

func (u *ProductService) FindProductByID(productID int64) (*model.Product, error) {
	if productID <= 0 {
		return nil, errors.New("invalid product ID")
	}

	// Call repository to find the product
	product, err := u.ProductRepository.FindProductByID(productID)
	if err != nil {
		log.Printf("error finding product with ID %d: %v", productID, err)
		return nil, err
	}

	return product, nil
}

func (u *ProductService) FindAllProduct() ([]model.Product, error) {
	// Call repository to find all products
	products, err := u.ProductRepository.FindAll()
	if err != nil {
		log.Printf("error finding all products: %v", err)
		return nil, err
	}

	return products, nil
}

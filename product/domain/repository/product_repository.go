package repository

import (
	"errors"
	"log"

	"github.com/jinzhu/gorm"
	"github.com/tongs-dev/shopping-platform/product/domain/model"
)

type IProductRepository interface {
	InitTable() error
	FindProductByID(int64) (*model.Product, error)
	CreateProduct(*model.Product) (int64, error)
	DeleteProductByID(int64) error
	UpdateProduct(*model.Product) error
	FindAll() ([]model.Product, error)
}

func NewProductRepository(db *gorm.DB) IProductRepository {
	return &ProductRepository{mysqlDb: db}
}

type ProductRepository struct {
	mysqlDb *gorm.DB
}

// InitTable initializes the product-related tables in the database.
func (u *ProductRepository) InitTable() error {
	if err := u.mysqlDb.CreateTable(&model.Product{}, &model.ProductSeo{}, &model.ProductImage{}, &model.ProductSize{}).Error; err != nil {
		log.Printf("Error initializing tables: %v", err)
		return err
	}
	return nil
}

// FindProductByID retrieves a product by its ID with related images, size, and SEO data.
func (u *ProductRepository) FindProductByID(productID int64) (product *model.Product, err error) {
	if productID <= 0 {
		return nil, errors.New("invalid product ID")
	}

	product = &model.Product{}
	err = u.mysqlDb.
		Preload("ProductImage").
		Preload("ProductSize").
		Preload("ProductSeo").
		First(product, productID).Error

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, errors.New("product not found")
		}
		log.Printf("Error finding product by ID %d: %v", productID, err)
		return nil, err
	}

	return product, nil
}

// CreateProduct inserts a new product into the database.
func (u *ProductRepository) CreateProduct(product *model.Product) (int64, error) {
	if err := u.mysqlDb.Create(product).Error; err != nil {
		log.Printf("Error creating product: %v", err)
		return 0, err
	}

	return product.ID, nil
}

// DeleteProductByID deletes a product and its associated data (images, sizes, SEO).
func (u *ProductRepository) DeleteProductByID(productID int64) error {
	tx := u.mysqlDb.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if tx.Error != nil {
		return tx.Error
	}

	// Delete product and its related data in a transactional manner
	if err := tx.Unscoped().Where("id = ?", productID).Delete(&model.Product{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Unscoped().Where("image_product_id = ?", productID).Delete(&model.ProductImage{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Unscoped().Where("size_product_id = ?", productID).Delete(&model.ProductSize{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Unscoped().Where("seo_product_id = ?", productID).Delete(&model.ProductSeo{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// UpdateProduct updates an existing product's data.
func (u *ProductRepository) UpdateProduct(product *model.Product) error {
	if err := u.mysqlDb.Model(product).Updates(product).Error; err != nil {
		log.Printf("Error updating product with ID %d: %v", product.ID, err)
		return err
	}

	return nil
}

// FindAll retrieves all products with related images, sizes, and SEO data.
func (u *ProductRepository) FindAll() (productAll []model.Product, err error) {
	err = u.mysqlDb.Preload("ProductImage").Preload("ProductSize").Preload("ProductSeo").
		Find(&productAll).Error
	if err != nil {
		log.Printf("Error retrieving all products: %v", err)
		return nil, err
	}
	return productAll, nil
}

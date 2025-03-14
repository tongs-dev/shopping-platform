package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/tongs-dev/shopping-platform/category/domain/model"
)

// ICategoryRepository defines the interface for interacting with the Category repository.
type ICategoryRepository interface {
	// InitTable initializes the Category table in the database.
	InitTable() error

	// FindCategoryByID retrieves a Category by its ID.
	FindCategoryByID(int64) (*model.Category, error)

	// CreateCategory inserts a new Category into the database.
	CreateCategory(*model.Category) (int64, error)

	// DeleteCategoryByID deletes a Category by its ID.
	DeleteCategoryByID(int64) error

	// UpdateCategory updates an existing Category's information.
	UpdateCategory(*model.Category) error

	// FindAll retrieves all Categories from the database.
	FindAll() ([]model.Category, error)

	// FindCategoryByName retrieves a Category by its name.
	FindCategoryByName(string) (*model.Category, error)

	// FindCategoryByLevel retrieves Categories by their level.
	FindCategoryByLevel(uint32) ([]model.Category, error)

	// FindCategoryByParent retrieves Categories by their parent category ID.
	FindCategoryByParent(int64) ([]model.Category, error)
}

// NewCategoryRepository creates and returns a new instance of CategoryRepository.
func NewCategoryRepository(db *gorm.DB) ICategoryRepository {
	return &CategoryRepository{mysqlDb: db}
}

// CategoryRepository implements the ICategoryRepository interface, handling
// interactions with the database using GORM.
type CategoryRepository struct {
	mysqlDb *gorm.DB
}

// InitTable initializes the Category table in the database if it does not already exist.
func (r *CategoryRepository) InitTable() error {
	// Creates the Category table based on the Category model
	err := r.mysqlDb.CreateTable(&model.Category{}).Error
	if err != nil {
		return err
	}
	return nil
}

// FindCategoryByID retrieves a Category by its ID from the database.
func (r *CategoryRepository) FindCategoryByID(categoryID int64) (*model.Category, error) {
	category := &model.Category{}
	// Retrieves the first category that matches the given category ID
	err := r.mysqlDb.First(category, categoryID).Error
	if err != nil {
		return nil, err
	}
	return category, nil
}

// CreateCategory inserts a new Category into the database.
func (r *CategoryRepository) CreateCategory(category *model.Category) (int64, error) {
	// Inserts the new category into the database
	err := r.mysqlDb.Create(category).Error
	if err != nil {
		return 0, err
	}
	return category.ID, nil
}

// DeleteCategoryByID deletes a Category from the database by its ID.
func (r *CategoryRepository) DeleteCategoryByID(categoryID int64) error {
	// Deletes the category with the given ID from the database
	err := r.mysqlDb.Where("id = ?", categoryID).Delete(&model.Category{}).Error
	if err != nil {
		return err
	}
	return nil
}

// UpdateCategory updates an existing Category's information in the database.
func (r *CategoryRepository) UpdateCategory(category *model.Category) error {
	// Updates the category record with new information
	err := r.mysqlDb.Model(category).Update(category).Error
	if err != nil {
		return err
	}
	return nil
}

// FindAll retrieves all Categories from the database.
func (r *CategoryRepository) FindAll() ([]model.Category, error) {
	var categories []model.Category
	// Retrieves all categories from the database
	err := r.mysqlDb.Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return categories, nil
}

// FindCategoryByName retrieves a Category by its name from the database.
func (r *CategoryRepository) FindCategoryByName(categoryName string) (*model.Category, error) {
	category := &model.Category{}
	// Retrieves the category that matches the provided category name
	err := r.mysqlDb.Where("category_name = ?", categoryName).Find(category).Error
	if err != nil {
		return nil, err
	}
	return category, nil
}

// FindCategoryByLevel retrieves Categories by their level from the database.
func (r *CategoryRepository) FindCategoryByLevel(level uint32) ([]model.Category, error) {
	var categories []model.Category
	// Retrieves all categories that match the provided level
	err := r.mysqlDb.Where("category_level = ?", level).Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return categories, nil
}

// FindCategoryByParent retrieves Categories by their parent category ID from the database.
func (r *CategoryRepository) FindCategoryByParent(parent int64) ([]model.Category, error) {
	var categories []model.Category
	// Retrieves all categories that belong to the given parent category
	err := r.mysqlDb.Where("category_parent = ?", parent).Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return categories, nil
}

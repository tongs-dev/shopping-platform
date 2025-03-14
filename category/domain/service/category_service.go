package service

import (
	"github.com/tongs-dev/shopping-platform/category/domain/model"
	"github.com/tongs-dev/shopping-platform/category/domain/repository"
)

// ICategoryService defines the interface for category data operations.
type ICategoryService interface {
	// AddCategory creates a new Category.
	AddCategory(*model.Category) (int64, error)

	// DeleteCategory removes a Category by its ID.
	DeleteCategory(int64) error

	// UpdateCategory updates an existing Category's information.
	UpdateCategory(*model.Category) error

	// FindCategoryByID retrieves a Category by its ID.
	FindCategoryByID(int64) (*model.Category, error)

	// FindAllCategory retrieves all Categories.
	FindAllCategory() ([]model.Category, error)

	// FindCategoryByName retrieves a Category by its name.
	FindCategoryByName(string) (*model.Category, error)

	// FindCategoryByLevel retrieves Categories by their level.
	FindCategoryByLevel(uint32) ([]model.Category, error)

	// FindCategoryByParent retrieves Categories by their parent category ID.
	FindCategoryByParent(int64) ([]model.Category, error)
}

// NewCategoryService creates and returns a new instance of CategoryService.
func NewCategoryService(categoryRepository repository.ICategoryRepository) ICategoryService {
	return &CategoryService{CategoryRepository: categoryRepository}
}

// CategoryService implements the ICategoryService interface and handles
// the logic for managing Categories by calling the repository methods.
type CategoryService struct {
	CategoryRepository repository.ICategoryRepository
}

// AddCategory creates a new Category in the repository.
func (u *CategoryService) AddCategory(category *model.Category) (int64, error) {
	return u.CategoryRepository.CreateCategory(category)
}

// DeleteCategory deletes a Category by its ID from the repository.
func (u *CategoryService) DeleteCategory(categoryID int64) error {
	return u.CategoryRepository.DeleteCategoryByID(categoryID)
}

// UpdateCategory updates an existing Category in the repository.
func (u *CategoryService) UpdateCategory(category *model.Category) error {
	return u.CategoryRepository.UpdateCategory(category)
}

// FindCategoryByID retrieves a Category by its ID from the repository.
func (u *CategoryService) FindCategoryByID(categoryID int64) (*model.Category, error) {
	return u.CategoryRepository.FindCategoryByID(categoryID)
}

// FindAllCategory retrieves all Categories from the repository.
func (u *CategoryService) FindAllCategory() ([]model.Category, error) {
	return u.CategoryRepository.FindAll()
}

// FindCategoryByName retrieves a Category by its name from the repository.
func (u *CategoryService) FindCategoryByName(categoryName string) (*model.Category, error) {
	return u.CategoryRepository.FindCategoryByName(categoryName)
}

// FindCategoryByLevel retrieves Categories by their level from the repository.
func (u *CategoryService) FindCategoryByLevel(level uint32) ([]model.Category, error) {
	return u.CategoryRepository.FindCategoryByLevel(level)
}

// FindCategoryByParent retrieves Categories by their parent category ID from the repository.
func (u *CategoryService) FindCategoryByParent(parent int64) ([]model.Category, error) {
	return u.CategoryRepository.FindCategoryByParent(parent)
}

package handler

import (
	"context"
	"errors"
	"github.com/prometheus/common/log"
	"github.com/tongs-dev/shopping-platform/category/common"
	"github.com/tongs-dev/shopping-platform/category/domain/model"
	"github.com/tongs-dev/shopping-platform/category/domain/service"
	categorypb "github.com/tongs-dev/shopping-platform/category/proto/category"
)

type CategoryHandler struct {
	CategoryService service.ICategoryService
}

// Helper function to handle error response
func handleErrorResponse(err error) error {
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

// Helper function to map category to response
func mapCategoryToResponse(category *model.Category, response interface{}) error {
	return common.SwapTo(category, response)
}

// CreateCategory provides a service to create a new category
func (c *CategoryHandler) CreateCategory(ctx context.Context, request *categorypb.CategoryRequest, response *categorypb.CreateCategoryResponse) error {
	category := &model.Category{}
	if err := common.SwapTo(request, category); err != nil {
		return handleErrorResponse(err)
	}

	categoryId, err := c.CategoryService.AddCategory(category)
	if err != nil {
		return handleErrorResponse(err)
	}

	response.Message = "Category created successfully"
	response.CategoryId = categoryId
	return nil
}

// UpdateCategory provides a service to update an existing category
func (c *CategoryHandler) UpdateCategory(ctx context.Context, request *categorypb.CategoryRequest, response *categorypb.UpdateCategoryResponse) error {
	category := &model.Category{}
	if err := common.SwapTo(request, category); err != nil {
		return handleErrorResponse(err)
	}

	if err := c.CategoryService.UpdateCategory(category); err != nil {
		return handleErrorResponse(err)
	}

	response.Message = "Category updated successfully"
	return nil
}

// DeleteCategory provides a service to delete a category by ID
func (c *CategoryHandler) DeleteCategory(ctx context.Context, request *categorypb.DeleteCategoryRequest, response *categorypb.DeleteCategoryResponse) error {
	if err := c.CategoryService.DeleteCategory(request.CategoryId); err != nil {
		return handleErrorResponse(err)
	}

	response.Message = "Category deleted successfully"
	return nil
}

// FindCategoryByName finds a category by its name
func (c *CategoryHandler) FindCategoryByName(ctx context.Context, request *categorypb.FindByNameRequest, response *categorypb.CategoryResponse) error {
	category, err := c.CategoryService.FindCategoryByName(request.CategoryName)
	if err != nil {
		return handleErrorResponse(err)
	}

	return mapCategoryToResponse(category, response)
}

// FindCategoryByID finds a category by its ID
func (c *CategoryHandler) FindCategoryByID(ctx context.Context, request *categorypb.FindByIdRequest, response *categorypb.CategoryResponse) error {
	category, err := c.CategoryService.FindCategoryByID(request.CategoryId)
	if err != nil {
		return handleErrorResponse(err)
	}

	return mapCategoryToResponse(category, response)
}

// FindCategoryByLevel finds categories by their level
func (c *CategoryHandler) FindCategoryByLevel(ctx context.Context, request *categorypb.FindByLevelRequest, response *categorypb.FindAllResponse) error {
	return c.findCategories(ctx, request.Level, response)
}

// FindCategoryByParent finds categories by their parent ID
func (c *CategoryHandler) FindCategoryByParent(ctx context.Context, request *categorypb.FindByParentRequest, response *categorypb.FindAllResponse) error {
	return c.findCategories(ctx, request.ParentId, response)
}

// FindAllCategory retrieves all categories
func (c *CategoryHandler) FindAllCategory(ctx context.Context, request *categorypb.FindAllRequest, response *categorypb.FindAllResponse) error {
	categorySlice, err := c.CategoryService.FindAllCategory()
	if err != nil {
		return handleErrorResponse(err)
	}

	return mapCategoriesToResponse(categorySlice, response)
}

// Utility function to map multiple categories to a response
func mapCategoriesToResponse(categorySlice []model.Category, response *categorypb.FindAllResponse) error {
	for _, cg := range categorySlice {
		cr := &categorypb.CategoryResponse{}
		if err := common.SwapTo(cg, cr); err != nil {
			log.Error(err)
			return err
		}
		response.Category = append(response.Category, cr)
	}
	return nil
}

// Helper function to reduce duplication in FindCategoryByLevel and FindCategoryByParent
func (c *CategoryHandler) findCategories(ctx context.Context, parentOrLevel interface{}, response *categorypb.FindAllResponse) error {
	var categorySlice []model.Category
	var err error

	switch v := parentOrLevel.(type) {
	case uint32:
		categorySlice, err = c.CategoryService.FindCategoryByLevel(v)
	case int64:
		categorySlice, err = c.CategoryService.FindCategoryByParent(v)
	default:
		err = errors.New("invalid parameter type")
	}

	if err != nil {
		return handleErrorResponse(err)
	}

	return mapCategoriesToResponse(categorySlice, response)
}

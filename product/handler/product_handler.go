package handler

import (
	"context"
	"fmt"

	"github.com/tongs-dev/shopping-platform/product/common"
	"github.com/tongs-dev/shopping-platform/product/domain/model"
	"github.com/tongs-dev/shopping-platform/product/domain/service"
	productpb "github.com/tongs-dev/shopping-platform/product/proto/product"
)

type ProductHandler struct {
	ProductService service.IProductService
}

// AddProduct handles adding a new product.
func (h *ProductHandler) AddProduct(ctx context.Context, request *productpb.ProductInfo, response *productpb.ResponseProduct) error {
	// Create a model object for Product
	productAdd := &model.Product{}

	// Convert gRPC request to model object
	if err := common.SwapTo(request, productAdd); err != nil {
		return fmt.Errorf("failed to convert request to product model: %v", err)
	}

	// Call service to add the product
	productID, err := h.ProductService.AddProduct(productAdd)
	if err != nil {
		return fmt.Errorf("failed to add product: %v", err)
	}

	// Set product ID in the response
	response.ProductId = productID
	return nil
}

// FindProductByID retrieves a product by ID.
func (h *ProductHandler) FindProductByID(ctx context.Context, request *productpb.RequestID, response *productpb.ProductInfo) error {
	// Fetch product data from service
	productData, err := h.ProductService.FindProductByID(request.ProductId)
	if err != nil {
		return err
	}

	// Map the product data to the response
	if err := common.SwapTo(productData, response); err != nil {
		return err
	}

	return nil
}

// UpdateProduct updates an existing product.
func (h *ProductHandler) UpdateProduct(ctx context.Context, request *productpb.ProductInfo, response *productpb.Response) error {
	// Convert request to model object
	productUpdate := &model.Product{}
	if err := common.SwapTo(request, productUpdate); err != nil {
		return fmt.Errorf("failed to convert request to product model: %v", err)
	}

	// Update product using service
	if err := h.ProductService.UpdateProduct(productUpdate); err != nil {
		return err
	}

	// Set success message in the response
	response.Msg = "Product updated successfully"
	return nil
}

// DeleteProductByID deletes a product by ID.
func (h *ProductHandler) DeleteProductByID(ctx context.Context, request *productpb.RequestID, response *productpb.Response) error {
	// Delete the product by ID using service
	if err := h.ProductService.DeleteProduct(request.ProductId); err != nil {
		return err
	}

	// Set success message in the response
	response.Msg = "Product deleted successfully"
	return nil
}

// FindAllProduct retrieves all products.
func (h *ProductHandler) FindAllProduct(ctx context.Context, request *productpb.RequestAll, response *productpb.AllProduct) error {
	// Fetch all products from the service
	productAll, err := h.ProductService.FindAllProduct()
	if err != nil {
		return err
	}

	// Convert products to gRPC response format
	for _, v := range productAll {
		productInfo := &productpb.ProductInfo{}
		if err := common.SwapTo(v, productInfo); err != nil {
			return fmt.Errorf("failed to convert product to product info: %v", err)
		}
		response.ProductInfo = append(response.ProductInfo, productInfo)
	}

	return nil
}

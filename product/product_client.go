package main

import (
	"context"
	"fmt"
	"log"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	opentracingPlugin "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	"github.com/opentracing/opentracing-go"

	"github.com/tongs-dev/shopping-platform/product/common"
	go_micro_service_product "github.com/tongs-dev/shopping-platform/product/proto/product"
)

// This is for testing.
func main() {
	// Set up Consul registry
	consulRegistry := registry.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"127.0.0.1:8500", // Consul address
		}
	})

	// Initialize tracing
	tracer, io, err := common.NewTracer("go.micro.service.product.client", "localhost:6831")
	if err != nil {
		log.Fatal("Failed to initialize tracer:", err)
	}
	defer io.Close() // Close the tracer on exit
	opentracing.SetGlobalTracer(tracer)

	// Create new Micro service
	service := micro.NewService(
		micro.Name("go.micro.service.product"),
		micro.Version("latest"),
		micro.Address("127.0.0.1:8085"),
		micro.Registry(consulRegistry),
		micro.WrapHandler(opentracingPlugin.NewHandlerWrapper(opentracing.GlobalTracer())),
	)

	// Initialize ProductService client
	productService := go_micro_service_product.NewProductService("go.micro.service.product", service.Client())

	// Prepare the product to be added
	productAdd := &go_micro_service_product.ProductInfo{
		ProductName:        "product a",
		ProductSku:         "123abc",
		ProductPrice:       1.1,
		ProductDescription: "desc a",
		ProductCategoryId:  1,
		ProductImage: []*go_micro_service_product.ProductImage{
			{
				ImageName: "image a",
				ImageCode: "image01",
				ImageUrl:  "image01",
			},
			{
				ImageName: "image b",
				ImageCode: "image02",
				ImageUrl:  "image02",
			},
		},
		ProductSize: []*go_micro_service_product.ProductSize{
			{
				SizeName: "image size",
				SizeCode: "image-size-code",
			},
		},
		ProductSeo: &go_micro_service_product.ProductSeo{
			SeoTitle:       "image-seo",
			SeoKeywords:    "image-seo",
			SeoDescription: "image-seo",
			SeoCode:        "image-seo",
		},
	}

	// Add the product using the product service
	response, err := productService.AddProduct(context.TODO(), productAdd)
	if err != nil {
		log.Println("Failed to add product:", err)
		return
	}

	// Print the response after adding the product
	fmt.Println("Product added successfully:", response)
}

package repository

import (
	"fmt"
	"log"
	"math/rand"
	"testing"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // Import MySQL dialect
	"github.com/stretchr/testify/assert"
	"github.com/tongs-dev/shopping-platform/product/domain/model"
)

// TestProductRepository tests the ProductRepository methods using MySQL database.
func TestProductRepository(t *testing.T) {
	// Initializes database and repository
	db := setupTestDB(t)
	repo := &ProductRepository{mysqlDb: db}

	t.Run("FindProductByID", func(t *testing.T) {
		product := mockProduct("Test Product")

		_, err := repo.CreateProduct(product)
		assert.NoError(t, err)

		product, err = repo.FindProductByID(product.ID)
		assert.NoError(t, err)
		assert.Equal(t, "Test Product", product.ProductName)
	})

	t.Run("FindProductByID - Not Found", func(t *testing.T) {
		product, err := repo.FindProductByID(99999) // Using a non-existent product ID

		assert.Nil(t, product)
		assert.Error(t, err)
		assert.Equal(t, "product not found", err.Error())
	})

	t.Run("CreateProduct", func(t *testing.T) {
		product := mockProduct()

		id, err := repo.CreateProduct(product)
		assert.NoError(t, err)
		assert.NotZero(t, id)
	})

	t.Run("DeleteProductByID", func(t *testing.T) {
		product := mockProduct()
		id, err := repo.CreateProduct(product)
		assert.NoError(t, err)

		err = repo.DeleteProductByID(id)
		assert.NoError(t, err)

		_, err = repo.FindProductByID(id)
		assert.Error(t, err)
	})

	t.Run("UpdateProduct", func(t *testing.T) {
		product := mockProduct("Old Product Name")
		id, err := repo.CreateProduct(product)
		assert.NoError(t, err)

		product.ProductName = "Updated Product Name"
		err = repo.UpdateProduct(product)
		assert.NoError(t, err)

		updatedProduct, err := repo.FindProductByID(id)
		assert.NoError(t, err)
		assert.Equal(t, "Updated Product Name", updatedProduct.ProductName)
	})

	t.Run("FindAll", func(t *testing.T) {
		clearTable(t, db)

		product1, product2 := mockProduct(), mockProduct()

		// Create two products
		_, err := repo.CreateProduct(product1)
		assert.NoError(t, err)
		_, err = repo.CreateProduct(product2)
		assert.NoError(t, err)

		// Find all products
		products, err := repo.FindAll()
		assert.NoError(t, err)
		assert.Len(t, products, 2)
	})
}

// setupTestDB initializes a real MySQL test database for unit tests.
func setupTestDB(t *testing.T) *gorm.DB {
	dsn := "root:123456@tcp(localhost:3306)/productdb?charset=utf8mb4&parseTime=True&loc=Local"

	// Opens MySQL connection
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to MySQL: %v", err)
	}

	// Automatically migrate the Product model (creating the table)
	err = db.AutoMigrate(&model.Product{}, &model.ProductImage{}, &model.ProductSize{}, &model.ProductSeo{}).Error
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	assert.NoError(t, err, "Failed to migrate test table")

	fmt.Println("MySQL test database setup complete")
	return db
}

// clearTable clears the products table before each test
func clearTable(t *testing.T, db *gorm.DB) {
	// List of tables to be dropped
	tables := []string{"products", "product_images", "product_sizes", "product_seos"}

	// Loop through and drop each table
	for _, table := range tables {
		err := db.Exec(fmt.Sprintf("TRUNCATE TABLE %s", table)).Error
		assert.NoError(t, err, fmt.Sprintf("Failed to drop '%s' table", table))
	}
}

func mockProduct(name ...string) *model.Product {
	// Set a default value if the name is not provided
	defaultName := "Default Product Name"
	if len(name) > 0 && name[0] != "" {
		defaultName = name[0]
	}

	return &model.Product{
		ProductName: defaultName,
		ProductSku:  generateRandomString(10),
	}
}

func generateRandomString(length int) string {
	// Create a new random source with the current Unix timestamp
	randSource := rand.NewSource(time.Now().UnixNano())
	r := rand.New(randSource)

	// Define the characters to use in the random string
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[r.Intn(len(charset))]
	}
	return string(result)
}

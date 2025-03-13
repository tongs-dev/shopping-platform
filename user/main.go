package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/micro/go-micro/v2"
	"log"
	"os"
	"shopping-platform/common/domain/repository"
	"shopping-platform/common/domain/service"
	"shopping-platform/common/handler"
	userpb "shopping-platform/common/proto/user"
)

func main() {
	// Read environment variables for database config
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// MySQL's connection string
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	// 1. Connect to MySQL database
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to MySQL: %v", err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalf("Failed to close DB connection: %v", err)
		}
	}()

	// 2. Configure GORM settings
	db.SingularTable(true) // Use singular table names (e.g., "user" instead of "users")

	// 3. Run database migrations (only once)
	// userRepo := repository.NewUserRepository(db)
	// err = userRepo.InitTable()
	// if err != nil {
	// 	log.Fatalf("Failed to initialize database tables: %v", err)
	// }

	// 4. Create a new microservice instance with name and version
	srv := micro.NewService(
		micro.Name("go.micro.service.user"),
		micro.Version("latest"),
	)

	// 5. Initialize the microservice
	srv.Init()

	// 6. Create a user service instance
	userDataService := service.NewUserDataService(repository.NewUserRepository(db))

	// 7. Register the user handler with the microservice
	if err := userpb.RegisterUserHandler(srv.Server(), &handler.User{UserDataService: userDataService}); err != nil {
		log.Fatalf("Failed to register user service: %v", err)
	}

	// 8. Run the microservice
	log.Println("Starting User Service...")
	if err := srv.Run(); err != nil {
		log.Fatalf("Service failed: %v", err)
	}
}

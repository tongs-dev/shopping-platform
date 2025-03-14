package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/micro/go-micro/v2"
	"github.com/tongs-dev/shopping-platform/user/domain/repository"
	"github.com/tongs-dev/shopping-platform/user/domain/service"
	"github.com/tongs-dev/shopping-platform/user/handler"
	userpb "github.com/tongs-dev/shopping-platform/user/proto/user"
	"github.com/tongs-dev/shopping-platform/user/util"
	"log"
)

func main() {
	// Read environment variables for database config
	dbHost := util.GetEnv("DB_HOST", "localhost")
	dbPort := util.GetEnv("DB_PORT", "3306")
	dbUser := util.GetEnv("DB_USER", "root")
	dbPassword := util.GetEnv("DB_PASSWORD", "123456")
	dbName := util.GetEnv("DB_NAME", "userdb")

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
	if err := userpb.RegisterUserHandler(srv.Server(), &handler.UserHandler{UserDataService: userDataService}); err != nil {
		log.Fatalf("Failed to register user service: %v", err)
	}

	// 8. Run the microservice
	log.Println("Starting User Service...")
	if err := srv.Run(); err != nil {
		log.Fatalf("Service failed: %v", err)
	}
}

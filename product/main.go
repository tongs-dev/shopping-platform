package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
	opentracingPlugin "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	"github.com/opentracing/opentracing-go"
	"log"
	"os"

	"github.com/tongs-dev/shopping-platform/product/common"
	"github.com/tongs-dev/shopping-platform/product/domain/repository"
	productService "github.com/tongs-dev/shopping-platform/product/domain/service"
	"github.com/tongs-dev/shopping-platform/product/handler"
	productpb "github.com/tongs-dev/shopping-platform/product/proto/product"
)

// setupConsulConfig loads the Consul configuration
func setupConsulConfig() (config.Config, error) {
	consulConfig, err := common.GetConsulConfig("127.0.0.1", 8500, "/micro/config")
	if err != nil {
		log.Printf("Error connecting to Consul: %v", err)
		return nil, err
	}
	return consulConfig, nil
}

// setupConsulRegistry sets up the Consul registry
func setupConsulRegistry() registry.Registry {
	return consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"127.0.0.1:8500",
		}
	})
}

// setupMySQLConnection establishes the MySQL connection
func setupMySQLConnection(config config.Config) (*gorm.DB, error) {
	mysqlInfo, err := common.GetMysqlFromConsul(config, "mysql")
	if err != nil {
		log.Fatalf("Error getting MySQL config: %v", err)
		return nil, err
	}

	dsn := mysqlInfo.User + ":" + mysqlInfo.Pwd + "@/" + mysqlInfo.Database + "?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error connecting to MySQL: %v", err)
		return nil, err
	}

	// Ensure singular table naming convention
	db.SingularTable(true)
	return db, nil
}

// setupService initializes the microservice with Consul registry and config
func setupService(consul registry.Registry) micro.Service {
	return micro.NewService(
		micro.Name("go.micro.service.product"),
		micro.Version("latest"),
		micro.Address("127.0.0.1:8085"),
		micro.Registry(consul),
		micro.WrapHandler(opentracingPlugin.NewHandlerWrapper(opentracing.GlobalTracer())),
	)
}

func main() {
	// Setup configuration and service
	consulConfig, err := setupConsulConfig()
	if err != nil {
		log.Fatal("Failed to set up Consul config")
		os.Exit(1)
	}
	consulRegistry := setupConsulRegistry()

	// Set up tracing
	tracer, io, err := common.NewTracer("go.micro.service.product", "localhost:6831")
	if err != nil {
		log.Fatalf("Failed to initialize tracer: %v", err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(tracer)

	service := setupService(consulRegistry)

	// Establish MySQL connection
	db, err := setupMySQLConnection(consulConfig)
	if err != nil {
		log.Fatal("Failed to connect to MySQL")
		os.Exit(1)
	}
	defer db.Close()

	// Initialise service
	service.Init()

	// Set up the category data service
	categoryDataService := productService.NewProductService(repository.NewProductRepository(db))

	// Register the handler
	err = productpb.RegisterProductHandler(service.Server(), &handler.ProductHandler{ProductService: categoryDataService})
	if err != nil {
		log.Fatalf("Error registering category handler: %v", err)
	}

	// Run the service
	if err := service.Run(); err != nil {
		log.Fatalf("Error running the service: %v", err)
	}
}

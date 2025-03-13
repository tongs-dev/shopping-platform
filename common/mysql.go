package common

import (
	"fmt"
	"github.com/micro/go-micro/v2/config"
	"log"
)

// 创建结构体
type MysqlConfig struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Pwd      string `json:"pwd"`
	Database string `json:"database"`
	Port     int64  `json:"port"`
}

// GetMysqlFromConsul retrieves MySQL configuration from Consul using the provided config.Config object.
func GetMysqlFromConsul(config config.Config, path ...string) (*MysqlConfig, error) {
	mysqlConfig := &MysqlConfig{}

	// Retrieve the configuration value
	value := config.Get(path...)

	// Check if the value is empty or nil
	if len(value.Bytes()) == 0 {
		log.Printf("MySQL config not found at path: %v, using default config", path)
		return nil, fmt.Errorf("MySQL config not found at path: %v", path)
	}

	// Scan the configuration into the struct
	if err := value.Scan(mysqlConfig); err != nil {
		log.Printf("Failed to load MySQL config from Consul: %v", err)
		return nil, fmt.Errorf("failed to scan MySQL config: %w", err)
	}

	return mysqlConfig, nil
}

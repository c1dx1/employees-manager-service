package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	PostgresHost        string
	PostgresPort        string
	PostgresUser        string
	PostgresPassword    string
	PostgresDB          string
	ServicesNetworkType string
	EmployeePort        string
}

func LoadConfig() (*Config, error) {
	viper.SetConfigFile("./config/config.env")
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	config := &Config{
		PostgresHost:        viper.GetString("POSTGRES_HOST"),
		PostgresPort:        viper.GetString("POSTGRES_PORT"),
		PostgresUser:        viper.GetString("POSTGRES_USER"),
		PostgresPassword:    viper.GetString("POSTGRES_PASSWORD"),
		PostgresDB:          viper.GetString("POSTGRES_DB"),
		ServicesNetworkType: viper.GetString("SERVICES_NETWORK_TYPE"),
		EmployeePort:        viper.GetString("EMPLOYEE_PORT"),
	}
	return config, nil
}

func (cfg *Config) PostgresURL() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresHost, cfg.PostgresPort, cfg.PostgresDB)
}

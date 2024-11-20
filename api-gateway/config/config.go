package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Address      string
	GatewayPort  string
	EmployeePort string
}

func LoadConfig() (*Config, error) {
	viper.SetConfigFile("./config/config.env")
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	config := &Config{
		Address:      viper.GetString("ADDRESS"),
		GatewayPort:  viper.GetString("GATEWAY_PORT"),
		EmployeePort: viper.GetString("EMPLOYEE_PORT"),
	}
	return config, nil
}

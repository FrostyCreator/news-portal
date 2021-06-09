package config

import (
	"fmt"
	"github.com/FrostyCreator/news-portal/internal/logger"
	"github.com/spf13/viper"
	"time"
)

type (
	Config struct {
		PostgreSQL PostgreSQLConfig
		HTTP       HTTPConfig
	}

	PostgreSQLConfig struct {
		Host     string
		User     string
		Password string
		DBName   string `mapstructure:"database"`
	}

	HTTPConfig struct {
		Port               string        `mapstructure:"port"`
		ReadTimeout        time.Duration `mapstructure:"readTimeout"`
		WriteTimeout       time.Duration `mapstructure:"writeTimeout"`
		MaxHeaderMegabytes int           `mapstructure:"maxHeaderBytes"`
	}
)

func (c *Config) String() string {
	return fmt.Sprintf("PostgreSQL configs:\n\tHost: %s\n\tUser: %s\n\tDB name: %s\n"+
		"HTTP configs:\n\tPort: %s\n\tRead timeout:%s\n\tWrite timeout: %s",
		c.PostgreSQL.Host, c.PostgreSQL.User, c.PostgreSQL.DBName, c.HTTP.Port, c.HTTP.ReadTimeout, c.HTTP.WriteTimeout)
}

func GetConf(configPath string) (*Config, error) {
	config := new(Config)

	if err := setFromConfigFile(config, configPath); err != nil {
		return nil, err
	}

	if err := setFromEnv(config); err != nil {
		return nil, err
	}

	return config, nil
}

func setFromConfigFile(config *Config, configPath string) error {
	viper.AddConfigPath(configPath)
	viper.SetConfigName("main")
	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		logger.LogErrorf("Error reading config file, %s", err)
		return err
	}

	if err := viper.Unmarshal(config); err != nil {
		return err
	}

	return nil
}

func setFromEnv(config *Config) error {
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("dotenv")

	if err := viper.ReadInConfig(); err != nil {
		logger.LogErrorf("Error reading config file, %s", err)
		return err
	}

	config.PostgreSQL.Host = viper.GetString("DATABASE_HOST")
	config.PostgreSQL.User = viper.GetString("DATABASE_USER")
	config.PostgreSQL.Password = viper.GetString("DATABASE_PASSWORD")

	return nil
}

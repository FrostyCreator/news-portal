package config

import (
	"fmt"
	"github.com/FrostyCreator/news-portal/pkg/logger"
	"github.com/spf13/viper"
	"os"
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
		URI      string
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

	envVarExist := os.Getenv("ENV_VAR_EXISTS")

	if envVarExist == "true" {
		if err := setConfFromEnvVar(config); err != nil {
			return nil, err
		}
	} else {
		if err := setConfFromEnvFile(config); err != nil {
			return nil, err
		}
	}

	if err := setBasicConfigs(config, configPath); err != nil {
		return nil, err
	}

	return config, nil
}

func setConfFromEnvFile(config *Config) error {
	viper.AddConfigPath("./")
	viper.SetConfigType("env")
	viper.SetConfigName(".env")
	viper.SetEnvPrefix("database")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	config.PostgreSQL.Host = viper.GetString("DATABASE_HOST")
	config.PostgreSQL.User = viper.GetString("DATABASE_USER")
	config.PostgreSQL.Password = viper.GetString("DATABASE_PASSWORD")
	config.PostgreSQL.URI = viper.GetString("DATABASE_URI")

	return nil
}

func setConfFromEnvVar(config *Config) error {
	config.PostgreSQL.Host = os.Getenv("DATABASE_HOST")
	config.PostgreSQL.User = os.Getenv("DATABASE_USER")
	config.PostgreSQL.Password = os.Getenv("DATABASE_PASSWORD")
	config.PostgreSQL.URI = viper.GetString("DATABASE_URI")

	return nil
}

func setBasicConfigs(config *Config, configPath string) error {
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

	portFromEnv := os.Getenv("PORT")
	if portFromEnv != "" {
		config.HTTP.Port = portFromEnv
	}

	return nil
}

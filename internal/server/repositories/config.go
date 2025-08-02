package repositories

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"

	"github.com/B-AJ-Amar/gTunnel/internal/server/models"
)

// in case of deployment on a serveless env like vercel we cant use the filesystem -> we cant use config files
// so we use environment variables instead
// this is controlled by the `GTUNNEL_USE_ENV` environment variable
// if `GTUNNEL_USE_ENV` is set to true, we will use environment variables for configuration

const (
	configName = "config"
	configType = "yaml"
	appName    = "gtunnel"
)

type ServerConfigRepository interface {
	InitConfig() error
	Load() (*models.ServerConfig, error)
	Save(config *models.ServerConfig) error
	SetConfig(config *models.ServerConfig) error
	UpdateAccessToken(token string) error
	UpdatePort(port int) error
	SetConfigValue(key string, value interface{}) error
	GetConfigPath() string
}

type ServerConfigRepo struct {
	configPath string
	useEnv     bool
}

func NewServerConfigRepo() ServerConfigRepository {
	configDir, err := os.UserConfigDir()
	if err != nil {
		configDir = "."
	}

	configPath := filepath.Join(configDir, appName)
	return &ServerConfigRepo{
		configPath: configPath,
		useEnv:     os.Getenv("GTUNNEL_USE_ENV") == "true",
	}
}

func (r *ServerConfigRepo) InitConfig() error {
	if r.useEnv {
		fmt.Println("Running in USE_ENV mode — loading config from environment variables")
		viper.AutomaticEnv()
		return nil
	}

	if err := os.MkdirAll(r.configPath, 0755); err != nil {
		return fmt.Errorf("could not create config directory: %w", err)
	}

	viper.AddConfigPath(r.configPath)
	viper.SetConfigName(configName)
	viper.SetConfigType(configType)

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("No config found. Creating default config...")
			defaultConfig := &models.ServerConfig{
				AccessToken: "",
				Port:        7205, // Set default port
			}
			if err := r.SetConfig(defaultConfig); err != nil {
				return fmt.Errorf("could not create default config: %w", err)
			}
			fmt.Printf("Default config created at: %s\n", filepath.Join(r.configPath, configName+"."+configType))
		} else {
			return fmt.Errorf("error reading config: %w", err)
		}
	}

	return nil
}

func (r *ServerConfigRepo) Load() (*models.ServerConfig, error) {
	var config models.ServerConfig

	// Initialize config first to read the config file
	if err := r.InitConfig(); err != nil {
		return nil, fmt.Errorf("failed to initialize config: %w", err)
	}

	if r.useEnv {
		viper.AutomaticEnv()
		_ = viper.BindEnv("access_token", "GTUNNEL_ACCESS_TOKEN")
		_ = viper.BindEnv("port", "GTUNNEL_PORT")
		// Set default values for environment mode
		viper.SetDefault("port", 7205)
	} else {
		// Only set default if not in env mode and value is not in config
		if !viper.IsSet("port") {
			viper.SetDefault("port", 7205)
		}
	}

	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("could not unmarshal config: %w", err)
	}

	return &config, nil
}

func (r *ServerConfigRepo) Save(config *models.ServerConfig) error {
	return r.SetConfig(config)
}

func (r *ServerConfigRepo) SetConfig(config *models.ServerConfig) error {
	if r.useEnv {
		return fmt.Errorf("cannot modify config when USE_ENV is true")
	}
	if config == nil {
		return fmt.Errorf("config is nil")
	}

	viper.Set("access_token", config.AccessToken)
	viper.Set("port", config.Port)

	if err := viper.WriteConfig(); err != nil {
		if err := viper.SafeWriteConfig(); err != nil {
			return fmt.Errorf("could not write config: %w", err)
		}
	}

	return nil
}

func (r *ServerConfigRepo) UpdateAccessToken(token string) error {
	if r.useEnv {
		return fmt.Errorf("cannot update access token in USE_ENV mode")
	}
	config, err := r.Load()
	if err != nil {
		return err
	}
	config.AccessToken = token
	return r.SetConfig(config)
}

func (r *ServerConfigRepo) UpdatePort(port int) error {
	if r.useEnv {
		return fmt.Errorf("cannot update port in USE_ENV mode")
	}
	config, err := r.Load()
	if err != nil {
		return err
	}
	config.Port = port
	return r.SetConfig(config)
}

func (r *ServerConfigRepo) SetConfigValue(key string, value interface{}) error {
	if r.useEnv {
		return fmt.Errorf("cannot set config value when USE_ENV is true")
	}

	viper.Set(key, value)
	if err := viper.WriteConfig(); err != nil {
		return fmt.Errorf("could not write config: %w", err)
	}
	return nil
}

func (r *ServerConfigRepo) GetConfigPath() string {
	return filepath.Join(r.configPath, configName+"."+configType)
}

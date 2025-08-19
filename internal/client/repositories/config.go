package repositories

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/B-AJ-Amar/gTunnel/internal/client/models"
	"github.com/spf13/viper"
)

const (
	configName = "config"
	configType = "yaml"
	appName    = "gtunnel"
)

type ClientConfigRepository interface {
	InitConfig() error

	Load() (*models.ClientConfig, error)

	Save(config *models.ClientConfig) error

	SetConfig(config *models.ClientConfig) error

	UpdateAccessToken(token string) error

	UpdateServerURL(url string) error

	SetConfigValue(key string, value interface{}) error

	GetConfigPath() string
}

type ClientConfigRepo struct {
	configPath string
}

func NewClientConfigRepo() ClientConfigRepository {
	// Determine platform-specific config path
	configDir, err := os.UserConfigDir()
	if err != nil {
		// Fallback to current directory if UserConfigDir fails
		configDir = "."
	}

	configPath := filepath.Join(configDir, appName)
	return &ClientConfigRepo{configPath: configPath}
}

func (r *ClientConfigRepo) InitConfig() error {
	if err := os.MkdirAll(r.configPath, 0755); err != nil {
		return fmt.Errorf("could not create config directory: %w", err)
	}

	viper.AddConfigPath(r.configPath)
	viper.SetConfigName(configName)
	viper.SetConfigType(configType)

	// Try to read existing config
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("No existing config file found. Creating default config...")

			// Set default values
			defaultConfig := &models.ClientConfig{
				AccessToken: "",
				ServerURL:   "localhost:5780",
			}

			// Create the config file with defaults
			if err := r.SetConfig(defaultConfig); err != nil {
				return fmt.Errorf("could not create default config: %w", err)
			}

			fmt.Printf("Default config created at: %s\n", filepath.Join(r.configPath, configName+"."+configType))
		} else {
			return fmt.Errorf("error reading config file: %w", err)
		}
	}

	return nil
}

func (r *ClientConfigRepo) Load() (*models.ClientConfig, error) {
	var config models.ClientConfig

	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("could not unmarshal config: %w", err)
	}

	return &config, nil
}

func (r *ClientConfigRepo) Save(config *models.ClientConfig) error {
	return r.SetConfig(config)
}

func (r *ClientConfigRepo) SetConfig(config *models.ClientConfig) error {
	if config == nil {
		return fmt.Errorf("config is nil")
	}

	viper.Set("access_token", config.AccessToken)
	viper.Set("server_url", config.ServerURL)

	// Write the config to file
	if err := viper.WriteConfig(); err != nil {
		// If WriteConfig fails, try SafeWriteConfig (for first time creation)
		if err := viper.SafeWriteConfig(); err != nil {
			return fmt.Errorf("could not write config file: %w", err)
		}
	}

	return nil
}

func (r *ClientConfigRepo) UpdateAccessToken(token string) error {
	config, err := r.Load()
	if err != nil {
		return err
	}
	config.AccessToken = token
	return r.SetConfig(config)
}

func (r *ClientConfigRepo) UpdateServerURL(url string) error {
	config, err := r.Load()
	if err != nil {
		return err
	}
	config.ServerURL = url
	return r.SetConfig(config)
}

func (r *ClientConfigRepo) SetConfigValue(key string, value interface{}) error {
	viper.Set(key, value)

	if err := viper.WriteConfig(); err != nil {
		return fmt.Errorf("could not write config file: %w", err)
	}

	return nil
}

func (r *ClientConfigRepo) GetConfigPath() string {
	return filepath.Join(r.configPath, configName+"."+configType)
}

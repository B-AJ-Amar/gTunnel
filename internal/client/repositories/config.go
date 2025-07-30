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

type ConfigRepository interface {
	// creates config file if it doesn't exist
	InitConfig() error

	// loads the configuration from file into the ClientConfig struct
	Load() (*models.ClientConfig, error)

	// saves the entire configuration to file
	Save(config *models.ClientConfig) error

	SetConfig(config *models.ClientConfig) error

	UpdateAccessToken(token string) error

	UpdateServerURL(url string) error

	SetConfigValue(key string, value interface{}) error

	GetConfigPath() string
}

type ConfigRepo struct {
	configPath string
}

// NewConfigRepo creates a new ConfigRepository implementation using Viper
func NewConfigRepo() ConfigRepository {
	// Determine platform-specific config path
	configDir, err := os.UserConfigDir()
	if err != nil {
		// Fallback to current directory if UserConfigDir fails
		configDir = "."
	}

	configPath := filepath.Join(configDir, appName)
	return &ConfigRepo{configPath: configPath}
}

func (r *ConfigRepo) InitConfig() error {
	// Create the config directory if it doesn't exist
	if err := os.MkdirAll(r.configPath, 0755); err != nil {
		return fmt.Errorf("could not create config directory: %w", err)
	}

	viper.AddConfigPath(r.configPath)
	viper.SetConfigName(configName)
	viper.SetConfigType(configType)

	// Try to read existing config
	if err := viper.ReadInConfig(); err != nil {
		// Config file doesn't exist, create it with default values
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("No existing config file found. Creating default config...")

			// Set default values
			defaultConfig := &models.ClientConfig{
				AccessToken: "",
				ServerURL:   "localhost:8080/___gTl___/ws",
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

func (r *ConfigRepo) Load() (*models.ClientConfig, error) {
	var config models.ClientConfig

	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("could not unmarshal config: %w", err)
	}

	return &config, nil
}

func (r *ConfigRepo) Save(config *models.ClientConfig) error {
	return r.SetConfig(config)
}

func (r *ConfigRepo) SetConfig(config *models.ClientConfig) error {
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

func (r *ConfigRepo) UpdateAccessToken(token string) error {
	config, err := r.Load()
	if err != nil {
		return err
	}
	config.AccessToken = token
	return r.SetConfig(config)
}

func (r *ConfigRepo) UpdateServerURL(url string) error {
	config, err := r.Load()
	if err != nil {
		return err
	}
	config.ServerURL = url
	return r.SetConfig(config)
}

func (r *ConfigRepo) SetConfigValue(key string, value interface{}) error {
	viper.Set(key, value)

	if err := viper.WriteConfig(); err != nil {
		return fmt.Errorf("could not write config file: %w", err)
	}

	return nil
}

func (r *ConfigRepo) GetConfigPath() string {
	return filepath.Join(r.configPath, configName+"."+configType)
}

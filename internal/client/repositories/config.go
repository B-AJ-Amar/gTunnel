package repositories

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"

	"github.com/B-AJ-Amar/gTunnel/internal/client/models"
)

type ConfigRepo struct {
	configPath string
}


func NewConfigRepo() *ConfigRepo {
	configPath := filepath.Join(".", "gtunnel.json")
	return &ConfigRepo{configPath: configPath}
}

func (r *ConfigRepo) Load() (*models.ClientConfig, error) {
	file, err := os.Open(r.configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var config models.ClientConfig
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return nil, err
	}
	return &config, nil
}

func (r *ConfigRepo) Save(config *models.ClientConfig) error {
	if config == nil {
		return errors.New("config is nil")
	}
	file, err := os.Create(r.configPath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(config)
}

func (r *ConfigRepo) UpdateAccessToken(token string) error {
	config, err := r.Load()
	if err != nil {
		return err
	}
	config.AccessToken = token
	return r.Save(config)
}

func (r *ConfigRepo) UpdateServerURL(url string) error {
	config, err := r.Load()
	if err != nil {
		return err
	}
	config.ServerURL = url
	return r.Save(config)
}
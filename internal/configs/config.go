package configs

import (
	"fmt"
	"os"

	"github.com/TrNgTien/new-feed-go/configs"
	"gopkg.in/yaml.v2"
)

type ConfigFilePath string

type Config struct {
	Log      Log      `yaml:"log"`
	Database Database `yaml:"database"`
	Cache    Cache    `yaml:"cache"`
	Storage  Storage  `yaml:"storage"`
}

func NewConfig(filePath ConfigFilePath) (Config, error) {
	var (
		configBytes = configs.DefaultConfigBytes
		config      = Config{}
		err         error
	)

	if filePath != "" {
		configBytes, err = os.ReadFile(string(filePath))
		if err != nil {
			return Config{}, fmt.Errorf("failed to read YAML file: %w", err)
		}
	}

	err = yaml.Unmarshal(configBytes, &config)
	if err != nil {
		return Config{}, fmt.Errorf("failed to unmarshal YAML: %w", err)
	}

	return config, nil
}

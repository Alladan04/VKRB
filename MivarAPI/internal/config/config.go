package configer

import (
	"errors"
	"fmt"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type Config struct {
	MatrixFilePath []string      `yaml:"matrix_file_path"`
	InitTimeout    time.Duration `yaml:"init_timeout"`
}

func LoadConfig(configPath string) (*Config, error) {
	// 1. Проверяем существование файла
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("config file %s does not exist", configPath)
	}

	// 2. Читаем файл
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	// 3. Парсим YAML
	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("failed to parse config: %w", err)
	}

	// 4. Валидация
	if len(cfg.MatrixFilePath) == 0 {
		return nil, errors.New("matrix_file_path cannot be empty")
	}

	if cfg.InitTimeout <= 0 {
		return nil, errors.New("init_timeout must be positive")
	}

	return &cfg, nil
}

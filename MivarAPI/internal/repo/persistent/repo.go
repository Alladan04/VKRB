package persistent

import (
	"fmt"
	"os"

	configer "mivar_robot_api/internal/config"
	"mivar_robot_api/utils"
)

type Repo struct {
	cfg configer.Config
}

func New(config configer.Config) *Repo {
	return &Repo{
		cfg: config,
	}
}

// GetLabirintByID Читает матрицу из файла. Если файла или пути в конфиге нет - вернет ошибку
func (r *Repo) GetLabirintByID(id string) ([][]uint8, error) {
	var filePath string
	for _, modelCfg := range r.cfg.Model {
		if modelCfg.ModelID == id {
			filePath = modelCfg.FilePath
			break
		}
	}

	if filePath == "" {
		return nil, fmt.Errorf("filePathIsEmpty")
	}

	matrix, err := utils.ReadMatrixFromFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("readMatrixFromFile: %w", err)
	}

	return matrix, nil
}

// GetModelByID Читает модель из файла. Если файла или пути в конфиге нет - вернет ошибку
func (r *Repo) GetModelByID(id string) ([]byte, error) {
	var filePath string
	for _, modelCfg := range r.cfg.Model {
		if modelCfg.ModelID == id {
			filePath = modelCfg.ModelXmlPath
			break
		}
	}

	if filePath == "" {
		return nil, fmt.Errorf("filePathIsEmpty")
	}

	model, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("readModelFromFile: %w", err)
	}

	return model, nil
}

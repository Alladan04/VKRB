package calculate_path

import (
	"context"

	"mivar_robot_api/internal/client/dto"
	"mivar_robot_api/internal/entity"
	"mivar_robot_api/pkg/generator"
)

type ModelGenerator interface {
	GenerateModelFromLabyrinth(matrixHardCoded [][]uint8, modelID string) ([]byte, error)
	MarshalModel(model generator.Model) ([]byte, error)
}

type Repo interface {
	GetModelFromCache(key string) ([]byte, error)
	UpsertModelToCache(data []byte, key string)
	UpsertLabirintToCache(data [][]uint8, key string)
	GetLabirintFromCache(key string) ([][]uint8, error)
}

type Client interface {
	AddModel(ctx context.Context, in dto.AddModelRequest) (*dto.AddModelResponse, error)
	DeleteModel(ctx context.Context, modelID string) error
}

type ModelManager interface {
	GetExitsByModelID(modelID string) ([]entity.Point, error)
}

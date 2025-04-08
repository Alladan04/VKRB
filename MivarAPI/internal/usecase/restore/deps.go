package restore

import (
	"context"

	"mivar_robot_api/internal/client/dto"
)

type Repo interface {
	UpsertModelToCache(data []byte, key string)
	UpsertLabirintToCache(data [][]uint8, key string)
}

type Client interface {
	AddModel(ctx context.Context, in dto.AddModelRequest) (*dto.AddModelResponse, error)
	DeleteModel(ctx context.Context, modelID string) error
}

type PersistentRepo interface {
	GetLabirintByID(id string) ([][]uint8, error)
	GetModelByID(id string) ([]byte, error)
}

package model_manager

import (
	"context"

	"mivar_robot_api/internal/client/dto"
	"mivar_robot_api/pkg/generator"
)

type ModelGenerator interface {
	UnmarshalModel(xmlData []byte) (generator.Model, error)
	GetParameterIDsByCoordinates(model generator.Model, coordinates []generator.Coordinate) ([]string, error)
	GetCoordinatesByParameterIDs(model generator.Model, ids []string) (map[string]generator.Coordinate, error)
	GenerateModelFromLabyrinth(matrixHardCoded [][]uint8, modelID string) ([]byte, error)
}

type ModelRepo interface {
	GetFromCache(key string) ([]byte, error)
	AddToCache(data []byte, key string)
}

type MivarClient interface {
	AddModel(ctx context.Context, in dto.AddModelRequest) (*dto.AddModelResponse, error)
}

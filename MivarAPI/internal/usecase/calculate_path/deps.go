package calculate_path

import (
	"context"

	"mivar_robot_api/internal/client/dto"
	"mivar_robot_api/internal/entity"
	"mivar_robot_api/pkg/generator"
)

type ModelGenerator interface {
	UnmarshalModel(xmlData []byte) (generator.Model, error)
	GetParameterIDsByCoordinates(model generator.Model, coordinates []generator.Coordinate) ([]string, error)
	GetCoordinatesByParameterIDs(model generator.Model, ids []string) (map[string]generator.Coordinate, error)
}

type ModelRepo interface {
	GetFromCache(key string) ([]byte, error)
}

type Client interface {
	CalculatePath(ctx context.Context, in dto.CalculateRrequest) (*dto.CalculateResponse, error)
}

type ModelManager interface {
	GetExitsByModelID(modelID string) ([]entity.Point, error)
}

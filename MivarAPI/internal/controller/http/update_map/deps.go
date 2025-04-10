package update_map

import (
	"context"

	"mivar_robot_api/internal/entity"
)

type Usecase interface {
	UpdateMap(
		ctx context.Context,
		points []entity.Point,
		modelID string,
	) ([][]uint8, error)
}

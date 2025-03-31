package calc_path

import (
	"context"

	"mivar_robot_api/internal/entity"
)

type Usecase interface {
	CalculatePath(
		ctx context.Context,
		start entity.Point,
		end []entity.Point,
		modelID string,
	) ([]entity.Transition, int64, error)
}

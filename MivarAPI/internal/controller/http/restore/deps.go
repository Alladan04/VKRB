package restore

import (
	"context"
)

type Usecase interface {
	RestoreModel(
		ctx context.Context,
		modelID string,
	) ([][]uint8, error)
}
